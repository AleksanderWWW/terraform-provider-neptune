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

// NewListProjectServiceAccountsParams creates a new ListProjectServiceAccountsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListProjectServiceAccountsParams() *ListProjectServiceAccountsParams {
	return &ListProjectServiceAccountsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListProjectServiceAccountsParamsWithTimeout creates a new ListProjectServiceAccountsParams object
// with the ability to set a timeout on a request.
func NewListProjectServiceAccountsParamsWithTimeout(timeout time.Duration) *ListProjectServiceAccountsParams {
	return &ListProjectServiceAccountsParams{
		timeout: timeout,
	}
}

// NewListProjectServiceAccountsParamsWithContext creates a new ListProjectServiceAccountsParams object
// with the ability to set a context for a request.
func NewListProjectServiceAccountsParamsWithContext(ctx context.Context) *ListProjectServiceAccountsParams {
	return &ListProjectServiceAccountsParams{
		Context: ctx,
	}
}

// NewListProjectServiceAccountsParamsWithHTTPClient creates a new ListProjectServiceAccountsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListProjectServiceAccountsParamsWithHTTPClient(client *http.Client) *ListProjectServiceAccountsParams {
	return &ListProjectServiceAccountsParams{
		HTTPClient: client,
	}
}

/*
ListProjectServiceAccountsParams contains all the parameters to send to the API endpoint

	for the list project service accounts operation.

	Typically these are written to a http.Request.
*/
type ListProjectServiceAccountsParams struct {

	// ProjectIdentifier.
	ProjectIdentifier string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list project service accounts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListProjectServiceAccountsParams) WithDefaults() *ListProjectServiceAccountsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list project service accounts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListProjectServiceAccountsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list project service accounts params
func (o *ListProjectServiceAccountsParams) WithTimeout(timeout time.Duration) *ListProjectServiceAccountsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list project service accounts params
func (o *ListProjectServiceAccountsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list project service accounts params
func (o *ListProjectServiceAccountsParams) WithContext(ctx context.Context) *ListProjectServiceAccountsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list project service accounts params
func (o *ListProjectServiceAccountsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list project service accounts params
func (o *ListProjectServiceAccountsParams) WithHTTPClient(client *http.Client) *ListProjectServiceAccountsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list project service accounts params
func (o *ListProjectServiceAccountsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProjectIdentifier adds the projectIdentifier to the list project service accounts params
func (o *ListProjectServiceAccountsParams) WithProjectIdentifier(projectIdentifier string) *ListProjectServiceAccountsParams {
	o.SetProjectIdentifier(projectIdentifier)
	return o
}

// SetProjectIdentifier adds the projectIdentifier to the list project service accounts params
func (o *ListProjectServiceAccountsParams) SetProjectIdentifier(projectIdentifier string) {
	o.ProjectIdentifier = projectIdentifier
}

// WriteToRequest writes these params to a swagger request
func (o *ListProjectServiceAccountsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
