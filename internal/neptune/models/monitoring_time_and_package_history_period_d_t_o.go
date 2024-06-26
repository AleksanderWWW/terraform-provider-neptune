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

// MonitoringTimeAndPackageHistoryPeriodDTO monitoring time and package history period d t o
//
// swagger:model MonitoringTimeAndPackageHistoryPeriodDTO
type MonitoringTimeAndPackageHistoryPeriodDTO struct {

	// period end
	// Required: true
	// Format: date-time
	PeriodEnd *strfmt.DateTime `json:"periodEnd"`

	// period start
	// Required: true
	// Format: date-time
	PeriodStart *strfmt.DateTime `json:"periodStart"`

	// top up seconds
	// Required: true
	TopUpSeconds *int64 `json:"topUpSeconds"`

	// used seconds
	// Required: true
	UsedSeconds *int64 `json:"usedSeconds"`
}

// Validate validates this monitoring time and package history period d t o
func (m *MonitoringTimeAndPackageHistoryPeriodDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePeriodEnd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePeriodStart(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTopUpSeconds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsedSeconds(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MonitoringTimeAndPackageHistoryPeriodDTO) validatePeriodEnd(formats strfmt.Registry) error {

	if err := validate.Required("periodEnd", "body", m.PeriodEnd); err != nil {
		return err
	}

	if err := validate.FormatOf("periodEnd", "body", "date-time", m.PeriodEnd.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MonitoringTimeAndPackageHistoryPeriodDTO) validatePeriodStart(formats strfmt.Registry) error {

	if err := validate.Required("periodStart", "body", m.PeriodStart); err != nil {
		return err
	}

	if err := validate.FormatOf("periodStart", "body", "date-time", m.PeriodStart.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MonitoringTimeAndPackageHistoryPeriodDTO) validateTopUpSeconds(formats strfmt.Registry) error {

	if err := validate.Required("topUpSeconds", "body", m.TopUpSeconds); err != nil {
		return err
	}

	return nil
}

func (m *MonitoringTimeAndPackageHistoryPeriodDTO) validateUsedSeconds(formats strfmt.Registry) error {

	if err := validate.Required("usedSeconds", "body", m.UsedSeconds); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this monitoring time and package history period d t o based on context it is used
func (m *MonitoringTimeAndPackageHistoryPeriodDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MonitoringTimeAndPackageHistoryPeriodDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MonitoringTimeAndPackageHistoryPeriodDTO) UnmarshalBinary(b []byte) error {
	var res MonitoringTimeAndPackageHistoryPeriodDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
