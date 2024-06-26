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
	"github.com/go-openapi/swag"
)

// NewListUsersParams creates a new ListUsersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListUsersParams() *ListUsersParams {
	return &ListUsersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListUsersParamsWithTimeout creates a new ListUsersParams object
// with the ability to set a timeout on a request.
func NewListUsersParamsWithTimeout(timeout time.Duration) *ListUsersParams {
	return &ListUsersParams{
		timeout: timeout,
	}
}

// NewListUsersParamsWithContext creates a new ListUsersParams object
// with the ability to set a context for a request.
func NewListUsersParamsWithContext(ctx context.Context) *ListUsersParams {
	return &ListUsersParams{
		Context: ctx,
	}
}

// NewListUsersParamsWithHTTPClient creates a new ListUsersParams object
// with the ability to set a custom HTTPClient for a request.
func NewListUsersParamsWithHTTPClient(client *http.Client) *ListUsersParams {
	return &ListUsersParams{
		HTTPClient: client,
	}
}

/*
ListUsersParams contains all the parameters to send to the API endpoint

	for the list users operation.

	Typically these are written to a http.Request.
*/
type ListUsersParams struct {

	// Username.
	Username []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list users params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListUsersParams) WithDefaults() *ListUsersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list users params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListUsersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list users params
func (o *ListUsersParams) WithTimeout(timeout time.Duration) *ListUsersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list users params
func (o *ListUsersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list users params
func (o *ListUsersParams) WithContext(ctx context.Context) *ListUsersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list users params
func (o *ListUsersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list users params
func (o *ListUsersParams) WithHTTPClient(client *http.Client) *ListUsersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list users params
func (o *ListUsersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUsername adds the username to the list users params
func (o *ListUsersParams) WithUsername(username []string) *ListUsersParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the list users params
func (o *ListUsersParams) SetUsername(username []string) {
	o.Username = username
}

// WriteToRequest writes these params to a swagger request
func (o *ListUsersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Username != nil {

		// binding items for username
		joinedUsername := o.bindParamUsername(reg)

		// query array param username
		if err := r.SetQueryParam("username", joinedUsername...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamListUsers binds the parameter username
func (o *ListUsersParams) bindParamUsername(formats strfmt.Registry) []string {
	usernameIR := o.Username

	var usernameIC []string
	for _, usernameIIR := range usernameIR { // explode []string

		usernameIIV := usernameIIR // string as string
		usernameIC = append(usernameIC, usernameIIV)
	}

	// items.CollectionFormat: "multi"
	usernameIS := swag.JoinByFormat(usernameIC, "multi")

	return usernameIS
}
