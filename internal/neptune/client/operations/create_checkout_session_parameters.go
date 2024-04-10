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

	"client/models"
)

// NewCreateCheckoutSessionParams creates a new CreateCheckoutSessionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateCheckoutSessionParams() *CreateCheckoutSessionParams {
	return &CreateCheckoutSessionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateCheckoutSessionParamsWithTimeout creates a new CreateCheckoutSessionParams object
// with the ability to set a timeout on a request.
func NewCreateCheckoutSessionParamsWithTimeout(timeout time.Duration) *CreateCheckoutSessionParams {
	return &CreateCheckoutSessionParams{
		timeout: timeout,
	}
}

// NewCreateCheckoutSessionParamsWithContext creates a new CreateCheckoutSessionParams object
// with the ability to set a context for a request.
func NewCreateCheckoutSessionParamsWithContext(ctx context.Context) *CreateCheckoutSessionParams {
	return &CreateCheckoutSessionParams{
		Context: ctx,
	}
}

// NewCreateCheckoutSessionParamsWithHTTPClient creates a new CreateCheckoutSessionParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateCheckoutSessionParamsWithHTTPClient(client *http.Client) *CreateCheckoutSessionParams {
	return &CreateCheckoutSessionParams{
		HTTPClient: client,
	}
}

/*
CreateCheckoutSessionParams contains all the parameters to send to the API endpoint

	for the create checkout session operation.

	Typically these are written to a http.Request.
*/
type CreateCheckoutSessionParams struct {

	// OrganizationIdentifier.
	OrganizationIdentifier string

	// Purchase.
	Purchase *models.UpgradePurchaseRequestDTO

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create checkout session params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateCheckoutSessionParams) WithDefaults() *CreateCheckoutSessionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create checkout session params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateCheckoutSessionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create checkout session params
func (o *CreateCheckoutSessionParams) WithTimeout(timeout time.Duration) *CreateCheckoutSessionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create checkout session params
func (o *CreateCheckoutSessionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create checkout session params
func (o *CreateCheckoutSessionParams) WithContext(ctx context.Context) *CreateCheckoutSessionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create checkout session params
func (o *CreateCheckoutSessionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create checkout session params
func (o *CreateCheckoutSessionParams) WithHTTPClient(client *http.Client) *CreateCheckoutSessionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create checkout session params
func (o *CreateCheckoutSessionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganizationIdentifier adds the organizationIdentifier to the create checkout session params
func (o *CreateCheckoutSessionParams) WithOrganizationIdentifier(organizationIdentifier string) *CreateCheckoutSessionParams {
	o.SetOrganizationIdentifier(organizationIdentifier)
	return o
}

// SetOrganizationIdentifier adds the organizationIdentifier to the create checkout session params
func (o *CreateCheckoutSessionParams) SetOrganizationIdentifier(organizationIdentifier string) {
	o.OrganizationIdentifier = organizationIdentifier
}

// WithPurchase adds the purchase to the create checkout session params
func (o *CreateCheckoutSessionParams) WithPurchase(purchase *models.UpgradePurchaseRequestDTO) *CreateCheckoutSessionParams {
	o.SetPurchase(purchase)
	return o
}

// SetPurchase adds the purchase to the create checkout session params
func (o *CreateCheckoutSessionParams) SetPurchase(purchase *models.UpgradePurchaseRequestDTO) {
	o.Purchase = purchase
}

// WriteToRequest writes these params to a swagger request
func (o *CreateCheckoutSessionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param organizationIdentifier
	qrOrganizationIdentifier := o.OrganizationIdentifier
	qOrganizationIdentifier := qrOrganizationIdentifier
	if qOrganizationIdentifier != "" {

		if err := r.SetQueryParam("organizationIdentifier", qOrganizationIdentifier); err != nil {
			return err
		}
	}
	if o.Purchase != nil {
		if err := r.SetBodyParam(o.Purchase); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
