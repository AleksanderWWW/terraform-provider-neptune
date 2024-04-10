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

// NewListPublicProjectsParams creates a new ListPublicProjectsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListPublicProjectsParams() *ListPublicProjectsParams {
	return &ListPublicProjectsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListPublicProjectsParamsWithTimeout creates a new ListPublicProjectsParams object
// with the ability to set a timeout on a request.
func NewListPublicProjectsParamsWithTimeout(timeout time.Duration) *ListPublicProjectsParams {
	return &ListPublicProjectsParams{
		timeout: timeout,
	}
}

// NewListPublicProjectsParamsWithContext creates a new ListPublicProjectsParams object
// with the ability to set a context for a request.
func NewListPublicProjectsParamsWithContext(ctx context.Context) *ListPublicProjectsParams {
	return &ListPublicProjectsParams{
		Context: ctx,
	}
}

// NewListPublicProjectsParamsWithHTTPClient creates a new ListPublicProjectsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListPublicProjectsParamsWithHTTPClient(client *http.Client) *ListPublicProjectsParams {
	return &ListPublicProjectsParams{
		HTTPClient: client,
	}
}

/*
ListPublicProjectsParams contains all the parameters to send to the API endpoint

	for the list public projects operation.

	Typically these are written to a http.Request.
*/
type ListPublicProjectsParams struct {

	// Limit.
	//
	// Format: int32
	Limit *int32

	// Offset.
	//
	// Format: int32
	Offset *int32

	// SearchTerm.
	SearchTerm *string

	// SortBy.
	SortBy []string

	// SortDirection.
	SortDirection []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list public projects params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListPublicProjectsParams) WithDefaults() *ListPublicProjectsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list public projects params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListPublicProjectsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list public projects params
func (o *ListPublicProjectsParams) WithTimeout(timeout time.Duration) *ListPublicProjectsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list public projects params
func (o *ListPublicProjectsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list public projects params
func (o *ListPublicProjectsParams) WithContext(ctx context.Context) *ListPublicProjectsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list public projects params
func (o *ListPublicProjectsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list public projects params
func (o *ListPublicProjectsParams) WithHTTPClient(client *http.Client) *ListPublicProjectsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list public projects params
func (o *ListPublicProjectsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the list public projects params
func (o *ListPublicProjectsParams) WithLimit(limit *int32) *ListPublicProjectsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the list public projects params
func (o *ListPublicProjectsParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithOffset adds the offset to the list public projects params
func (o *ListPublicProjectsParams) WithOffset(offset *int32) *ListPublicProjectsParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the list public projects params
func (o *ListPublicProjectsParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithSearchTerm adds the searchTerm to the list public projects params
func (o *ListPublicProjectsParams) WithSearchTerm(searchTerm *string) *ListPublicProjectsParams {
	o.SetSearchTerm(searchTerm)
	return o
}

// SetSearchTerm adds the searchTerm to the list public projects params
func (o *ListPublicProjectsParams) SetSearchTerm(searchTerm *string) {
	o.SearchTerm = searchTerm
}

// WithSortBy adds the sortBy to the list public projects params
func (o *ListPublicProjectsParams) WithSortBy(sortBy []string) *ListPublicProjectsParams {
	o.SetSortBy(sortBy)
	return o
}

// SetSortBy adds the sortBy to the list public projects params
func (o *ListPublicProjectsParams) SetSortBy(sortBy []string) {
	o.SortBy = sortBy
}

// WithSortDirection adds the sortDirection to the list public projects params
func (o *ListPublicProjectsParams) WithSortDirection(sortDirection []string) *ListPublicProjectsParams {
	o.SetSortDirection(sortDirection)
	return o
}

// SetSortDirection adds the sortDirection to the list public projects params
func (o *ListPublicProjectsParams) SetSortDirection(sortDirection []string) {
	o.SortDirection = sortDirection
}

// WriteToRequest writes these params to a swagger request
func (o *ListPublicProjectsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Limit != nil {

		// query param limit
		var qrLimit int32

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt32(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int32

		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt32(qrOffset)
		if qOffset != "" {

			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}
	}

	if o.SearchTerm != nil {

		// query param searchTerm
		var qrSearchTerm string

		if o.SearchTerm != nil {
			qrSearchTerm = *o.SearchTerm
		}
		qSearchTerm := qrSearchTerm
		if qSearchTerm != "" {

			if err := r.SetQueryParam("searchTerm", qSearchTerm); err != nil {
				return err
			}
		}
	}

	if o.SortBy != nil {

		// binding items for sortBy
		joinedSortBy := o.bindParamSortBy(reg)

		// query array param sortBy
		if err := r.SetQueryParam("sortBy", joinedSortBy...); err != nil {
			return err
		}
	}

	if o.SortDirection != nil {

		// binding items for sortDirection
		joinedSortDirection := o.bindParamSortDirection(reg)

		// query array param sortDirection
		if err := r.SetQueryParam("sortDirection", joinedSortDirection...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamListPublicProjects binds the parameter sortBy
func (o *ListPublicProjectsParams) bindParamSortBy(formats strfmt.Registry) []string {
	sortByIR := o.SortBy

	var sortByIC []string
	for _, sortByIIR := range sortByIR { // explode []string

		sortByIIV := sortByIIR // string as string
		sortByIC = append(sortByIC, sortByIIV)
	}

	// items.CollectionFormat: "multi"
	sortByIS := swag.JoinByFormat(sortByIC, "multi")

	return sortByIS
}

// bindParamListPublicProjects binds the parameter sortDirection
func (o *ListPublicProjectsParams) bindParamSortDirection(formats strfmt.Registry) []string {
	sortDirectionIR := o.SortDirection

	var sortDirectionIC []string
	for _, sortDirectionIIR := range sortDirectionIR { // explode []string

		sortDirectionIIV := sortDirectionIIR // string as string
		sortDirectionIC = append(sortDirectionIC, sortDirectionIIV)
	}

	// items.CollectionFormat: "multi"
	sortDirectionIS := swag.JoinByFormat(sortDirectionIC, "multi")

	return sortDirectionIS
}