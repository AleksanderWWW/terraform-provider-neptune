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

// NewSetStorageLimitParams creates a new SetStorageLimitParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSetStorageLimitParams() *SetStorageLimitParams {
	return &SetStorageLimitParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSetStorageLimitParamsWithTimeout creates a new SetStorageLimitParams object
// with the ability to set a timeout on a request.
func NewSetStorageLimitParamsWithTimeout(timeout time.Duration) *SetStorageLimitParams {
	return &SetStorageLimitParams{
		timeout: timeout,
	}
}

// NewSetStorageLimitParamsWithContext creates a new SetStorageLimitParams object
// with the ability to set a context for a request.
func NewSetStorageLimitParamsWithContext(ctx context.Context) *SetStorageLimitParams {
	return &SetStorageLimitParams{
		Context: ctx,
	}
}

// NewSetStorageLimitParamsWithHTTPClient creates a new SetStorageLimitParams object
// with the ability to set a custom HTTPClient for a request.
func NewSetStorageLimitParamsWithHTTPClient(client *http.Client) *SetStorageLimitParams {
	return &SetStorageLimitParams{
		HTTPClient: client,
	}
}

/*
SetStorageLimitParams contains all the parameters to send to the API endpoint

	for the set storage limit operation.

	Typically these are written to a http.Request.
*/
type SetStorageLimitParams struct {

	// ProjectIdentifier.
	ProjectIdentifier string

	// StorageLimit.
	StorageLimit *models.SetProjectStorageLimitDTO

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the set storage limit params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetStorageLimitParams) WithDefaults() *SetStorageLimitParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the set storage limit params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetStorageLimitParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the set storage limit params
func (o *SetStorageLimitParams) WithTimeout(timeout time.Duration) *SetStorageLimitParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set storage limit params
func (o *SetStorageLimitParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set storage limit params
func (o *SetStorageLimitParams) WithContext(ctx context.Context) *SetStorageLimitParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set storage limit params
func (o *SetStorageLimitParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set storage limit params
func (o *SetStorageLimitParams) WithHTTPClient(client *http.Client) *SetStorageLimitParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set storage limit params
func (o *SetStorageLimitParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProjectIdentifier adds the projectIdentifier to the set storage limit params
func (o *SetStorageLimitParams) WithProjectIdentifier(projectIdentifier string) *SetStorageLimitParams {
	o.SetProjectIdentifier(projectIdentifier)
	return o
}

// SetProjectIdentifier adds the projectIdentifier to the set storage limit params
func (o *SetStorageLimitParams) SetProjectIdentifier(projectIdentifier string) {
	o.ProjectIdentifier = projectIdentifier
}

// WithStorageLimit adds the storageLimit to the set storage limit params
func (o *SetStorageLimitParams) WithStorageLimit(storageLimit *models.SetProjectStorageLimitDTO) *SetStorageLimitParams {
	o.SetStorageLimit(storageLimit)
	return o
}

// SetStorageLimit adds the storageLimit to the set storage limit params
func (o *SetStorageLimitParams) SetStorageLimit(storageLimit *models.SetProjectStorageLimitDTO) {
	o.StorageLimit = storageLimit
}

// WriteToRequest writes these params to a swagger request
func (o *SetStorageLimitParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param projectIdentifier
	qrProjectIdentifier := o.ProjectIdentifier
	qProjectIdentifier := qrProjectIdentifier
	if qProjectIdentifier != "" {

		if err := r.SetQueryParam("projectIdentifier", qProjectIdentifier); err != nil {
			return err
		}
	}
	if o.StorageLimit != nil {
		if err := r.SetBodyParam(o.StorageLimit); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
