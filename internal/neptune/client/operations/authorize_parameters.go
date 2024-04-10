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

// NewAuthorizeParams creates a new AuthorizeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAuthorizeParams() *AuthorizeParams {
	return &AuthorizeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAuthorizeParamsWithTimeout creates a new AuthorizeParams object
// with the ability to set a timeout on a request.
func NewAuthorizeParamsWithTimeout(timeout time.Duration) *AuthorizeParams {
	return &AuthorizeParams{
		timeout: timeout,
	}
}

// NewAuthorizeParamsWithContext creates a new AuthorizeParams object
// with the ability to set a context for a request.
func NewAuthorizeParamsWithContext(ctx context.Context) *AuthorizeParams {
	return &AuthorizeParams{
		Context: ctx,
	}
}

// NewAuthorizeParamsWithHTTPClient creates a new AuthorizeParams object
// with the ability to set a custom HTTPClient for a request.
func NewAuthorizeParamsWithHTTPClient(client *http.Client) *AuthorizeParams {
	return &AuthorizeParams{
		HTTPClient: client,
	}
}

/*
AuthorizeParams contains all the parameters to send to the API endpoint

	for the authorize operation.

	Typically these are written to a http.Request.
*/
type AuthorizeParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the authorize params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AuthorizeParams) WithDefaults() *AuthorizeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the authorize params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AuthorizeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the authorize params
func (o *AuthorizeParams) WithTimeout(timeout time.Duration) *AuthorizeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the authorize params
func (o *AuthorizeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the authorize params
func (o *AuthorizeParams) WithContext(ctx context.Context) *AuthorizeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the authorize params
func (o *AuthorizeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the authorize params
func (o *AuthorizeParams) WithHTTPClient(client *http.Client) *AuthorizeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the authorize params
func (o *AuthorizeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *AuthorizeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}