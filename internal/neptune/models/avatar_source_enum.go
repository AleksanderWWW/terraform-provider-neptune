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

// AvatarSourceEnum avatar source enum
//
// swagger:model AvatarSourceEnum
type AvatarSourceEnum string

func NewAvatarSourceEnum(value AvatarSourceEnum) *AvatarSourceEnum {
	return &value
}

// Pointer returns a pointer to a freshly-allocated AvatarSourceEnum.
func (m AvatarSourceEnum) Pointer() *AvatarSourceEnum {
	return &m
}

const (

	// AvatarSourceEnumDefault captures enum value "default"
	AvatarSourceEnumDefault AvatarSourceEnum = "default"

	// AvatarSourceEnumThirdParty captures enum value "thirdParty"
	AvatarSourceEnumThirdParty AvatarSourceEnum = "thirdParty"

	// AvatarSourceEnumUser captures enum value "user"
	AvatarSourceEnumUser AvatarSourceEnum = "user"

	// AvatarSourceEnumInherited captures enum value "inherited"
	AvatarSourceEnumInherited AvatarSourceEnum = "inherited"
)

// for schema
var avatarSourceEnumEnum []interface{}

func init() {
	var res []AvatarSourceEnum
	if err := json.Unmarshal([]byte(`["default","thirdParty","user","inherited"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		avatarSourceEnumEnum = append(avatarSourceEnumEnum, v)
	}
}

func (m AvatarSourceEnum) validateAvatarSourceEnumEnum(path, location string, value AvatarSourceEnum) error {
	if err := validate.EnumCase(path, location, value, avatarSourceEnumEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this avatar source enum
func (m AvatarSourceEnum) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateAvatarSourceEnumEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this avatar source enum based on context it is used
func (m AvatarSourceEnum) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}