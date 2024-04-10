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

// PricingPlanDefinitionDTO pricing plan definition d t o
//
// swagger:model PricingPlanDefinitionDTO
type PricingPlanDefinitionDTO struct {

	// display name
	// Required: true
	DisplayName *string `json:"displayName"`

	// plan
	// Required: true
	Plan *PricingPlanDTO `json:"plan"`
}

// Validate validates this pricing plan definition d t o
func (m *PricingPlanDefinitionDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDisplayName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlan(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PricingPlanDefinitionDTO) validateDisplayName(formats strfmt.Registry) error {

	if err := validate.Required("displayName", "body", m.DisplayName); err != nil {
		return err
	}

	return nil
}

func (m *PricingPlanDefinitionDTO) validatePlan(formats strfmt.Registry) error {

	if err := validate.Required("plan", "body", m.Plan); err != nil {
		return err
	}

	if err := validate.Required("plan", "body", m.Plan); err != nil {
		return err
	}

	if m.Plan != nil {
		if err := m.Plan.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("plan")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("plan")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this pricing plan definition d t o based on the context it is used
func (m *PricingPlanDefinitionDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePlan(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PricingPlanDefinitionDTO) contextValidatePlan(ctx context.Context, formats strfmt.Registry) error {

	if m.Plan != nil {

		if err := m.Plan.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("plan")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("plan")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PricingPlanDefinitionDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PricingPlanDefinitionDTO) UnmarshalBinary(b []byte) error {
	var res PricingPlanDefinitionDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
