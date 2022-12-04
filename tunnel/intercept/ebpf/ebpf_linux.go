/*
	Copyright NetFoundry Inc.

	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

	https://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
*/

package ebpf

import (
	"context"
	"fmt"
	"github.com/michaelquigley/pfxlog"
	"github.com/openziti/edge/tunnel"
	"github.com/openziti/edge/tunnel/dns"
	"github.com/openziti/edge/tunnel/entities"
	"github.com/openziti/edge/tunnel/intercept"
	"github.com/openziti/edge/tunnel/router"
	"github.com/openziti/edge/tunnel/udp_vconn"
	"github.com/openziti/foundation/v2/info"
	"github.com/openziti/foundation/v2/mempool"
	"github.com/openziti/foundation/v2/stringz"
	"github.com/openziti/sdk-golang/ziti/edge/impl"
	cmap "github.com/orcaman/concurrent-map/v2"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"golang.org/x/sys/unix"
	"net"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
)

// https://github.com/torvalds/linux/blob/master/Documentation/networking/tproxy.txt

// Configure listening sockets with options that must be set before the socket is bound to an address (IP_TRANSPARENT).
var listenConfig = net.ListenConfig{
	Control: func(network, address string, c syscall.RawConn) error {
		var sockOptErr error
		controlErr := c.Control(func(sockFd uintptr) {
			// - https://www.kernel.org/doc/Documentation/networking/tproxy.txt
			if err := unix.SetsockoptInt(int(sockFd), unix.IPPROTO_IP, unix.IP_TRANSPARENT, 1); err != nil {
				sockOptErr = fmt.Errorf("error setting IP_TRANSPARENT socket option: %v", err)
				return
			}
			if err := unix.SetsockoptInt(int(sockFd), unix.SOL_SOCKET, unix.SO_REUSEADDR, 1); err != nil {
				sockOptErr = fmt.Errorf("error setting SO_REUSEADDR socket option: %v", err)
				return
			}

			if err := unix.SetsockoptInt(int(sockFd), syscall.SOL_IP, unix.IP_RECVORIGDSTADDR, 1); err != nil {
				sockOptErr = fmt.Errorf("error setting SO_REUSEADDR socket option: %v", err)
				return
			}
		})
		if controlErr != nil {
			return fmt.Errorf("error invoking listener socket control function: %v", controlErr)
		}
		return sockOptErr
	},
}

func New() (intercept.Interceptor, error) {
	cmd := exec.Command("map_update", "-V")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil, errors.Wrap(err, "ebpf: failed to verify map_update binary")
	}
	pfxlog.Logger().Infof("%v", string(out))
	_, err = os.Stat("/sys/fs/bpf/tc/globals/zt_tproxy_map")
	if err != nil {
		return nil, errors.Wrap(err, "ebpf: failed to verify hash map zt_tproxy_map exists")
	}
	_, err = os.Stat("/sys/fs/bpf/tc/globals/ifindex_ip_map")
	if err != nil {
		return nil, errors.Wrap(err, "ebpf: failed to verify array map ifindex_ip_map exists")
	}

	return &interceptor{
		serviceProxies: cmap.New[*eBpf](),
	}, nil
}

type alwaysRemoveAddressTracker struct{}

func (a alwaysRemoveAddressTracker) AddAddress(string) {}

func (a alwaysRemoveAddressTracker) RemoveAddress(string) bool {
	return true
}

type interceptor struct {
	serviceProxies cmap.ConcurrentMap[*eBpf]
}

func (self *interceptor) Stop() {
	self.serviceProxies.IterCb(func(key string, proxy *eBpf) {
		proxy.Stop(alwaysRemoveAddressTracker{})
	})
	self.serviceProxies.Clear()
}

func (self *interceptor) Intercept(service *entities.Service, resolver dns.Resolver, tracker intercept.AddressTracker) error {
	ebpf, err := self.newEbpf(service, resolver, tracker)
	if err != nil {
		return err
	}
	self.serviceProxies.Set(service.Name, ebpf)
	return nil
}

func (self *interceptor) StopIntercepting(serviceName string, tracker intercept.AddressTracker) error {
	if proxy, found := self.serviceProxies.Get(serviceName); found {
		proxy.Stop(tracker)
		self.serviceProxies.Remove(serviceName)
	}
	return nil
}

func (self *interceptor) newEbpf(service *entities.Service, resolver dns.Resolver, tracker intercept.AddressTracker) (*eBpf, error) {
	t := &eBpf{
		interceptor: self,
		service:     service,
		tracker:     tracker,
		resolver:    resolver,
	}

	config := service.InterceptV1Config

	if config == nil {
		return nil, errors.Errorf("service %v has no intercept information", service.Name)
	}

	if stringz.Contains(config.Protocols, "tcp") {
		tcpLn, err := listenConfig.Listen(context.Background(), "tcp", "127.0.0.1:")
		if err != nil {
			return nil, errors.Wrapf(err, "failed to create TCP listener for service: %v", service.Name)
		}
		logrus.Infof("ebpf-tproxy listening on tcp:%s", tcpLn.Addr().String())
		t.tcpLn = tcpLn
	}

	if stringz.Contains(config.Protocols, "udp") {
		packetLn, err := listenConfig.ListenPacket(context.Background(), "udp", "127.0.0.1:")
		if err != nil {
			return nil, errors.Wrapf(err, "failed to create UDP listener for service: %v", service.Name)
		}
		udpLn, ok := packetLn.(*net.UDPConn)
		if !ok {
			return nil, errors.New("failed to create UDP listener. listener was not net.UDPConn")
		}
		logrus.Infof("ebpf-tproxy listening on udp:%s, remoteAddr: %v", udpLn.LocalAddr(), udpLn.RemoteAddr())
		t.udpLn = udpLn
	}

	if t.tcpLn == nil && t.udpLn == nil {
		return nil, errors.Errorf("service %v has no supported protocols (tcp, udp). Serivce protocols: %+v", service.Name, config.Protocols)
	}

	if t.tcpLn != nil {
		go t.acceptTCP()
	}

	if t.udpLn != nil {
		go t.acceptUDP()
	}

	return t, t.Intercept(resolver, tracker)
}

type eBpf struct {
	interceptor *interceptor
	service     *entities.Service
	addresses   []*intercept.InterceptAddress
	tcpLn       net.Listener
	udpLn       *net.UDPConn
	tracker     intercept.AddressTracker
	resolver    dns.Resolver
}

func (self *eBpf) acceptTCP() {
	log := pfxlog.Logger()
	for {
		client, err := self.tcpLn.Accept()
		if err != nil {
			log.Errorf("error while accepting: %v", err)
		}
		if client == nil {
			log.Info("shutting down")
			return
		}
		log.Infof("received connection: %s --> %s", client.LocalAddr().String(), client.RemoteAddr().String())
		dstIp, dstPort := tunnel.GetIpAndPort(client.LocalAddr())
		dstHostname, _ := self.resolver.Lookup(client.LocalAddr().(*net.TCPAddr).IP)
		sourceAddr := self.service.GetSourceAddr(client.RemoteAddr(), client.LocalAddr())
		appInfo := tunnel.GetAppInfo("tcp", dstHostname, dstIp, dstPort, sourceAddr)
		identity := self.service.GetDialIdentity(client.RemoteAddr(), client.LocalAddr())
		go tunnel.DialAndRun(self.service, identity, client, appInfo, true)
	}
}

func (self *eBpf) acceptUDP() {
	vconnMgr := udp_vconn.NewManager(self.service.GetFabricProvider(), udp_vconn.NewUnlimitedConnectionPolicy(), udp_vconn.NewDefaultExpirationPolicy())
	self.generateReadEvents(vconnMgr)
}

func (self *eBpf) generateReadEvents(manager udp_vconn.Manager) {
	oobSize := 1600
	bufPool := mempool.NewPool(16, info.MaxUdpPacketSize+oobSize)
	log := pfxlog.Logger()

	for {
		pooled := bufPool.AcquireBuffer()
		oob := pooled.Buf[info.MaxUdpPacketSize:]
		pooled.Buf = pooled.Buf[:info.MaxUdpPacketSize]
		log.Debugf("waiting for datagram")
		n, oobn, _, srcAddr, err := self.udpLn.ReadMsgUDP(pooled.Buf, oob)
		if err != nil {
			log.WithError(err).Error("failure while reading udp message. stopping UDP read loop")
			manager.QueueError(err)
			return
		}
		log.Debugf("received %d bytes from %s", n, srcAddr.String())
		pooled.Buf = pooled.Buf[:n]
		event := &udpReadEvent{
			interceptor: self,
			buf:         pooled,
			oob:         oob[:oobn],
			srcAddr:     srcAddr,
		}
		manager.QueueEvent(event)
	}
}

type udpReadEvent struct {
	interceptor *eBpf
	buf         *mempool.DefaultPooledBuffer
	oob         []byte
	srcAddr     net.Addr
}

func (event *udpReadEvent) Handle(manager udp_vconn.Manager) error {
	writeQueue := manager.GetWriteQueue(event.srcAddr)

	if writeQueue == nil {
		log := pfxlog.Logger()
		origDest, err := getOriginalDest(event.oob)
		if err != nil {
			event.buf.Release()
			return fmt.Errorf("error while getting original destination packet: %v", err)
		}
		log.Infof("received datagram from %v (original dest %v). Creating udp listen socket on original dest", event.srcAddr, origDest)
		packetConn, err := listenConfig.ListenPacket(context.Background(), "udp", origDest.String())
		if err != nil {
			event.buf.Release()
			return err
		}
		writeConn := packetConn.(*net.UDPConn)
		writeQueue, err = manager.CreateWriteQueue(origDest, event.srcAddr, event.interceptor.service, writeConn)
		if err != nil {
			event.buf.Release()
			return err
		}
	}

	pfxlog.Logger().Debugf("received %v bytes for conn %v -> %v", len(event.buf.Buf), writeQueue.LocalAddr().String(), writeQueue.Service())
	writeQueue.Accept(event.buf)

	return nil
}

func getOriginalDest(oob []byte) (*net.UDPAddr, error) {
	cmsgs, err := syscall.ParseSocketControlMessage(oob)
	if err != nil {
		return nil, err
	}
	for _, cmsg := range cmsgs {
		if cmsg.Header.Level == syscall.SOL_IP && cmsg.Header.Type == syscall.IP_ORIGDSTADDR {
			ip := cmsg.Data[4:8]
			port := int(cmsg.Data[2])<<8 + int(cmsg.Data[3])
			return &net.UDPAddr{IP: ip, Port: port}, nil
		}
	}
	return nil, fmt.Errorf("original destination not found in out of band data")
}

func (self *eBpf) Stop(tracker intercept.AddressTracker) {
	log := pfxlog.Logger().WithField("service", self.service.Name)
	if self.tcpLn != nil {
		if err := self.tcpLn.Close(); err != nil {
			log.WithError(err).Error("failed to close TCP listener")
		}
	}

	if self.udpLn != nil {
		if err := self.udpLn.Close(); err != nil {
			log.WithError(err).Error("failed to close UDP listener")
		}
	}

	err := self.StopIntercepting(tracker)
	if err != nil {
		log.WithError(err).Error("failed to clean up intercept configuration")
	}
}

func (self *eBpf) tcpPort() IPPortAddr {
	if self.tcpLn != nil {
		return (*TCPIPPortAddr)(self.tcpLn.Addr().(*net.TCPAddr))
	}
	logrus.Errorf("invalid state: no tcp listener for tproxy[%s]", self.service.Name)
	return nil
}

func (self *eBpf) udpPort() IPPortAddr {
	if self.udpLn != nil {
		return (*UDPIPPortAddr)(self.udpLn.LocalAddr().(*net.UDPAddr))
	}

	logrus.Errorf("invalid state: no udp listener for tproxy[%s]", self.service.Name)
	return nil
}

func (self *eBpf) Intercept(resolver dns.Resolver, tracker intercept.AddressTracker) error {
	service := self.service
	if service.InterceptV1Config == nil {
		return errors.Errorf("no client configuration for service %v", service.Name)
	}

	config := service.InterceptV1Config
	logrus.Debugf("service %v using intercept.v1", service.Name)
	var ports []IPPortAddr
	for _, p := range config.Protocols {
		if p == "tcp" {
			logrus.Debugf("service %v intercepting tcp", service.Name)
			ports = append(ports, self.tcpPort())
		} else if p == "udp" {
			logrus.Debugf("service %v intercepting udp", service.Name)
			ports = append(ports, self.udpPort())
		}
	}

	return self.intercept(service, resolver, ports, tracker)
}

func (self *eBpf) Apply(addr *intercept.InterceptAddress) {
	logrus.Debugf("for service %v, intercepting proto: %v, cidr: %v, ports: %v:%v", self.service.Name, addr.Proto(), addr.IpNet(), addr.LowPort(), addr.HighPort())

	var port IPPortAddr
	switch addr.Proto() {
	case "tcp":
		port = self.tcpPort()
	case "udp":
		port = self.udpPort()
	default:
		logrus.Errorf("unknown proto[%s] for tproxy[%s]", addr.Proto(), self.service.Name)
		return
	}
	if err := self.addInterceptAddr(addr, self.service, port, self.tracker); err != nil {
		logrus.Debugf("failed for service %v, intercepting proto: %v, cidr: %v, ports: %v:%v", self.service.Name, addr.Proto(), addr.IpNet(), addr.LowPort(), addr.HighPort())

		// do we undo the previous succesful ones?
		// only fail at end and return all that failed?
	}
}

func (self *eBpf) intercept(service *entities.Service, resolver dns.Resolver, ports []IPPortAddr, tracker intercept.AddressTracker) error {
	var protocols []string
	for _, p := range ports {
		protocols = append(protocols, p.GetProtocol())
	}

	err := intercept.GetInterceptAddresses(service, protocols, resolver, self)
	if err != nil {
		return err
	}

	return nil
}

func (self *eBpf) addInterceptAddr(interceptAddr *intercept.InterceptAddress, service *entities.Service, port IPPortAddr, tracker intercept.AddressTracker) error {
	ipNet := interceptAddr.IpNet()
	if err := router.AddLocalAddress(ipNet, "lo"); err != nil {
		return errors.Wrapf(err, "failed to add local route %v", ipNet)
	}
	tracker.AddAddress(ipNet.String())
	self.addresses = append(self.addresses, interceptAddr)
	ipNetList := strings.Split(ipNet.String(), "/")
	low_port := strconv.Itoa(int(interceptAddr.LowPort()))
	high_port := strconv.Itoa(int(interceptAddr.HighPort()))
	tproxy_port := strconv.Itoa(int(port.GetPort()))
	pfxlog.Logger().WithField("dst_ip", ipNet).WithField("protocol", interceptAddr.Proto()).WithField("low-port",
		interceptAddr.LowPort()).WithField("high-port", interceptAddr.HighPort()).WithField("tproxy-ip",
		port.GetIP()).WithField("tproxy-port", port.GetPort()).Info("setting up intercept: ", interceptAddr.Proto())
	cmd := exec.Command("map_update", "-I", "-c", ipNetList[0], "-m", ipNetList[1], "-l", low_port, "-h",
		high_port, "-t", tproxy_port, "-p", interceptAddr.Proto())
	out, err := cmd.CombinedOutput()
	if err != nil {
		pfxlog.Logger().Infof("Failed to insert entry to ebpf hash table for %v : %v", ipNet.String(), string(out))
	} else {
		pfxlog.Logger().Infof("Updated ebpf zt_tproxy_map: map_update -I -c %v -m %v -l %v -h %v -t %v -p %v",
			ipNetList[0], ipNetList[1], low_port, high_port, tproxy_port, interceptAddr.Proto())
	}
	return nil
}

func (self *eBpf) StopIntercepting(tracker intercept.AddressTracker) error {
	var errorList []error

	log := pfxlog.Logger().WithField("sevice", self.service.Name)

	for _, addr := range self.addresses {
		ipNetList := strings.Split(addr.IpNet().String(), "/")
		log.Infof("removing service entry from ebpf zt_tproxy_map: dst_prefix: %v dest mask: %v low-port: %v, high-port: %v", ipNetList[0], ipNetList[1], addr.LowPort(), addr.HighPort())
		cmd := exec.Command("map_update", "-D", "-c", ipNetList[0], "-m", ipNetList[1], "-l", strconv.Itoa(int(addr.LowPort())), "-p", addr.Proto())
		out, err := cmd.CombinedOutput()
		if err != nil {
			pfxlog.Logger().Infof("Failed to remove entry from ebpf hash table for %v port: %v Protocol: %v : %v", addr.IpNet().String(), addr.LowPort(), addr.Proto(), string(out))
		} else {
			pfxlog.Logger().Infof("Updated ebpf zt_tproxy_map: map_update -D -c %v -m %v -l %v -p %v", ipNetList[0], ipNetList[1], addr.LowPort(), addr.Proto())
		}
		ipNet := addr.IpNet()
		if tracker.RemoveAddress(ipNet.String()) {
			err := router.RemoveLocalAddress(ipNet, "lo")
			if err != nil {
				errorList = append(errorList, err)
				log.WithError(err).Errorf("failed to remove route %v for service %s", ipNet, self.service.Name)
			}
		}
	}
	if len(errorList) == 0 {
		return nil
	}
	if len(errorList) == 1 {
		return errorList[0]
	}
	return impl.MultipleErrors(errorList)
}

type IPPortAddr interface {
	GetIP() net.IP
	GetPort() int
	GetProtocol() string
}

type UDPIPPortAddr net.UDPAddr

func (addr *UDPIPPortAddr) GetIP() net.IP {
	return addr.IP
}

func (addr *UDPIPPortAddr) GetPort() int {
	return addr.Port
}

func (addr *UDPIPPortAddr) GetProtocol() string {
	return "udp"
}

type TCPIPPortAddr net.TCPAddr

func (addr *TCPIPPortAddr) GetIP() net.IP {
	return addr.IP
}

func (addr *TCPIPPortAddr) GetPort() int {
	return addr.Port
}

func (addr *TCPIPPortAddr) GetProtocol() string {
	return "tcp"
}
