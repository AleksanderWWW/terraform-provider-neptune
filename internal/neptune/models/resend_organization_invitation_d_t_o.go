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

// ResendOrganizationInvitationDTO resend organization invitation d t o
//
// swagger:model ResendOrganizationInvitationDTO
type ResendOrganizationInvitationDTO struct {

	// invitee email
	// Required: true
	InviteeEmail *string `json:"inviteeEmail"`

	// organization identifier
	// Required: true
	OrganizationIdentifier *string `json:"organizationIdentifier"`
}

// Validate validates this resend organization invitation d t o
func (m *ResendOrganizationInvitationDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInviteeEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganizationIdentifier(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResendOrganizationInvitationDTO) validateInviteeEmail(formats strfmt.Registry) error {

	if err := validate.Required("inviteeEmail", "body", m.InviteeEmail); err != nil {
		return err
	}

	return nil
}

func (m *ResendOrganizationInvitationDTO) validateOrganizationIdentifier(formats strfmt.Registry) error {

	if err := validate.Required("organizationIdentifier", "body", m.OrganizationIdentifier); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this resend organization invitation d t o based on context it is used
func (m *ResendOrganizationInvitationDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ResendOrganizationInvitationDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResendOrganizationInvitationDTO) UnmarshalBinary(b []byte) error {
	var res ResendOrganizationInvitationDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
