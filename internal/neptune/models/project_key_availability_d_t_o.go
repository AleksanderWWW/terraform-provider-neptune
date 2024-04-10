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

// ProjectKeyAvailabilityDTO project key availability d t o
//
// swagger:model ProjectKeyAvailabilityDTO
type ProjectKeyAvailabilityDTO struct {

	// is key available
	// Required: true
	IsKeyAvailable *bool `json:"isKeyAvailable"`
}

// Validate validates this project key availability d t o
func (m *ProjectKeyAvailabilityDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIsKeyAvailable(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProjectKeyAvailabilityDTO) validateIsKeyAvailable(formats strfmt.Registry) error {

	if err := validate.Required("isKeyAvailable", "body", m.IsKeyAvailable); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this project key availability d t o based on context it is used
func (m *ProjectKeyAvailabilityDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ProjectKeyAvailabilityDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProjectKeyAvailabilityDTO) UnmarshalBinary(b []byte) error {
	var res ProjectKeyAvailabilityDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
