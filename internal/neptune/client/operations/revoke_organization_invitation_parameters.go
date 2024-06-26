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

// NewRevokeOrganizationInvitationParams creates a new RevokeOrganizationInvitationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRevokeOrganizationInvitationParams() *RevokeOrganizationInvitationParams {
	return &RevokeOrganizationInvitationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRevokeOrganizationInvitationParamsWithTimeout creates a new RevokeOrganizationInvitationParams object
// with the ability to set a timeout on a request.
func NewRevokeOrganizationInvitationParamsWithTimeout(timeout time.Duration) *RevokeOrganizationInvitationParams {
	return &RevokeOrganizationInvitationParams{
		timeout: timeout,
	}
}

// NewRevokeOrganizationInvitationParamsWithContext creates a new RevokeOrganizationInvitationParams object
// with the ability to set a context for a request.
func NewRevokeOrganizationInvitationParamsWithContext(ctx context.Context) *RevokeOrganizationInvitationParams {
	return &RevokeOrganizationInvitationParams{
		Context: ctx,
	}
}

// NewRevokeOrganizationInvitationParamsWithHTTPClient creates a new RevokeOrganizationInvitationParams object
// with the ability to set a custom HTTPClient for a request.
func NewRevokeOrganizationInvitationParamsWithHTTPClient(client *http.Client) *RevokeOrganizationInvitationParams {
	return &RevokeOrganizationInvitationParams{
		HTTPClient: client,
	}
}

/*
RevokeOrganizationInvitationParams contains all the parameters to send to the API endpoint

	for the revoke organization invitation operation.

	Typically these are written to a http.Request.
*/
type RevokeOrganizationInvitationParams struct {

	// InvitationID.
	//
	// Format: uuid
	InvitationID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the revoke organization invitation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RevokeOrganizationInvitationParams) WithDefaults() *RevokeOrganizationInvitationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the revoke organization invitation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RevokeOrganizationInvitationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the revoke organization invitation params
func (o *RevokeOrganizationInvitationParams) WithTimeout(timeout time.Duration) *RevokeOrganizationInvitationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the revoke organization invitation params
func (o *RevokeOrganizationInvitationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the revoke organization invitation params
func (o *RevokeOrganizationInvitationParams) WithContext(ctx context.Context) *RevokeOrganizationInvitationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the revoke organization invitation params
func (o *RevokeOrganizationInvitationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the revoke organization invitation params
func (o *RevokeOrganizationInvitationParams) WithHTTPClient(client *http.Client) *RevokeOrganizationInvitationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the revoke organization invitation params
func (o *RevokeOrganizationInvitationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInvitationID adds the invitationID to the revoke organization invitation params
func (o *RevokeOrganizationInvitationParams) WithInvitationID(invitationID strfmt.UUID) *RevokeOrganizationInvitationParams {
	o.SetInvitationID(invitationID)
	return o
}

// SetInvitationID adds the invitationId to the revoke organization invitation params
func (o *RevokeOrganizationInvitationParams) SetInvitationID(invitationID strfmt.UUID) {
	o.InvitationID = invitationID
}

// WriteToRequest writes these params to a swagger request
func (o *RevokeOrganizationInvitationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
