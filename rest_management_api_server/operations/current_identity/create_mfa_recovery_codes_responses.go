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

package current_identity

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openziti/edge/rest_model"
)

// CreateMfaRecoveryCodesOKCode is the HTTP code returned for type CreateMfaRecoveryCodesOK
const CreateMfaRecoveryCodesOKCode int = 200

/*CreateMfaRecoveryCodesOK The recovery codes of an MFA enrollment

swagger:response createMfaRecoveryCodesOK
*/
type CreateMfaRecoveryCodesOK struct {

	/*
	  In: Body
	*/
	Payload *rest_model.DetailMfaRecoveryCodesEnvelope `json:"body,omitempty"`
}

// NewCreateMfaRecoveryCodesOK creates CreateMfaRecoveryCodesOK with default headers values
func NewCreateMfaRecoveryCodesOK() *CreateMfaRecoveryCodesOK {

	return &CreateMfaRecoveryCodesOK{}
}

// WithPayload adds the payload to the create mfa recovery codes o k response
func (o *CreateMfaRecoveryCodesOK) WithPayload(payload *rest_model.DetailMfaRecoveryCodesEnvelope) *CreateMfaRecoveryCodesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create mfa recovery codes o k response
func (o *CreateMfaRecoveryCodesOK) SetPayload(payload *rest_model.DetailMfaRecoveryCodesEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateMfaRecoveryCodesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateMfaRecoveryCodesUnauthorizedCode is the HTTP code returned for type CreateMfaRecoveryCodesUnauthorized
const CreateMfaRecoveryCodesUnauthorizedCode int = 401

/*CreateMfaRecoveryCodesUnauthorized The currently supplied session does not have the correct access rights to request this resource

swagger:response createMfaRecoveryCodesUnauthorized
*/
type CreateMfaRecoveryCodesUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewCreateMfaRecoveryCodesUnauthorized creates CreateMfaRecoveryCodesUnauthorized with default headers values
func NewCreateMfaRecoveryCodesUnauthorized() *CreateMfaRecoveryCodesUnauthorized {

	return &CreateMfaRecoveryCodesUnauthorized{}
}

// WithPayload adds the payload to the create mfa recovery codes unauthorized response
func (o *CreateMfaRecoveryCodesUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *CreateMfaRecoveryCodesUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create mfa recovery codes unauthorized response
func (o *CreateMfaRecoveryCodesUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateMfaRecoveryCodesUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateMfaRecoveryCodesNotFoundCode is the HTTP code returned for type CreateMfaRecoveryCodesNotFound
const CreateMfaRecoveryCodesNotFoundCode int = 404

/*CreateMfaRecoveryCodesNotFound The requested resource does not exist

swagger:response createMfaRecoveryCodesNotFound
*/
type CreateMfaRecoveryCodesNotFound struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewCreateMfaRecoveryCodesNotFound creates CreateMfaRecoveryCodesNotFound with default headers values
func NewCreateMfaRecoveryCodesNotFound() *CreateMfaRecoveryCodesNotFound {

	return &CreateMfaRecoveryCodesNotFound{}
}

// WithPayload adds the payload to the create mfa recovery codes not found response
func (o *CreateMfaRecoveryCodesNotFound) WithPayload(payload *rest_model.APIErrorEnvelope) *CreateMfaRecoveryCodesNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create mfa recovery codes not found response
func (o *CreateMfaRecoveryCodesNotFound) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateMfaRecoveryCodesNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
