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

// ProjectKeyDTO project key d t o
//
// swagger:model ProjectKeyDTO
type ProjectKeyDTO struct {

	// proposal
	// Required: true
	Proposal *string `json:"proposal"`
}

// Validate validates this project key d t o
func (m *ProjectKeyDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProposal(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProjectKeyDTO) validateProposal(formats strfmt.Registry) error {

	if err := validate.Required("proposal", "body", m.Proposal); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this project key d t o based on context it is used
func (m *ProjectKeyDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ProjectKeyDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProjectKeyDTO) UnmarshalBinary(b []byte) error {
	var res ProjectKeyDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}