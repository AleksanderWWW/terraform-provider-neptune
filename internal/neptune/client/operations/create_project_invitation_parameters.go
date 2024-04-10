// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"terraform-provider-neptune/internal/neptune/models"
)

// NewCreateProjectInvitationParams creates a new CreateProjectInvitationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateProjectInvitationParams() *CreateProjectInvitationParams {
	return &CreateProjectInvitationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateProjectInvitationParamsWithTimeout creates a new CreateProjectInvitationParams object
// with the ability to set a timeout on a request.
func NewCreateProjectInvitationParamsWithTimeout(timeout time.Duration) *CreateProjectInvitationParams {
	return &CreateProjectInvitationParams{
		timeout: timeout,
	}
}

// NewCreateProjectInvitationParamsWithContext creates a new CreateProjectInvitationParams object
// with the ability to set a context for a request.
func NewCreateProjectInvitationParamsWithContext(ctx context.Context) *CreateProjectInvitationParams {
	return &CreateProjectInvitationParams{
		Context: ctx,
	}
}

// NewCreateProjectInvitationParamsWithHTTPClient creates a new CreateProjectInvitationParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateProjectInvitationParamsWithHTTPClient(client *http.Client) *CreateProjectInvitationParams {
	return &CreateProjectInvitationParams{
		HTTPClient: client,
	}
}

/*
CreateProjectInvitationParams contains all the parameters to send to the API endpoint

	for the create project invitation operation.

	Typically these are written to a http.Request.
*/
type CreateProjectInvitationParams struct {

	// NewProjectInvitation.
	NewProjectInvitation *models.NewProjectInvitationDTO

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create project invitation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateProjectInvitationParams) WithDefaults() *CreateProjectInvitationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create project invitation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateProjectInvitationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create project invitation params
func (o *CreateProjectInvitationParams) WithTimeout(timeout time.Duration) *CreateProjectInvitationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create project invitation params
func (o *CreateProjectInvitationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create project invitation params
func (o *CreateProjectInvitationParams) WithContext(ctx context.Context) *CreateProjectInvitationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create project invitation params
func (o *CreateProjectInvitationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create project invitation params
func (o *CreateProjectInvitationParams) WithHTTPClient(client *http.Client) *CreateProjectInvitationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create project invitation params
func (o *CreateProjectInvitationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNewProjectInvitation adds the newProjectInvitation to the create project invitation params
func (o *CreateProjectInvitationParams) WithNewProjectInvitation(newProjectInvitation *models.NewProjectInvitationDTO) *CreateProjectInvitationParams {
	o.SetNewProjectInvitation(newProjectInvitation)
	return o
}

// SetNewProjectInvitation adds the newProjectInvitation to the create project invitation params
func (o *CreateProjectInvitationParams) SetNewProjectInvitation(newProjectInvitation *models.NewProjectInvitationDTO) {
	o.NewProjectInvitation = newProjectInvitation
}

// WriteToRequest writes these params to a swagger request
func (o *CreateProjectInvitationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.NewProjectInvitation != nil {
		if err := r.SetBodyParam(o.NewProjectInvitation); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
