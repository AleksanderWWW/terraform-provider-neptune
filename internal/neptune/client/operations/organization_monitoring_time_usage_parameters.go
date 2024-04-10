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

// NewOrganizationMonitoringTimeUsageParams creates a new OrganizationMonitoringTimeUsageParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewOrganizationMonitoringTimeUsageParams() *OrganizationMonitoringTimeUsageParams {
	return &OrganizationMonitoringTimeUsageParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewOrganizationMonitoringTimeUsageParamsWithTimeout creates a new OrganizationMonitoringTimeUsageParams object
// with the ability to set a timeout on a request.
func NewOrganizationMonitoringTimeUsageParamsWithTimeout(timeout time.Duration) *OrganizationMonitoringTimeUsageParams {
	return &OrganizationMonitoringTimeUsageParams{
		timeout: timeout,
	}
}

// NewOrganizationMonitoringTimeUsageParamsWithContext creates a new OrganizationMonitoringTimeUsageParams object
// with the ability to set a context for a request.
func NewOrganizationMonitoringTimeUsageParamsWithContext(ctx context.Context) *OrganizationMonitoringTimeUsageParams {
	return &OrganizationMonitoringTimeUsageParams{
		Context: ctx,
	}
}

// NewOrganizationMonitoringTimeUsageParamsWithHTTPClient creates a new OrganizationMonitoringTimeUsageParams object
// with the ability to set a custom HTTPClient for a request.
func NewOrganizationMonitoringTimeUsageParamsWithHTTPClient(client *http.Client) *OrganizationMonitoringTimeUsageParams {
	return &OrganizationMonitoringTimeUsageParams{
		HTTPClient: client,
	}
}

/*
OrganizationMonitoringTimeUsageParams contains all the parameters to send to the API endpoint

	for the organization monitoring time usage operation.

	Typically these are written to a http.Request.
*/
type OrganizationMonitoringTimeUsageParams struct {

	// DateMax.
	//
	// Format: date-time
	DateMax *strfmt.DateTime

	// DateMin.
	//
	// Format: date-time
	DateMin *strfmt.DateTime

	// OrganizationIdentifier.
	OrganizationIdentifier string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the organization monitoring time usage params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OrganizationMonitoringTimeUsageParams) WithDefaults() *OrganizationMonitoringTimeUsageParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the organization monitoring time usage params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OrganizationMonitoringTimeUsageParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the organization monitoring time usage params
func (o *OrganizationMonitoringTimeUsageParams) WithTimeout(timeout time.Duration) *OrganizationMonitoringTimeUsageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the organization monitoring time usage params
func (o *OrganizationMonitoringTimeUsageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the organization monitoring time usage params
func (o *OrganizationMonitoringTimeUsageParams) WithContext(ctx context.Context) *OrganizationMonitoringTimeUsageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the organization monitoring time usage params
func (o *OrganizationMonitoringTimeUsageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the organization monitoring time usage params
func (o *OrganizationMonitoringTimeUsageParams) WithHTTPClient(client *http.Client) *OrganizationMonitoringTimeUsageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the organization monitoring time usage params
func (o *OrganizationMonitoringTimeUsageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDateMax adds the dateMax to the organization monitoring time usage params
func (o *OrganizationMonitoringTimeUsageParams) WithDateMax(dateMax *strfmt.DateTime) *OrganizationMonitoringTimeUsageParams {
	o.SetDateMax(dateMax)
	return o
}

// SetDateMax adds the dateMax to the organization monitoring time usage params
func (o *OrganizationMonitoringTimeUsageParams) SetDateMax(dateMax *strfmt.DateTime) {
	o.DateMax = dateMax
}

// WithDateMin adds the dateMin to the organization monitoring time usage params
func (o *OrganizationMonitoringTimeUsageParams) WithDateMin(dateMin *strfmt.DateTime) *OrganizationMonitoringTimeUsageParams {
	o.SetDateMin(dateMin)
	return o
}

// SetDateMin adds the dateMin to the organization monitoring time usage params
func (o *OrganizationMonitoringTimeUsageParams) SetDateMin(dateMin *strfmt.DateTime) {
	o.DateMin = dateMin
}

// WithOrganizationIdentifier adds the organizationIdentifier to the organization monitoring time usage params
func (o *OrganizationMonitoringTimeUsageParams) WithOrganizationIdentifier(organizationIdentifier string) *OrganizationMonitoringTimeUsageParams {
	o.SetOrganizationIdentifier(organizationIdentifier)
	return o
}

// SetOrganizationIdentifier adds the organizationIdentifier to the organization monitoring time usage params
func (o *OrganizationMonitoringTimeUsageParams) SetOrganizationIdentifier(organizationIdentifier string) {
	o.OrganizationIdentifier = organizationIdentifier
}

// WriteToRequest writes these params to a swagger request
func (o *OrganizationMonitoringTimeUsageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DateMax != nil {

		// query param dateMax
		var qrDateMax strfmt.DateTime

		if o.DateMax != nil {
			qrDateMax = *o.DateMax
		}
		qDateMax := qrDateMax.String()
		if qDateMax != "" {

			if err := r.SetQueryParam("dateMax", qDateMax); err != nil {
				return err
			}
		}
	}

	if o.DateMin != nil {

		// query param dateMin
		var qrDateMin strfmt.DateTime

		if o.DateMin != nil {
			qrDateMin = *o.DateMin
		}
		qDateMin := qrDateMin.String()
		if qDateMin != "" {

			if err := r.SetQueryParam("dateMin", qDateMin); err != nil {
				return err
			}
		}
	}

	// query param organizationIdentifier
	qrOrganizationIdentifier := o.OrganizationIdentifier
	qOrganizationIdentifier := qrOrganizationIdentifier
	if qOrganizationIdentifier != "" {

		if err := r.SetQueryParam("organizationIdentifier", qOrganizationIdentifier); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
