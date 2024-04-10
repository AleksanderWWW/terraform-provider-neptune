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

// ProjectStorageLimitTraitDTO project storage limit trait d t o
//
// swagger:model ProjectStorageLimitTraitDTO
type ProjectStorageLimitTraitDTO struct {

	// max storage size
	// Required: true
	MaxStorageSize *int64 `json:"maxStorageSize"`
}

// Validate validates this project storage limit trait d t o
func (m *ProjectStorageLimitTraitDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMaxStorageSize(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProjectStorageLimitTraitDTO) validateMaxStorageSize(formats strfmt.Registry) error {

	if err := validate.Required("maxStorageSize", "body", m.MaxStorageSize); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this project storage limit trait d t o based on context it is used
func (m *ProjectStorageLimitTraitDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ProjectStorageLimitTraitDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProjectStorageLimitTraitDTO) UnmarshalBinary(b []byte) error {
	var res ProjectStorageLimitTraitDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}