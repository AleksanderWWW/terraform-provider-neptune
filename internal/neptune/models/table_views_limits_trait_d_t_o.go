// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TableViewsLimitsTraitDTO table views limits trait d t o
//
// swagger:model TableViewsLimitsTraitDTO
type TableViewsLimitsTraitDTO struct {

	// max table views per project
	MaxTableViewsPerProject int32 `json:"maxTableViewsPerProject,omitempty"`
}

// Validate validates this table views limits trait d t o
func (m *TableViewsLimitsTraitDTO) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this table views limits trait d t o based on context it is used
func (m *TableViewsLimitsTraitDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TableViewsLimitsTraitDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TableViewsLimitsTraitDTO) UnmarshalBinary(b []byte) error {
	var res TableViewsLimitsTraitDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
