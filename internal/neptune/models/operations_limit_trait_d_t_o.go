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

// OperationsLimitTraitDTO operations limit trait d t o
//
// swagger:model OperationsLimitTraitDTO
type OperationsLimitTraitDTO struct {

	// limits
	// Required: true
	Limits []*OperationsLimitDTO `json:"limits"`
}

// Validate validates this operations limit trait d t o
func (m *OperationsLimitTraitDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLimits(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OperationsLimitTraitDTO) validateLimits(formats strfmt.Registry) error {

	if err := validate.Required("limits", "body", m.Limits); err != nil {
		return err
	}

	for i := 0; i < len(m.Limits); i++ {
		if swag.IsZero(m.Limits[i]) { // not required
			continue
		}

		if m.Limits[i] != nil {
			if err := m.Limits[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("limits" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("limits" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this operations limit trait d t o based on the context it is used
func (m *OperationsLimitTraitDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLimits(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OperationsLimitTraitDTO) contextValidateLimits(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Limits); i++ {

		if m.Limits[i] != nil {

			if swag.IsZero(m.Limits[i]) { // not required
				return nil
			}

			if err := m.Limits[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("limits" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("limits" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *OperationsLimitTraitDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OperationsLimitTraitDTO) UnmarshalBinary(b []byte) error {
	var res OperationsLimitTraitDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
