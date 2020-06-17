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

package rest_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// IdentityCreate An identity to create
//
// swagger:model identityCreate
type IdentityCreate struct {

	// enrollment
	Enrollment *IdentityCreateEnrollment `json:"enrollment,omitempty"`

	// is admin
	// Required: true
	IsAdmin *bool `json:"isAdmin"`

	// name
	// Required: true
	Name *string `json:"name"`

	// role attributes
	RoleAttributes Attributes `json:"roleAttributes"`

	// tags
	Tags Tags `json:"tags"`

	// type
	// Required: true
	Type IdentityType `json:"type"`
}

// Validate validates this identity create
func (m *IdentityCreate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnrollment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsAdmin(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoleAttributes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IdentityCreate) validateEnrollment(formats strfmt.Registry) error {

	if swag.IsZero(m.Enrollment) { // not required
		return nil
	}

	if m.Enrollment != nil {
		if err := m.Enrollment.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("enrollment")
			}
			return err
		}
	}

	return nil
}

func (m *IdentityCreate) validateIsAdmin(formats strfmt.Registry) error {

	if err := validate.Required("isAdmin", "body", m.IsAdmin); err != nil {
		return err
	}

	return nil
}

func (m *IdentityCreate) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *IdentityCreate) validateRoleAttributes(formats strfmt.Registry) error {

	if swag.IsZero(m.RoleAttributes) { // not required
		return nil
	}

	if err := m.RoleAttributes.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("roleAttributes")
		}
		return err
	}

	return nil
}

func (m *IdentityCreate) validateTags(formats strfmt.Registry) error {

	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	if err := m.Tags.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("tags")
		}
		return err
	}

	return nil
}

func (m *IdentityCreate) validateType(formats strfmt.Registry) error {

	if err := m.Type.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("type")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IdentityCreate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IdentityCreate) UnmarshalBinary(b []byte) error {
	var res IdentityCreate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IdentityCreateEnrollment identity create enrollment
//
// swagger:model IdentityCreateEnrollment
type IdentityCreateEnrollment struct {

	// ott
	Ott bool `json:"ott,omitempty"`

	// ottca
	// Format: uuid
	Ottca strfmt.UUID `json:"ottca,omitempty"`

	// updb
	Updb string `json:"updb,omitempty"`
}

// Validate validates this identity create enrollment
func (m *IdentityCreateEnrollment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOttca(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IdentityCreateEnrollment) validateOttca(formats strfmt.Registry) error {

	if swag.IsZero(m.Ottca) { // not required
		return nil
	}

	if err := validate.FormatOf("enrollment"+"."+"ottca", "body", "uuid", m.Ottca.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IdentityCreateEnrollment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IdentityCreateEnrollment) UnmarshalBinary(b []byte) error {
	var res IdentityCreateEnrollment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}