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

// OperationsLimitDTO operations limit d t o
//
// swagger:model OperationsLimitDTO
type OperationsLimitDTO struct {

	// duration seconds
	// Required: true
	DurationSeconds *int64 `json:"durationSeconds"`

	// name
	// Required: true
	Name *OperationLimitNameDTO `json:"name"`

	// value
	Value int64 `json:"value,omitempty"`
}

// Validate validates this operations limit d t o
func (m *OperationsLimitDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDurationSeconds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OperationsLimitDTO) validateDurationSeconds(formats strfmt.Registry) error {

	if err := validate.Required("durationSeconds", "body", m.DurationSeconds); err != nil {
		return err
	}

	return nil
}

func (m *OperationsLimitDTO) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if m.Name != nil {
		if err := m.Name.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("name")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("name")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this operations limit d t o based on the context it is used
func (m *OperationsLimitDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OperationsLimitDTO) contextValidateName(ctx context.Context, formats strfmt.Registry) error {

	if m.Name != nil {

		if err := m.Name.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("name")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("name")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OperationsLimitDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OperationsLimitDTO) UnmarshalBinary(b []byte) error {
	var res OperationsLimitDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
