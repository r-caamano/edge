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

package current_api_session

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openziti/edge/rest_model"
)

// CreateCurrentAPISessionCertificateCreatedCode is the HTTP code returned for type CreateCurrentAPISessionCertificateCreated
const CreateCurrentAPISessionCertificateCreatedCode int = 201

/*CreateCurrentAPISessionCertificateCreated A response of a create API Session certificate

swagger:response createCurrentApiSessionCertificateCreated
*/
type CreateCurrentAPISessionCertificateCreated struct {

	/*
	  In: Body
	*/
	Payload *rest_model.CreateCurrentAPISessionCertificateEnvelope `json:"body,omitempty"`
}

// NewCreateCurrentAPISessionCertificateCreated creates CreateCurrentAPISessionCertificateCreated with default headers values
func NewCreateCurrentAPISessionCertificateCreated() *CreateCurrentAPISessionCertificateCreated {

	return &CreateCurrentAPISessionCertificateCreated{}
}

// WithPayload adds the payload to the create current Api session certificate created response
func (o *CreateCurrentAPISessionCertificateCreated) WithPayload(payload *rest_model.CreateCurrentAPISessionCertificateEnvelope) *CreateCurrentAPISessionCertificateCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create current Api session certificate created response
func (o *CreateCurrentAPISessionCertificateCreated) SetPayload(payload *rest_model.CreateCurrentAPISessionCertificateEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateCurrentAPISessionCertificateCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateCurrentAPISessionCertificateBadRequestCode is the HTTP code returned for type CreateCurrentAPISessionCertificateBadRequest
const CreateCurrentAPISessionCertificateBadRequestCode int = 400

/*CreateCurrentAPISessionCertificateBadRequest The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information

swagger:response createCurrentApiSessionCertificateBadRequest
*/
type CreateCurrentAPISessionCertificateBadRequest struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewCreateCurrentAPISessionCertificateBadRequest creates CreateCurrentAPISessionCertificateBadRequest with default headers values
func NewCreateCurrentAPISessionCertificateBadRequest() *CreateCurrentAPISessionCertificateBadRequest {

	return &CreateCurrentAPISessionCertificateBadRequest{}
}

// WithPayload adds the payload to the create current Api session certificate bad request response
func (o *CreateCurrentAPISessionCertificateBadRequest) WithPayload(payload *rest_model.APIErrorEnvelope) *CreateCurrentAPISessionCertificateBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create current Api session certificate bad request response
func (o *CreateCurrentAPISessionCertificateBadRequest) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateCurrentAPISessionCertificateBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateCurrentAPISessionCertificateUnauthorizedCode is the HTTP code returned for type CreateCurrentAPISessionCertificateUnauthorized
const CreateCurrentAPISessionCertificateUnauthorizedCode int = 401

/*CreateCurrentAPISessionCertificateUnauthorized The currently supplied session does not have the correct access rights to request this resource

swagger:response createCurrentApiSessionCertificateUnauthorized
*/
type CreateCurrentAPISessionCertificateUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewCreateCurrentAPISessionCertificateUnauthorized creates CreateCurrentAPISessionCertificateUnauthorized with default headers values
func NewCreateCurrentAPISessionCertificateUnauthorized() *CreateCurrentAPISessionCertificateUnauthorized {

	return &CreateCurrentAPISessionCertificateUnauthorized{}
}

// WithPayload adds the payload to the create current Api session certificate unauthorized response
func (o *CreateCurrentAPISessionCertificateUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *CreateCurrentAPISessionCertificateUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create current Api session certificate unauthorized response
func (o *CreateCurrentAPISessionCertificateUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateCurrentAPISessionCertificateUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
