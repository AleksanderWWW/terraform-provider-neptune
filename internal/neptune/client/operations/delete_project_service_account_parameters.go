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

// NewDeleteProjectServiceAccountParams creates a new DeleteProjectServiceAccountParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteProjectServiceAccountParams() *DeleteProjectServiceAccountParams {
	return &DeleteProjectServiceAccountParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteProjectServiceAccountParamsWithTimeout creates a new DeleteProjectServiceAccountParams object
// with the ability to set a timeout on a request.
func NewDeleteProjectServiceAccountParamsWithTimeout(timeout time.Duration) *DeleteProjectServiceAccountParams {
	return &DeleteProjectServiceAccountParams{
		timeout: timeout,
	}
}

// NewDeleteProjectServiceAccountParamsWithContext creates a new DeleteProjectServiceAccountParams object
// with the ability to set a context for a request.
func NewDeleteProjectServiceAccountParamsWithContext(ctx context.Context) *DeleteProjectServiceAccountParams {
	return &DeleteProjectServiceAccountParams{
		Context: ctx,
	}
}

// NewDeleteProjectServiceAccountParamsWithHTTPClient creates a new DeleteProjectServiceAccountParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteProjectServiceAccountParamsWithHTTPClient(client *http.Client) *DeleteProjectServiceAccountParams {
	return &DeleteProjectServiceAccountParams{
		HTTPClient: client,
	}
}

/*
DeleteProjectServiceAccountParams contains all the parameters to send to the API endpoint

	for the delete project service account operation.

	Typically these are written to a http.Request.
*/
type DeleteProjectServiceAccountParams struct {

	// ProjectIdentifier.
	ProjectIdentifier string

	// ServiceAccountID.
	ServiceAccountID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete project service account params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteProjectServiceAccountParams) WithDefaults() *DeleteProjectServiceAccountParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete project service account params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteProjectServiceAccountParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete project service account params
func (o *DeleteProjectServiceAccountParams) WithTimeout(timeout time.Duration) *DeleteProjectServiceAccountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete project service account params
func (o *DeleteProjectServiceAccountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete project service account params
func (o *DeleteProjectServiceAccountParams) WithContext(ctx context.Context) *DeleteProjectServiceAccountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete project service account params
func (o *DeleteProjectServiceAccountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete project service account params
func (o *DeleteProjectServiceAccountParams) WithHTTPClient(client *http.Client) *DeleteProjectServiceAccountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete project service account params
func (o *DeleteProjectServiceAccountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProjectIdentifier adds the projectIdentifier to the delete project service account params
func (o *DeleteProjectServiceAccountParams) WithProjectIdentifier(projectIdentifier string) *DeleteProjectServiceAccountParams {
	o.SetProjectIdentifier(projectIdentifier)
	return o
}

// SetProjectIdentifier adds the projectIdentifier to the delete project service account params
func (o *DeleteProjectServiceAccountParams) SetProjectIdentifier(projectIdentifier string) {
	o.ProjectIdentifier = projectIdentifier
}

// WithServiceAccountID adds the serviceAccountID to the delete project service account params
func (o *DeleteProjectServiceAccountParams) WithServiceAccountID(serviceAccountID string) *DeleteProjectServiceAccountParams {
	o.SetServiceAccountID(serviceAccountID)
	return o
}

// SetServiceAccountID adds the serviceAccountId to the delete project service account params
func (o *DeleteProjectServiceAccountParams) SetServiceAccountID(serviceAccountID string) {
	o.ServiceAccountID = serviceAccountID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteProjectServiceAccountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param serviceAccountId
	if err := r.SetPathParam("serviceAccountId", o.ServiceAccountID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
