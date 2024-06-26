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

// ProjectListDTO project list d t o
//
// swagger:model ProjectListDTO
type ProjectListDTO struct {

	// entries
	// Required: true
	Entries []*ProjectWithRoleDTO `json:"entries"`

	// matching item count
	// Required: true
	MatchingItemCount *int32 `json:"matchingItemCount"`

	// total item count
	// Required: true
	TotalItemCount *int32 `json:"totalItemCount"`
}

// Validate validates this project list d t o
func (m *ProjectListDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEntries(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMatchingItemCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotalItemCount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProjectListDTO) validateEntries(formats strfmt.Registry) error {

	if err := validate.Required("entries", "body", m.Entries); err != nil {
		return err
	}

	for i := 0; i < len(m.Entries); i++ {
		if swag.IsZero(m.Entries[i]) { // not required
			continue
		}

		if m.Entries[i] != nil {
			if err := m.Entries[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("entries" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("entries" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ProjectListDTO) validateMatchingItemCount(formats strfmt.Registry) error {

	if err := validate.Required("matchingItemCount", "body", m.MatchingItemCount); err != nil {
		return err
	}

	return nil
}

func (m *ProjectListDTO) validateTotalItemCount(formats strfmt.Registry) error {

	if err := validate.Required("totalItemCount", "body", m.TotalItemCount); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this project list d t o based on the context it is used
func (m *ProjectListDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEntries(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProjectListDTO) contextValidateEntries(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Entries); i++ {

		if m.Entries[i] != nil {

			if swag.IsZero(m.Entries[i]) { // not required
				return nil
			}

			if err := m.Entries[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("entries" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("entries" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProjectListDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProjectListDTO) UnmarshalBinary(b []byte) error {
	var res ProjectListDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
