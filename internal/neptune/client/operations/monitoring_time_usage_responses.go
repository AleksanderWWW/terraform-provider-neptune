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

// MonitoringTimeUsageReader is a Reader for the MonitoringTimeUsage structure.
type MonitoringTimeUsageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MonitoringTimeUsageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMonitoringTimeUsageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewMonitoringTimeUsageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewMonitoringTimeUsageUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewMonitoringTimeUsageForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewMonitoringTimeUsageNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewMonitoringTimeUsageRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewMonitoringTimeUsageUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /api/backend/v1/resources/usage] monitoringTimeUsage", response, response.Code())
	}
}

// NewMonitoringTimeUsageOK creates a MonitoringTimeUsageOK with default headers values
func NewMonitoringTimeUsageOK() *MonitoringTimeUsageOK {
	return &MonitoringTimeUsageOK{}
}

/*
MonitoringTimeUsageOK describes a response with status code 200, with default header values.

OK
*/
type MonitoringTimeUsageOK struct {
	Payload *models.MonitoringTimeUsageDTO
}

// IsSuccess returns true when this monitoring time usage o k response has a 2xx status code
func (o *MonitoringTimeUsageOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this monitoring time usage o k response has a 3xx status code
func (o *MonitoringTimeUsageOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this monitoring time usage o k response has a 4xx status code
func (o *MonitoringTimeUsageOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this monitoring time usage o k response has a 5xx status code
func (o *MonitoringTimeUsageOK) IsServerError() bool {
	return false
}

// IsCode returns true when this monitoring time usage o k response a status code equal to that given
func (o *MonitoringTimeUsageOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the monitoring time usage o k response
func (o *MonitoringTimeUsageOK) Code() int {
	return 200
}

func (o *MonitoringTimeUsageOK) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/resources/usage][%d] monitoringTimeUsageOK  %+v", 200, o.Payload)
}

func (o *MonitoringTimeUsageOK) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/resources/usage][%d] monitoringTimeUsageOK  %+v", 200, o.Payload)
}

func (o *MonitoringTimeUsageOK) GetPayload() *models.MonitoringTimeUsageDTO {
	return o.Payload
}

func (o *MonitoringTimeUsageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MonitoringTimeUsageDTO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMonitoringTimeUsageBadRequest creates a MonitoringTimeUsageBadRequest with default headers values
func NewMonitoringTimeUsageBadRequest() *MonitoringTimeUsageBadRequest {
	return &MonitoringTimeUsageBadRequest{}
}

/*
MonitoringTimeUsageBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type MonitoringTimeUsageBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this monitoring time usage bad request response has a 2xx status code
func (o *MonitoringTimeUsageBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this monitoring time usage bad request response has a 3xx status code
func (o *MonitoringTimeUsageBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this monitoring time usage bad request response has a 4xx status code
func (o *MonitoringTimeUsageBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this monitoring time usage bad request response has a 5xx status code
func (o *MonitoringTimeUsageBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this monitoring time usage bad request response a status code equal to that given
func (o *MonitoringTimeUsageBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the monitoring time usage bad request response
func (o *MonitoringTimeUsageBadRequest) Code() int {
	return 400
}

func (o *MonitoringTimeUsageBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/resources/usage][%d] monitoringTimeUsageBadRequest  %+v", 400, o.Payload)
}

func (o *MonitoringTimeUsageBadRequest) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/resources/usage][%d] monitoringTimeUsageBadRequest  %+v", 400, o.Payload)
}

func (o *MonitoringTimeUsageBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *MonitoringTimeUsageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMonitoringTimeUsageUnauthorized creates a MonitoringTimeUsageUnauthorized with default headers values
func NewMonitoringTimeUsageUnauthorized() *MonitoringTimeUsageUnauthorized {
	return &MonitoringTimeUsageUnauthorized{}
}

/*
MonitoringTimeUsageUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type MonitoringTimeUsageUnauthorized struct {
}

// IsSuccess returns true when this monitoring time usage unauthorized response has a 2xx status code
func (o *MonitoringTimeUsageUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this monitoring time usage unauthorized response has a 3xx status code
func (o *MonitoringTimeUsageUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this monitoring time usage unauthorized response has a 4xx status code
func (o *MonitoringTimeUsageUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this monitoring time usage unauthorized response has a 5xx status code
func (o *MonitoringTimeUsageUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this monitoring time usage unauthorized response a status code equal to that given
func (o *MonitoringTimeUsageUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the monitoring time usage unauthorized response
func (o *MonitoringTimeUsageUnauthorized) Code() int {
	return 401
}

func (o *MonitoringTimeUsageUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/resources/usage][%d] monitoringTimeUsageUnauthorized ", 401)
}

func (o *MonitoringTimeUsageUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/resources/usage][%d] monitoringTimeUsageUnauthorized ", 401)
}

func (o *MonitoringTimeUsageUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewMonitoringTimeUsageForbidden creates a MonitoringTimeUsageForbidden with default headers values
func NewMonitoringTimeUsageForbidden() *MonitoringTimeUsageForbidden {
	return &MonitoringTimeUsageForbidden{}
}

/*
MonitoringTimeUsageForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type MonitoringTimeUsageForbidden struct {
}

// IsSuccess returns true when this monitoring time usage forbidden response has a 2xx status code
func (o *MonitoringTimeUsageForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this monitoring time usage forbidden response has a 3xx status code
func (o *MonitoringTimeUsageForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this monitoring time usage forbidden response has a 4xx status code
func (o *MonitoringTimeUsageForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this monitoring time usage forbidden response has a 5xx status code
func (o *MonitoringTimeUsageForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this monitoring time usage forbidden response a status code equal to that given
func (o *MonitoringTimeUsageForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the monitoring time usage forbidden response
func (o *MonitoringTimeUsageForbidden) Code() int {
	return 403
}

func (o *MonitoringTimeUsageForbidden) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/resources/usage][%d] monitoringTimeUsageForbidden ", 403)
}

func (o *MonitoringTimeUsageForbidden) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/resources/usage][%d] monitoringTimeUsageForbidden ", 403)
}

func (o *MonitoringTimeUsageForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewMonitoringTimeUsageNotFound creates a MonitoringTimeUsageNotFound with default headers values
func NewMonitoringTimeUsageNotFound() *MonitoringTimeUsageNotFound {
	return &MonitoringTimeUsageNotFound{}
}

/*
MonitoringTimeUsageNotFound describes a response with status code 404, with default header values.

Not Found
*/
type MonitoringTimeUsageNotFound struct {
}

// IsSuccess returns true when this monitoring time usage not found response has a 2xx status code
func (o *MonitoringTimeUsageNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this monitoring time usage not found response has a 3xx status code
func (o *MonitoringTimeUsageNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this monitoring time usage not found response has a 4xx status code
func (o *MonitoringTimeUsageNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this monitoring time usage not found response has a 5xx status code
func (o *MonitoringTimeUsageNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this monitoring time usage not found response a status code equal to that given
func (o *MonitoringTimeUsageNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the monitoring time usage not found response
func (o *MonitoringTimeUsageNotFound) Code() int {
	return 404
}

func (o *MonitoringTimeUsageNotFound) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/resources/usage][%d] monitoringTimeUsageNotFound ", 404)
}

func (o *MonitoringTimeUsageNotFound) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/resources/usage][%d] monitoringTimeUsageNotFound ", 404)
}

func (o *MonitoringTimeUsageNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewMonitoringTimeUsageRequestTimeout creates a MonitoringTimeUsageRequestTimeout with default headers values
func NewMonitoringTimeUsageRequestTimeout() *MonitoringTimeUsageRequestTimeout {
	return &MonitoringTimeUsageRequestTimeout{}
}

/*
MonitoringTimeUsageRequestTimeout describes a response with status code 408, with default header values.

Request Timeout
*/
type MonitoringTimeUsageRequestTimeout struct {
}

// IsSuccess returns true when this monitoring time usage request timeout response has a 2xx status code
func (o *MonitoringTimeUsageRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this monitoring time usage request timeout response has a 3xx status code
func (o *MonitoringTimeUsageRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this monitoring time usage request timeout response has a 4xx status code
func (o *MonitoringTimeUsageRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this monitoring time usage request timeout response has a 5xx status code
func (o *MonitoringTimeUsageRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this monitoring time usage request timeout response a status code equal to that given
func (o *MonitoringTimeUsageRequestTimeout) IsCode(code int) bool {
	return code == 408
}

// Code gets the status code for the monitoring time usage request timeout response
func (o *MonitoringTimeUsageRequestTimeout) Code() int {
	return 408
}

func (o *MonitoringTimeUsageRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/resources/usage][%d] monitoringTimeUsageRequestTimeout ", 408)
}

func (o *MonitoringTimeUsageRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/resources/usage][%d] monitoringTimeUsageRequestTimeout ", 408)
}

func (o *MonitoringTimeUsageRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewMonitoringTimeUsageUnprocessableEntity creates a MonitoringTimeUsageUnprocessableEntity with default headers values
func NewMonitoringTimeUsageUnprocessableEntity() *MonitoringTimeUsageUnprocessableEntity {
	return &MonitoringTimeUsageUnprocessableEntity{}
}

/*
MonitoringTimeUsageUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type MonitoringTimeUsageUnprocessableEntity struct {
}

// IsSuccess returns true when this monitoring time usage unprocessable entity response has a 2xx status code
func (o *MonitoringTimeUsageUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this monitoring time usage unprocessable entity response has a 3xx status code
func (o *MonitoringTimeUsageUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this monitoring time usage unprocessable entity response has a 4xx status code
func (o *MonitoringTimeUsageUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this monitoring time usage unprocessable entity response has a 5xx status code
func (o *MonitoringTimeUsageUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this monitoring time usage unprocessable entity response a status code equal to that given
func (o *MonitoringTimeUsageUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the monitoring time usage unprocessable entity response
func (o *MonitoringTimeUsageUnprocessableEntity) Code() int {
	return 422
}

func (o *MonitoringTimeUsageUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/resources/usage][%d] monitoringTimeUsageUnprocessableEntity ", 422)
}

func (o *MonitoringTimeUsageUnprocessableEntity) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/resources/usage][%d] monitoringTimeUsageUnprocessableEntity ", 422)
}

func (o *MonitoringTimeUsageUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}