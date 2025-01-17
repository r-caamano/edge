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

package edge_router

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openziti/edge/rest_model"
)

// CreateEdgeRouterCreatedCode is the HTTP code returned for type CreateEdgeRouterCreated
const CreateEdgeRouterCreatedCode int = 201

/*CreateEdgeRouterCreated The create request was successful and the resource has been added at the following location

swagger:response createEdgeRouterCreated
*/
type CreateEdgeRouterCreated struct {

	/*
	  In: Body
	*/
	Payload *rest_model.CreateEnvelope `json:"body,omitempty"`
}

// NewCreateEdgeRouterCreated creates CreateEdgeRouterCreated with default headers values
func NewCreateEdgeRouterCreated() *CreateEdgeRouterCreated {

	return &CreateEdgeRouterCreated{}
}

// WithPayload adds the payload to the create edge router created response
func (o *CreateEdgeRouterCreated) WithPayload(payload *rest_model.CreateEnvelope) *CreateEdgeRouterCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create edge router created response
func (o *CreateEdgeRouterCreated) SetPayload(payload *rest_model.CreateEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateEdgeRouterCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateEdgeRouterBadRequestCode is the HTTP code returned for type CreateEdgeRouterBadRequest
const CreateEdgeRouterBadRequestCode int = 400

/*CreateEdgeRouterBadRequest The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information

swagger:response createEdgeRouterBadRequest
*/
type CreateEdgeRouterBadRequest struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewCreateEdgeRouterBadRequest creates CreateEdgeRouterBadRequest with default headers values
func NewCreateEdgeRouterBadRequest() *CreateEdgeRouterBadRequest {

	return &CreateEdgeRouterBadRequest{}
}

// WithPayload adds the payload to the create edge router bad request response
func (o *CreateEdgeRouterBadRequest) WithPayload(payload *rest_model.APIErrorEnvelope) *CreateEdgeRouterBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create edge router bad request response
func (o *CreateEdgeRouterBadRequest) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateEdgeRouterBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateEdgeRouterUnauthorizedCode is the HTTP code returned for type CreateEdgeRouterUnauthorized
const CreateEdgeRouterUnauthorizedCode int = 401

/*CreateEdgeRouterUnauthorized The currently supplied session does not have the correct access rights to request this resource

swagger:response createEdgeRouterUnauthorized
*/
type CreateEdgeRouterUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewCreateEdgeRouterUnauthorized creates CreateEdgeRouterUnauthorized with default headers values
func NewCreateEdgeRouterUnauthorized() *CreateEdgeRouterUnauthorized {

	return &CreateEdgeRouterUnauthorized{}
}

// WithPayload adds the payload to the create edge router unauthorized response
func (o *CreateEdgeRouterUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *CreateEdgeRouterUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create edge router unauthorized response
func (o *CreateEdgeRouterUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateEdgeRouterUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
