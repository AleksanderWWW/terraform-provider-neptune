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

// V1TeamPricingDTO v1 team pricing d t o
//
// swagger:model V1TeamPricingDTO
type V1TeamPricingDTO struct {

	// number of users
	// Required: true
	NumberOfUsers *int64 `json:"numberOfUsers"`

	// user price in cents
	// Required: true
	UserPriceInCents *int64 `json:"userPriceInCents"`
}

// Validate validates this v1 team pricing d t o
func (m *V1TeamPricingDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNumberOfUsers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserPriceInCents(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1TeamPricingDTO) validateNumberOfUsers(formats strfmt.Registry) error {

	if err := validate.Required("numberOfUsers", "body", m.NumberOfUsers); err != nil {
		return err
	}

	return nil
}

func (m *V1TeamPricingDTO) validateUserPriceInCents(formats strfmt.Registry) error {

	if err := validate.Required("userPriceInCents", "body", m.UserPriceInCents); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 team pricing d t o based on context it is used
func (m *V1TeamPricingDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1TeamPricingDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1TeamPricingDTO) UnmarshalBinary(b []byte) error {
	var res V1TeamPricingDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
