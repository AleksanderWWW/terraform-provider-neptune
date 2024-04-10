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

// ProjectMonitoringTimeUsageReader is a Reader for the ProjectMonitoringTimeUsage structure.
type ProjectMonitoringTimeUsageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectMonitoringTimeUsageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectMonitoringTimeUsageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewProjectMonitoringTimeUsageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewProjectMonitoringTimeUsageUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewProjectMonitoringTimeUsageForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewProjectMonitoringTimeUsageNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewProjectMonitoringTimeUsageRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewProjectMonitoringTimeUsageUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /api/backend/v1/resources/projectUsageHistory] projectMonitoringTimeUsage", response, response.Code())
	}
}

// NewProjectMonitoringTimeUsageOK creates a ProjectMonitoringTimeUsageOK with default headers values
func NewProjectMonitoringTimeUsageOK() *ProjectMonitoringTimeUsageOK {
	return &ProjectMonitoringTimeUsageOK{}
}

/*
ProjectMonitoringTimeUsageOK describes a response with status code 200, with default header values.

OK
*/
type ProjectMonitoringTimeUsageOK struct {
	Payload *models.MonitoringTimeHistoryDTO
}

// IsSuccess returns true when this project monitoring time usage o k response has a 2xx status code
func (o *ProjectMonitoringTimeUsageOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this project monitoring time usage o k response has a 3xx status code
func (o *ProjectMonitoringTimeUsageOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project monitoring time usage o k response has a 4xx status code
func (o *ProjectMonitoringTimeUsageOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this project monitoring time usage o k response has a 5xx status code
func (o *ProjectMonitoringTimeUsageOK) IsServerError() bool {
	return false
}

// IsCode returns true when this project monitoring time usage o k response a status code equal to that given
func (o *ProjectMonitoringTimeUsageOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the project monitoring time usage o k response
func (o *ProjectMonitoringTimeUsageOK) Code() int {
	return 200
}

func (o *ProjectMonitoringTimeUsageOK) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/resources/projectUsageHistory][%d] projectMonitoringTimeUsageOK  %+v", 200, o.Payload)
}

func (o *ProjectMonitoringTimeUsageOK) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/resources/projectUsageHistory][%d] projectMonitoringTimeUsageOK  %+v", 200, o.Payload)
}

func (o *ProjectMonitoringTimeUsageOK) GetPayload() *models.MonitoringTimeHistoryDTO {
	return o.Payload
}

func (o *ProjectMonitoringTimeUsageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MonitoringTimeHistoryDTO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectMonitoringTimeUsageBadRequest creates a ProjectMonitoringTimeUsageBadRequest with default headers values
func NewProjectMonitoringTimeUsageBadRequest() *ProjectMonitoringTimeUsageBadRequest {
	return &ProjectMonitoringTimeUsageBadRequest{}
}

/*
ProjectMonitoringTimeUsageBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ProjectMonitoringTimeUsageBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this project monitoring time usage bad request response has a 2xx status code
func (o *ProjectMonitoringTimeUsageBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this project monitoring time usage bad request response has a 3xx status code
func (o *ProjectMonitoringTimeUsageBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project monitoring time usage bad request response has a 4xx status code
func (o *ProjectMonitoringTimeUsageBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this project monitoring time usage bad request response has a 5xx status code
func (o *ProjectMonitoringTimeUsageBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this project monitoring time usage bad request response a status code equal to that given
func (o *ProjectMonitoringTimeUsageBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the project monitoring time usage bad request response
func (o *ProjectMonitoringTimeUsageBadRequest) Code() int {
	return 400
}

func (o *ProjectMonitoringTimeUsageBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/resources/projectUsageHistory][%d] projectMonitoringTimeUsageBadRequest  %+v", 400, o.Payload)
}

func (o *ProjectMonitoringTimeUsageBadRequest) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/resources/projectUsageHistory][%d] projectMonitoringTimeUsageBadRequest  %+v", 400, o.Payload)
}

func (o *ProjectMonitoringTimeUsageBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ProjectMonitoringTimeUsageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectMonitoringTimeUsageUnauthorized creates a ProjectMonitoringTimeUsageUnauthorized with default headers values
func NewProjectMonitoringTimeUsageUnauthorized() *ProjectMonitoringTimeUsageUnauthorized {
	return &ProjectMonitoringTimeUsageUnauthorized{}
}

/*
ProjectMonitoringTimeUsageUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ProjectMonitoringTimeUsageUnauthorized struct {
}

// IsSuccess returns true when this project monitoring time usage unauthorized response has a 2xx status code
func (o *ProjectMonitoringTimeUsageUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this project monitoring time usage unauthorized response has a 3xx status code
func (o *ProjectMonitoringTimeUsageUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project monitoring time usage unauthorized response has a 4xx status code
func (o *ProjectMonitoringTimeUsageUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this project monitoring time usage unauthorized response has a 5xx status code
func (o *ProjectMonitoringTimeUsageUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this project monitoring time usage unauthorized response a status code equal to that given
func (o *ProjectMonitoringTimeUsageUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the project monitoring time usage unauthorized response
func (o *ProjectMonitoringTimeUsageUnauthorized) Code() int {
	return 401
}

func (o *ProjectMonitoringTimeUsageUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/resources/projectUsageHistory][%d] projectMonitoringTimeUsageUnauthorized ", 401)
}

func (o *ProjectMonitoringTimeUsageUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/resources/projectUsageHistory][%d] projectMonitoringTimeUsageUnauthorized ", 401)
}

func (o *ProjectMonitoringTimeUsageUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewProjectMonitoringTimeUsageForbidden creates a ProjectMonitoringTimeUsageForbidden with default headers values
func NewProjectMonitoringTimeUsageForbidden() *ProjectMonitoringTimeUsageForbidden {
	return &ProjectMonitoringTimeUsageForbidden{}
}

/*
ProjectMonitoringTimeUsageForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ProjectMonitoringTimeUsageForbidden struct {
}

// IsSuccess returns true when this project monitoring time usage forbidden response has a 2xx status code
func (o *ProjectMonitoringTimeUsageForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this project monitoring time usage forbidden response has a 3xx status code
func (o *ProjectMonitoringTimeUsageForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project monitoring time usage forbidden response has a 4xx status code
func (o *ProjectMonitoringTimeUsageForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this project monitoring time usage forbidden response has a 5xx status code
func (o *ProjectMonitoringTimeUsageForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this project monitoring time usage forbidden response a status code equal to that given
func (o *ProjectMonitoringTimeUsageForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the project monitoring time usage forbidden response
func (o *ProjectMonitoringTimeUsageForbidden) Code() int {
	return 403
}

func (o *ProjectMonitoringTimeUsageForbidden) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/resources/projectUsageHistory][%d] projectMonitoringTimeUsageForbidden ", 403)
}

func (o *ProjectMonitoringTimeUsageForbidden) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/resources/projectUsageHistory][%d] projectMonitoringTimeUsageForbidden ", 403)
}

func (o *ProjectMonitoringTimeUsageForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewProjectMonitoringTimeUsageNotFound creates a ProjectMonitoringTimeUsageNotFound with default headers values
func NewProjectMonitoringTimeUsageNotFound() *ProjectMonitoringTimeUsageNotFound {
	return &ProjectMonitoringTimeUsageNotFound{}
}

/*
ProjectMonitoringTimeUsageNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ProjectMonitoringTimeUsageNotFound struct {
}

// IsSuccess returns true when this project monitoring time usage not found response has a 2xx status code
func (o *ProjectMonitoringTimeUsageNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this project monitoring time usage not found response has a 3xx status code
func (o *ProjectMonitoringTimeUsageNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project monitoring time usage not found response has a 4xx status code
func (o *ProjectMonitoringTimeUsageNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this project monitoring time usage not found response has a 5xx status code
func (o *ProjectMonitoringTimeUsageNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this project monitoring time usage not found response a status code equal to that given
func (o *ProjectMonitoringTimeUsageNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the project monitoring time usage not found response
func (o *ProjectMonitoringTimeUsageNotFound) Code() int {
	return 404
}

func (o *ProjectMonitoringTimeUsageNotFound) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/resources/projectUsageHistory][%d] projectMonitoringTimeUsageNotFound ", 404)
}

func (o *ProjectMonitoringTimeUsageNotFound) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/resources/projectUsageHistory][%d] projectMonitoringTimeUsageNotFound ", 404)
}

func (o *ProjectMonitoringTimeUsageNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewProjectMonitoringTimeUsageRequestTimeout creates a ProjectMonitoringTimeUsageRequestTimeout with default headers values
func NewProjectMonitoringTimeUsageRequestTimeout() *ProjectMonitoringTimeUsageRequestTimeout {
	return &ProjectMonitoringTimeUsageRequestTimeout{}
}

/*
ProjectMonitoringTimeUsageRequestTimeout describes a response with status code 408, with default header values.

Request Timeout
*/
type ProjectMonitoringTimeUsageRequestTimeout struct {
}

// IsSuccess returns true when this project monitoring time usage request timeout response has a 2xx status code
func (o *ProjectMonitoringTimeUsageRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this project monitoring time usage request timeout response has a 3xx status code
func (o *ProjectMonitoringTimeUsageRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project monitoring time usage request timeout response has a 4xx status code
func (o *ProjectMonitoringTimeUsageRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this project monitoring time usage request timeout response has a 5xx status code
func (o *ProjectMonitoringTimeUsageRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this project monitoring time usage request timeout response a status code equal to that given
func (o *ProjectMonitoringTimeUsageRequestTimeout) IsCode(code int) bool {
	return code == 408
}

// Code gets the status code for the project monitoring time usage request timeout response
func (o *ProjectMonitoringTimeUsageRequestTimeout) Code() int {
	return 408
}

func (o *ProjectMonitoringTimeUsageRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/resources/projectUsageHistory][%d] projectMonitoringTimeUsageRequestTimeout ", 408)
}

func (o *ProjectMonitoringTimeUsageRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/resources/projectUsageHistory][%d] projectMonitoringTimeUsageRequestTimeout ", 408)
}

func (o *ProjectMonitoringTimeUsageRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewProjectMonitoringTimeUsageUnprocessableEntity creates a ProjectMonitoringTimeUsageUnprocessableEntity with default headers values
func NewProjectMonitoringTimeUsageUnprocessableEntity() *ProjectMonitoringTimeUsageUnprocessableEntity {
	return &ProjectMonitoringTimeUsageUnprocessableEntity{}
}

/*
ProjectMonitoringTimeUsageUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type ProjectMonitoringTimeUsageUnprocessableEntity struct {
}

// IsSuccess returns true when this project monitoring time usage unprocessable entity response has a 2xx status code
func (o *ProjectMonitoringTimeUsageUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this project monitoring time usage unprocessable entity response has a 3xx status code
func (o *ProjectMonitoringTimeUsageUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project monitoring time usage unprocessable entity response has a 4xx status code
func (o *ProjectMonitoringTimeUsageUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this project monitoring time usage unprocessable entity response has a 5xx status code
func (o *ProjectMonitoringTimeUsageUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this project monitoring time usage unprocessable entity response a status code equal to that given
func (o *ProjectMonitoringTimeUsageUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the project monitoring time usage unprocessable entity response
func (o *ProjectMonitoringTimeUsageUnprocessableEntity) Code() int {
	return 422
}

func (o *ProjectMonitoringTimeUsageUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/resources/projectUsageHistory][%d] projectMonitoringTimeUsageUnprocessableEntity ", 422)
}

func (o *ProjectMonitoringTimeUsageUnprocessableEntity) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/resources/projectUsageHistory][%d] projectMonitoringTimeUsageUnprocessableEntity ", 422)
}

func (o *ProjectMonitoringTimeUsageUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}