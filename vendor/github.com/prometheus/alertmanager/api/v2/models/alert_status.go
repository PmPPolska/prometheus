// Code generated by go-swagger; DO NOT EDIT.

// Copyright Prometheus Team
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AlertStatus alert status
//
// swagger:model alertStatus
type AlertStatus struct {

	// inhibited by
	// Required: true
	InhibitedBy []string `json:"inhibitedBy"`

	// silenced by
	// Required: true
	SilencedBy []string `json:"silencedBy"`

	// state
	// Required: true
	// Enum: [unprocessed active suppressed]
	State *string `json:"state"`
}

// Validate validates this alert status
func (m *AlertStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInhibitedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSilencedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AlertStatus) validateInhibitedBy(formats strfmt.Registry) error {

	if err := validate.Required("inhibitedBy", "body", m.InhibitedBy); err != nil {
		return err
	}

	return nil
}

func (m *AlertStatus) validateSilencedBy(formats strfmt.Registry) error {

	if err := validate.Required("silencedBy", "body", m.SilencedBy); err != nil {
		return err
	}

	return nil
}

var alertStatusTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["unprocessed","active","suppressed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		alertStatusTypeStatePropEnum = append(alertStatusTypeStatePropEnum, v)
	}
}

const (

	// AlertStatusStateUnprocessed captures enum value "unprocessed"
	AlertStatusStateUnprocessed string = "unprocessed"

	// AlertStatusStateActive captures enum value "active"
	AlertStatusStateActive string = "active"

	// AlertStatusStateSuppressed captures enum value "suppressed"
	AlertStatusStateSuppressed string = "suppressed"
)

// prop value enum
func (m *AlertStatus) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, alertStatusTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AlertStatus) validateState(formats strfmt.Registry) error {

	if err := validate.Required("state", "body", m.State); err != nil {
		return err
	}

	// value enum
	if err := m.validateStateEnum("state", "body", *m.State); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AlertStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AlertStatus) UnmarshalBinary(b []byte) error {
	var res AlertStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}