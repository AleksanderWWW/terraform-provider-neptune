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

// ProjectVisibilityRestrictedTraitDTO project visibility restricted trait d t o
//
// swagger:model ProjectVisibilityRestrictedTraitDTO
type ProjectVisibilityRestrictedTraitDTO struct {

	// disabled by admin
	// Required: true
	DisabledByAdmin []ProjectVisibilityDTO `json:"disabledByAdmin"`

	// limited to
	// Required: true
	LimitedTo []ProjectVisibilityDTO `json:"limitedTo"`

	// limited to by plan
	// Required: true
	LimitedToByPlan []ProjectVisibilityDTO `json:"limitedToByPlan"`
}

// Validate validates this project visibility restricted trait d t o
func (m *ProjectVisibilityRestrictedTraitDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDisabledByAdmin(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLimitedTo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLimitedToByPlan(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProjectVisibilityRestrictedTraitDTO) validateDisabledByAdmin(formats strfmt.Registry) error {

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

func (m *ProjectVisibilityRestrictedTraitDTO) validateLimitedTo(formats strfmt.Registry) error {

	if err := validate.Required("limitedTo", "body", m.LimitedTo); err != nil {
		return err
	}

	for i := 0; i < len(m.LimitedTo); i++ {

		if err := m.LimitedTo[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("limitedTo" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("limitedTo" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *ProjectVisibilityRestrictedTraitDTO) validateLimitedToByPlan(formats strfmt.Registry) error {

	if err := validate.Required("limitedToByPlan", "body", m.LimitedToByPlan); err != nil {
		return err
	}

	for i := 0; i < len(m.LimitedToByPlan); i++ {

		if err := m.LimitedToByPlan[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("limitedToByPlan" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("limitedToByPlan" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// ContextValidate validate this project visibility restricted trait d t o based on the context it is used
func (m *ProjectVisibilityRestrictedTraitDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDisabledByAdmin(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLimitedTo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLimitedToByPlan(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProjectVisibilityRestrictedTraitDTO) contextValidateDisabledByAdmin(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ProjectVisibilityRestrictedTraitDTO) contextValidateLimitedTo(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.LimitedTo); i++ {

		if swag.IsZero(m.LimitedTo[i]) { // not required
			return nil
		}

		if err := m.LimitedTo[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("limitedTo" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("limitedTo" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *ProjectVisibilityRestrictedTraitDTO) contextValidateLimitedToByPlan(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.LimitedToByPlan); i++ {

		if swag.IsZero(m.LimitedToByPlan[i]) { // not required
			return nil
		}

		if err := m.LimitedToByPlan[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("limitedToByPlan" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("limitedToByPlan" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProjectVisibilityRestrictedTraitDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProjectVisibilityRestrictedTraitDTO) UnmarshalBinary(b []byte) error {
	var res ProjectVisibilityRestrictedTraitDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}