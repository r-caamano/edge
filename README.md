# Ziti Edge

This is a modified version of the openziti edge designed to work with the eBPF Interception FW https://github.com/r-caamano/ebpf-tproxy-splicer it is currently aligned with tagged release v0.26.10 of https://github.com/openziti/ziti.

Ziti is a modern, programmable network overlay with associated edge components, for application-embedded, zero trust network connectivity, written by developers for developers. Ziti allows developers to take control of networking while with secure connectivity and advanced security concepts such as Zero Trust.

The Ziti Edge implements the application-embedded connectivity and zero trust components of the Ziti ecosystem.

How to build:
     
Make sure that `go` is in your path:

```
$ go version
go version go1.19.2 linux/amd64
```

Consider using a separate directory to contain the `GOPATH` for your `ziti-edge` development (instead of just using `~/go`). i.e.  `~/local/ziti`

Essentially:

```
mkdir -p ~/local/ziti
export GOPATH=~/local/ziti
```

Include `GOPATH/bin` in your shell's `PATH`:

```
$ export PATH=$GOPATH/bin:$PATH
```

Next clone the repositories. i.e in `~/repos`.

```
mkdir ~/repos
cd ~/repos
git clone https://github.com/openziti/ziti.git
git clone https://github.com/r-caamano/edge.git
```

Update the `go.mod` file in the root of `ziti repo`, using the `replace` directive to point the `ziti-cmd` build at our local `ziti-edge` development tree.

```
cd ~/repos/ziti

git fetch -a -t
git checkout tags/v0.26.10 -b v0.26.10

vi go.mod
```

The top of the file should look like this:

```
module github.com/openziti/ziti

go 1.19

require (
```

Add a `replace` line, like:

```
module github.com/openziti/ziti

go 1.19

replace github.com/openziti/edge => ../edge

require (
```
	
With that change made, you can change the contents of your local clone of `ziti-edge`, and builds of `ziti` will use your local changes, rather than the version it pulled into `GOPATH` from GitHub.

Checkout the v0.26.10 branch of the modified edge:

```
cd ~/repos/edge
git fetch -a
git switch v0.26.10
```

Build the tree:

```
$ cd ~/repos/ziti
$ go install ./...
```

The binaries will be placed in `$GOPATH/bin`.
     
