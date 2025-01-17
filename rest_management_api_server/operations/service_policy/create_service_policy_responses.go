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

// CreateServicePolicyCreatedCode is the HTTP code returned for type CreateServicePolicyCreated
const CreateServicePolicyCreatedCode int = 201

/*CreateServicePolicyCreated The create request was successful and the resource has been added at the following location

swagger:response createServicePolicyCreated
*/
type CreateServicePolicyCreated struct {

	/*
	  In: Body
	*/
	Payload *rest_model.CreateEnvelope `json:"body,omitempty"`
}

// NewCreateServicePolicyCreated creates CreateServicePolicyCreated with default headers values
func NewCreateServicePolicyCreated() *CreateServicePolicyCreated {

	return &CreateServicePolicyCreated{}
}

// WithPayload adds the payload to the create service policy created response
func (o *CreateServicePolicyCreated) WithPayload(payload *rest_model.CreateEnvelope) *CreateServicePolicyCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create service policy created response
func (o *CreateServicePolicyCreated) SetPayload(payload *rest_model.CreateEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateServicePolicyCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateServicePolicyBadRequestCode is the HTTP code returned for type CreateServicePolicyBadRequest
const CreateServicePolicyBadRequestCode int = 400

/*CreateServicePolicyBadRequest The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information

swagger:response createServicePolicyBadRequest
*/
type CreateServicePolicyBadRequest struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewCreateServicePolicyBadRequest creates CreateServicePolicyBadRequest with default headers values
func NewCreateServicePolicyBadRequest() *CreateServicePolicyBadRequest {

	return &CreateServicePolicyBadRequest{}
}

// WithPayload adds the payload to the create service policy bad request response
func (o *CreateServicePolicyBadRequest) WithPayload(payload *rest_model.APIErrorEnvelope) *CreateServicePolicyBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create service policy bad request response
func (o *CreateServicePolicyBadRequest) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateServicePolicyBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateServicePolicyUnauthorizedCode is the HTTP code returned for type CreateServicePolicyUnauthorized
const CreateServicePolicyUnauthorizedCode int = 401

/*CreateServicePolicyUnauthorized The currently supplied session does not have the correct access rights to request this resource

swagger:response createServicePolicyUnauthorized
*/
type CreateServicePolicyUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewCreateServicePolicyUnauthorized creates CreateServicePolicyUnauthorized with default headers values
func NewCreateServicePolicyUnauthorized() *CreateServicePolicyUnauthorized {

	return &CreateServicePolicyUnauthorized{}
}

// WithPayload adds the payload to the create service policy unauthorized response
func (o *CreateServicePolicyUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *CreateServicePolicyUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create service policy unauthorized response
func (o *CreateServicePolicyUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateServicePolicyUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
