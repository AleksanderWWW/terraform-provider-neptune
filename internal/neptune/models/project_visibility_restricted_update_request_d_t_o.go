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

// ProjectVisibilityRestrictedUpdateRequestDTO project visibility restricted update request d t o
//
// swagger:model ProjectVisibilityRestrictedUpdateRequestDTO
type ProjectVisibilityRestrictedUpdateRequestDTO struct {

	// default value
	// Required: true
	DefaultValue *ProjectVisibilityDTO `json:"defaultValue"`

	// disabled by admin
	// Required: true
	DisabledByAdmin []ProjectVisibilityDTO `json:"disabledByAdmin"`
}

// Validate validates this project visibility restricted update request d t o
func (m *ProjectVisibilityRestrictedUpdateRequestDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDefaultValue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisabledByAdmin(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProjectVisibilityRestrictedUpdateRequestDTO) validateDefaultValue(formats strfmt.Registry) error {

	if err := validate.Required("defaultValue", "body", m.DefaultValue); err != nil {
		return err
	}

	if err := validate.Required("defaultValue", "body", m.DefaultValue); err != nil {
		return err
	}

	if m.DefaultValue != nil {
		if err := m.DefaultValue.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("defaultValue")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("defaultValue")
			}
			return err
		}
	}

	return nil
}

func (m *ProjectVisibilityRestrictedUpdateRequestDTO) validateDisabledByAdmin(formats strfmt.Registry) error {

	if err := validate.Required("disabledByAdmin", "body", m.DisabledByAdmin); err != nil {
		return err
	}

	for i := 0; i < len(m.DisabledByAdmin); i++ {

		if err := m.DisabledByAdmin[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("disabledByAdmin" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("disabledByAdmin" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// ContextValidate validate this project visibility restricted update request d t o based on the context it is used
func (m *ProjectVisibilityRestrictedUpdateRequestDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDefaultValue(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDisabledByAdmin(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProjectVisibilityRestrictedUpdateRequestDTO) contextValidateDefaultValue(ctx context.Context, formats strfmt.Registry) error {

	if m.DefaultValue != nil {

		if err := m.DefaultValue.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("defaultValue")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("defaultValue")
			}
			return err
		}
	}

	return nil
}

func (m *ProjectVisibilityRestrictedUpdateRequestDTO) contextValidateDisabledByAdmin(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.DisabledByAdmin); i++ {

		if swag.IsZero(m.DisabledByAdmin[i]) { // not required
			return nil
		}

		if err := m.DisabledByAdmin[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("disabledByAdmin" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("disabledByAdmin" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProjectVisibilityRestrictedUpdateRequestDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProjectVisibilityRestrictedUpdateRequestDTO) UnmarshalBinary(b []byte) error {
	var res ProjectVisibilityRestrictedUpdateRequestDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
