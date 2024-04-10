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

// MultiPartUploadConfigDTO multi part upload config d t o
//
// swagger:model MultiPartUploadConfigDTO
type MultiPartUploadConfigDTO struct {

	// enabled
	// Required: true
	Enabled *bool `json:"enabled"`

	// max chunk count
	// Required: true
	MaxChunkCount *int32 `json:"maxChunkCount"`

	// max chunk size
	// Required: true
	MaxChunkSize *int32 `json:"maxChunkSize"`

	// max single part size
	// Required: true
	MaxSinglePartSize *int32 `json:"maxSinglePartSize"`

	// min chunk size
	// Required: true
	MinChunkSize *int32 `json:"minChunkSize"`
}

// Validate validates this multi part upload config d t o
func (m *MultiPartUploadConfigDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaxChunkCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaxChunkSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaxSinglePartSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMinChunkSize(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MultiPartUploadConfigDTO) validateEnabled(formats strfmt.Registry) error {

	if err := validate.Required("enabled", "body", m.Enabled); err != nil {
		return err
	}

	return nil
}

func (m *MultiPartUploadConfigDTO) validateMaxChunkCount(formats strfmt.Registry) error {

	if err := validate.Required("maxChunkCount", "body", m.MaxChunkCount); err != nil {
		return err
	}

	return nil
}

func (m *MultiPartUploadConfigDTO) validateMaxChunkSize(formats strfmt.Registry) error {

	if err := validate.Required("maxChunkSize", "body", m.MaxChunkSize); err != nil {
		return err
	}

	return nil
}

func (m *MultiPartUploadConfigDTO) validateMaxSinglePartSize(formats strfmt.Registry) error {

	if err := validate.Required("maxSinglePartSize", "body", m.MaxSinglePartSize); err != nil {
		return err
	}

	return nil
}

func (m *MultiPartUploadConfigDTO) validateMinChunkSize(formats strfmt.Registry) error {

	if err := validate.Required("minChunkSize", "body", m.MinChunkSize); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this multi part upload config d t o based on context it is used
func (m *MultiPartUploadConfigDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MultiPartUploadConfigDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MultiPartUploadConfigDTO) UnmarshalBinary(b []byte) error {
	var res MultiPartUploadConfigDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
