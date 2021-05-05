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

package current_identity

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

// NewDetailMfaQrCodeParams creates a new DetailMfaQrCodeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDetailMfaQrCodeParams() *DetailMfaQrCodeParams {
	return &DetailMfaQrCodeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDetailMfaQrCodeParamsWithTimeout creates a new DetailMfaQrCodeParams object
// with the ability to set a timeout on a request.
func NewDetailMfaQrCodeParamsWithTimeout(timeout time.Duration) *DetailMfaQrCodeParams {
	return &DetailMfaQrCodeParams{
		timeout: timeout,
	}
}

// NewDetailMfaQrCodeParamsWithContext creates a new DetailMfaQrCodeParams object
// with the ability to set a context for a request.
func NewDetailMfaQrCodeParamsWithContext(ctx context.Context) *DetailMfaQrCodeParams {
	return &DetailMfaQrCodeParams{
		Context: ctx,
	}
}

// NewDetailMfaQrCodeParamsWithHTTPClient creates a new DetailMfaQrCodeParams object
// with the ability to set a custom HTTPClient for a request.
func NewDetailMfaQrCodeParamsWithHTTPClient(client *http.Client) *DetailMfaQrCodeParams {
	return &DetailMfaQrCodeParams{
		HTTPClient: client,
	}
}

/* DetailMfaQrCodeParams contains all the parameters to send to the API endpoint
   for the detail mfa qr code operation.

   Typically these are written to a http.Request.
*/
type DetailMfaQrCodeParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the detail mfa qr code params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DetailMfaQrCodeParams) WithDefaults() *DetailMfaQrCodeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the detail mfa qr code params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DetailMfaQrCodeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the detail mfa qr code params
func (o *DetailMfaQrCodeParams) WithTimeout(timeout time.Duration) *DetailMfaQrCodeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the detail mfa qr code params
func (o *DetailMfaQrCodeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the detail mfa qr code params
func (o *DetailMfaQrCodeParams) WithContext(ctx context.Context) *DetailMfaQrCodeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the detail mfa qr code params
func (o *DetailMfaQrCodeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the detail mfa qr code params
func (o *DetailMfaQrCodeParams) WithHTTPClient(client *http.Client) *DetailMfaQrCodeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the detail mfa qr code params
func (o *DetailMfaQrCodeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *DetailMfaQrCodeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
