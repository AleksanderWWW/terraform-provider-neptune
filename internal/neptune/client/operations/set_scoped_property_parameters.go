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

// NewSetScopedPropertyParams creates a new SetScopedPropertyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSetScopedPropertyParams() *SetScopedPropertyParams {
	return &SetScopedPropertyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSetScopedPropertyParamsWithTimeout creates a new SetScopedPropertyParams object
// with the ability to set a timeout on a request.
func NewSetScopedPropertyParamsWithTimeout(timeout time.Duration) *SetScopedPropertyParams {
	return &SetScopedPropertyParams{
		timeout: timeout,
	}
}

// NewSetScopedPropertyParamsWithContext creates a new SetScopedPropertyParams object
// with the ability to set a context for a request.
func NewSetScopedPropertyParamsWithContext(ctx context.Context) *SetScopedPropertyParams {
	return &SetScopedPropertyParams{
		Context: ctx,
	}
}

// NewSetScopedPropertyParamsWithHTTPClient creates a new SetScopedPropertyParams object
// with the ability to set a custom HTTPClient for a request.
func NewSetScopedPropertyParamsWithHTTPClient(client *http.Client) *SetScopedPropertyParams {
	return &SetScopedPropertyParams{
		HTTPClient: client,
	}
}

/*
SetScopedPropertyParams contains all the parameters to send to the API endpoint

	for the set scoped property operation.

	Typically these are written to a http.Request.
*/
type SetScopedPropertyParams struct {

	// PropertyBody.
	PropertyBody *models.PropertyDTO

	// PropertyType.
	PropertyType string

	// Scope.
	Scope string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the set scoped property params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetScopedPropertyParams) WithDefaults() *SetScopedPropertyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the set scoped property params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetScopedPropertyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the set scoped property params
func (o *SetScopedPropertyParams) WithTimeout(timeout time.Duration) *SetScopedPropertyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set scoped property params
func (o *SetScopedPropertyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set scoped property params
func (o *SetScopedPropertyParams) WithContext(ctx context.Context) *SetScopedPropertyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set scoped property params
func (o *SetScopedPropertyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set scoped property params
func (o *SetScopedPropertyParams) WithHTTPClient(client *http.Client) *SetScopedPropertyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set scoped property params
func (o *SetScopedPropertyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPropertyBody adds the propertyBody to the set scoped property params
func (o *SetScopedPropertyParams) WithPropertyBody(propertyBody *models.PropertyDTO) *SetScopedPropertyParams {
	o.SetPropertyBody(propertyBody)
	return o
}

// SetPropertyBody adds the propertyBody to the set scoped property params
func (o *SetScopedPropertyParams) SetPropertyBody(propertyBody *models.PropertyDTO) {
	o.PropertyBody = propertyBody
}

// WithPropertyType adds the propertyType to the set scoped property params
func (o *SetScopedPropertyParams) WithPropertyType(propertyType string) *SetScopedPropertyParams {
	o.SetPropertyType(propertyType)
	return o
}

// SetPropertyType adds the propertyType to the set scoped property params
func (o *SetScopedPropertyParams) SetPropertyType(propertyType string) {
	o.PropertyType = propertyType
}

// WithScope adds the scope to the set scoped property params
func (o *SetScopedPropertyParams) WithScope(scope string) *SetScopedPropertyParams {
	o.SetScope(scope)
	return o
}

// SetScope adds the scope to the set scoped property params
func (o *SetScopedPropertyParams) SetScope(scope string) {
	o.Scope = scope
}

// WriteToRequest writes these params to a swagger request
func (o *SetScopedPropertyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.PropertyBody != nil {
		if err := r.SetBodyParam(o.PropertyBody); err != nil {
			return err
		}
	}

	// query param propertyType
	qrPropertyType := o.PropertyType
	qPropertyType := qrPropertyType
	if qPropertyType != "" {

		if err := r.SetQueryParam("propertyType", qPropertyType); err != nil {
			return err
		}
	}

	// query param scope
	qrScope := o.Scope
	qScope := qrScope
	if qScope != "" {

		if err := r.SetQueryParam("scope", qScope); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
