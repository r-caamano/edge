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

package external_jwt_signer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/openziti/edge/rest_model"
)

// NewCreateExternalJWTSignerParams creates a new CreateExternalJWTSignerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateExternalJWTSignerParams() *CreateExternalJWTSignerParams {
	return &CreateExternalJWTSignerParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateExternalJWTSignerParamsWithTimeout creates a new CreateExternalJWTSignerParams object
// with the ability to set a timeout on a request.
func NewCreateExternalJWTSignerParamsWithTimeout(timeout time.Duration) *CreateExternalJWTSignerParams {
	return &CreateExternalJWTSignerParams{
		timeout: timeout,
	}
}

// NewCreateExternalJWTSignerParamsWithContext creates a new CreateExternalJWTSignerParams object
// with the ability to set a context for a request.
func NewCreateExternalJWTSignerParamsWithContext(ctx context.Context) *CreateExternalJWTSignerParams {
	return &CreateExternalJWTSignerParams{
		Context: ctx,
	}
}

// NewCreateExternalJWTSignerParamsWithHTTPClient creates a new CreateExternalJWTSignerParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateExternalJWTSignerParamsWithHTTPClient(client *http.Client) *CreateExternalJWTSignerParams {
	return &CreateExternalJWTSignerParams{
		HTTPClient: client,
	}
}

/* CreateExternalJWTSignerParams contains all the parameters to send to the API endpoint
   for the create external Jwt signer operation.

   Typically these are written to a http.Request.
*/
type CreateExternalJWTSignerParams struct {

	/* ExternalJWTSigner.

	   An External JWT Signer to create
	*/
	ExternalJWTSigner *rest_model.ExternalJWTSignerCreate

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create external Jwt signer params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateExternalJWTSignerParams) WithDefaults() *CreateExternalJWTSignerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create external Jwt signer params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateExternalJWTSignerParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create external Jwt signer params
func (o *CreateExternalJWTSignerParams) WithTimeout(timeout time.Duration) *CreateExternalJWTSignerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create external Jwt signer params
func (o *CreateExternalJWTSignerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create external Jwt signer params
func (o *CreateExternalJWTSignerParams) WithContext(ctx context.Context) *CreateExternalJWTSignerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create external Jwt signer params
func (o *CreateExternalJWTSignerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create external Jwt signer params
func (o *CreateExternalJWTSignerParams) WithHTTPClient(client *http.Client) *CreateExternalJWTSignerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create external Jwt signer params
func (o *CreateExternalJWTSignerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExternalJWTSigner adds the externalJWTSigner to the create external Jwt signer params
func (o *CreateExternalJWTSignerParams) WithExternalJWTSigner(externalJWTSigner *rest_model.ExternalJWTSignerCreate) *CreateExternalJWTSignerParams {
	o.SetExternalJWTSigner(externalJWTSigner)
	return o
}

// SetExternalJWTSigner adds the externalJwtSigner to the create external Jwt signer params
func (o *CreateExternalJWTSignerParams) SetExternalJWTSigner(externalJWTSigner *rest_model.ExternalJWTSignerCreate) {
	o.ExternalJWTSigner = externalJWTSigner
}

// WriteToRequest writes these params to a swagger request
func (o *CreateExternalJWTSignerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.ExternalJWTSigner != nil {
		if err := r.SetBodyParam(o.ExternalJWTSigner); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
