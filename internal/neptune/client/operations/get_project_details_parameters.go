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

// NewGetProjectDetailsParams creates a new GetProjectDetailsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetProjectDetailsParams() *GetProjectDetailsParams {
	return &GetProjectDetailsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetProjectDetailsParamsWithTimeout creates a new GetProjectDetailsParams object
// with the ability to set a timeout on a request.
func NewGetProjectDetailsParamsWithTimeout(timeout time.Duration) *GetProjectDetailsParams {
	return &GetProjectDetailsParams{
		timeout: timeout,
	}
}

// NewGetProjectDetailsParamsWithContext creates a new GetProjectDetailsParams object
// with the ability to set a context for a request.
func NewGetProjectDetailsParamsWithContext(ctx context.Context) *GetProjectDetailsParams {
	return &GetProjectDetailsParams{
		Context: ctx,
	}
}

// NewGetProjectDetailsParamsWithHTTPClient creates a new GetProjectDetailsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetProjectDetailsParamsWithHTTPClient(client *http.Client) *GetProjectDetailsParams {
	return &GetProjectDetailsParams{
		HTTPClient: client,
	}
}

/*
GetProjectDetailsParams contains all the parameters to send to the API endpoint

	for the get project details operation.

	Typically these are written to a http.Request.
*/
type GetProjectDetailsParams struct {

	// ProjectIdentifier.
	ProjectIdentifier string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get project details params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProjectDetailsParams) WithDefaults() *GetProjectDetailsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get project details params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProjectDetailsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get project details params
func (o *GetProjectDetailsParams) WithTimeout(timeout time.Duration) *GetProjectDetailsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get project details params
func (o *GetProjectDetailsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get project details params
func (o *GetProjectDetailsParams) WithContext(ctx context.Context) *GetProjectDetailsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get project details params
func (o *GetProjectDetailsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get project details params
func (o *GetProjectDetailsParams) WithHTTPClient(client *http.Client) *GetProjectDetailsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get project details params
func (o *GetProjectDetailsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProjectIdentifier adds the projectIdentifier to the get project details params
func (o *GetProjectDetailsParams) WithProjectIdentifier(projectIdentifier string) *GetProjectDetailsParams {
	o.SetProjectIdentifier(projectIdentifier)
	return o
}

// SetProjectIdentifier adds the projectIdentifier to the get project details params
func (o *GetProjectDetailsParams) SetProjectIdentifier(projectIdentifier string) {
	o.ProjectIdentifier = projectIdentifier
}

// WriteToRequest writes these params to a swagger request
func (o *GetProjectDetailsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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