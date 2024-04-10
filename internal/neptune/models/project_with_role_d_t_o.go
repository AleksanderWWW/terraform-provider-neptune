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

// ProjectWithRoleDTO project with role d t o
//
// swagger:model ProjectWithRoleDTO
type ProjectWithRoleDTO struct {

	// archived
	// Required: true
	Archived *bool `json:"archived"`

	// avatar source
	// Required: true
	AvatarSource *AvatarSourceEnum `json:"avatarSource"`

	// avatar Url
	// Required: true
	AvatarURL *string `json:"avatarUrl"`

	// background Url
	BackgroundURL string `json:"backgroundUrl,omitempty"`

	// code access
	// Required: true
	CodeAccess *ProjectCodeAccessDTO `json:"codeAccess"`

	// description
	Description string `json:"description,omitempty"`

	// display class
	DisplayClass string `json:"displayClass,omitempty"`

	// featured
	// Required: true
	Featured *bool `json:"featured"`

	// id
	// Required: true
	// Format: uuid
	ID *strfmt.UUID `json:"id"`

	// last activity
	// Required: true
	// Format: date-time
	LastActivity *strfmt.DateTime `json:"lastActivity"`

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

	// time of creation
	// Required: true
	// Format: date-time
	TimeOfCreation *strfmt.DateTime `json:"timeOfCreation"`

	// user count
	// Required: true
	UserCount *int32 `json:"userCount"`

	// user role in organization
	UserRoleInOrganization OrganizationRoleDTO `json:"userRoleInOrganization,omitempty"`

	// user role in project
	// Required: true
	UserRoleInProject *ProjectRoleDTO `json:"userRoleInProject"`

	// version
	// Required: true
	Version *int32 `json:"version"`

	// visibility
	// Required: true
	Visibility *ProjectVisibilityDTO `json:"visibility"`
}

// Validate validates this project with role d t o
func (m *ProjectWithRoleDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateArchived(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAvatarSource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAvatarURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCodeAccess(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFeatured(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastActivity(formats); err != nil {
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

	if err := m.validateTimeOfCreation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserRoleInOrganization(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserRoleInProject(formats); err != nil {
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

func (m *ProjectWithRoleDTO) validateArchived(formats strfmt.Registry) error {

	if err := validate.Required("archived", "body", m.Archived); err != nil {
		return err
	}

	return nil
}

func (m *ProjectWithRoleDTO) validateAvatarSource(formats strfmt.Registry) error {

	if err := validate.Required("avatarSource", "body", m.AvatarSource); err != nil {
		return err
	}

	if err := validate.Required("avatarSource", "body", m.AvatarSource); err != nil {
		return err
	}

	if m.AvatarSource != nil {
		if err := m.AvatarSource.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("avatarSource")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("avatarSource")
			}
			return err
		}
	}

	return nil
}

func (m *ProjectWithRoleDTO) validateAvatarURL(formats strfmt.Registry) error {

	if err := validate.Required("avatarUrl", "body", m.AvatarURL); err != nil {
		return err
	}

	return nil
}

func (m *ProjectWithRoleDTO) validateCodeAccess(formats strfmt.Registry) error {

	if err := validate.Required("codeAccess", "body", m.CodeAccess); err != nil {
		return err
	}

	if err := validate.Required("codeAccess", "body", m.CodeAccess); err != nil {
		return err
	}

	if m.CodeAccess != nil {
		if err := m.CodeAccess.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("codeAccess")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("codeAccess")
			}
			return err
		}
	}

	return nil
}

func (m *ProjectWithRoleDTO) validateFeatured(formats strfmt.Registry) error {

	if err := validate.Required("featured", "body", m.Featured); err != nil {
		return err
	}

	return nil
}

func (m *ProjectWithRoleDTO) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ProjectWithRoleDTO) validateLastActivity(formats strfmt.Registry) error {

	if err := validate.Required("lastActivity", "body", m.LastActivity); err != nil {
		return err
	}

	if err := validate.FormatOf("lastActivity", "body", "date-time", m.LastActivity.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ProjectWithRoleDTO) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *ProjectWithRoleDTO) validateOrganizationID(formats strfmt.Registry) error {

	if err := validate.Required("organizationId", "body", m.OrganizationID); err != nil {
		return err
	}

	if err := validate.FormatOf("organizationId", "body", "uuid", m.OrganizationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ProjectWithRoleDTO) validateOrganizationName(formats strfmt.Registry) error {

	if err := validate.Required("organizationName", "body", m.OrganizationName); err != nil {
		return err
	}

	return nil
}

func (m *ProjectWithRoleDTO) validateOrganizationType(formats strfmt.Registry) error {

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

func (m *ProjectWithRoleDTO) validateProjectKey(formats strfmt.Registry) error {

	if err := validate.Required("projectKey", "body", m.ProjectKey); err != nil {
		return err
	}

	return nil
}

func (m *ProjectWithRoleDTO) validateTimeOfCreation(formats strfmt.Registry) error {

	if err := validate.Required("timeOfCreation", "body", m.TimeOfCreation); err != nil {
		return err
	}

	if err := validate.FormatOf("timeOfCreation", "body", "date-time", m.TimeOfCreation.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ProjectWithRoleDTO) validateUserCount(formats strfmt.Registry) error {

	if err := validate.Required("userCount", "body", m.UserCount); err != nil {
		return err
	}

	return nil
}

func (m *ProjectWithRoleDTO) validateUserRoleInOrganization(formats strfmt.Registry) error {
	if swag.IsZero(m.UserRoleInOrganization) { // not required
		return nil
	}

	if err := m.UserRoleInOrganization.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("userRoleInOrganization")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("userRoleInOrganization")
		}
		return err
	}

	return nil
}

func (m *ProjectWithRoleDTO) validateUserRoleInProject(formats strfmt.Registry) error {

	if err := validate.Required("userRoleInProject", "body", m.UserRoleInProject); err != nil {
		return err
	}

	if err := validate.Required("userRoleInProject", "body", m.UserRoleInProject); err != nil {
		return err
	}

	if m.UserRoleInProject != nil {
		if err := m.UserRoleInProject.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("userRoleInProject")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("userRoleInProject")
			}
			return err
		}
	}

	return nil
}

func (m *ProjectWithRoleDTO) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

func (m *ProjectWithRoleDTO) validateVisibility(formats strfmt.Registry) error {

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

// ContextValidate validate this project with role d t o based on the context it is used
func (m *ProjectWithRoleDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAvatarSource(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCodeAccess(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOrganizationType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUserRoleInOrganization(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUserRoleInProject(ctx, formats); err != nil {
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

func (m *ProjectWithRoleDTO) contextValidateAvatarSource(ctx context.Context, formats strfmt.Registry) error {

	if m.AvatarSource != nil {

		if err := m.AvatarSource.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("avatarSource")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("avatarSource")
			}
			return err
		}
	}

	return nil
}

func (m *ProjectWithRoleDTO) contextValidateCodeAccess(ctx context.Context, formats strfmt.Registry) error {

	if m.CodeAccess != nil {

		if err := m.CodeAccess.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("codeAccess")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("codeAccess")
			}
			return err
		}
	}

	return nil
}

func (m *ProjectWithRoleDTO) contextValidateOrganizationType(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ProjectWithRoleDTO) contextValidateUserRoleInOrganization(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.UserRoleInOrganization) { // not required
		return nil
	}

	if err := m.UserRoleInOrganization.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("userRoleInOrganization")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("userRoleInOrganization")
		}
		return err
	}

	return nil
}

func (m *ProjectWithRoleDTO) contextValidateUserRoleInProject(ctx context.Context, formats strfmt.Registry) error {

	if m.UserRoleInProject != nil {

		if err := m.UserRoleInProject.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("userRoleInProject")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("userRoleInProject")
			}
			return err
		}
	}

	return nil
}

func (m *ProjectWithRoleDTO) contextValidateVisibility(ctx context.Context, formats strfmt.Registry) error {

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
func (m *ProjectWithRoleDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProjectWithRoleDTO) UnmarshalBinary(b []byte) error {
	var res ProjectWithRoleDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}