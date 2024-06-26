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

// NewListProjectsParams creates a new ListProjectsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListProjectsParams() *ListProjectsParams {
	return &ListProjectsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListProjectsParamsWithTimeout creates a new ListProjectsParams object
// with the ability to set a timeout on a request.
func NewListProjectsParamsWithTimeout(timeout time.Duration) *ListProjectsParams {
	return &ListProjectsParams{
		timeout: timeout,
	}
}

// NewListProjectsParamsWithContext creates a new ListProjectsParams object
// with the ability to set a context for a request.
func NewListProjectsParamsWithContext(ctx context.Context) *ListProjectsParams {
	return &ListProjectsParams{
		Context: ctx,
	}
}

// NewListProjectsParamsWithHTTPClient creates a new ListProjectsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListProjectsParamsWithHTTPClient(client *http.Client) *ListProjectsParams {
	return &ListProjectsParams{
		HTTPClient: client,
	}
}

/*
ListProjectsParams contains all the parameters to send to the API endpoint

	for the list projects operation.

	Typically these are written to a http.Request.
*/
type ListProjectsParams struct {

	// Limit.
	//
	// Format: int32
	Limit *int32

	// Offset.
	//
	// Format: int32
	Offset *int32

	// OrganizationIdentifier.
	OrganizationIdentifier *string

	// SearchTerm.
	SearchTerm *string

	// SortBy.
	SortBy []string

	// SortDirection.
	SortDirection []string

	// UserRelation.
	UserRelation *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list projects params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListProjectsParams) WithDefaults() *ListProjectsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list projects params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListProjectsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list projects params
func (o *ListProjectsParams) WithTimeout(timeout time.Duration) *ListProjectsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list projects params
func (o *ListProjectsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list projects params
func (o *ListProjectsParams) WithContext(ctx context.Context) *ListProjectsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list projects params
func (o *ListProjectsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list projects params
func (o *ListProjectsParams) WithHTTPClient(client *http.Client) *ListProjectsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list projects params
func (o *ListProjectsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the list projects params
func (o *ListProjectsParams) WithLimit(limit *int32) *ListProjectsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the list projects params
func (o *ListProjectsParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithOffset adds the offset to the list projects params
func (o *ListProjectsParams) WithOffset(offset *int32) *ListProjectsParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the list projects params
func (o *ListProjectsParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithOrganizationIdentifier adds the organizationIdentifier to the list projects params
func (o *ListProjectsParams) WithOrganizationIdentifier(organizationIdentifier *string) *ListProjectsParams {
	o.SetOrganizationIdentifier(organizationIdentifier)
	return o
}

// SetOrganizationIdentifier adds the organizationIdentifier to the list projects params
func (o *ListProjectsParams) SetOrganizationIdentifier(organizationIdentifier *string) {
	o.OrganizationIdentifier = organizationIdentifier
}

// WithSearchTerm adds the searchTerm to the list projects params
func (o *ListProjectsParams) WithSearchTerm(searchTerm *string) *ListProjectsParams {
	o.SetSearchTerm(searchTerm)
	return o
}

// SetSearchTerm adds the searchTerm to the list projects params
func (o *ListProjectsParams) SetSearchTerm(searchTerm *string) {
	o.SearchTerm = searchTerm
}

// WithSortBy adds the sortBy to the list projects params
func (o *ListProjectsParams) WithSortBy(sortBy []string) *ListProjectsParams {
	o.SetSortBy(sortBy)
	return o
}

// SetSortBy adds the sortBy to the list projects params
func (o *ListProjectsParams) SetSortBy(sortBy []string) {
	o.SortBy = sortBy
}

// WithSortDirection adds the sortDirection to the list projects params
func (o *ListProjectsParams) WithSortDirection(sortDirection []string) *ListProjectsParams {
	o.SetSortDirection(sortDirection)
	return o
}

// SetSortDirection adds the sortDirection to the list projects params
func (o *ListProjectsParams) SetSortDirection(sortDirection []string) {
	o.SortDirection = sortDirection
}

// WithUserRelation adds the userRelation to the list projects params
func (o *ListProjectsParams) WithUserRelation(userRelation *string) *ListProjectsParams {
	o.SetUserRelation(userRelation)
	return o
}

// SetUserRelation adds the userRelation to the list projects params
func (o *ListProjectsParams) SetUserRelation(userRelation *string) {
	o.UserRelation = userRelation
}

// WriteToRequest writes these params to a swagger request
func (o *ListProjectsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.OrganizationIdentifier != nil {

		// query param organizationIdentifier
		var qrOrganizationIdentifier string

		if o.OrganizationIdentifier != nil {
			qrOrganizationIdentifier = *o.OrganizationIdentifier
		}
		qOrganizationIdentifier := qrOrganizationIdentifier
		if qOrganizationIdentifier != "" {

			if err := r.SetQueryParam("organizationIdentifier", qOrganizationIdentifier); err != nil {
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

	if o.UserRelation != nil {

		// query param userRelation
		var qrUserRelation string

		if o.UserRelation != nil {
			qrUserRelation = *o.UserRelation
		}
		qUserRelation := qrUserRelation
		if qUserRelation != "" {

			if err := r.SetQueryParam("userRelation", qUserRelation); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamListProjects binds the parameter sortBy
func (o *ListProjectsParams) bindParamSortBy(formats strfmt.Registry) []string {
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

// bindParamListProjects binds the parameter sortDirection
func (o *ListProjectsParams) bindParamSortDirection(formats strfmt.Registry) []string {
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
