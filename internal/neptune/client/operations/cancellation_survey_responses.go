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

// CancellationSurveyReader is a Reader for the CancellationSurvey structure.
type CancellationSurveyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CancellationSurveyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCancellationSurveyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCancellationSurveyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCancellationSurveyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCancellationSurveyForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCancellationSurveyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewCancellationSurveyRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewCancellationSurveyUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /api/backend/v1/organizations2/cancellationSurvey] cancellationSurvey", response, response.Code())
	}
}

// NewCancellationSurveyOK creates a CancellationSurveyOK with default headers values
func NewCancellationSurveyOK() *CancellationSurveyOK {
	return &CancellationSurveyOK{}
}

/*
CancellationSurveyOK describes a response with status code 200, with default header values.

No response
*/
type CancellationSurveyOK struct {
}

// IsSuccess returns true when this cancellation survey o k response has a 2xx status code
func (o *CancellationSurveyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cancellation survey o k response has a 3xx status code
func (o *CancellationSurveyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cancellation survey o k response has a 4xx status code
func (o *CancellationSurveyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cancellation survey o k response has a 5xx status code
func (o *CancellationSurveyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cancellation survey o k response a status code equal to that given
func (o *CancellationSurveyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the cancellation survey o k response
func (o *CancellationSurveyOK) Code() int {
	return 200
}

func (o *CancellationSurveyOK) Error() string {
	return fmt.Sprintf("[POST /api/backend/v1/organizations2/cancellationSurvey][%d] cancellationSurveyOK ", 200)
}

func (o *CancellationSurveyOK) String() string {
	return fmt.Sprintf("[POST /api/backend/v1/organizations2/cancellationSurvey][%d] cancellationSurveyOK ", 200)
}

func (o *CancellationSurveyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCancellationSurveyBadRequest creates a CancellationSurveyBadRequest with default headers values
func NewCancellationSurveyBadRequest() *CancellationSurveyBadRequest {
	return &CancellationSurveyBadRequest{}
}

/*
CancellationSurveyBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CancellationSurveyBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this cancellation survey bad request response has a 2xx status code
func (o *CancellationSurveyBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cancellation survey bad request response has a 3xx status code
func (o *CancellationSurveyBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cancellation survey bad request response has a 4xx status code
func (o *CancellationSurveyBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this cancellation survey bad request response has a 5xx status code
func (o *CancellationSurveyBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this cancellation survey bad request response a status code equal to that given
func (o *CancellationSurveyBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the cancellation survey bad request response
func (o *CancellationSurveyBadRequest) Code() int {
	return 400
}

func (o *CancellationSurveyBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/backend/v1/organizations2/cancellationSurvey][%d] cancellationSurveyBadRequest  %+v", 400, o.Payload)
}

func (o *CancellationSurveyBadRequest) String() string {
	return fmt.Sprintf("[POST /api/backend/v1/organizations2/cancellationSurvey][%d] cancellationSurveyBadRequest  %+v", 400, o.Payload)
}

func (o *CancellationSurveyBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CancellationSurveyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancellationSurveyUnauthorized creates a CancellationSurveyUnauthorized with default headers values
func NewCancellationSurveyUnauthorized() *CancellationSurveyUnauthorized {
	return &CancellationSurveyUnauthorized{}
}

/*
CancellationSurveyUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CancellationSurveyUnauthorized struct {
}

// IsSuccess returns true when this cancellation survey unauthorized response has a 2xx status code
func (o *CancellationSurveyUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cancellation survey unauthorized response has a 3xx status code
func (o *CancellationSurveyUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cancellation survey unauthorized response has a 4xx status code
func (o *CancellationSurveyUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this cancellation survey unauthorized response has a 5xx status code
func (o *CancellationSurveyUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this cancellation survey unauthorized response a status code equal to that given
func (o *CancellationSurveyUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the cancellation survey unauthorized response
func (o *CancellationSurveyUnauthorized) Code() int {
	return 401
}

func (o *CancellationSurveyUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/backend/v1/organizations2/cancellationSurvey][%d] cancellationSurveyUnauthorized ", 401)
}

func (o *CancellationSurveyUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/backend/v1/organizations2/cancellationSurvey][%d] cancellationSurveyUnauthorized ", 401)
}

func (o *CancellationSurveyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCancellationSurveyForbidden creates a CancellationSurveyForbidden with default headers values
func NewCancellationSurveyForbidden() *CancellationSurveyForbidden {
	return &CancellationSurveyForbidden{}
}

/*
CancellationSurveyForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CancellationSurveyForbidden struct {
}

// IsSuccess returns true when this cancellation survey forbidden response has a 2xx status code
func (o *CancellationSurveyForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cancellation survey forbidden response has a 3xx status code
func (o *CancellationSurveyForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cancellation survey forbidden response has a 4xx status code
func (o *CancellationSurveyForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this cancellation survey forbidden response has a 5xx status code
func (o *CancellationSurveyForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this cancellation survey forbidden response a status code equal to that given
func (o *CancellationSurveyForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the cancellation survey forbidden response
func (o *CancellationSurveyForbidden) Code() int {
	return 403
}

func (o *CancellationSurveyForbidden) Error() string {
	return fmt.Sprintf("[POST /api/backend/v1/organizations2/cancellationSurvey][%d] cancellationSurveyForbidden ", 403)
}

func (o *CancellationSurveyForbidden) String() string {
	return fmt.Sprintf("[POST /api/backend/v1/organizations2/cancellationSurvey][%d] cancellationSurveyForbidden ", 403)
}

func (o *CancellationSurveyForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCancellationSurveyNotFound creates a CancellationSurveyNotFound with default headers values
func NewCancellationSurveyNotFound() *CancellationSurveyNotFound {
	return &CancellationSurveyNotFound{}
}

/*
CancellationSurveyNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CancellationSurveyNotFound struct {
}

// IsSuccess returns true when this cancellation survey not found response has a 2xx status code
func (o *CancellationSurveyNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cancellation survey not found response has a 3xx status code
func (o *CancellationSurveyNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cancellation survey not found response has a 4xx status code
func (o *CancellationSurveyNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this cancellation survey not found response has a 5xx status code
func (o *CancellationSurveyNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this cancellation survey not found response a status code equal to that given
func (o *CancellationSurveyNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the cancellation survey not found response
func (o *CancellationSurveyNotFound) Code() int {
	return 404
}

func (o *CancellationSurveyNotFound) Error() string {
	return fmt.Sprintf("[POST /api/backend/v1/organizations2/cancellationSurvey][%d] cancellationSurveyNotFound ", 404)
}

func (o *CancellationSurveyNotFound) String() string {
	return fmt.Sprintf("[POST /api/backend/v1/organizations2/cancellationSurvey][%d] cancellationSurveyNotFound ", 404)
}

func (o *CancellationSurveyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCancellationSurveyRequestTimeout creates a CancellationSurveyRequestTimeout with default headers values
func NewCancellationSurveyRequestTimeout() *CancellationSurveyRequestTimeout {
	return &CancellationSurveyRequestTimeout{}
}

/*
CancellationSurveyRequestTimeout describes a response with status code 408, with default header values.

Request Timeout
*/
type CancellationSurveyRequestTimeout struct {
}

// IsSuccess returns true when this cancellation survey request timeout response has a 2xx status code
func (o *CancellationSurveyRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cancellation survey request timeout response has a 3xx status code
func (o *CancellationSurveyRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cancellation survey request timeout response has a 4xx status code
func (o *CancellationSurveyRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this cancellation survey request timeout response has a 5xx status code
func (o *CancellationSurveyRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this cancellation survey request timeout response a status code equal to that given
func (o *CancellationSurveyRequestTimeout) IsCode(code int) bool {
	return code == 408
}

// Code gets the status code for the cancellation survey request timeout response
func (o *CancellationSurveyRequestTimeout) Code() int {
	return 408
}

func (o *CancellationSurveyRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/backend/v1/organizations2/cancellationSurvey][%d] cancellationSurveyRequestTimeout ", 408)
}

func (o *CancellationSurveyRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/backend/v1/organizations2/cancellationSurvey][%d] cancellationSurveyRequestTimeout ", 408)
}

func (o *CancellationSurveyRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCancellationSurveyUnprocessableEntity creates a CancellationSurveyUnprocessableEntity with default headers values
func NewCancellationSurveyUnprocessableEntity() *CancellationSurveyUnprocessableEntity {
	return &CancellationSurveyUnprocessableEntity{}
}

/*
CancellationSurveyUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type CancellationSurveyUnprocessableEntity struct {
}

// IsSuccess returns true when this cancellation survey unprocessable entity response has a 2xx status code
func (o *CancellationSurveyUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cancellation survey unprocessable entity response has a 3xx status code
func (o *CancellationSurveyUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cancellation survey unprocessable entity response has a 4xx status code
func (o *CancellationSurveyUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this cancellation survey unprocessable entity response has a 5xx status code
func (o *CancellationSurveyUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this cancellation survey unprocessable entity response a status code equal to that given
func (o *CancellationSurveyUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the cancellation survey unprocessable entity response
func (o *CancellationSurveyUnprocessableEntity) Code() int {
	return 422
}

func (o *CancellationSurveyUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /api/backend/v1/organizations2/cancellationSurvey][%d] cancellationSurveyUnprocessableEntity ", 422)
}

func (o *CancellationSurveyUnprocessableEntity) String() string {
	return fmt.Sprintf("[POST /api/backend/v1/organizations2/cancellationSurvey][%d] cancellationSurveyUnprocessableEntity ", 422)
}

func (o *CancellationSurveyUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
