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

// GraduatedPricingDTO graduated pricing d t o
//
// swagger:model GraduatedPricingDTO
type GraduatedPricingDTO struct {

	// current value
	// Required: true
	CurrentValue *int64 `json:"currentValue"`

	// price tiers
	// Required: true
	PriceTiers []*PriceTierDTO `json:"priceTiers"`

	// step size
	// Required: true
	StepSize *int64 `json:"stepSize"`

	// unit symbol
	// Required: true
	UnitSymbol *string `json:"unitSymbol"`
}

// Validate validates this graduated pricing d t o
func (m *GraduatedPricingDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCurrentValue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePriceTiers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStepSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUnitSymbol(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GraduatedPricingDTO) validateCurrentValue(formats strfmt.Registry) error {

	if err := validate.Required("currentValue", "body", m.CurrentValue); err != nil {
		return err
	}

	return nil
}

func (m *GraduatedPricingDTO) validatePriceTiers(formats strfmt.Registry) error {

	if err := validate.Required("priceTiers", "body", m.PriceTiers); err != nil {
		return err
	}

	for i := 0; i < len(m.PriceTiers); i++ {
		if swag.IsZero(m.PriceTiers[i]) { // not required
			continue
		}

		if m.PriceTiers[i] != nil {
			if err := m.PriceTiers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("priceTiers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("priceTiers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *GraduatedPricingDTO) validateStepSize(formats strfmt.Registry) error {

	if err := validate.Required("stepSize", "body", m.StepSize); err != nil {
		return err
	}

	return nil
}

func (m *GraduatedPricingDTO) validateUnitSymbol(formats strfmt.Registry) error {

	if err := validate.Required("unitSymbol", "body", m.UnitSymbol); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this graduated pricing d t o based on the context it is used
func (m *GraduatedPricingDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePriceTiers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GraduatedPricingDTO) contextValidatePriceTiers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PriceTiers); i++ {

		if m.PriceTiers[i] != nil {

			if swag.IsZero(m.PriceTiers[i]) { // not required
				return nil
			}

			if err := m.PriceTiers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("priceTiers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("priceTiers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *GraduatedPricingDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GraduatedPricingDTO) UnmarshalBinary(b []byte) error {
	var res GraduatedPricingDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
