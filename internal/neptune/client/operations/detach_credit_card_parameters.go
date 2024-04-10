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

// NewDetachCreditCardParams creates a new DetachCreditCardParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDetachCreditCardParams() *DetachCreditCardParams {
	return &DetachCreditCardParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDetachCreditCardParamsWithTimeout creates a new DetachCreditCardParams object
// with the ability to set a timeout on a request.
func NewDetachCreditCardParamsWithTimeout(timeout time.Duration) *DetachCreditCardParams {
	return &DetachCreditCardParams{
		timeout: timeout,
	}
}

// NewDetachCreditCardParamsWithContext creates a new DetachCreditCardParams object
// with the ability to set a context for a request.
func NewDetachCreditCardParamsWithContext(ctx context.Context) *DetachCreditCardParams {
	return &DetachCreditCardParams{
		Context: ctx,
	}
}

// NewDetachCreditCardParamsWithHTTPClient creates a new DetachCreditCardParams object
// with the ability to set a custom HTTPClient for a request.
func NewDetachCreditCardParamsWithHTTPClient(client *http.Client) *DetachCreditCardParams {
	return &DetachCreditCardParams{
		HTTPClient: client,
	}
}

/*
DetachCreditCardParams contains all the parameters to send to the API endpoint

	for the detach credit card operation.

	Typically these are written to a http.Request.
*/
type DetachCreditCardParams struct {

	// OrganizationIdentifier.
	OrganizationIdentifier string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the detach credit card params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DetachCreditCardParams) WithDefaults() *DetachCreditCardParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the detach credit card params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DetachCreditCardParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the detach credit card params
func (o *DetachCreditCardParams) WithTimeout(timeout time.Duration) *DetachCreditCardParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the detach credit card params
func (o *DetachCreditCardParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the detach credit card params
func (o *DetachCreditCardParams) WithContext(ctx context.Context) *DetachCreditCardParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the detach credit card params
func (o *DetachCreditCardParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the detach credit card params
func (o *DetachCreditCardParams) WithHTTPClient(client *http.Client) *DetachCreditCardParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the detach credit card params
func (o *DetachCreditCardParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganizationIdentifier adds the organizationIdentifier to the detach credit card params
func (o *DetachCreditCardParams) WithOrganizationIdentifier(organizationIdentifier string) *DetachCreditCardParams {
	o.SetOrganizationIdentifier(organizationIdentifier)
	return o
}

// SetOrganizationIdentifier adds the organizationIdentifier to the detach credit card params
func (o *DetachCreditCardParams) SetOrganizationIdentifier(organizationIdentifier string) {
	o.OrganizationIdentifier = organizationIdentifier
}

// WriteToRequest writes these params to a swagger request
func (o *DetachCreditCardParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param organizationIdentifier
	if err := r.SetPathParam("organizationIdentifier", o.OrganizationIdentifier); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}