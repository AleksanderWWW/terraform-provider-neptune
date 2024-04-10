// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"client/models"
)

// MonitoringTimeStatusReader is a Reader for the MonitoringTimeStatus structure.
type MonitoringTimeStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MonitoringTimeStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMonitoringTimeStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewMonitoringTimeStatusBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewMonitoringTimeStatusUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewMonitoringTimeStatusForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewMonitoringTimeStatusNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewMonitoringTimeStatusRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewMonitoringTimeStatusUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /api/backend/v1/resources/status] monitoringTimeStatus", response, response.Code())
	}
}

// NewMonitoringTimeStatusOK creates a MonitoringTimeStatusOK with default headers values
func NewMonitoringTimeStatusOK() *MonitoringTimeStatusOK {
	return &MonitoringTimeStatusOK{}
}

/*
MonitoringTimeStatusOK describes a response with status code 200, with default header values.

OK
*/
type MonitoringTimeStatusOK struct {
	Payload *models.MonitoringTimeStatusDTO
}

// IsSuccess returns true when this monitoring time status o k response has a 2xx status code
func (o *MonitoringTimeStatusOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this monitoring time status o k response has a 3xx status code
func (o *MonitoringTimeStatusOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this monitoring time status o k response has a 4xx status code
func (o *MonitoringTimeStatusOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this monitoring time status o k response has a 5xx status code
func (o *MonitoringTimeStatusOK) IsServerError() bool {
	return false
}

// IsCode returns true when this monitoring time status o k response a status code equal to that given
func (o *MonitoringTimeStatusOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the monitoring time status o k response
func (o *MonitoringTimeStatusOK) Code() int {
	return 200
}

func (o *MonitoringTimeStatusOK) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/resources/status][%d] monitoringTimeStatusOK  %+v", 200, o.Payload)
}

func (o *MonitoringTimeStatusOK) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/resources/status][%d] monitoringTimeStatusOK  %+v", 200, o.Payload)
}

func (o *MonitoringTimeStatusOK) GetPayload() *models.MonitoringTimeStatusDTO {
	return o.Payload
}

func (o *MonitoringTimeStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MonitoringTimeStatusDTO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMonitoringTimeStatusBadRequest creates a MonitoringTimeStatusBadRequest with default headers values
func NewMonitoringTimeStatusBadRequest() *MonitoringTimeStatusBadRequest {
	return &MonitoringTimeStatusBadRequest{}
}

/*
MonitoringTimeStatusBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type MonitoringTimeStatusBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this monitoring time status bad request response has a 2xx status code
func (o *MonitoringTimeStatusBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this monitoring time status bad request response has a 3xx status code
func (o *MonitoringTimeStatusBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this monitoring time status bad request response has a 4xx status code
func (o *MonitoringTimeStatusBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this monitoring time status bad request response has a 5xx status code
func (o *MonitoringTimeStatusBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this monitoring time status bad request response a status code equal to that given
func (o *MonitoringTimeStatusBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the monitoring time status bad request response
func (o *MonitoringTimeStatusBadRequest) Code() int {
	return 400
}

func (o *MonitoringTimeStatusBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/resources/status][%d] monitoringTimeStatusBadRequest  %+v", 400, o.Payload)
}

func (o *MonitoringTimeStatusBadRequest) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/resources/status][%d] monitoringTimeStatusBadRequest  %+v", 400, o.Payload)
}

func (o *MonitoringTimeStatusBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *MonitoringTimeStatusBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMonitoringTimeStatusUnauthorized creates a MonitoringTimeStatusUnauthorized with default headers values
func NewMonitoringTimeStatusUnauthorized() *MonitoringTimeStatusUnauthorized {
	return &MonitoringTimeStatusUnauthorized{}
}

/*
MonitoringTimeStatusUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type MonitoringTimeStatusUnauthorized struct {
}

// IsSuccess returns true when this monitoring time status unauthorized response has a 2xx status code
func (o *MonitoringTimeStatusUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this monitoring time status unauthorized response has a 3xx status code
func (o *MonitoringTimeStatusUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this monitoring time status unauthorized response has a 4xx status code
func (o *MonitoringTimeStatusUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this monitoring time status unauthorized response has a 5xx status code
func (o *MonitoringTimeStatusUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this monitoring time status unauthorized response a status code equal to that given
func (o *MonitoringTimeStatusUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the monitoring time status unauthorized response
func (o *MonitoringTimeStatusUnauthorized) Code() int {
	return 401
}

func (o *MonitoringTimeStatusUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/resources/status][%d] monitoringTimeStatusUnauthorized ", 401)
}

func (o *MonitoringTimeStatusUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/resources/status][%d] monitoringTimeStatusUnauthorized ", 401)
}

func (o *MonitoringTimeStatusUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewMonitoringTimeStatusForbidden creates a MonitoringTimeStatusForbidden with default headers values
func NewMonitoringTimeStatusForbidden() *MonitoringTimeStatusForbidden {
	return &MonitoringTimeStatusForbidden{}
}

/*
MonitoringTimeStatusForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type MonitoringTimeStatusForbidden struct {
}

// IsSuccess returns true when this monitoring time status forbidden response has a 2xx status code
func (o *MonitoringTimeStatusForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this monitoring time status forbidden response has a 3xx status code
func (o *MonitoringTimeStatusForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this monitoring time status forbidden response has a 4xx status code
func (o *MonitoringTimeStatusForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this monitoring time status forbidden response has a 5xx status code
func (o *MonitoringTimeStatusForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this monitoring time status forbidden response a status code equal to that given
func (o *MonitoringTimeStatusForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the monitoring time status forbidden response
func (o *MonitoringTimeStatusForbidden) Code() int {
	return 403
}

func (o *MonitoringTimeStatusForbidden) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/resources/status][%d] monitoringTimeStatusForbidden ", 403)
}

func (o *MonitoringTimeStatusForbidden) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/resources/status][%d] monitoringTimeStatusForbidden ", 403)
}

func (o *MonitoringTimeStatusForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewMonitoringTimeStatusNotFound creates a MonitoringTimeStatusNotFound with default headers values
func NewMonitoringTimeStatusNotFound() *MonitoringTimeStatusNotFound {
	return &MonitoringTimeStatusNotFound{}
}

/*
MonitoringTimeStatusNotFound describes a response with status code 404, with default header values.

Not Found
*/
type MonitoringTimeStatusNotFound struct {
}

// IsSuccess returns true when this monitoring time status not found response has a 2xx status code
func (o *MonitoringTimeStatusNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this monitoring time status not found response has a 3xx status code
func (o *MonitoringTimeStatusNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this monitoring time status not found response has a 4xx status code
func (o *MonitoringTimeStatusNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this monitoring time status not found response has a 5xx status code
func (o *MonitoringTimeStatusNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this monitoring time status not found response a status code equal to that given
func (o *MonitoringTimeStatusNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the monitoring time status not found response
func (o *MonitoringTimeStatusNotFound) Code() int {
	return 404
}

func (o *MonitoringTimeStatusNotFound) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/resources/status][%d] monitoringTimeStatusNotFound ", 404)
}

func (o *MonitoringTimeStatusNotFound) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/resources/status][%d] monitoringTimeStatusNotFound ", 404)
}

func (o *MonitoringTimeStatusNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewMonitoringTimeStatusRequestTimeout creates a MonitoringTimeStatusRequestTimeout with default headers values
func NewMonitoringTimeStatusRequestTimeout() *MonitoringTimeStatusRequestTimeout {
	return &MonitoringTimeStatusRequestTimeout{}
}

/*
MonitoringTimeStatusRequestTimeout describes a response with status code 408, with default header values.

Request Timeout
*/
type MonitoringTimeStatusRequestTimeout struct {
}

// IsSuccess returns true when this monitoring time status request timeout response has a 2xx status code
func (o *MonitoringTimeStatusRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this monitoring time status request timeout response has a 3xx status code
func (o *MonitoringTimeStatusRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this monitoring time status request timeout response has a 4xx status code
func (o *MonitoringTimeStatusRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this monitoring time status request timeout response has a 5xx status code
func (o *MonitoringTimeStatusRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this monitoring time status request timeout response a status code equal to that given
func (o *MonitoringTimeStatusRequestTimeout) IsCode(code int) bool {
	return code == 408
}

// Code gets the status code for the monitoring time status request timeout response
func (o *MonitoringTimeStatusRequestTimeout) Code() int {
	return 408
}

func (o *MonitoringTimeStatusRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/resources/status][%d] monitoringTimeStatusRequestTimeout ", 408)
}

func (o *MonitoringTimeStatusRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/resources/status][%d] monitoringTimeStatusRequestTimeout ", 408)
}

func (o *MonitoringTimeStatusRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewMonitoringTimeStatusUnprocessableEntity creates a MonitoringTimeStatusUnprocessableEntity with default headers values
func NewMonitoringTimeStatusUnprocessableEntity() *MonitoringTimeStatusUnprocessableEntity {
	return &MonitoringTimeStatusUnprocessableEntity{}
}

/*
MonitoringTimeStatusUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type MonitoringTimeStatusUnprocessableEntity struct {
}

// IsSuccess returns true when this monitoring time status unprocessable entity response has a 2xx status code
func (o *MonitoringTimeStatusUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this monitoring time status unprocessable entity response has a 3xx status code
func (o *MonitoringTimeStatusUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this monitoring time status unprocessable entity response has a 4xx status code
func (o *MonitoringTimeStatusUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this monitoring time status unprocessable entity response has a 5xx status code
func (o *MonitoringTimeStatusUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this monitoring time status unprocessable entity response a status code equal to that given
func (o *MonitoringTimeStatusUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the monitoring time status unprocessable entity response
func (o *MonitoringTimeStatusUnprocessableEntity) Code() int {
	return 422
}

func (o *MonitoringTimeStatusUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/resources/status][%d] monitoringTimeStatusUnprocessableEntity ", 422)
}

func (o *MonitoringTimeStatusUnprocessableEntity) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/resources/status][%d] monitoringTimeStatusUnprocessableEntity ", 422)
}

func (o *MonitoringTimeStatusUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
