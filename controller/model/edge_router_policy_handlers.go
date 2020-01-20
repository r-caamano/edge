/*
	Copyright 2019 Netfoundry, Inc.

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

package model

import (
	"github.com/netfoundry/ziti-edge/controller/persistence"
	"github.com/netfoundry/ziti-foundation/storage/boltz"
	"go.etcd.io/bbolt"
)

func NewEdgeRouterPolicyHandler(env Env) *EdgeRouterPolicyHandler {
	handler := &EdgeRouterPolicyHandler{
		baseHandler: baseHandler{
			env:   env,
			store: env.GetStores().EdgeRouterPolicy,
		},
	}
	handler.impl = handler
	return handler
}

type EdgeRouterPolicyHandler struct {
	baseHandler
}

func (handler *EdgeRouterPolicyHandler) NewModelEntity() BaseModelEntity {
	return &EdgeRouterPolicy{}
}

func (handler *EdgeRouterPolicyHandler) Create(edgeRouterPolicy *EdgeRouterPolicy) (string, error) {
	return handler.createEntity(edgeRouterPolicy, nil)
}

func (handler *EdgeRouterPolicyHandler) Read(id string) (*EdgeRouterPolicy, error) {
	modelEntity := &EdgeRouterPolicy{}
	if err := handler.readEntity(id, modelEntity); err != nil {
		return nil, err
	}
	return modelEntity, nil
}

func (handler *EdgeRouterPolicyHandler) readInTx(tx *bbolt.Tx, id string) (*EdgeRouterPolicy, error) {
	modelEntity := &EdgeRouterPolicy{}
	if err := handler.readEntityInTx(tx, id, modelEntity); err != nil {
		return nil, err
	}
	return modelEntity, nil
}

func (handler *EdgeRouterPolicyHandler) Update(edgeRouterPolicy *EdgeRouterPolicy) error {
	return handler.updateEntity(edgeRouterPolicy, nil, nil)
}

func (handler *EdgeRouterPolicyHandler) Patch(edgeRouterPolicy *EdgeRouterPolicy, checker boltz.FieldChecker) error {
	return handler.patchEntity(edgeRouterPolicy, checker, nil)
}

func (handler *EdgeRouterPolicyHandler) Delete(id string) error {
	return handler.deleteEntity(id, nil, nil)
}

func (handler *EdgeRouterPolicyHandler) List(queryOptions *QueryOptions) (*EdgeRouterPolicyListResult, error) {
	result := &EdgeRouterPolicyListResult{handler: handler}
	err := handler.parseAndList(queryOptions, result.collect)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (handler *EdgeRouterPolicyHandler) ListEdgeRouters(id string) ([]*EdgeRouter, error) {
	var result []*EdgeRouter
	err := handler.CollectEdgeRouters(id, func(entity BaseModelEntity) {
		result = append(result, entity.(*EdgeRouter))
	})

	if err != nil {
		return nil, err
	}
	return result, nil
}

func (handler *EdgeRouterPolicyHandler) CollectEdgeRouters(id string, collector func(entity BaseModelEntity)) error {
	return handler.collectAssociated(id, persistence.EntityTypeEdgeRouters, handler.env.GetHandlers().EdgeRouter, collector)
}

func (handler *EdgeRouterPolicyHandler) CollectIdentities(id string, collector func(entity BaseModelEntity)) error {
	return handler.collectAssociated(id, persistence.EntityTypeIdentities, handler.env.GetHandlers().Identity, collector)
}

type EdgeRouterPolicyListResult struct {
	handler            *EdgeRouterPolicyHandler
	EdgeRouterPolicies []*EdgeRouterPolicy
	QueryMetaData
}

func (result *EdgeRouterPolicyListResult) collect(tx *bbolt.Tx, ids []string, queryMetaData *QueryMetaData) error {
	result.QueryMetaData = *queryMetaData
	for _, key := range ids {
		entity, err := result.handler.readInTx(tx, key)
		if err != nil {
			return err
		}
		result.EdgeRouterPolicies = append(result.EdgeRouterPolicies, entity)
	}
	return nil
}