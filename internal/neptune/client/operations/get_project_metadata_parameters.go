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
)

// NewGetProjectMetadataParams creates a new GetProjectMetadataParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetProjectMetadataParams() *GetProjectMetadataParams {
	return &GetProjectMetadataParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetProjectMetadataParamsWithTimeout creates a new GetProjectMetadataParams object
// with the ability to set a timeout on a request.
func NewGetProjectMetadataParamsWithTimeout(timeout time.Duration) *GetProjectMetadataParams {
	return &GetProjectMetadataParams{
		timeout: timeout,
	}
}

// NewGetProjectMetadataParamsWithContext creates a new GetProjectMetadataParams object
// with the ability to set a context for a request.
func NewGetProjectMetadataParamsWithContext(ctx context.Context) *GetProjectMetadataParams {
	return &GetProjectMetadataParams{
		Context: ctx,
	}
}

// NewGetProjectMetadataParamsWithHTTPClient creates a new GetProjectMetadataParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetProjectMetadataParamsWithHTTPClient(client *http.Client) *GetProjectMetadataParams {
	return &GetProjectMetadataParams{
		HTTPClient: client,
	}
}

/*
GetProjectMetadataParams contains all the parameters to send to the API endpoint

	for the get project metadata operation.

	Typically these are written to a http.Request.
*/
type GetProjectMetadataParams struct {

	// ProjectIdentifier.
	ProjectIdentifier string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get project metadata params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProjectMetadataParams) WithDefaults() *GetProjectMetadataParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get project metadata params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProjectMetadataParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get project metadata params
func (o *GetProjectMetadataParams) WithTimeout(timeout time.Duration) *GetProjectMetadataParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get project metadata params
func (o *GetProjectMetadataParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get project metadata params
func (o *GetProjectMetadataParams) WithContext(ctx context.Context) *GetProjectMetadataParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get project metadata params
func (o *GetProjectMetadataParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get project metadata params
func (o *GetProjectMetadataParams) WithHTTPClient(client *http.Client) *GetProjectMetadataParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get project metadata params
func (o *GetProjectMetadataParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProjectIdentifier adds the projectIdentifier to the get project metadata params
func (o *GetProjectMetadataParams) WithProjectIdentifier(projectIdentifier string) *GetProjectMetadataParams {
	o.SetProjectIdentifier(projectIdentifier)
	return o
}

// SetProjectIdentifier adds the projectIdentifier to the get project metadata params
func (o *GetProjectMetadataParams) SetProjectIdentifier(projectIdentifier string) {
	o.ProjectIdentifier = projectIdentifier
}

// WriteToRequest writes these params to a swagger request
func (o *GetProjectMetadataParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param projectIdentifier
	qrProjectIdentifier := o.ProjectIdentifier
	qProjectIdentifier := qrProjectIdentifier
	if qProjectIdentifier != "" {

		if err := r.SetQueryParam("projectIdentifier", qProjectIdentifier); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
