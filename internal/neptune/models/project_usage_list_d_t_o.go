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

// ProjectUsageListDTO project usage list d t o
//
// swagger:model ProjectUsageListDTO
type ProjectUsageListDTO struct {

	// projects
	// Required: true
	Projects []*ProjectUsageDTO `json:"projects"`
}

// Validate validates this project usage list d t o
func (m *ProjectUsageListDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProjects(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProjectUsageListDTO) validateProjects(formats strfmt.Registry) error {

	if err := validate.Required("projects", "body", m.Projects); err != nil {
		return err
	}

	for i := 0; i < len(m.Projects); i++ {
		if swag.IsZero(m.Projects[i]) { // not required
			continue
		}

		if m.Projects[i] != nil {
			if err := m.Projects[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("projects" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("projects" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this project usage list d t o based on the context it is used
func (m *ProjectUsageListDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateProjects(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProjectUsageListDTO) contextValidateProjects(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Projects); i++ {

		if m.Projects[i] != nil {

			if swag.IsZero(m.Projects[i]) { // not required
				return nil
			}

			if err := m.Projects[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("projects" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("projects" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProjectUsageListDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProjectUsageListDTO) UnmarshalBinary(b []byte) error {
	var res ProjectUsageListDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
