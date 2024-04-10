// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"terraform-provider-neptune/internal/neptune/models"
)

// UpdateVisibilityRestrictionsReader is a Reader for the UpdateVisibilityRestrictions structure.
type UpdateVisibilityRestrictionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateVisibilityRestrictionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateVisibilityRestrictionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateVisibilityRestrictionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateVisibilityRestrictionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateVisibilityRestrictionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateVisibilityRestrictionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewUpdateVisibilityRestrictionsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewUpdateVisibilityRestrictionsUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /api/backend/v1/organizations2/{organizationIdentifier}/updateVisibilityRestrictions] updateVisibilityRestrictions", response, response.Code())
	}
}

// NewUpdateVisibilityRestrictionsOK creates a UpdateVisibilityRestrictionsOK with default headers values
func NewUpdateVisibilityRestrictionsOK() *UpdateVisibilityRestrictionsOK {
	return &UpdateVisibilityRestrictionsOK{}
}

/*
UpdateVisibilityRestrictionsOK describes a response with status code 200, with default header values.

No response
*/
type UpdateVisibilityRestrictionsOK struct {
}

// IsSuccess returns true when this update visibility restrictions o k response has a 2xx status code
func (o *UpdateVisibilityRestrictionsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update visibility restrictions o k response has a 3xx status code
func (o *UpdateVisibilityRestrictionsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update visibility restrictions o k response has a 4xx status code
func (o *UpdateVisibilityRestrictionsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update visibility restrictions o k response has a 5xx status code
func (o *UpdateVisibilityRestrictionsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update visibility restrictions o k response a status code equal to that given
func (o *UpdateVisibilityRestrictionsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update visibility restrictions o k response
func (o *UpdateVisibilityRestrictionsOK) Code() int {
	return 200
}

func (o *UpdateVisibilityRestrictionsOK) Error() string {
	return fmt.Sprintf("[POST /api/backend/v1/organizations2/{organizationIdentifier}/updateVisibilityRestrictions][%d] updateVisibilityRestrictionsOK ", 200)
}

func (o *UpdateVisibilityRestrictionsOK) String() string {
	return fmt.Sprintf("[POST /api/backend/v1/organizations2/{organizationIdentifier}/updateVisibilityRestrictions][%d] updateVisibilityRestrictionsOK ", 200)
}

func (o *UpdateVisibilityRestrictionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateVisibilityRestrictionsBadRequest creates a UpdateVisibilityRestrictionsBadRequest with default headers values
func NewUpdateVisibilityRestrictionsBadRequest() *UpdateVisibilityRestrictionsBadRequest {
	return &UpdateVisibilityRestrictionsBadRequest{}
}

/*
UpdateVisibilityRestrictionsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UpdateVisibilityRestrictionsBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this update visibility restrictions bad request response has a 2xx status code
func (o *UpdateVisibilityRestrictionsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update visibility restrictions bad request response has a 3xx status code
func (o *UpdateVisibilityRestrictionsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update visibility restrictions bad request response has a 4xx status code
func (o *UpdateVisibilityRestrictionsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update visibility restrictions bad request response has a 5xx status code
func (o *UpdateVisibilityRestrictionsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update visibility restrictions bad request response a status code equal to that given
func (o *UpdateVisibilityRestrictionsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update visibility restrictions bad request response
func (o *UpdateVisibilityRestrictionsBadRequest) Code() int {
	return 400
}

func (o *UpdateVisibilityRestrictionsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/backend/v1/organizations2/{organizationIdentifier}/updateVisibilityRestrictions][%d] updateVisibilityRestrictionsBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateVisibilityRestrictionsBadRequest) String() string {
	return fmt.Sprintf("[POST /api/backend/v1/organizations2/{organizationIdentifier}/updateVisibilityRestrictions][%d] updateVisibilityRestrictionsBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateVisibilityRestrictionsBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateVisibilityRestrictionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateVisibilityRestrictionsUnauthorized creates a UpdateVisibilityRestrictionsUnauthorized with default headers values
func NewUpdateVisibilityRestrictionsUnauthorized() *UpdateVisibilityRestrictionsUnauthorized {
	return &UpdateVisibilityRestrictionsUnauthorized{}
}

/*
UpdateVisibilityRestrictionsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type UpdateVisibilityRestrictionsUnauthorized struct {
}

// IsSuccess returns true when this update visibility restrictions unauthorized response has a 2xx status code
func (o *UpdateVisibilityRestrictionsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update visibility restrictions unauthorized response has a 3xx status code
func (o *UpdateVisibilityRestrictionsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update visibility restrictions unauthorized response has a 4xx status code
func (o *UpdateVisibilityRestrictionsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update visibility restrictions unauthorized response has a 5xx status code
func (o *UpdateVisibilityRestrictionsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update visibility restrictions unauthorized response a status code equal to that given
func (o *UpdateVisibilityRestrictionsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the update visibility restrictions unauthorized response
func (o *UpdateVisibilityRestrictionsUnauthorized) Code() int {
	return 401
}

func (o *UpdateVisibilityRestrictionsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/backend/v1/organizations2/{organizationIdentifier}/updateVisibilityRestrictions][%d] updateVisibilityRestrictionsUnauthorized ", 401)
}

func (o *UpdateVisibilityRestrictionsUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/backend/v1/organizations2/{organizationIdentifier}/updateVisibilityRestrictions][%d] updateVisibilityRestrictionsUnauthorized ", 401)
}

func (o *UpdateVisibilityRestrictionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateVisibilityRestrictionsForbidden creates a UpdateVisibilityRestrictionsForbidden with default headers values
func NewUpdateVisibilityRestrictionsForbidden() *UpdateVisibilityRestrictionsForbidden {
	return &UpdateVisibilityRestrictionsForbidden{}
}

/*
UpdateVisibilityRestrictionsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UpdateVisibilityRestrictionsForbidden struct {
}

// IsSuccess returns true when this update visibility restrictions forbidden response has a 2xx status code
func (o *UpdateVisibilityRestrictionsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update visibility restrictions forbidden response has a 3xx status code
func (o *UpdateVisibilityRestrictionsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update visibility restrictions forbidden response has a 4xx status code
func (o *UpdateVisibilityRestrictionsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update visibility restrictions forbidden response has a 5xx status code
func (o *UpdateVisibilityRestrictionsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update visibility restrictions forbidden response a status code equal to that given
func (o *UpdateVisibilityRestrictionsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the update visibility restrictions forbidden response
func (o *UpdateVisibilityRestrictionsForbidden) Code() int {
	return 403
}

func (o *UpdateVisibilityRestrictionsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/backend/v1/organizations2/{organizationIdentifier}/updateVisibilityRestrictions][%d] updateVisibilityRestrictionsForbidden ", 403)
}

func (o *UpdateVisibilityRestrictionsForbidden) String() string {
	return fmt.Sprintf("[POST /api/backend/v1/organizations2/{organizationIdentifier}/updateVisibilityRestrictions][%d] updateVisibilityRestrictionsForbidden ", 403)
}

func (o *UpdateVisibilityRestrictionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateVisibilityRestrictionsNotFound creates a UpdateVisibilityRestrictionsNotFound with default headers values
func NewUpdateVisibilityRestrictionsNotFound() *UpdateVisibilityRestrictionsNotFound {
	return &UpdateVisibilityRestrictionsNotFound{}
}

/*
UpdateVisibilityRestrictionsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UpdateVisibilityRestrictionsNotFound struct {
}

// IsSuccess returns true when this update visibility restrictions not found response has a 2xx status code
func (o *UpdateVisibilityRestrictionsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update visibility restrictions not found response has a 3xx status code
func (o *UpdateVisibilityRestrictionsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update visibility restrictions not found response has a 4xx status code
func (o *UpdateVisibilityRestrictionsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update visibility restrictions not found response has a 5xx status code
func (o *UpdateVisibilityRestrictionsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update visibility restrictions not found response a status code equal to that given
func (o *UpdateVisibilityRestrictionsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update visibility restrictions not found response
func (o *UpdateVisibilityRestrictionsNotFound) Code() int {
	return 404
}

func (o *UpdateVisibilityRestrictionsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/backend/v1/organizations2/{organizationIdentifier}/updateVisibilityRestrictions][%d] updateVisibilityRestrictionsNotFound ", 404)
}

func (o *UpdateVisibilityRestrictionsNotFound) String() string {
	return fmt.Sprintf("[POST /api/backend/v1/organizations2/{organizationIdentifier}/updateVisibilityRestrictions][%d] updateVisibilityRestrictionsNotFound ", 404)
}

func (o *UpdateVisibilityRestrictionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateVisibilityRestrictionsRequestTimeout creates a UpdateVisibilityRestrictionsRequestTimeout with default headers values
func NewUpdateVisibilityRestrictionsRequestTimeout() *UpdateVisibilityRestrictionsRequestTimeout {
	return &UpdateVisibilityRestrictionsRequestTimeout{}
}

/*
UpdateVisibilityRestrictionsRequestTimeout describes a response with status code 408, with default header values.

Request Timeout
*/
type UpdateVisibilityRestrictionsRequestTimeout struct {
}

// IsSuccess returns true when this update visibility restrictions request timeout response has a 2xx status code
func (o *UpdateVisibilityRestrictionsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update visibility restrictions request timeout response has a 3xx status code
func (o *UpdateVisibilityRestrictionsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update visibility restrictions request timeout response has a 4xx status code
func (o *UpdateVisibilityRestrictionsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this update visibility restrictions request timeout response has a 5xx status code
func (o *UpdateVisibilityRestrictionsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this update visibility restrictions request timeout response a status code equal to that given
func (o *UpdateVisibilityRestrictionsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

// Code gets the status code for the update visibility restrictions request timeout response
func (o *UpdateVisibilityRestrictionsRequestTimeout) Code() int {
	return 408
}

func (o *UpdateVisibilityRestrictionsRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/backend/v1/organizations2/{organizationIdentifier}/updateVisibilityRestrictions][%d] updateVisibilityRestrictionsRequestTimeout ", 408)
}

func (o *UpdateVisibilityRestrictionsRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/backend/v1/organizations2/{organizationIdentifier}/updateVisibilityRestrictions][%d] updateVisibilityRestrictionsRequestTimeout ", 408)
}

func (o *UpdateVisibilityRestrictionsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateVisibilityRestrictionsUnprocessableEntity creates a UpdateVisibilityRestrictionsUnprocessableEntity with default headers values
func NewUpdateVisibilityRestrictionsUnprocessableEntity() *UpdateVisibilityRestrictionsUnprocessableEntity {
	return &UpdateVisibilityRestrictionsUnprocessableEntity{}
}

/*
UpdateVisibilityRestrictionsUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type UpdateVisibilityRestrictionsUnprocessableEntity struct {
}

// IsSuccess returns true when this update visibility restrictions unprocessable entity response has a 2xx status code
func (o *UpdateVisibilityRestrictionsUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update visibility restrictions unprocessable entity response has a 3xx status code
func (o *UpdateVisibilityRestrictionsUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update visibility restrictions unprocessable entity response has a 4xx status code
func (o *UpdateVisibilityRestrictionsUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this update visibility restrictions unprocessable entity response has a 5xx status code
func (o *UpdateVisibilityRestrictionsUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this update visibility restrictions unprocessable entity response a status code equal to that given
func (o *UpdateVisibilityRestrictionsUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the update visibility restrictions unprocessable entity response
func (o *UpdateVisibilityRestrictionsUnprocessableEntity) Code() int {
	return 422
}

func (o *UpdateVisibilityRestrictionsUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /api/backend/v1/organizations2/{organizationIdentifier}/updateVisibilityRestrictions][%d] updateVisibilityRestrictionsUnprocessableEntity ", 422)
}

func (o *UpdateVisibilityRestrictionsUnprocessableEntity) String() string {
	return fmt.Sprintf("[POST /api/backend/v1/organizations2/{organizationIdentifier}/updateVisibilityRestrictions][%d] updateVisibilityRestrictionsUnprocessableEntity ", 422)
}

func (o *UpdateVisibilityRestrictionsUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}