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

// NewDeleteUserProfileLinkParams creates a new DeleteUserProfileLinkParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteUserProfileLinkParams() *DeleteUserProfileLinkParams {
	return &DeleteUserProfileLinkParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteUserProfileLinkParamsWithTimeout creates a new DeleteUserProfileLinkParams object
// with the ability to set a timeout on a request.
func NewDeleteUserProfileLinkParamsWithTimeout(timeout time.Duration) *DeleteUserProfileLinkParams {
	return &DeleteUserProfileLinkParams{
		timeout: timeout,
	}
}

// NewDeleteUserProfileLinkParamsWithContext creates a new DeleteUserProfileLinkParams object
// with the ability to set a context for a request.
func NewDeleteUserProfileLinkParamsWithContext(ctx context.Context) *DeleteUserProfileLinkParams {
	return &DeleteUserProfileLinkParams{
		Context: ctx,
	}
}

// NewDeleteUserProfileLinkParamsWithHTTPClient creates a new DeleteUserProfileLinkParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteUserProfileLinkParamsWithHTTPClient(client *http.Client) *DeleteUserProfileLinkParams {
	return &DeleteUserProfileLinkParams{
		HTTPClient: client,
	}
}

/*
DeleteUserProfileLinkParams contains all the parameters to send to the API endpoint

	for the delete user profile link operation.

	Typically these are written to a http.Request.
*/
type DeleteUserProfileLinkParams struct {

	// Link.
	Link *models.LinkDTO

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete user profile link params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteUserProfileLinkParams) WithDefaults() *DeleteUserProfileLinkParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete user profile link params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteUserProfileLinkParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete user profile link params
func (o *DeleteUserProfileLinkParams) WithTimeout(timeout time.Duration) *DeleteUserProfileLinkParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete user profile link params
func (o *DeleteUserProfileLinkParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete user profile link params
func (o *DeleteUserProfileLinkParams) WithContext(ctx context.Context) *DeleteUserProfileLinkParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete user profile link params
func (o *DeleteUserProfileLinkParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete user profile link params
func (o *DeleteUserProfileLinkParams) WithHTTPClient(client *http.Client) *DeleteUserProfileLinkParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete user profile link params
func (o *DeleteUserProfileLinkParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLink adds the link to the delete user profile link params
func (o *DeleteUserProfileLinkParams) WithLink(link *models.LinkDTO) *DeleteUserProfileLinkParams {
	o.SetLink(link)
	return o
}

// SetLink adds the link to the delete user profile link params
func (o *DeleteUserProfileLinkParams) SetLink(link *models.LinkDTO) {
	o.Link = link
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteUserProfileLinkParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Link != nil {
		if err := r.SetBodyParam(o.Link); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
