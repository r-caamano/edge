// Code generated by go-swagger; DO NOT EDIT.

//
// Copyright NetFoundry, Inc.
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
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewDeleteCurrentAPISessionCertificateParams creates a new DeleteCurrentAPISessionCertificateParams object
// with the default values initialized.
func NewDeleteCurrentAPISessionCertificateParams() *DeleteCurrentAPISessionCertificateParams {
	var ()
	return &DeleteCurrentAPISessionCertificateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteCurrentAPISessionCertificateParamsWithTimeout creates a new DeleteCurrentAPISessionCertificateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteCurrentAPISessionCertificateParamsWithTimeout(timeout time.Duration) *DeleteCurrentAPISessionCertificateParams {
	var ()
	return &DeleteCurrentAPISessionCertificateParams{

		timeout: timeout,
	}
}

// NewDeleteCurrentAPISessionCertificateParamsWithContext creates a new DeleteCurrentAPISessionCertificateParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteCurrentAPISessionCertificateParamsWithContext(ctx context.Context) *DeleteCurrentAPISessionCertificateParams {
	var ()
	return &DeleteCurrentAPISessionCertificateParams{

		Context: ctx,
	}
}

// NewDeleteCurrentAPISessionCertificateParamsWithHTTPClient creates a new DeleteCurrentAPISessionCertificateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteCurrentAPISessionCertificateParamsWithHTTPClient(client *http.Client) *DeleteCurrentAPISessionCertificateParams {
	var ()
	return &DeleteCurrentAPISessionCertificateParams{
		HTTPClient: client,
	}
}

/*DeleteCurrentAPISessionCertificateParams contains all the parameters to send to the API endpoint
for the delete current Api session certificate operation typically these are written to a http.Request
*/
type DeleteCurrentAPISessionCertificateParams struct {

	/*ID
	  The id of the requested resource

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete current Api session certificate params
func (o *DeleteCurrentAPISessionCertificateParams) WithTimeout(timeout time.Duration) *DeleteCurrentAPISessionCertificateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete current Api session certificate params
func (o *DeleteCurrentAPISessionCertificateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete current Api session certificate params
func (o *DeleteCurrentAPISessionCertificateParams) WithContext(ctx context.Context) *DeleteCurrentAPISessionCertificateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete current Api session certificate params
func (o *DeleteCurrentAPISessionCertificateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete current Api session certificate params
func (o *DeleteCurrentAPISessionCertificateParams) WithHTTPClient(client *http.Client) *DeleteCurrentAPISessionCertificateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete current Api session certificate params
func (o *DeleteCurrentAPISessionCertificateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete current Api session certificate params
func (o *DeleteCurrentAPISessionCertificateParams) WithID(id string) *DeleteCurrentAPISessionCertificateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete current Api session certificate params
func (o *DeleteCurrentAPISessionCertificateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteCurrentAPISessionCertificateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}