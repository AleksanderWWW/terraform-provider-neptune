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

// MonitoringTimeTopProjects30DaysReader is a Reader for the MonitoringTimeTopProjects30Days structure.
type MonitoringTimeTopProjects30DaysReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MonitoringTimeTopProjects30DaysReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMonitoringTimeTopProjects30DaysOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewMonitoringTimeTopProjects30DaysBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewMonitoringTimeTopProjects30DaysUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewMonitoringTimeTopProjects30DaysForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewMonitoringTimeTopProjects30DaysNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewMonitoringTimeTopProjects30DaysRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewMonitoringTimeTopProjects30DaysUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /api/backend/v1/resources/topProjects30Days] monitoringTimeTopProjects30Days", response, response.Code())
	}
}

// NewMonitoringTimeTopProjects30DaysOK creates a MonitoringTimeTopProjects30DaysOK with default headers values
func NewMonitoringTimeTopProjects30DaysOK() *MonitoringTimeTopProjects30DaysOK {
	return &MonitoringTimeTopProjects30DaysOK{}
}

/*
MonitoringTimeTopProjects30DaysOK describes a response with status code 200, with default header values.

OK
*/
type MonitoringTimeTopProjects30DaysOK struct {
	Payload *models.ProjectUsageListDTO
}

// IsSuccess returns true when this monitoring time top projects30 days o k response has a 2xx status code
func (o *MonitoringTimeTopProjects30DaysOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this monitoring time top projects30 days o k response has a 3xx status code
func (o *MonitoringTimeTopProjects30DaysOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this monitoring time top projects30 days o k response has a 4xx status code
func (o *MonitoringTimeTopProjects30DaysOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this monitoring time top projects30 days o k response has a 5xx status code
func (o *MonitoringTimeTopProjects30DaysOK) IsServerError() bool {
	return false
}

// IsCode returns true when this monitoring time top projects30 days o k response a status code equal to that given
func (o *MonitoringTimeTopProjects30DaysOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the monitoring time top projects30 days o k response
func (o *MonitoringTimeTopProjects30DaysOK) Code() int {
	return 200
}

func (o *MonitoringTimeTopProjects30DaysOK) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/resources/topProjects30Days][%d] monitoringTimeTopProjects30DaysOK  %+v", 200, o.Payload)
}

func (o *MonitoringTimeTopProjects30DaysOK) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/resources/topProjects30Days][%d] monitoringTimeTopProjects30DaysOK  %+v", 200, o.Payload)
}

func (o *MonitoringTimeTopProjects30DaysOK) GetPayload() *models.ProjectUsageListDTO {
	return o.Payload
}

func (o *MonitoringTimeTopProjects30DaysOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProjectUsageListDTO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMonitoringTimeTopProjects30DaysBadRequest creates a MonitoringTimeTopProjects30DaysBadRequest with default headers values
func NewMonitoringTimeTopProjects30DaysBadRequest() *MonitoringTimeTopProjects30DaysBadRequest {
	return &MonitoringTimeTopProjects30DaysBadRequest{}
}

/*
MonitoringTimeTopProjects30DaysBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type MonitoringTimeTopProjects30DaysBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this monitoring time top projects30 days bad request response has a 2xx status code
func (o *MonitoringTimeTopProjects30DaysBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this monitoring time top projects30 days bad request response has a 3xx status code
func (o *MonitoringTimeTopProjects30DaysBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this monitoring time top projects30 days bad request response has a 4xx status code
func (o *MonitoringTimeTopProjects30DaysBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this monitoring time top projects30 days bad request response has a 5xx status code
func (o *MonitoringTimeTopProjects30DaysBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this monitoring time top projects30 days bad request response a status code equal to that given
func (o *MonitoringTimeTopProjects30DaysBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the monitoring time top projects30 days bad request response
func (o *MonitoringTimeTopProjects30DaysBadRequest) Code() int {
	return 400
}

func (o *MonitoringTimeTopProjects30DaysBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/resources/topProjects30Days][%d] monitoringTimeTopProjects30DaysBadRequest  %+v", 400, o.Payload)
}

func (o *MonitoringTimeTopProjects30DaysBadRequest) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/resources/topProjects30Days][%d] monitoringTimeTopProjects30DaysBadRequest  %+v", 400, o.Payload)
}

func (o *MonitoringTimeTopProjects30DaysBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *MonitoringTimeTopProjects30DaysBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMonitoringTimeTopProjects30DaysUnauthorized creates a MonitoringTimeTopProjects30DaysUnauthorized with default headers values
func NewMonitoringTimeTopProjects30DaysUnauthorized() *MonitoringTimeTopProjects30DaysUnauthorized {
	return &MonitoringTimeTopProjects30DaysUnauthorized{}
}

/*
MonitoringTimeTopProjects30DaysUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type MonitoringTimeTopProjects30DaysUnauthorized struct {
}

// IsSuccess returns true when this monitoring time top projects30 days unauthorized response has a 2xx status code
func (o *MonitoringTimeTopProjects30DaysUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this monitoring time top projects30 days unauthorized response has a 3xx status code
func (o *MonitoringTimeTopProjects30DaysUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this monitoring time top projects30 days unauthorized response has a 4xx status code
func (o *MonitoringTimeTopProjects30DaysUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this monitoring time top projects30 days unauthorized response has a 5xx status code
func (o *MonitoringTimeTopProjects30DaysUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this monitoring time top projects30 days unauthorized response a status code equal to that given
func (o *MonitoringTimeTopProjects30DaysUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the monitoring time top projects30 days unauthorized response
func (o *MonitoringTimeTopProjects30DaysUnauthorized) Code() int {
	return 401
}

func (o *MonitoringTimeTopProjects30DaysUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/resources/topProjects30Days][%d] monitoringTimeTopProjects30DaysUnauthorized ", 401)
}

func (o *MonitoringTimeTopProjects30DaysUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/resources/topProjects30Days][%d] monitoringTimeTopProjects30DaysUnauthorized ", 401)
}

func (o *MonitoringTimeTopProjects30DaysUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewMonitoringTimeTopProjects30DaysForbidden creates a MonitoringTimeTopProjects30DaysForbidden with default headers values
func NewMonitoringTimeTopProjects30DaysForbidden() *MonitoringTimeTopProjects30DaysForbidden {
	return &MonitoringTimeTopProjects30DaysForbidden{}
}

/*
MonitoringTimeTopProjects30DaysForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type MonitoringTimeTopProjects30DaysForbidden struct {
}

// IsSuccess returns true when this monitoring time top projects30 days forbidden response has a 2xx status code
func (o *MonitoringTimeTopProjects30DaysForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this monitoring time top projects30 days forbidden response has a 3xx status code
func (o *MonitoringTimeTopProjects30DaysForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this monitoring time top projects30 days forbidden response has a 4xx status code
func (o *MonitoringTimeTopProjects30DaysForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this monitoring time top projects30 days forbidden response has a 5xx status code
func (o *MonitoringTimeTopProjects30DaysForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this monitoring time top projects30 days forbidden response a status code equal to that given
func (o *MonitoringTimeTopProjects30DaysForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the monitoring time top projects30 days forbidden response
func (o *MonitoringTimeTopProjects30DaysForbidden) Code() int {
	return 403
}

func (o *MonitoringTimeTopProjects30DaysForbidden) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/resources/topProjects30Days][%d] monitoringTimeTopProjects30DaysForbidden ", 403)
}

func (o *MonitoringTimeTopProjects30DaysForbidden) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/resources/topProjects30Days][%d] monitoringTimeTopProjects30DaysForbidden ", 403)
}

func (o *MonitoringTimeTopProjects30DaysForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewMonitoringTimeTopProjects30DaysNotFound creates a MonitoringTimeTopProjects30DaysNotFound with default headers values
func NewMonitoringTimeTopProjects30DaysNotFound() *MonitoringTimeTopProjects30DaysNotFound {
	return &MonitoringTimeTopProjects30DaysNotFound{}
}

/*
MonitoringTimeTopProjects30DaysNotFound describes a response with status code 404, with default header values.

Not Found
*/
type MonitoringTimeTopProjects30DaysNotFound struct {
}

// IsSuccess returns true when this monitoring time top projects30 days not found response has a 2xx status code
func (o *MonitoringTimeTopProjects30DaysNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this monitoring time top projects30 days not found response has a 3xx status code
func (o *MonitoringTimeTopProjects30DaysNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this monitoring time top projects30 days not found response has a 4xx status code
func (o *MonitoringTimeTopProjects30DaysNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this monitoring time top projects30 days not found response has a 5xx status code
func (o *MonitoringTimeTopProjects30DaysNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this monitoring time top projects30 days not found response a status code equal to that given
func (o *MonitoringTimeTopProjects30DaysNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the monitoring time top projects30 days not found response
func (o *MonitoringTimeTopProjects30DaysNotFound) Code() int {
	return 404
}

func (o *MonitoringTimeTopProjects30DaysNotFound) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/resources/topProjects30Days][%d] monitoringTimeTopProjects30DaysNotFound ", 404)
}

func (o *MonitoringTimeTopProjects30DaysNotFound) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/resources/topProjects30Days][%d] monitoringTimeTopProjects30DaysNotFound ", 404)
}

func (o *MonitoringTimeTopProjects30DaysNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewMonitoringTimeTopProjects30DaysRequestTimeout creates a MonitoringTimeTopProjects30DaysRequestTimeout with default headers values
func NewMonitoringTimeTopProjects30DaysRequestTimeout() *MonitoringTimeTopProjects30DaysRequestTimeout {
	return &MonitoringTimeTopProjects30DaysRequestTimeout{}
}

/*
MonitoringTimeTopProjects30DaysRequestTimeout describes a response with status code 408, with default header values.

Request Timeout
*/
type MonitoringTimeTopProjects30DaysRequestTimeout struct {
}

// IsSuccess returns true when this monitoring time top projects30 days request timeout response has a 2xx status code
func (o *MonitoringTimeTopProjects30DaysRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this monitoring time top projects30 days request timeout response has a 3xx status code
func (o *MonitoringTimeTopProjects30DaysRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this monitoring time top projects30 days request timeout response has a 4xx status code
func (o *MonitoringTimeTopProjects30DaysRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this monitoring time top projects30 days request timeout response has a 5xx status code
func (o *MonitoringTimeTopProjects30DaysRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this monitoring time top projects30 days request timeout response a status code equal to that given
func (o *MonitoringTimeTopProjects30DaysRequestTimeout) IsCode(code int) bool {
	return code == 408
}

// Code gets the status code for the monitoring time top projects30 days request timeout response
func (o *MonitoringTimeTopProjects30DaysRequestTimeout) Code() int {
	return 408
}

func (o *MonitoringTimeTopProjects30DaysRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/resources/topProjects30Days][%d] monitoringTimeTopProjects30DaysRequestTimeout ", 408)
}

func (o *MonitoringTimeTopProjects30DaysRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/resources/topProjects30Days][%d] monitoringTimeTopProjects30DaysRequestTimeout ", 408)
}

func (o *MonitoringTimeTopProjects30DaysRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewMonitoringTimeTopProjects30DaysUnprocessableEntity creates a MonitoringTimeTopProjects30DaysUnprocessableEntity with default headers values
func NewMonitoringTimeTopProjects30DaysUnprocessableEntity() *MonitoringTimeTopProjects30DaysUnprocessableEntity {
	return &MonitoringTimeTopProjects30DaysUnprocessableEntity{}
}

/*
MonitoringTimeTopProjects30DaysUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type MonitoringTimeTopProjects30DaysUnprocessableEntity struct {
}

// IsSuccess returns true when this monitoring time top projects30 days unprocessable entity response has a 2xx status code
func (o *MonitoringTimeTopProjects30DaysUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this monitoring time top projects30 days unprocessable entity response has a 3xx status code
func (o *MonitoringTimeTopProjects30DaysUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this monitoring time top projects30 days unprocessable entity response has a 4xx status code
func (o *MonitoringTimeTopProjects30DaysUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this monitoring time top projects30 days unprocessable entity response has a 5xx status code
func (o *MonitoringTimeTopProjects30DaysUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this monitoring time top projects30 days unprocessable entity response a status code equal to that given
func (o *MonitoringTimeTopProjects30DaysUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the monitoring time top projects30 days unprocessable entity response
func (o *MonitoringTimeTopProjects30DaysUnprocessableEntity) Code() int {
	return 422
}

func (o *MonitoringTimeTopProjects30DaysUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/resources/topProjects30Days][%d] monitoringTimeTopProjects30DaysUnprocessableEntity ", 422)
}

func (o *MonitoringTimeTopProjects30DaysUnprocessableEntity) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/resources/topProjects30Days][%d] monitoringTimeTopProjects30DaysUnprocessableEntity ", 422)
}

func (o *MonitoringTimeTopProjects30DaysUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
