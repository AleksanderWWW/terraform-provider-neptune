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

// NewProjectServiceAccountDTO new project service account d t o
//
// swagger:model NewProjectServiceAccountDTO
type NewProjectServiceAccountDTO struct {

	// role
	// Required: true
	Role *ProjectRoleDTO `json:"role"`

	// service account Id
	// Required: true
	ServiceAccountID *string `json:"serviceAccountId"`
}

// Validate validates this new project service account d t o
func (m *NewProjectServiceAccountDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRole(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceAccountID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewProjectServiceAccountDTO) validateRole(formats strfmt.Registry) error {

	if err := validate.Required("role", "body", m.Role); err != nil {
		return err
	}

	if err := validate.Required("role", "body", m.Role); err != nil {
		return err
	}

	if m.Role != nil {
		if err := m.Role.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("role")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("role")
			}
			return err
		}
	}

	return nil
}

func (m *NewProjectServiceAccountDTO) validateServiceAccountID(formats strfmt.Registry) error {

	if err := validate.Required("serviceAccountId", "body", m.ServiceAccountID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this new project service account d t o based on the context it is used
func (m *NewProjectServiceAccountDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRole(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewProjectServiceAccountDTO) contextValidateRole(ctx context.Context, formats strfmt.Registry) error {

	if m.Role != nil {

		if err := m.Role.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("role")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("role")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NewProjectServiceAccountDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewProjectServiceAccountDTO) UnmarshalBinary(b []byte) error {
	var res NewProjectServiceAccountDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}