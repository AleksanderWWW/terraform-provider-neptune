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

// ProjectMetadataDTO project metadata d t o
//
// swagger:model ProjectMetadataDTO
type ProjectMetadataDTO struct {

	// archived
	// Required: true
	Archived *bool `json:"archived"`

	// id
	// Required: true
	// Format: uuid
	ID *strfmt.UUID `json:"id"`

	// name
	// Required: true
	Name *string `json:"name"`

	// organization Id
	// Required: true
	// Format: uuid
	OrganizationID *strfmt.UUID `json:"organizationId"`

	// organization name
	// Required: true
	OrganizationName *string `json:"organizationName"`

	// organization type
	// Required: true
	OrganizationType *OrganizationTypeDTO `json:"organizationType"`

	// project key
	// Required: true
	ProjectKey *string `json:"projectKey"`

	// storage usage
	// Required: true
	StorageUsage *int64 `json:"storageUsage"`

	// time of creation
	// Required: true
	// Format: date-time
	TimeOfCreation *strfmt.DateTime `json:"timeOfCreation"`

	// traits
	// Required: true
	Traits *ProjectTraitsSetDTO `json:"traits"`

	// version
	// Required: true
	Version *int32 `json:"version"`

	// visibility
	// Required: true
	Visibility *ProjectVisibilityDTO `json:"visibility"`
}

// Validate validates this project metadata d t o
func (m *ProjectMetadataDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateArchived(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganizationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganizationName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganizationType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjectKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStorageUsage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimeOfCreation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTraits(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVisibility(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProjectMetadataDTO) validateArchived(formats strfmt.Registry) error {

	if err := validate.Required("archived", "body", m.Archived); err != nil {
		return err
	}

	return nil
}

func (m *ProjectMetadataDTO) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ProjectMetadataDTO) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *ProjectMetadataDTO) validateOrganizationID(formats strfmt.Registry) error {

	if err := validate.Required("organizationId", "body", m.OrganizationID); err != nil {
		return err
	}

	if err := validate.FormatOf("organizationId", "body", "uuid", m.OrganizationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ProjectMetadataDTO) validateOrganizationName(formats strfmt.Registry) error {

	if err := validate.Required("organizationName", "body", m.OrganizationName); err != nil {
		return err
	}

	return nil
}

func (m *ProjectMetadataDTO) validateOrganizationType(formats strfmt.Registry) error {

	if err := validate.Required("organizationType", "body", m.OrganizationType); err != nil {
		return err
	}

	if err := validate.Required("organizationType", "body", m.OrganizationType); err != nil {
		return err
	}

	if m.OrganizationType != nil {
		if err := m.OrganizationType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("organizationType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("organizationType")
			}
			return err
		}
	}

	return nil
}

func (m *ProjectMetadataDTO) validateProjectKey(formats strfmt.Registry) error {

	if err := validate.Required("projectKey", "body", m.ProjectKey); err != nil {
		return err
	}

	return nil
}

func (m *ProjectMetadataDTO) validateStorageUsage(formats strfmt.Registry) error {

	if err := validate.Required("storageUsage", "body", m.StorageUsage); err != nil {
		return err
	}

	return nil
}

func (m *ProjectMetadataDTO) validateTimeOfCreation(formats strfmt.Registry) error {

	if err := validate.Required("timeOfCreation", "body", m.TimeOfCreation); err != nil {
		return err
	}

	if err := validate.FormatOf("timeOfCreation", "body", "date-time", m.TimeOfCreation.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ProjectMetadataDTO) validateTraits(formats strfmt.Registry) error {

	if err := validate.Required("traits", "body", m.Traits); err != nil {
		return err
	}

	if m.Traits != nil {
		if err := m.Traits.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("traits")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("traits")
			}
			return err
		}
	}

	return nil
}

func (m *ProjectMetadataDTO) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

func (m *ProjectMetadataDTO) validateVisibility(formats strfmt.Registry) error {

	if err := validate.Required("visibility", "body", m.Visibility); err != nil {
		return err
	}

	if err := validate.Required("visibility", "body", m.Visibility); err != nil {
		return err
	}

	if m.Visibility != nil {
		if err := m.Visibility.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("visibility")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("visibility")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this project metadata d t o based on the context it is used
func (m *ProjectMetadataDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOrganizationType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTraits(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVisibility(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProjectMetadataDTO) contextValidateOrganizationType(ctx context.Context, formats strfmt.Registry) error {

	if m.OrganizationType != nil {

		if err := m.OrganizationType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("organizationType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("organizationType")
			}
			return err
		}
	}

	return nil
}

func (m *ProjectMetadataDTO) contextValidateTraits(ctx context.Context, formats strfmt.Registry) error {

	if m.Traits != nil {

		if err := m.Traits.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("traits")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("traits")
			}
			return err
		}
	}

	return nil
}

func (m *ProjectMetadataDTO) contextValidateVisibility(ctx context.Context, formats strfmt.Registry) error {

	if m.Visibility != nil {

		if err := m.Visibility.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("visibility")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("visibility")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProjectMetadataDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProjectMetadataDTO) UnmarshalBinary(b []byte) error {
	var res ProjectMetadataDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}