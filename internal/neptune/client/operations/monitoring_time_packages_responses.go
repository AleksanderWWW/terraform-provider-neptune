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

// MonitoringTimePackagesReader is a Reader for the MonitoringTimePackages structure.
type MonitoringTimePackagesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MonitoringTimePackagesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMonitoringTimePackagesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewMonitoringTimePackagesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewMonitoringTimePackagesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewMonitoringTimePackagesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewMonitoringTimePackagesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewMonitoringTimePackagesRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewMonitoringTimePackagesUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /api/backend/v1/resources/packages] monitoringTimePackages", response, response.Code())
	}
}

// NewMonitoringTimePackagesOK creates a MonitoringTimePackagesOK with default headers values
func NewMonitoringTimePackagesOK() *MonitoringTimePackagesOK {
	return &MonitoringTimePackagesOK{}
}

/*
MonitoringTimePackagesOK describes a response with status code 200, with default header values.

OK
*/
type MonitoringTimePackagesOK struct {
	Payload *models.MonitoringTimePackageListDTO
}

// IsSuccess returns true when this monitoring time packages o k response has a 2xx status code
func (o *MonitoringTimePackagesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this monitoring time packages o k response has a 3xx status code
func (o *MonitoringTimePackagesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this monitoring time packages o k response has a 4xx status code
func (o *MonitoringTimePackagesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this monitoring time packages o k response has a 5xx status code
func (o *MonitoringTimePackagesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this monitoring time packages o k response a status code equal to that given
func (o *MonitoringTimePackagesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the monitoring time packages o k response
func (o *MonitoringTimePackagesOK) Code() int {
	return 200
}

func (o *MonitoringTimePackagesOK) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/resources/packages][%d] monitoringTimePackagesOK  %+v", 200, o.Payload)
}

func (o *MonitoringTimePackagesOK) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/resources/packages][%d] monitoringTimePackagesOK  %+v", 200, o.Payload)
}

func (o *MonitoringTimePackagesOK) GetPayload() *models.MonitoringTimePackageListDTO {
	return o.Payload
}

func (o *MonitoringTimePackagesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MonitoringTimePackageListDTO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMonitoringTimePackagesBadRequest creates a MonitoringTimePackagesBadRequest with default headers values
func NewMonitoringTimePackagesBadRequest() *MonitoringTimePackagesBadRequest {
	return &MonitoringTimePackagesBadRequest{}
}

/*
MonitoringTimePackagesBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type MonitoringTimePackagesBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this monitoring time packages bad request response has a 2xx status code
func (o *MonitoringTimePackagesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this monitoring time packages bad request response has a 3xx status code
func (o *MonitoringTimePackagesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this monitoring time packages bad request response has a 4xx status code
func (o *MonitoringTimePackagesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this monitoring time packages bad request response has a 5xx status code
func (o *MonitoringTimePackagesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this monitoring time packages bad request response a status code equal to that given
func (o *MonitoringTimePackagesBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the monitoring time packages bad request response
func (o *MonitoringTimePackagesBadRequest) Code() int {
	return 400
}

func (o *MonitoringTimePackagesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/resources/packages][%d] monitoringTimePackagesBadRequest  %+v", 400, o.Payload)
}

func (o *MonitoringTimePackagesBadRequest) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/resources/packages][%d] monitoringTimePackagesBadRequest  %+v", 400, o.Payload)
}

func (o *MonitoringTimePackagesBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *MonitoringTimePackagesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMonitoringTimePackagesUnauthorized creates a MonitoringTimePackagesUnauthorized with default headers values
func NewMonitoringTimePackagesUnauthorized() *MonitoringTimePackagesUnauthorized {
	return &MonitoringTimePackagesUnauthorized{}
}

/*
MonitoringTimePackagesUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type MonitoringTimePackagesUnauthorized struct {
}

// IsSuccess returns true when this monitoring time packages unauthorized response has a 2xx status code
func (o *MonitoringTimePackagesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this monitoring time packages unauthorized response has a 3xx status code
func (o *MonitoringTimePackagesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this monitoring time packages unauthorized response has a 4xx status code
func (o *MonitoringTimePackagesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this monitoring time packages unauthorized response has a 5xx status code
func (o *MonitoringTimePackagesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this monitoring time packages unauthorized response a status code equal to that given
func (o *MonitoringTimePackagesUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the monitoring time packages unauthorized response
func (o *MonitoringTimePackagesUnauthorized) Code() int {
	return 401
}

func (o *MonitoringTimePackagesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/resources/packages][%d] monitoringTimePackagesUnauthorized ", 401)
}

func (o *MonitoringTimePackagesUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/resources/packages][%d] monitoringTimePackagesUnauthorized ", 401)
}

func (o *MonitoringTimePackagesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewMonitoringTimePackagesForbidden creates a MonitoringTimePackagesForbidden with default headers values
func NewMonitoringTimePackagesForbidden() *MonitoringTimePackagesForbidden {
	return &MonitoringTimePackagesForbidden{}
}

/*
MonitoringTimePackagesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type MonitoringTimePackagesForbidden struct {
}

// IsSuccess returns true when this monitoring time packages forbidden response has a 2xx status code
func (o *MonitoringTimePackagesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this monitoring time packages forbidden response has a 3xx status code
func (o *MonitoringTimePackagesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this monitoring time packages forbidden response has a 4xx status code
func (o *MonitoringTimePackagesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this monitoring time packages forbidden response has a 5xx status code
func (o *MonitoringTimePackagesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this monitoring time packages forbidden response a status code equal to that given
func (o *MonitoringTimePackagesForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the monitoring time packages forbidden response
func (o *MonitoringTimePackagesForbidden) Code() int {
	return 403
}

func (o *MonitoringTimePackagesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/resources/packages][%d] monitoringTimePackagesForbidden ", 403)
}

func (o *MonitoringTimePackagesForbidden) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/resources/packages][%d] monitoringTimePackagesForbidden ", 403)
}

func (o *MonitoringTimePackagesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewMonitoringTimePackagesNotFound creates a MonitoringTimePackagesNotFound with default headers values
func NewMonitoringTimePackagesNotFound() *MonitoringTimePackagesNotFound {
	return &MonitoringTimePackagesNotFound{}
}

/*
MonitoringTimePackagesNotFound describes a response with status code 404, with default header values.

Not Found
*/
type MonitoringTimePackagesNotFound struct {
}

// IsSuccess returns true when this monitoring time packages not found response has a 2xx status code
func (o *MonitoringTimePackagesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this monitoring time packages not found response has a 3xx status code
func (o *MonitoringTimePackagesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this monitoring time packages not found response has a 4xx status code
func (o *MonitoringTimePackagesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this monitoring time packages not found response has a 5xx status code
func (o *MonitoringTimePackagesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this monitoring time packages not found response a status code equal to that given
func (o *MonitoringTimePackagesNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the monitoring time packages not found response
func (o *MonitoringTimePackagesNotFound) Code() int {
	return 404
}

func (o *MonitoringTimePackagesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/resources/packages][%d] monitoringTimePackagesNotFound ", 404)
}

func (o *MonitoringTimePackagesNotFound) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/resources/packages][%d] monitoringTimePackagesNotFound ", 404)
}

func (o *MonitoringTimePackagesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewMonitoringTimePackagesRequestTimeout creates a MonitoringTimePackagesRequestTimeout with default headers values
func NewMonitoringTimePackagesRequestTimeout() *MonitoringTimePackagesRequestTimeout {
	return &MonitoringTimePackagesRequestTimeout{}
}

/*
MonitoringTimePackagesRequestTimeout describes a response with status code 408, with default header values.

Request Timeout
*/
type MonitoringTimePackagesRequestTimeout struct {
}

// IsSuccess returns true when this monitoring time packages request timeout response has a 2xx status code
func (o *MonitoringTimePackagesRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this monitoring time packages request timeout response has a 3xx status code
func (o *MonitoringTimePackagesRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this monitoring time packages request timeout response has a 4xx status code
func (o *MonitoringTimePackagesRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this monitoring time packages request timeout response has a 5xx status code
func (o *MonitoringTimePackagesRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this monitoring time packages request timeout response a status code equal to that given
func (o *MonitoringTimePackagesRequestTimeout) IsCode(code int) bool {
	return code == 408
}

// Code gets the status code for the monitoring time packages request timeout response
func (o *MonitoringTimePackagesRequestTimeout) Code() int {
	return 408
}

func (o *MonitoringTimePackagesRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/resources/packages][%d] monitoringTimePackagesRequestTimeout ", 408)
}

func (o *MonitoringTimePackagesRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/resources/packages][%d] monitoringTimePackagesRequestTimeout ", 408)
}

func (o *MonitoringTimePackagesRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewMonitoringTimePackagesUnprocessableEntity creates a MonitoringTimePackagesUnprocessableEntity with default headers values
func NewMonitoringTimePackagesUnprocessableEntity() *MonitoringTimePackagesUnprocessableEntity {
	return &MonitoringTimePackagesUnprocessableEntity{}
}

/*
MonitoringTimePackagesUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type MonitoringTimePackagesUnprocessableEntity struct {
}

// IsSuccess returns true when this monitoring time packages unprocessable entity response has a 2xx status code
func (o *MonitoringTimePackagesUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this monitoring time packages unprocessable entity response has a 3xx status code
func (o *MonitoringTimePackagesUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this monitoring time packages unprocessable entity response has a 4xx status code
func (o *MonitoringTimePackagesUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this monitoring time packages unprocessable entity response has a 5xx status code
func (o *MonitoringTimePackagesUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this monitoring time packages unprocessable entity response a status code equal to that given
func (o *MonitoringTimePackagesUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the monitoring time packages unprocessable entity response
func (o *MonitoringTimePackagesUnprocessableEntity) Code() int {
	return 422
}

func (o *MonitoringTimePackagesUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/resources/packages][%d] monitoringTimePackagesUnprocessableEntity ", 422)
}

func (o *MonitoringTimePackagesUnprocessableEntity) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/resources/packages][%d] monitoringTimePackagesUnprocessableEntity ", 422)
}

func (o *MonitoringTimePackagesUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
