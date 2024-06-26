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

// WorkspaceStatusDTO workspace status d t o
//
// swagger:model WorkspaceStatusDTO
type WorkspaceStatusDTO struct {

	// active projects limit
	ActiveProjectsLimit int64 `json:"activeProjectsLimit,omitempty"`

	// active projects usage
	// Required: true
	ActiveProjectsUsage *int64 `json:"activeProjectsUsage"`

	// members count
	// Required: true
	MembersCount *int64 `json:"membersCount"`

	// members limit
	MembersLimit int64 `json:"membersLimit,omitempty"`

	// mode
	// Required: true
	Mode *WorkspaceModeDTO `json:"mode"`

	// monitoring seconds available
	MonitoringSecondsAvailable int64 `json:"monitoringSecondsAvailable,omitempty"`

	// storage bytes available
	// Required: true
	StorageBytesAvailable *int64 `json:"storageBytesAvailable"`

	// storage bytes limit
	// Required: true
	StorageBytesLimit *int64 `json:"storageBytesLimit"`

	// subscription
	Subscription *SubscriptionStateDTO `json:"subscription,omitempty"`

	// trial end date
	// Format: date-time
	TrialEndDate strfmt.DateTime `json:"trialEndDate,omitempty"`
}

// Validate validates this workspace status d t o
func (m *WorkspaceStatusDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActiveProjectsUsage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMembersCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStorageBytesAvailable(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStorageBytesLimit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubscription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrialEndDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WorkspaceStatusDTO) validateActiveProjectsUsage(formats strfmt.Registry) error {

	if err := validate.Required("activeProjectsUsage", "body", m.ActiveProjectsUsage); err != nil {
		return err
	}

	return nil
}

func (m *WorkspaceStatusDTO) validateMembersCount(formats strfmt.Registry) error {

	if err := validate.Required("membersCount", "body", m.MembersCount); err != nil {
		return err
	}

	return nil
}

func (m *WorkspaceStatusDTO) validateMode(formats strfmt.Registry) error {

	if err := validate.Required("mode", "body", m.Mode); err != nil {
		return err
	}

	if err := validate.Required("mode", "body", m.Mode); err != nil {
		return err
	}

	if m.Mode != nil {
		if err := m.Mode.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mode")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mode")
			}
			return err
		}
	}

	return nil
}

func (m *WorkspaceStatusDTO) validateStorageBytesAvailable(formats strfmt.Registry) error {

	if err := validate.Required("storageBytesAvailable", "body", m.StorageBytesAvailable); err != nil {
		return err
	}

	return nil
}

func (m *WorkspaceStatusDTO) validateStorageBytesLimit(formats strfmt.Registry) error {

	if err := validate.Required("storageBytesLimit", "body", m.StorageBytesLimit); err != nil {
		return err
	}

	return nil
}

func (m *WorkspaceStatusDTO) validateSubscription(formats strfmt.Registry) error {
	if swag.IsZero(m.Subscription) { // not required
		return nil
	}

	if m.Subscription != nil {
		if err := m.Subscription.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("subscription")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("subscription")
			}
			return err
		}
	}

	return nil
}

func (m *WorkspaceStatusDTO) validateTrialEndDate(formats strfmt.Registry) error {
	if swag.IsZero(m.TrialEndDate) { // not required
		return nil
	}

	if err := validate.FormatOf("trialEndDate", "body", "date-time", m.TrialEndDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this workspace status d t o based on the context it is used
func (m *WorkspaceStatusDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSubscription(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WorkspaceStatusDTO) contextValidateMode(ctx context.Context, formats strfmt.Registry) error {

	if m.Mode != nil {

		if err := m.Mode.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mode")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mode")
			}
			return err
		}
	}

	return nil
}

func (m *WorkspaceStatusDTO) contextValidateSubscription(ctx context.Context, formats strfmt.Registry) error {

	if m.Subscription != nil {

		if swag.IsZero(m.Subscription) { // not required
			return nil
		}

		if err := m.Subscription.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("subscription")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("subscription")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WorkspaceStatusDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WorkspaceStatusDTO) UnmarshalBinary(b []byte) error {
	var res WorkspaceStatusDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
