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

// ClientConfig client config
//
// swagger:model ClientConfig
type ClientConfig struct {

	// api Url
	// Required: true
	APIURL *string `json:"apiUrl"`

	// application Url
	// Required: true
	ApplicationURL *string `json:"applicationUrl"`

	// artifacts
	Artifacts *ArtifactsConfigDTO `json:"artifacts,omitempty"`

	// gzip upload
	// Required: true
	GzipUpload *bool `json:"gzipUpload"`

	// multi part upload
	MultiPartUpload *MultiPartUploadConfigDTO `json:"multiPartUpload,omitempty"`

	// py lib versions
	// Required: true
	PyLibVersions *ClientVersionsConfigDTO `json:"pyLibVersions"`
}

// Validate validates this client config
func (m *ClientConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAPIURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateApplicationURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateArtifacts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGzipUpload(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMultiPartUpload(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePyLibVersions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClientConfig) validateAPIURL(formats strfmt.Registry) error {

	if err := validate.Required("apiUrl", "body", m.APIURL); err != nil {
		return err
	}

	return nil
}

func (m *ClientConfig) validateApplicationURL(formats strfmt.Registry) error {

	if err := validate.Required("applicationUrl", "body", m.ApplicationURL); err != nil {
		return err
	}

	return nil
}

func (m *ClientConfig) validateArtifacts(formats strfmt.Registry) error {
	if swag.IsZero(m.Artifacts) { // not required
		return nil
	}

	if m.Artifacts != nil {
		if err := m.Artifacts.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("artifacts")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("artifacts")
			}
			return err
		}
	}

	return nil
}

func (m *ClientConfig) validateGzipUpload(formats strfmt.Registry) error {

	if err := validate.Required("gzipUpload", "body", m.GzipUpload); err != nil {
		return err
	}

	return nil
}

func (m *ClientConfig) validateMultiPartUpload(formats strfmt.Registry) error {
	if swag.IsZero(m.MultiPartUpload) { // not required
		return nil
	}

	if m.MultiPartUpload != nil {
		if err := m.MultiPartUpload.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("multiPartUpload")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("multiPartUpload")
			}
			return err
		}
	}

	return nil
}

func (m *ClientConfig) validatePyLibVersions(formats strfmt.Registry) error {

	if err := validate.Required("pyLibVersions", "body", m.PyLibVersions); err != nil {
		return err
	}

	if m.PyLibVersions != nil {
		if err := m.PyLibVersions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pyLibVersions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pyLibVersions")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this client config based on the context it is used
func (m *ClientConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateArtifacts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMultiPartUpload(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePyLibVersions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClientConfig) contextValidateArtifacts(ctx context.Context, formats strfmt.Registry) error {

	if m.Artifacts != nil {

		if swag.IsZero(m.Artifacts) { // not required
			return nil
		}

		if err := m.Artifacts.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("artifacts")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("artifacts")
			}
			return err
		}
	}

	return nil
}

func (m *ClientConfig) contextValidateMultiPartUpload(ctx context.Context, formats strfmt.Registry) error {

	if m.MultiPartUpload != nil {

		if swag.IsZero(m.MultiPartUpload) { // not required
			return nil
		}

		if err := m.MultiPartUpload.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("multiPartUpload")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("multiPartUpload")
			}
			return err
		}
	}

	return nil
}

func (m *ClientConfig) contextValidatePyLibVersions(ctx context.Context, formats strfmt.Registry) error {

	if m.PyLibVersions != nil {

		if err := m.PyLibVersions.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pyLibVersions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pyLibVersions")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClientConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClientConfig) UnmarshalBinary(b []byte) error {
	var res ClientConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}