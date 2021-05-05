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
	"bytes"
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PostureCheckProcessMultiUpdate posture check process multi update
//
// swagger:model postureCheckProcessMultiUpdate
type PostureCheckProcessMultiUpdate struct {
	nameField *string

	roleAttributesField Attributes

	tagsField Tags

	// processes
	// Required: true
	// Min Items: 1
	Processes []*ProcessMulti `json:"processes"`

	// semantic
	// Required: true
	Semantic *Semantic `json:"semantic"`
}

// Name gets the name of this subtype
func (m *PostureCheckProcessMultiUpdate) Name() *string {
	return m.nameField
}

// SetName sets the name of this subtype
func (m *PostureCheckProcessMultiUpdate) SetName(val *string) {
	m.nameField = val
}

// RoleAttributes gets the role attributes of this subtype
func (m *PostureCheckProcessMultiUpdate) RoleAttributes() Attributes {
	return m.roleAttributesField
}

// SetRoleAttributes sets the role attributes of this subtype
func (m *PostureCheckProcessMultiUpdate) SetRoleAttributes(val Attributes) {
	m.roleAttributesField = val
}

// Tags gets the tags of this subtype
func (m *PostureCheckProcessMultiUpdate) Tags() Tags {
	return m.tagsField
}

// SetTags sets the tags of this subtype
func (m *PostureCheckProcessMultiUpdate) SetTags(val Tags) {
	m.tagsField = val
}

// TypeID gets the type Id of this subtype
func (m *PostureCheckProcessMultiUpdate) TypeID() PostureCheckType {
	return "PROCESS_MULTI"
}

// SetTypeID sets the type Id of this subtype
func (m *PostureCheckProcessMultiUpdate) SetTypeID(val PostureCheckType) {
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *PostureCheckProcessMultiUpdate) UnmarshalJSON(raw []byte) error {
	var data struct {

		// processes
		// Required: true
		// Min Items: 1
		Processes []*ProcessMulti `json:"processes"`

		// semantic
		// Required: true
		Semantic *Semantic `json:"semantic"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var base struct {
		/* Just the base type fields. Used for unmashalling polymorphic types.*/

		Name *string `json:"name"`

		RoleAttributes Attributes `json:"roleAttributes"`

		Tags Tags `json:"tags"`

		TypeID PostureCheckType `json:"typeId,omitempty"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result PostureCheckProcessMultiUpdate

	result.nameField = base.Name

	result.roleAttributesField = base.RoleAttributes

	result.tagsField = base.Tags

	if base.TypeID != result.TypeID() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid typeId value: %q", base.TypeID)
	}

	result.Processes = data.Processes
	result.Semantic = data.Semantic

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m PostureCheckProcessMultiUpdate) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// processes
		// Required: true
		// Min Items: 1
		Processes []*ProcessMulti `json:"processes"`

		// semantic
		// Required: true
		Semantic *Semantic `json:"semantic"`
	}{

		Processes: m.Processes,

		Semantic: m.Semantic,
	})
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		Name *string `json:"name"`

		RoleAttributes Attributes `json:"roleAttributes"`

		Tags Tags `json:"tags"`

		TypeID PostureCheckType `json:"typeId,omitempty"`
	}{

		Name: m.Name(),

		RoleAttributes: m.RoleAttributes(),

		Tags: m.Tags(),

		TypeID: m.TypeID(),
	})
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this posture check process multi update
func (m *PostureCheckProcessMultiUpdate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoleAttributes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProcesses(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSemantic(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostureCheckProcessMultiUpdate) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name()); err != nil {
		return err
	}

	return nil
}

func (m *PostureCheckProcessMultiUpdate) validateRoleAttributes(formats strfmt.Registry) error {

	if swag.IsZero(m.RoleAttributes()) { // not required
		return nil
	}

	if err := m.RoleAttributes().Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("roleAttributes")
		}
		return err
	}

	return nil
}

func (m *PostureCheckProcessMultiUpdate) validateTags(formats strfmt.Registry) error {

	if swag.IsZero(m.Tags()) { // not required
		return nil
	}

	if m.Tags() != nil {
		if err := m.Tags().Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tags")
			}
			return err
		}
	}

	return nil
}

func (m *PostureCheckProcessMultiUpdate) validateProcesses(formats strfmt.Registry) error {

	if err := validate.Required("processes", "body", m.Processes); err != nil {
		return err
	}

	iProcessesSize := int64(len(m.Processes))

	if err := validate.MinItems("processes", "body", iProcessesSize, 1); err != nil {
		return err
	}

	for i := 0; i < len(m.Processes); i++ {
		if swag.IsZero(m.Processes[i]) { // not required
			continue
		}

		if m.Processes[i] != nil {
			if err := m.Processes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("processes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PostureCheckProcessMultiUpdate) validateSemantic(formats strfmt.Registry) error {

	if err := validate.Required("semantic", "body", m.Semantic); err != nil {
		return err
	}

	if err := validate.Required("semantic", "body", m.Semantic); err != nil {
		return err
	}

	if m.Semantic != nil {
		if err := m.Semantic.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("semantic")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this posture check process multi update based on the context it is used
func (m *PostureCheckProcessMultiUpdate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRoleAttributes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTags(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProcesses(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSemantic(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostureCheckProcessMultiUpdate) contextValidateRoleAttributes(ctx context.Context, formats strfmt.Registry) error {

	if err := m.RoleAttributes().ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("roleAttributes")
		}
		return err
	}

	return nil
}

func (m *PostureCheckProcessMultiUpdate) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Tags().ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("tags")
		}
		return err
	}

	return nil
}

func (m *PostureCheckProcessMultiUpdate) contextValidateTypeID(ctx context.Context, formats strfmt.Registry) error {

	if err := m.TypeID().ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("typeId")
		}
		return err
	}

	return nil
}

func (m *PostureCheckProcessMultiUpdate) contextValidateProcesses(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Processes); i++ {

		if m.Processes[i] != nil {
			if err := m.Processes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("processes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PostureCheckProcessMultiUpdate) contextValidateSemantic(ctx context.Context, formats strfmt.Registry) error {

	if m.Semantic != nil {
		if err := m.Semantic.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("semantic")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostureCheckProcessMultiUpdate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostureCheckProcessMultiUpdate) UnmarshalBinary(b []byte) error {
	var res PostureCheckProcessMultiUpdate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
