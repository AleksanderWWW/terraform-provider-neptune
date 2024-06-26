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

// NewIsOrganizationNameAvailableParams creates a new IsOrganizationNameAvailableParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIsOrganizationNameAvailableParams() *IsOrganizationNameAvailableParams {
	return &IsOrganizationNameAvailableParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIsOrganizationNameAvailableParamsWithTimeout creates a new IsOrganizationNameAvailableParams object
// with the ability to set a timeout on a request.
func NewIsOrganizationNameAvailableParamsWithTimeout(timeout time.Duration) *IsOrganizationNameAvailableParams {
	return &IsOrganizationNameAvailableParams{
		timeout: timeout,
	}
}

// NewIsOrganizationNameAvailableParamsWithContext creates a new IsOrganizationNameAvailableParams object
// with the ability to set a context for a request.
func NewIsOrganizationNameAvailableParamsWithContext(ctx context.Context) *IsOrganizationNameAvailableParams {
	return &IsOrganizationNameAvailableParams{
		Context: ctx,
	}
}

// NewIsOrganizationNameAvailableParamsWithHTTPClient creates a new IsOrganizationNameAvailableParams object
// with the ability to set a custom HTTPClient for a request.
func NewIsOrganizationNameAvailableParamsWithHTTPClient(client *http.Client) *IsOrganizationNameAvailableParams {
	return &IsOrganizationNameAvailableParams{
		HTTPClient: client,
	}
}

/*
IsOrganizationNameAvailableParams contains all the parameters to send to the API endpoint

	for the is organization name available operation.

	Typically these are written to a http.Request.
*/
type IsOrganizationNameAvailableParams struct {

	// OrganizationName.
	OrganizationName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the is organization name available params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IsOrganizationNameAvailableParams) WithDefaults() *IsOrganizationNameAvailableParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the is organization name available params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IsOrganizationNameAvailableParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the is organization name available params
func (o *IsOrganizationNameAvailableParams) WithTimeout(timeout time.Duration) *IsOrganizationNameAvailableParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the is organization name available params
func (o *IsOrganizationNameAvailableParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the is organization name available params
func (o *IsOrganizationNameAvailableParams) WithContext(ctx context.Context) *IsOrganizationNameAvailableParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the is organization name available params
func (o *IsOrganizationNameAvailableParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the is organization name available params
func (o *IsOrganizationNameAvailableParams) WithHTTPClient(client *http.Client) *IsOrganizationNameAvailableParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the is organization name available params
func (o *IsOrganizationNameAvailableParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganizationName adds the organizationName to the is organization name available params
func (o *IsOrganizationNameAvailableParams) WithOrganizationName(organizationName string) *IsOrganizationNameAvailableParams {
	o.SetOrganizationName(organizationName)
	return o
}

// SetOrganizationName adds the organizationName to the is organization name available params
func (o *IsOrganizationNameAvailableParams) SetOrganizationName(organizationName string) {
	o.OrganizationName = organizationName
}

// WriteToRequest writes these params to a swagger request
func (o *IsOrganizationNameAvailableParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param organizationName
	if err := r.SetPathParam("organizationName", o.OrganizationName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
