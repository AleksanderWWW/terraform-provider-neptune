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

// NewGlobalConfigurationParams creates a new GlobalConfigurationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGlobalConfigurationParams() *GlobalConfigurationParams {
	return &GlobalConfigurationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGlobalConfigurationParamsWithTimeout creates a new GlobalConfigurationParams object
// with the ability to set a timeout on a request.
func NewGlobalConfigurationParamsWithTimeout(timeout time.Duration) *GlobalConfigurationParams {
	return &GlobalConfigurationParams{
		timeout: timeout,
	}
}

// NewGlobalConfigurationParamsWithContext creates a new GlobalConfigurationParams object
// with the ability to set a context for a request.
func NewGlobalConfigurationParamsWithContext(ctx context.Context) *GlobalConfigurationParams {
	return &GlobalConfigurationParams{
		Context: ctx,
	}
}

// NewGlobalConfigurationParamsWithHTTPClient creates a new GlobalConfigurationParams object
// with the ability to set a custom HTTPClient for a request.
func NewGlobalConfigurationParamsWithHTTPClient(client *http.Client) *GlobalConfigurationParams {
	return &GlobalConfigurationParams{
		HTTPClient: client,
	}
}

/*
GlobalConfigurationParams contains all the parameters to send to the API endpoint

	for the global configuration operation.

	Typically these are written to a http.Request.
*/
type GlobalConfigurationParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the global configuration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GlobalConfigurationParams) WithDefaults() *GlobalConfigurationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the global configuration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GlobalConfigurationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the global configuration params
func (o *GlobalConfigurationParams) WithTimeout(timeout time.Duration) *GlobalConfigurationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the global configuration params
func (o *GlobalConfigurationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the global configuration params
func (o *GlobalConfigurationParams) WithContext(ctx context.Context) *GlobalConfigurationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the global configuration params
func (o *GlobalConfigurationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the global configuration params
func (o *GlobalConfigurationParams) WithHTTPClient(client *http.Client) *GlobalConfigurationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the global configuration params
func (o *GlobalConfigurationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GlobalConfigurationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
