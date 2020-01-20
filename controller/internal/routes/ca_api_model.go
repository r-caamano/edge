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

package routes

import (
	"fmt"
	"github.com/michaelquigley/pfxlog"
	"github.com/netfoundry/ziti-edge/controller/env"
	"github.com/netfoundry/ziti-edge/controller/model"
	"github.com/netfoundry/ziti-edge/controller/response"
	"github.com/netfoundry/ziti-foundation/util/stringz"
	"net/http"
)

const EntityNameCa = "cas"

type CaApiUpdate struct {
	Tags                      map[string]interface{} `json:"tags"`
	Name                      *string                `json:"name"`
	IsAutoCaEnrollmentEnabled *bool                  `json:"isAutoCaEnrollmentEnabled"`
	IsOttCaEnrollmentEnabled  *bool                  `json:"isOttCaEnrollmentEnabled"`
	IsAuthEnabled             *bool                  `json:"isAuthEnabled"`
}

func (i *CaApiUpdate) ToModel(id string) *model.Ca {
	result := &model.Ca{}
	result.Id = id
	result.Name = stringz.OrEmpty(i.Name)
	if i.Tags != nil {
		result.Tags = i.Tags
	}
	result.IsAutoCaEnrollmentEnabled = i.IsAutoCaEnrollmentEnabled != nil && *i.IsAutoCaEnrollmentEnabled
	result.IsOttCaEnrollmentEnabled = i.IsOttCaEnrollmentEnabled != nil && *i.IsOttCaEnrollmentEnabled
	result.IsAuthEnabled = i.IsAuthEnabled != nil && *i.IsAuthEnabled
	return result
}

type CaApiCreate struct {
	Tags                      map[string]interface{} `json:"tags"`
	Name                      *string                `json:"name"`
	CertPem                   *string                `json:"certPem"`
	IsAutoCaEnrollmentEnabled *bool                  `json:"isAutoCaEnrollmentEnabled"`
	IsOttCaEnrollmentEnabled  *bool                  `json:"isOttCaEnrollmentEnabled"`
	IsAuthEnabled             *bool                  `json:"isAuthEnabled"`
}

func (i *CaApiCreate) ToModel() *model.Ca {
	result := &model.Ca{}
	result.Name = stringz.OrEmpty(i.Name)
	result.Tags = i.Tags
	result.CertPem = stringz.OrEmpty(i.CertPem)
	result.IsAutoCaEnrollmentEnabled = i.IsAutoCaEnrollmentEnabled != nil && *i.IsAutoCaEnrollmentEnabled
	result.IsOttCaEnrollmentEnabled = i.IsOttCaEnrollmentEnabled != nil && *i.IsOttCaEnrollmentEnabled
	result.IsAuthEnabled = i.IsAuthEnabled != nil && *i.IsAuthEnabled
	return result
}

type CaApiList struct {
	*env.BaseApi
	Name                      *string `json:"name"`
	Fingerprint               *string `json:"fingerprint"`
	CertPem                   *string `json:"certPem"`
	IsVerified                *bool   `json:"isVerified"`
	VerificationToken         *string `json:"verificationToken"`
	IsAutoCaEnrollmentEnabled *bool   `json:"isAutoCaEnrollmentEnabled"`
	IsOttCaEnrollmentEnabled  *bool   `json:"isOttCaEnrollmentEnabled"`
	IsAuthEnabled             *bool   `json:"isAuthEnabled"`
}

func (CaApiList) BuildSelfLink(id string) *response.Link {
	return response.NewLink(fmt.Sprintf("./%s/%s", EntityNameCa, id))
}

func (e *CaApiList) GetSelfLink() *response.Link {
	return e.BuildSelfLink(e.Id)
}

func (e *CaApiList) PopulateLinks() {
	if e.Links == nil {
		e.Links = &response.Links{
			EntityNameSelf: e.GetSelfLink(),
		}

		if !*e.IsVerified {
			vl := response.NewLink(fmt.Sprintf("%s/verify", (*e.Links)[EntityNameSelf].Href))
			vl.Method = http.MethodPost
			(*e.Links)["verify"] = vl
		}

		if *e.IsAutoCaEnrollmentEnabled {
			jl := response.NewLink(fmt.Sprintf("%s/jwt", (*e.Links)[EntityNameSelf].Href))
			jl.Method = http.MethodGet
			(*e.Links)["jwt"] = jl
		}
	}
}

func (e *CaApiList) ToEntityApiRef() *EntityApiRef {
	e.PopulateLinks()
	return &EntityApiRef{
		Entity: EntityNameCa,
		Name:   e.Name,
		Id:     e.Id,
		Links:  e.Links,
	}
}

func MapCaToApiEntity(_ *env.AppEnv, _ *response.RequestContext, e model.BaseModelEntity) (BaseApiEntity, error) {
	i, ok := e.(*model.Ca)

	if !ok {
		err := fmt.Errorf("entity is not a CA \"%s\"", e.GetId())
		log := pfxlog.Logger()
		log.Error(err)
		return nil, err
	}

	al, err := MapCaToApiList(i)

	if err != nil {
		err := fmt.Errorf("could not convert to API entity \"%s\": %s", e.GetId(), err)
		log := pfxlog.Logger()
		log.Error(err)
		return nil, err
	}
	return al, nil
}
func MapCaToApiList(i *model.Ca) (*CaApiList, error) {
	ret := &CaApiList{
		BaseApi:                   env.FromBaseModelEntity(i),
		Name:                      &i.Name,
		VerificationToken:         &i.VerificationToken,
		IsVerified:                &i.IsVerified,
		CertPem:                   &i.CertPem,
		Fingerprint:               &i.Fingerprint,
		IsAuthEnabled:             &i.IsAuthEnabled,
		IsAutoCaEnrollmentEnabled: &i.IsAutoCaEnrollmentEnabled,
		IsOttCaEnrollmentEnabled:  &i.IsOttCaEnrollmentEnabled,
	}

	ret.PopulateLinks()

	return ret, nil
}