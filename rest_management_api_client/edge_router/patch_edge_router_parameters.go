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
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/openziti/edge/rest_model"
)

// NewPatchEdgeRouterParams creates a new PatchEdgeRouterParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchEdgeRouterParams() *PatchEdgeRouterParams {
	return &PatchEdgeRouterParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchEdgeRouterParamsWithTimeout creates a new PatchEdgeRouterParams object
// with the ability to set a timeout on a request.
func NewPatchEdgeRouterParamsWithTimeout(timeout time.Duration) *PatchEdgeRouterParams {
	return &PatchEdgeRouterParams{
		timeout: timeout,
	}
}

// NewPatchEdgeRouterParamsWithContext creates a new PatchEdgeRouterParams object
// with the ability to set a context for a request.
func NewPatchEdgeRouterParamsWithContext(ctx context.Context) *PatchEdgeRouterParams {
	return &PatchEdgeRouterParams{
		Context: ctx,
	}
}

// NewPatchEdgeRouterParamsWithHTTPClient creates a new PatchEdgeRouterParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchEdgeRouterParamsWithHTTPClient(client *http.Client) *PatchEdgeRouterParams {
	return &PatchEdgeRouterParams{
		HTTPClient: client,
	}
}

/* PatchEdgeRouterParams contains all the parameters to send to the API endpoint
   for the patch edge router operation.

   Typically these are written to a http.Request.
*/
type PatchEdgeRouterParams struct {

	/* EdgeRouter.

	   An edge router patch object
	*/
	EdgeRouter *rest_model.EdgeRouterPatch

	/* ID.

	   The id of the requested resource
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch edge router params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchEdgeRouterParams) WithDefaults() *PatchEdgeRouterParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch edge router params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchEdgeRouterParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch edge router params
func (o *PatchEdgeRouterParams) WithTimeout(timeout time.Duration) *PatchEdgeRouterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch edge router params
func (o *PatchEdgeRouterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch edge router params
func (o *PatchEdgeRouterParams) WithContext(ctx context.Context) *PatchEdgeRouterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch edge router params
func (o *PatchEdgeRouterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch edge router params
func (o *PatchEdgeRouterParams) WithHTTPClient(client *http.Client) *PatchEdgeRouterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch edge router params
func (o *PatchEdgeRouterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEdgeRouter adds the edgeRouter to the patch edge router params
func (o *PatchEdgeRouterParams) WithEdgeRouter(edgeRouter *rest_model.EdgeRouterPatch) *PatchEdgeRouterParams {
	o.SetEdgeRouter(edgeRouter)
	return o
}

// SetEdgeRouter adds the edgeRouter to the patch edge router params
func (o *PatchEdgeRouterParams) SetEdgeRouter(edgeRouter *rest_model.EdgeRouterPatch) {
	o.EdgeRouter = edgeRouter
}

// WithID adds the id to the patch edge router params
func (o *PatchEdgeRouterParams) WithID(id string) *PatchEdgeRouterParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the patch edge router params
func (o *PatchEdgeRouterParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PatchEdgeRouterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.EdgeRouter != nil {
		if err := r.SetBodyParam(o.EdgeRouter); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
