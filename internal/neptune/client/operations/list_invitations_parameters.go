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

// NewListInvitationsParams creates a new ListInvitationsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListInvitationsParams() *ListInvitationsParams {
	return &ListInvitationsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListInvitationsParamsWithTimeout creates a new ListInvitationsParams object
// with the ability to set a timeout on a request.
func NewListInvitationsParamsWithTimeout(timeout time.Duration) *ListInvitationsParams {
	return &ListInvitationsParams{
		timeout: timeout,
	}
}

// NewListInvitationsParamsWithContext creates a new ListInvitationsParams object
// with the ability to set a context for a request.
func NewListInvitationsParamsWithContext(ctx context.Context) *ListInvitationsParams {
	return &ListInvitationsParams{
		Context: ctx,
	}
}

// NewListInvitationsParamsWithHTTPClient creates a new ListInvitationsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListInvitationsParamsWithHTTPClient(client *http.Client) *ListInvitationsParams {
	return &ListInvitationsParams{
		HTTPClient: client,
	}
}

/*
ListInvitationsParams contains all the parameters to send to the API endpoint

	for the list invitations operation.

	Typically these are written to a http.Request.
*/
type ListInvitationsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list invitations params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListInvitationsParams) WithDefaults() *ListInvitationsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list invitations params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListInvitationsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list invitations params
func (o *ListInvitationsParams) WithTimeout(timeout time.Duration) *ListInvitationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list invitations params
func (o *ListInvitationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list invitations params
func (o *ListInvitationsParams) WithContext(ctx context.Context) *ListInvitationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list invitations params
func (o *ListInvitationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list invitations params
func (o *ListInvitationsParams) WithHTTPClient(client *http.Client) *ListInvitationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list invitations params
func (o *ListInvitationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListInvitationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
