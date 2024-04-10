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

	"terraform-provider-neptune/internal/neptune/models"
)

// NewAddOrganizationMemberParams creates a new AddOrganizationMemberParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddOrganizationMemberParams() *AddOrganizationMemberParams {
	return &AddOrganizationMemberParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddOrganizationMemberParamsWithTimeout creates a new AddOrganizationMemberParams object
// with the ability to set a timeout on a request.
func NewAddOrganizationMemberParamsWithTimeout(timeout time.Duration) *AddOrganizationMemberParams {
	return &AddOrganizationMemberParams{
		timeout: timeout,
	}
}

// NewAddOrganizationMemberParamsWithContext creates a new AddOrganizationMemberParams object
// with the ability to set a context for a request.
func NewAddOrganizationMemberParamsWithContext(ctx context.Context) *AddOrganizationMemberParams {
	return &AddOrganizationMemberParams{
		Context: ctx,
	}
}

// NewAddOrganizationMemberParamsWithHTTPClient creates a new AddOrganizationMemberParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddOrganizationMemberParamsWithHTTPClient(client *http.Client) *AddOrganizationMemberParams {
	return &AddOrganizationMemberParams{
		HTTPClient: client,
	}
}

/*
AddOrganizationMemberParams contains all the parameters to send to the API endpoint

	for the add organization member operation.

	Typically these are written to a http.Request.
*/
type AddOrganizationMemberParams struct {

	// Member.
	Member *models.NewOrganizationMemberDTO

	// OrganizationIdentifier.
	OrganizationIdentifier string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add organization member params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddOrganizationMemberParams) WithDefaults() *AddOrganizationMemberParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add organization member params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddOrganizationMemberParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add organization member params
func (o *AddOrganizationMemberParams) WithTimeout(timeout time.Duration) *AddOrganizationMemberParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add organization member params
func (o *AddOrganizationMemberParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add organization member params
func (o *AddOrganizationMemberParams) WithContext(ctx context.Context) *AddOrganizationMemberParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add organization member params
func (o *AddOrganizationMemberParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add organization member params
func (o *AddOrganizationMemberParams) WithHTTPClient(client *http.Client) *AddOrganizationMemberParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add organization member params
func (o *AddOrganizationMemberParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMember adds the member to the add organization member params
func (o *AddOrganizationMemberParams) WithMember(member *models.NewOrganizationMemberDTO) *AddOrganizationMemberParams {
	o.SetMember(member)
	return o
}

// SetMember adds the member to the add organization member params
func (o *AddOrganizationMemberParams) SetMember(member *models.NewOrganizationMemberDTO) {
	o.Member = member
}

// WithOrganizationIdentifier adds the organizationIdentifier to the add organization member params
func (o *AddOrganizationMemberParams) WithOrganizationIdentifier(organizationIdentifier string) *AddOrganizationMemberParams {
	o.SetOrganizationIdentifier(organizationIdentifier)
	return o
}

// SetOrganizationIdentifier adds the organizationIdentifier to the add organization member params
func (o *AddOrganizationMemberParams) SetOrganizationIdentifier(organizationIdentifier string) {
	o.OrganizationIdentifier = organizationIdentifier
}

// WriteToRequest writes these params to a swagger request
func (o *AddOrganizationMemberParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Member != nil {
		if err := r.SetBodyParam(o.Member); err != nil {
			return err
		}
	}

	// path param organizationIdentifier
	if err := r.SetPathParam("organizationIdentifier", o.OrganizationIdentifier); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
