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

// NewOrganizationInvitationEntryDTO new organization invitation entry d t o
//
// swagger:model NewOrganizationInvitationEntryDTO
type NewOrganizationInvitationEntryDTO struct {

	// add to all projects
	AddToAllProjects bool `json:"addToAllProjects,omitempty"`

	// add to projects
	AddToProjects []*NewOrganizationInvitationProjectEntryDTO `json:"addToProjects"`

	// invitation type
	// Required: true
	InvitationType *InvitationTypeEnumDTO `json:"invitationType"`

	// invitee
	// Required: true
	Invitee *string `json:"invitee"`

	// role grant
	// Required: true
	RoleGrant *OrganizationRoleDTO `json:"roleGrant"`
}

// Validate validates this new organization invitation entry d t o
func (m *NewOrganizationInvitationEntryDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddToProjects(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInvitationType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInvitee(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoleGrant(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewOrganizationInvitationEntryDTO) validateAddToProjects(formats strfmt.Registry) error {
	if swag.IsZero(m.AddToProjects) { // not required
		return nil
	}

	for i := 0; i < len(m.AddToProjects); i++ {
		if swag.IsZero(m.AddToProjects[i]) { // not required
			continue
		}

		if m.AddToProjects[i] != nil {
			if err := m.AddToProjects[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("addToProjects" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("addToProjects" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *NewOrganizationInvitationEntryDTO) validateInvitationType(formats strfmt.Registry) error {

	if err := validate.Required("invitationType", "body", m.InvitationType); err != nil {
		return err
	}

	if err := validate.Required("invitationType", "body", m.InvitationType); err != nil {
		return err
	}

	if m.InvitationType != nil {
		if err := m.InvitationType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("invitationType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("invitationType")
			}
			return err
		}
	}

	return nil
}

func (m *NewOrganizationInvitationEntryDTO) validateInvitee(formats strfmt.Registry) error {

	if err := validate.Required("invitee", "body", m.Invitee); err != nil {
		return err
	}

	return nil
}

func (m *NewOrganizationInvitationEntryDTO) validateRoleGrant(formats strfmt.Registry) error {

	if err := validate.Required("roleGrant", "body", m.RoleGrant); err != nil {
		return err
	}

	if err := validate.Required("roleGrant", "body", m.RoleGrant); err != nil {
		return err
	}

	if m.RoleGrant != nil {
		if err := m.RoleGrant.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("roleGrant")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("roleGrant")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this new organization invitation entry d t o based on the context it is used
func (m *NewOrganizationInvitationEntryDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAddToProjects(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInvitationType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRoleGrant(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewOrganizationInvitationEntryDTO) contextValidateAddToProjects(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AddToProjects); i++ {

		if m.AddToProjects[i] != nil {

			if swag.IsZero(m.AddToProjects[i]) { // not required
				return nil
			}

			if err := m.AddToProjects[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("addToProjects" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("addToProjects" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *NewOrganizationInvitationEntryDTO) contextValidateInvitationType(ctx context.Context, formats strfmt.Registry) error {

	if m.InvitationType != nil {

		if err := m.InvitationType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("invitationType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("invitationType")
			}
			return err
		}
	}

	return nil
}

func (m *NewOrganizationInvitationEntryDTO) contextValidateRoleGrant(ctx context.Context, formats strfmt.Registry) error {

	if m.RoleGrant != nil {

		if err := m.RoleGrant.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("roleGrant")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("roleGrant")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NewOrganizationInvitationEntryDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewOrganizationInvitationEntryDTO) UnmarshalBinary(b []byte) error {
	var res NewOrganizationInvitationEntryDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
