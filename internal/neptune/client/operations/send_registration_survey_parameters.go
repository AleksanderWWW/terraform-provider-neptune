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

// NewSendRegistrationSurveyParams creates a new SendRegistrationSurveyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSendRegistrationSurveyParams() *SendRegistrationSurveyParams {
	return &SendRegistrationSurveyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSendRegistrationSurveyParamsWithTimeout creates a new SendRegistrationSurveyParams object
// with the ability to set a timeout on a request.
func NewSendRegistrationSurveyParamsWithTimeout(timeout time.Duration) *SendRegistrationSurveyParams {
	return &SendRegistrationSurveyParams{
		timeout: timeout,
	}
}

// NewSendRegistrationSurveyParamsWithContext creates a new SendRegistrationSurveyParams object
// with the ability to set a context for a request.
func NewSendRegistrationSurveyParamsWithContext(ctx context.Context) *SendRegistrationSurveyParams {
	return &SendRegistrationSurveyParams{
		Context: ctx,
	}
}

// NewSendRegistrationSurveyParamsWithHTTPClient creates a new SendRegistrationSurveyParams object
// with the ability to set a custom HTTPClient for a request.
func NewSendRegistrationSurveyParamsWithHTTPClient(client *http.Client) *SendRegistrationSurveyParams {
	return &SendRegistrationSurveyParams{
		HTTPClient: client,
	}
}

/*
SendRegistrationSurveyParams contains all the parameters to send to the API endpoint

	for the send registration survey operation.

	Typically these are written to a http.Request.
*/
type SendRegistrationSurveyParams struct {

	// Survey.
	Survey *models.RegistrationSurveyDTO

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the send registration survey params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SendRegistrationSurveyParams) WithDefaults() *SendRegistrationSurveyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the send registration survey params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SendRegistrationSurveyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the send registration survey params
func (o *SendRegistrationSurveyParams) WithTimeout(timeout time.Duration) *SendRegistrationSurveyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the send registration survey params
func (o *SendRegistrationSurveyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the send registration survey params
func (o *SendRegistrationSurveyParams) WithContext(ctx context.Context) *SendRegistrationSurveyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the send registration survey params
func (o *SendRegistrationSurveyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the send registration survey params
func (o *SendRegistrationSurveyParams) WithHTTPClient(client *http.Client) *SendRegistrationSurveyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the send registration survey params
func (o *SendRegistrationSurveyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSurvey adds the survey to the send registration survey params
func (o *SendRegistrationSurveyParams) WithSurvey(survey *models.RegistrationSurveyDTO) *SendRegistrationSurveyParams {
	o.SetSurvey(survey)
	return o
}

// SetSurvey adds the survey to the send registration survey params
func (o *SendRegistrationSurveyParams) SetSurvey(survey *models.RegistrationSurveyDTO) {
	o.Survey = survey
}

// WriteToRequest writes these params to a swagger request
func (o *SendRegistrationSurveyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Survey != nil {
		if err := r.SetBodyParam(o.Survey); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
