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

// WorkspaceModeDTO workspace mode d t o
//
// swagger:model WorkspaceModeDTO
type WorkspaceModeDTO string

func NewWorkspaceModeDTO(value WorkspaceModeDTO) *WorkspaceModeDTO {
	return &value
}

// Pointer returns a pointer to a freshly-allocated WorkspaceModeDTO.
func (m WorkspaceModeDTO) Pointer() *WorkspaceModeDTO {
	return &m
}

const (

	// WorkspaceModeDTOReadOnly captures enum value "read_only"
	WorkspaceModeDTOReadOnly WorkspaceModeDTO = "read_only"

	// WorkspaceModeDTONormal captures enum value "normal"
	WorkspaceModeDTONormal WorkspaceModeDTO = "normal"
)

// for schema
var workspaceModeDTOEnum []interface{}

func init() {
	var res []WorkspaceModeDTO
	if err := json.Unmarshal([]byte(`["read_only","normal"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		workspaceModeDTOEnum = append(workspaceModeDTOEnum, v)
	}
}

func (m WorkspaceModeDTO) validateWorkspaceModeDTOEnum(path, location string, value WorkspaceModeDTO) error {
	if err := validate.EnumCase(path, location, value, workspaceModeDTOEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this workspace mode d t o
func (m WorkspaceModeDTO) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateWorkspaceModeDTOEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this workspace mode d t o based on context it is used
func (m WorkspaceModeDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
