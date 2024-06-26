// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// LoginActionsListDTO login actions list d t o
//
// swagger:model LoginActionsListDTO
type LoginActionsListDTO struct {

	// required actions
	// Required: true
	RequiredActions []LoginActionDTO `json:"requiredActions"`
}

// Validate validates this login actions list d t o
func (m *LoginActionsListDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRequiredActions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LoginActionsListDTO) validateRequiredActions(formats strfmt.Registry) error {

	if err := validate.Required("requiredActions", "body", m.RequiredActions); err != nil {
		return err
	}

	for i := 0; i < len(m.RequiredActions); i++ {

		if err := m.RequiredActions[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("requiredActions" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("requiredActions" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// ContextValidate validate this login actions list d t o based on the context it is used
func (m *LoginActionsListDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRequiredActions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LoginActionsListDTO) contextValidateRequiredActions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.RequiredActions); i++ {

		if swag.IsZero(m.RequiredActions[i]) { // not required
			return nil
		}

		if err := m.RequiredActions[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("requiredActions" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("requiredActions" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *LoginActionsListDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LoginActionsListDTO) UnmarshalBinary(b []byte) error {
	var res LoginActionsListDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
