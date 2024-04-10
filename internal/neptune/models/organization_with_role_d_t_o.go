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

// OrganizationWithRoleDTO organization with role d t o
//
// swagger:model OrganizationWithRoleDTO
type OrganizationWithRoleDTO struct {

	// avatar source
	// Required: true
	AvatarSource *AvatarSourceEnum `json:"avatarSource"`

	// avatar Url
	// Required: true
	AvatarURL *string `json:"avatarUrl"`

	// created
	// Required: true
	// Format: date-time
	Created *strfmt.DateTime `json:"created"`

	// id
	// Required: true
	// Format: uuid
	ID *strfmt.UUID `json:"id"`

	// info
	Info string `json:"info,omitempty"`

	// limits enforced
	// Required: true
	LimitsEnforced *bool `json:"limitsEnforced"`

	// name
	// Required: true
	Name *string `json:"name"`

	// organization
	// Required: true
	Organization *OrganizationDTO `json:"organization"`

	// organization type
	// Required: true
	OrganizationType *OrganizationTypeDTO `json:"organizationType"`

	// payment status
	// Required: true
	PaymentStatus *string `json:"paymentStatus"`

	// pricing plan
	// Required: true
	PricingPlan *PricingPlanDTO `json:"pricingPlan"`

	// user role
	// Required: true
	UserRole *OrganizationRoleDTO `json:"userRole"`
}

// Validate validates this organization with role d t o
func (m *OrganizationWithRoleDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAvatarSource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAvatarURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLimitsEnforced(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganization(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganizationType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePaymentStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePricingPlan(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserRole(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OrganizationWithRoleDTO) validateAvatarSource(formats strfmt.Registry) error {

	if err := validate.Required("avatarSource", "body", m.AvatarSource); err != nil {
		return err
	}

	if err := validate.Required("avatarSource", "body", m.AvatarSource); err != nil {
		return err
	}

	if m.AvatarSource != nil {
		if err := m.AvatarSource.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("avatarSource")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("avatarSource")
			}
			return err
		}
	}

	return nil
}

func (m *OrganizationWithRoleDTO) validateAvatarURL(formats strfmt.Registry) error {

	if err := validate.Required("avatarUrl", "body", m.AvatarURL); err != nil {
		return err
	}

	return nil
}

func (m *OrganizationWithRoleDTO) validateCreated(formats strfmt.Registry) error {

	if err := validate.Required("created", "body", m.Created); err != nil {
		return err
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *OrganizationWithRoleDTO) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *OrganizationWithRoleDTO) validateLimitsEnforced(formats strfmt.Registry) error {

	if err := validate.Required("limitsEnforced", "body", m.LimitsEnforced); err != nil {
		return err
	}

	return nil
}

func (m *OrganizationWithRoleDTO) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *OrganizationWithRoleDTO) validateOrganization(formats strfmt.Registry) error {

	if err := validate.Required("organization", "body", m.Organization); err != nil {
		return err
	}

	if m.Organization != nil {
		if err := m.Organization.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("organization")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("organization")
			}
			return err
		}
	}

	return nil
}

func (m *OrganizationWithRoleDTO) validateOrganizationType(formats strfmt.Registry) error {

	if err := validate.Required("organizationType", "body", m.OrganizationType); err != nil {
		return err
	}

	if err := validate.Required("organizationType", "body", m.OrganizationType); err != nil {
		return err
	}

	if m.OrganizationType != nil {
		if err := m.OrganizationType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("organizationType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("organizationType")
			}
			return err
		}
	}

	return nil
}

func (m *OrganizationWithRoleDTO) validatePaymentStatus(formats strfmt.Registry) error {

	if err := validate.Required("paymentStatus", "body", m.PaymentStatus); err != nil {
		return err
	}

	return nil
}

func (m *OrganizationWithRoleDTO) validatePricingPlan(formats strfmt.Registry) error {

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

func (m *OrganizationWithRoleDTO) validateUserRole(formats strfmt.Registry) error {

	if err := validate.Required("userRole", "body", m.UserRole); err != nil {
		return err
	}

	if err := validate.Required("userRole", "body", m.UserRole); err != nil {
		return err
	}

	if m.UserRole != nil {
		if err := m.UserRole.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("userRole")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("userRole")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this organization with role d t o based on the context it is used
func (m *OrganizationWithRoleDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAvatarSource(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOrganization(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOrganizationType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePricingPlan(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUserRole(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OrganizationWithRoleDTO) contextValidateAvatarSource(ctx context.Context, formats strfmt.Registry) error {

	if m.AvatarSource != nil {

		if err := m.AvatarSource.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("avatarSource")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("avatarSource")
			}
			return err
		}
	}

	return nil
}

func (m *OrganizationWithRoleDTO) contextValidateOrganization(ctx context.Context, formats strfmt.Registry) error {

	if m.Organization != nil {

		if err := m.Organization.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("organization")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("organization")
			}
			return err
		}
	}

	return nil
}

func (m *OrganizationWithRoleDTO) contextValidateOrganizationType(ctx context.Context, formats strfmt.Registry) error {

	if m.OrganizationType != nil {

		if err := m.OrganizationType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("organizationType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("organizationType")
			}
			return err
		}
	}

	return nil
}

func (m *OrganizationWithRoleDTO) contextValidatePricingPlan(ctx context.Context, formats strfmt.Registry) error {

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

func (m *OrganizationWithRoleDTO) contextValidateUserRole(ctx context.Context, formats strfmt.Registry) error {

	if m.UserRole != nil {

		if err := m.UserRole.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("userRole")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("userRole")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OrganizationWithRoleDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OrganizationWithRoleDTO) UnmarshalBinary(b []byte) error {
	var res OrganizationWithRoleDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
