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

package auth_policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openziti/edge/rest_model"
)

// DeleteAuthPolicyOKCode is the HTTP code returned for type DeleteAuthPolicyOK
const DeleteAuthPolicyOKCode int = 200

/*DeleteAuthPolicyOK The delete request was successful and the resource has been removed

swagger:response deleteAuthPolicyOK
*/
type DeleteAuthPolicyOK struct {

	/*
	  In: Body
	*/
	Payload *rest_model.Empty `json:"body,omitempty"`
}

// NewDeleteAuthPolicyOK creates DeleteAuthPolicyOK with default headers values
func NewDeleteAuthPolicyOK() *DeleteAuthPolicyOK {

	return &DeleteAuthPolicyOK{}
}

// WithPayload adds the payload to the delete auth policy o k response
func (o *DeleteAuthPolicyOK) WithPayload(payload *rest_model.Empty) *DeleteAuthPolicyOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete auth policy o k response
func (o *DeleteAuthPolicyOK) SetPayload(payload *rest_model.Empty) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteAuthPolicyOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteAuthPolicyBadRequestCode is the HTTP code returned for type DeleteAuthPolicyBadRequest
const DeleteAuthPolicyBadRequestCode int = 400

/*DeleteAuthPolicyBadRequest The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information

swagger:response deleteAuthPolicyBadRequest
*/
type DeleteAuthPolicyBadRequest struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewDeleteAuthPolicyBadRequest creates DeleteAuthPolicyBadRequest with default headers values
func NewDeleteAuthPolicyBadRequest() *DeleteAuthPolicyBadRequest {

	return &DeleteAuthPolicyBadRequest{}
}

// WithPayload adds the payload to the delete auth policy bad request response
func (o *DeleteAuthPolicyBadRequest) WithPayload(payload *rest_model.APIErrorEnvelope) *DeleteAuthPolicyBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete auth policy bad request response
func (o *DeleteAuthPolicyBadRequest) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteAuthPolicyBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteAuthPolicyUnauthorizedCode is the HTTP code returned for type DeleteAuthPolicyUnauthorized
const DeleteAuthPolicyUnauthorizedCode int = 401

/*DeleteAuthPolicyUnauthorized The currently supplied session does not have the correct access rights to request this resource

swagger:response deleteAuthPolicyUnauthorized
*/
type DeleteAuthPolicyUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewDeleteAuthPolicyUnauthorized creates DeleteAuthPolicyUnauthorized with default headers values
func NewDeleteAuthPolicyUnauthorized() *DeleteAuthPolicyUnauthorized {

	return &DeleteAuthPolicyUnauthorized{}
}

// WithPayload adds the payload to the delete auth policy unauthorized response
func (o *DeleteAuthPolicyUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *DeleteAuthPolicyUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete auth policy unauthorized response
func (o *DeleteAuthPolicyUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteAuthPolicyUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
