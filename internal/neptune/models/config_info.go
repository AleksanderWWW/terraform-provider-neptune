// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ConfigInfo config info
//
// swagger:model ConfigInfo
type ConfigInfo struct {

	// max form content size
	// Required: true
	MaxFormContentSize *int32 `json:"maxFormContentSize"`
}

// Validate validates this config info
func (m *ConfigInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMaxFormContentSize(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigInfo) validateMaxFormContentSize(formats strfmt.Registry) error {

	if err := validate.Required("maxFormContentSize", "body", m.MaxFormContentSize); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this config info based on context it is used
func (m *ConfigInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ConfigInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigInfo) UnmarshalBinary(b []byte) error {
	var res ConfigInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}