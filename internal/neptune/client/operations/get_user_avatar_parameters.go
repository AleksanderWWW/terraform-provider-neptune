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

// NewGetUserAvatarParams creates a new GetUserAvatarParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetUserAvatarParams() *GetUserAvatarParams {
	return &GetUserAvatarParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetUserAvatarParamsWithTimeout creates a new GetUserAvatarParams object
// with the ability to set a timeout on a request.
func NewGetUserAvatarParamsWithTimeout(timeout time.Duration) *GetUserAvatarParams {
	return &GetUserAvatarParams{
		timeout: timeout,
	}
}

// NewGetUserAvatarParamsWithContext creates a new GetUserAvatarParams object
// with the ability to set a context for a request.
func NewGetUserAvatarParamsWithContext(ctx context.Context) *GetUserAvatarParams {
	return &GetUserAvatarParams{
		Context: ctx,
	}
}

// NewGetUserAvatarParamsWithHTTPClient creates a new GetUserAvatarParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetUserAvatarParamsWithHTTPClient(client *http.Client) *GetUserAvatarParams {
	return &GetUserAvatarParams{
		HTTPClient: client,
	}
}

/*
GetUserAvatarParams contains all the parameters to send to the API endpoint

	for the get user avatar operation.

	Typically these are written to a http.Request.
*/
type GetUserAvatarParams struct {

	// Username.
	Username string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get user avatar params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUserAvatarParams) WithDefaults() *GetUserAvatarParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get user avatar params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUserAvatarParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get user avatar params
func (o *GetUserAvatarParams) WithTimeout(timeout time.Duration) *GetUserAvatarParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get user avatar params
func (o *GetUserAvatarParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get user avatar params
func (o *GetUserAvatarParams) WithContext(ctx context.Context) *GetUserAvatarParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get user avatar params
func (o *GetUserAvatarParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get user avatar params
func (o *GetUserAvatarParams) WithHTTPClient(client *http.Client) *GetUserAvatarParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get user avatar params
func (o *GetUserAvatarParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUsername adds the username to the get user avatar params
func (o *GetUserAvatarParams) WithUsername(username string) *GetUserAvatarParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the get user avatar params
func (o *GetUserAvatarParams) SetUsername(username string) {
	o.Username = username
}

// WriteToRequest writes these params to a swagger request
func (o *GetUserAvatarParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param username
	if err := r.SetPathParam("username", o.Username); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
