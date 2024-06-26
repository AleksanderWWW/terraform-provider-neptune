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

// NewRevokeAPITokenParams creates a new RevokeAPITokenParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRevokeAPITokenParams() *RevokeAPITokenParams {
	return &RevokeAPITokenParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRevokeAPITokenParamsWithTimeout creates a new RevokeAPITokenParams object
// with the ability to set a timeout on a request.
func NewRevokeAPITokenParamsWithTimeout(timeout time.Duration) *RevokeAPITokenParams {
	return &RevokeAPITokenParams{
		timeout: timeout,
	}
}

// NewRevokeAPITokenParamsWithContext creates a new RevokeAPITokenParams object
// with the ability to set a context for a request.
func NewRevokeAPITokenParamsWithContext(ctx context.Context) *RevokeAPITokenParams {
	return &RevokeAPITokenParams{
		Context: ctx,
	}
}

// NewRevokeAPITokenParamsWithHTTPClient creates a new RevokeAPITokenParams object
// with the ability to set a custom HTTPClient for a request.
func NewRevokeAPITokenParamsWithHTTPClient(client *http.Client) *RevokeAPITokenParams {
	return &RevokeAPITokenParams{
		HTTPClient: client,
	}
}

/*
RevokeAPITokenParams contains all the parameters to send to the API endpoint

	for the revoke Api token operation.

	Typically these are written to a http.Request.
*/
type RevokeAPITokenParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the revoke Api token params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RevokeAPITokenParams) WithDefaults() *RevokeAPITokenParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the revoke Api token params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RevokeAPITokenParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the revoke Api token params
func (o *RevokeAPITokenParams) WithTimeout(timeout time.Duration) *RevokeAPITokenParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the revoke Api token params
func (o *RevokeAPITokenParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the revoke Api token params
func (o *RevokeAPITokenParams) WithContext(ctx context.Context) *RevokeAPITokenParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the revoke Api token params
func (o *RevokeAPITokenParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the revoke Api token params
func (o *RevokeAPITokenParams) WithHTTPClient(client *http.Client) *RevokeAPITokenParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the revoke Api token params
func (o *RevokeAPITokenParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *RevokeAPITokenParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
