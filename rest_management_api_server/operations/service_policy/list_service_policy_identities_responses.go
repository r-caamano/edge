// Code generated by go-swagger; DO NOT EDIT.

//
// Copyright NetFoundry Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// __          __              _
// \ \        / /             (_)
//  \ \  /\  / /_ _ _ __ _ __  _ _ __   __ _
//   \ \/  \/ / _` | '__| '_ \| | '_ \ / _` |
//    \  /\  / (_| | |  | | | | | | | | (_| | : This file is generated, do not edit it.
//     \/  \/ \__,_|_|  |_| |_|_|_| |_|\__, |
//                                      __/ |
//                                     |___/

package service_policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openziti/edge/rest_model"
)

// ListServicePolicyIdentitiesOKCode is the HTTP code returned for type ListServicePolicyIdentitiesOK
const ListServicePolicyIdentitiesOKCode int = 200

/*ListServicePolicyIdentitiesOK A list of identities

swagger:response listServicePolicyIdentitiesOK
*/
type ListServicePolicyIdentitiesOK struct {

	/*
	  In: Body
	*/
	Payload *rest_model.ListIdentitiesEnvelope `json:"body,omitempty"`
}

// NewListServicePolicyIdentitiesOK creates ListServicePolicyIdentitiesOK with default headers values
func NewListServicePolicyIdentitiesOK() *ListServicePolicyIdentitiesOK {

	return &ListServicePolicyIdentitiesOK{}
}

// WithPayload adds the payload to the list service policy identities o k response
func (o *ListServicePolicyIdentitiesOK) WithPayload(payload *rest_model.ListIdentitiesEnvelope) *ListServicePolicyIdentitiesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list service policy identities o k response
func (o *ListServicePolicyIdentitiesOK) SetPayload(payload *rest_model.ListIdentitiesEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListServicePolicyIdentitiesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListServicePolicyIdentitiesBadRequestCode is the HTTP code returned for type ListServicePolicyIdentitiesBadRequest
const ListServicePolicyIdentitiesBadRequestCode int = 400

/*ListServicePolicyIdentitiesBadRequest The requested resource does not exist

swagger:response listServicePolicyIdentitiesBadRequest
*/
type ListServicePolicyIdentitiesBadRequest struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewListServicePolicyIdentitiesBadRequest creates ListServicePolicyIdentitiesBadRequest with default headers values
func NewListServicePolicyIdentitiesBadRequest() *ListServicePolicyIdentitiesBadRequest {

	return &ListServicePolicyIdentitiesBadRequest{}
}

// WithPayload adds the payload to the list service policy identities bad request response
func (o *ListServicePolicyIdentitiesBadRequest) WithPayload(payload *rest_model.APIErrorEnvelope) *ListServicePolicyIdentitiesBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list service policy identities bad request response
func (o *ListServicePolicyIdentitiesBadRequest) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListServicePolicyIdentitiesBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListServicePolicyIdentitiesUnauthorizedCode is the HTTP code returned for type ListServicePolicyIdentitiesUnauthorized
const ListServicePolicyIdentitiesUnauthorizedCode int = 401

/*ListServicePolicyIdentitiesUnauthorized The currently supplied session does not have the correct access rights to request this resource

swagger:response listServicePolicyIdentitiesUnauthorized
*/
type ListServicePolicyIdentitiesUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewListServicePolicyIdentitiesUnauthorized creates ListServicePolicyIdentitiesUnauthorized with default headers values
func NewListServicePolicyIdentitiesUnauthorized() *ListServicePolicyIdentitiesUnauthorized {

	return &ListServicePolicyIdentitiesUnauthorized{}
}

// WithPayload adds the payload to the list service policy identities unauthorized response
func (o *ListServicePolicyIdentitiesUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *ListServicePolicyIdentitiesUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list service policy identities unauthorized response
func (o *ListServicePolicyIdentitiesUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListServicePolicyIdentitiesUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
