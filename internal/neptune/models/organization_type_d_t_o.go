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

// OrganizationTypeDTO organization type d t o
//
// swagger:model OrganizationTypeDTO
type OrganizationTypeDTO string

func NewOrganizationTypeDTO(value OrganizationTypeDTO) *OrganizationTypeDTO {
	return &value
}

// Pointer returns a pointer to a freshly-allocated OrganizationTypeDTO.
func (m OrganizationTypeDTO) Pointer() *OrganizationTypeDTO {
	return &m
}

const (

	// OrganizationTypeDTOIndividual captures enum value "individual"
	OrganizationTypeDTOIndividual OrganizationTypeDTO = "individual"

	// OrganizationTypeDTOTeam captures enum value "team"
	OrganizationTypeDTOTeam OrganizationTypeDTO = "team"
)

// for schema
var organizationTypeDTOEnum []interface{}

func init() {
	var res []OrganizationTypeDTO
	if err := json.Unmarshal([]byte(`["individual","team"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		organizationTypeDTOEnum = append(organizationTypeDTOEnum, v)
	}
}

func (m OrganizationTypeDTO) validateOrganizationTypeDTOEnum(path, location string, value OrganizationTypeDTO) error {
	if err := validate.EnumCase(path, location, value, organizationTypeDTOEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this organization type d t o
func (m OrganizationTypeDTO) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateOrganizationTypeDTOEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this organization type d t o based on context it is used
func (m OrganizationTypeDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
