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

// NewUpdateActiveProjectsQuotaParams creates a new UpdateActiveProjectsQuotaParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateActiveProjectsQuotaParams() *UpdateActiveProjectsQuotaParams {
	return &UpdateActiveProjectsQuotaParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateActiveProjectsQuotaParamsWithTimeout creates a new UpdateActiveProjectsQuotaParams object
// with the ability to set a timeout on a request.
func NewUpdateActiveProjectsQuotaParamsWithTimeout(timeout time.Duration) *UpdateActiveProjectsQuotaParams {
	return &UpdateActiveProjectsQuotaParams{
		timeout: timeout,
	}
}

// NewUpdateActiveProjectsQuotaParamsWithContext creates a new UpdateActiveProjectsQuotaParams object
// with the ability to set a context for a request.
func NewUpdateActiveProjectsQuotaParamsWithContext(ctx context.Context) *UpdateActiveProjectsQuotaParams {
	return &UpdateActiveProjectsQuotaParams{
		Context: ctx,
	}
}

// NewUpdateActiveProjectsQuotaParamsWithHTTPClient creates a new UpdateActiveProjectsQuotaParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateActiveProjectsQuotaParamsWithHTTPClient(client *http.Client) *UpdateActiveProjectsQuotaParams {
	return &UpdateActiveProjectsQuotaParams{
		HTTPClient: client,
	}
}

/*
UpdateActiveProjectsQuotaParams contains all the parameters to send to the API endpoint

	for the update active projects quota operation.

	Typically these are written to a http.Request.
*/
type UpdateActiveProjectsQuotaParams struct {

	// OrganizationIdentifier.
	OrganizationIdentifier string

	// Quota.
	//
	// Format: int32
	Quota int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update active projects quota params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateActiveProjectsQuotaParams) WithDefaults() *UpdateActiveProjectsQuotaParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update active projects quota params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateActiveProjectsQuotaParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update active projects quota params
func (o *UpdateActiveProjectsQuotaParams) WithTimeout(timeout time.Duration) *UpdateActiveProjectsQuotaParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update active projects quota params
func (o *UpdateActiveProjectsQuotaParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update active projects quota params
func (o *UpdateActiveProjectsQuotaParams) WithContext(ctx context.Context) *UpdateActiveProjectsQuotaParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update active projects quota params
func (o *UpdateActiveProjectsQuotaParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update active projects quota params
func (o *UpdateActiveProjectsQuotaParams) WithHTTPClient(client *http.Client) *UpdateActiveProjectsQuotaParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update active projects quota params
func (o *UpdateActiveProjectsQuotaParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganizationIdentifier adds the organizationIdentifier to the update active projects quota params
func (o *UpdateActiveProjectsQuotaParams) WithOrganizationIdentifier(organizationIdentifier string) *UpdateActiveProjectsQuotaParams {
	o.SetOrganizationIdentifier(organizationIdentifier)
	return o
}

// SetOrganizationIdentifier adds the organizationIdentifier to the update active projects quota params
func (o *UpdateActiveProjectsQuotaParams) SetOrganizationIdentifier(organizationIdentifier string) {
	o.OrganizationIdentifier = organizationIdentifier
}

// WithQuota adds the quota to the update active projects quota params
func (o *UpdateActiveProjectsQuotaParams) WithQuota(quota int32) *UpdateActiveProjectsQuotaParams {
	o.SetQuota(quota)
	return o
}

// SetQuota adds the quota to the update active projects quota params
func (o *UpdateActiveProjectsQuotaParams) SetQuota(quota int32) {
	o.Quota = quota
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateActiveProjectsQuotaParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param organizationIdentifier
	if err := r.SetPathParam("organizationIdentifier", o.OrganizationIdentifier); err != nil {
		return err
	}

	// query param quota
	qrQuota := o.Quota
	qQuota := swag.FormatInt32(qrQuota)
	if qQuota != "" {

		if err := r.SetQueryParam("quota", qQuota); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}