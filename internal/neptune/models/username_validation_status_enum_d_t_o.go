// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// UsernameValidationStatusEnumDTO username validation status enum d t o
//
// swagger:model UsernameValidationStatusEnumDTO
type UsernameValidationStatusEnumDTO string

func NewUsernameValidationStatusEnumDTO(value UsernameValidationStatusEnumDTO) *UsernameValidationStatusEnumDTO {
	return &value
}

// Pointer returns a pointer to a freshly-allocated UsernameValidationStatusEnumDTO.
func (m UsernameValidationStatusEnumDTO) Pointer() *UsernameValidationStatusEnumDTO {
	return &m
}

const (

	// UsernameValidationStatusEnumDTOAvailable captures enum value "available"
	UsernameValidationStatusEnumDTOAvailable UsernameValidationStatusEnumDTO = "available"

	// UsernameValidationStatusEnumDTOInvalid captures enum value "invalid"
	UsernameValidationStatusEnumDTOInvalid UsernameValidationStatusEnumDTO = "invalid"

	// UsernameValidationStatusEnumDTOUnavailable captures enum value "unavailable"
	UsernameValidationStatusEnumDTOUnavailable UsernameValidationStatusEnumDTO = "unavailable"
)

// for schema
var usernameValidationStatusEnumDTOEnum []interface{}

func init() {
	var res []UsernameValidationStatusEnumDTO
	if err := json.Unmarshal([]byte(`["available","invalid","unavailable"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		usernameValidationStatusEnumDTOEnum = append(usernameValidationStatusEnumDTOEnum, v)
	}
}

func (m UsernameValidationStatusEnumDTO) validateUsernameValidationStatusEnumDTOEnum(path, location string, value UsernameValidationStatusEnumDTO) error {
	if err := validate.EnumCase(path, location, value, usernameValidationStatusEnumDTOEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this username validation status enum d t o
func (m UsernameValidationStatusEnumDTO) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateUsernameValidationStatusEnumDTOEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this username validation status enum d t o based on context it is used
func (m UsernameValidationStatusEnumDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
