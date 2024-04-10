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

// NewUpdateLastViewedParams creates a new UpdateLastViewedParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateLastViewedParams() *UpdateLastViewedParams {
	return &UpdateLastViewedParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateLastViewedParamsWithTimeout creates a new UpdateLastViewedParams object
// with the ability to set a timeout on a request.
func NewUpdateLastViewedParamsWithTimeout(timeout time.Duration) *UpdateLastViewedParams {
	return &UpdateLastViewedParams{
		timeout: timeout,
	}
}

// NewUpdateLastViewedParamsWithContext creates a new UpdateLastViewedParams object
// with the ability to set a context for a request.
func NewUpdateLastViewedParamsWithContext(ctx context.Context) *UpdateLastViewedParams {
	return &UpdateLastViewedParams{
		Context: ctx,
	}
}

// NewUpdateLastViewedParamsWithHTTPClient creates a new UpdateLastViewedParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateLastViewedParamsWithHTTPClient(client *http.Client) *UpdateLastViewedParams {
	return &UpdateLastViewedParams{
		HTTPClient: client,
	}
}

/*
UpdateLastViewedParams contains all the parameters to send to the API endpoint

	for the update last viewed operation.

	Typically these are written to a http.Request.
*/
type UpdateLastViewedParams struct {

	// ProjectID.
	//
	// Format: uuid
	ProjectID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update last viewed params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateLastViewedParams) WithDefaults() *UpdateLastViewedParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update last viewed params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateLastViewedParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update last viewed params
func (o *UpdateLastViewedParams) WithTimeout(timeout time.Duration) *UpdateLastViewedParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update last viewed params
func (o *UpdateLastViewedParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update last viewed params
func (o *UpdateLastViewedParams) WithContext(ctx context.Context) *UpdateLastViewedParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update last viewed params
func (o *UpdateLastViewedParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update last viewed params
func (o *UpdateLastViewedParams) WithHTTPClient(client *http.Client) *UpdateLastViewedParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update last viewed params
func (o *UpdateLastViewedParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProjectID adds the projectID to the update last viewed params
func (o *UpdateLastViewedParams) WithProjectID(projectID strfmt.UUID) *UpdateLastViewedParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the update last viewed params
func (o *UpdateLastViewedParams) SetProjectID(projectID strfmt.UUID) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateLastViewedParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param projectId
	qrProjectID := o.ProjectID
	qProjectID := qrProjectID.String()
	if qProjectID != "" {

		if err := r.SetQueryParam("projectId", qProjectID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
