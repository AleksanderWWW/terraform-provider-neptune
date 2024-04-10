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

// NewAcceptProjectInvitationParams creates a new AcceptProjectInvitationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAcceptProjectInvitationParams() *AcceptProjectInvitationParams {
	return &AcceptProjectInvitationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAcceptProjectInvitationParamsWithTimeout creates a new AcceptProjectInvitationParams object
// with the ability to set a timeout on a request.
func NewAcceptProjectInvitationParamsWithTimeout(timeout time.Duration) *AcceptProjectInvitationParams {
	return &AcceptProjectInvitationParams{
		timeout: timeout,
	}
}

// NewAcceptProjectInvitationParamsWithContext creates a new AcceptProjectInvitationParams object
// with the ability to set a context for a request.
func NewAcceptProjectInvitationParamsWithContext(ctx context.Context) *AcceptProjectInvitationParams {
	return &AcceptProjectInvitationParams{
		Context: ctx,
	}
}

// NewAcceptProjectInvitationParamsWithHTTPClient creates a new AcceptProjectInvitationParams object
// with the ability to set a custom HTTPClient for a request.
func NewAcceptProjectInvitationParamsWithHTTPClient(client *http.Client) *AcceptProjectInvitationParams {
	return &AcceptProjectInvitationParams{
		HTTPClient: client,
	}
}

/*
AcceptProjectInvitationParams contains all the parameters to send to the API endpoint

	for the accept project invitation operation.

	Typically these are written to a http.Request.
*/
type AcceptProjectInvitationParams struct {

	// InvitationID.
	//
	// Format: uuid
	InvitationID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the accept project invitation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AcceptProjectInvitationParams) WithDefaults() *AcceptProjectInvitationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the accept project invitation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AcceptProjectInvitationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the accept project invitation params
func (o *AcceptProjectInvitationParams) WithTimeout(timeout time.Duration) *AcceptProjectInvitationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the accept project invitation params
func (o *AcceptProjectInvitationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the accept project invitation params
func (o *AcceptProjectInvitationParams) WithContext(ctx context.Context) *AcceptProjectInvitationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the accept project invitation params
func (o *AcceptProjectInvitationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the accept project invitation params
func (o *AcceptProjectInvitationParams) WithHTTPClient(client *http.Client) *AcceptProjectInvitationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the accept project invitation params
func (o *AcceptProjectInvitationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInvitationID adds the invitationID to the accept project invitation params
func (o *AcceptProjectInvitationParams) WithInvitationID(invitationID strfmt.UUID) *AcceptProjectInvitationParams {
	o.SetInvitationID(invitationID)
	return o
}

// SetInvitationID adds the invitationId to the accept project invitation params
func (o *AcceptProjectInvitationParams) SetInvitationID(invitationID strfmt.UUID) {
	o.InvitationID = invitationID
}

// WriteToRequest writes these params to a swagger request
func (o *AcceptProjectInvitationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param invitationId
	if err := r.SetPathParam("invitationId", o.InvitationID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}