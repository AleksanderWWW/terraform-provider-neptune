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

// UpgradePurchaseRequestDTO upgrade purchase request d t o
//
// swagger:model UpgradePurchaseRequestDTO
type UpgradePurchaseRequestDTO struct {

	// pricing plan
	// Required: true
	PricingPlan *PricingPlanDTO `json:"pricingPlan"`

	// session params
	// Required: true
	SessionParams *CreateSessionParamsDTO `json:"sessionParams"`
}

// Validate validates this upgrade purchase request d t o
func (m *UpgradePurchaseRequestDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePricingPlan(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSessionParams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpgradePurchaseRequestDTO) validatePricingPlan(formats strfmt.Registry) error {

	if err := validate.Required("pricingPlan", "body", m.PricingPlan); err != nil {
		return err
	}

	if err := validate.Required("pricingPlan", "body", m.PricingPlan); err != nil {
		return err
	}

	if m.PricingPlan != nil {
		if err := m.PricingPlan.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pricingPlan")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pricingPlan")
			}
			return err
		}
	}

	return nil
}

func (m *UpgradePurchaseRequestDTO) validateSessionParams(formats strfmt.Registry) error {

	if err := validate.Required("sessionParams", "body", m.SessionParams); err != nil {
		return err
	}

	if m.SessionParams != nil {
		if err := m.SessionParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sessionParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sessionParams")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this upgrade purchase request d t o based on the context it is used
func (m *UpgradePurchaseRequestDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePricingPlan(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSessionParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpgradePurchaseRequestDTO) contextValidatePricingPlan(ctx context.Context, formats strfmt.Registry) error {

	if m.PricingPlan != nil {

		if err := m.PricingPlan.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pricingPlan")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pricingPlan")
			}
			return err
		}
	}

	return nil
}

func (m *UpgradePurchaseRequestDTO) contextValidateSessionParams(ctx context.Context, formats strfmt.Registry) error {

	if m.SessionParams != nil {

		if err := m.SessionParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sessionParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sessionParams")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpgradePurchaseRequestDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpgradePurchaseRequestDTO) UnmarshalBinary(b []byte) error {
	var res UpgradePurchaseRequestDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
