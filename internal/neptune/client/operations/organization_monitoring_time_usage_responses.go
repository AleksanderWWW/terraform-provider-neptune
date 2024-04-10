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

// OrganizationMonitoringTimeUsageReader is a Reader for the OrganizationMonitoringTimeUsage structure.
type OrganizationMonitoringTimeUsageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OrganizationMonitoringTimeUsageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOrganizationMonitoringTimeUsageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewOrganizationMonitoringTimeUsageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewOrganizationMonitoringTimeUsageUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewOrganizationMonitoringTimeUsageForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewOrganizationMonitoringTimeUsageNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewOrganizationMonitoringTimeUsageRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewOrganizationMonitoringTimeUsageUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /api/backend/v1/resources/organizationUsageHistory] organizationMonitoringTimeUsage", response, response.Code())
	}
}

// NewOrganizationMonitoringTimeUsageOK creates a OrganizationMonitoringTimeUsageOK with default headers values
func NewOrganizationMonitoringTimeUsageOK() *OrganizationMonitoringTimeUsageOK {
	return &OrganizationMonitoringTimeUsageOK{}
}

/*
OrganizationMonitoringTimeUsageOK describes a response with status code 200, with default header values.

OK
*/
type OrganizationMonitoringTimeUsageOK struct {
	Payload *models.MonitoringTimeAndPackageHistoryDTO
}

// IsSuccess returns true when this organization monitoring time usage o k response has a 2xx status code
func (o *OrganizationMonitoringTimeUsageOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this organization monitoring time usage o k response has a 3xx status code
func (o *OrganizationMonitoringTimeUsageOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this organization monitoring time usage o k response has a 4xx status code
func (o *OrganizationMonitoringTimeUsageOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this organization monitoring time usage o k response has a 5xx status code
func (o *OrganizationMonitoringTimeUsageOK) IsServerError() bool {
	return false
}

// IsCode returns true when this organization monitoring time usage o k response a status code equal to that given
func (o *OrganizationMonitoringTimeUsageOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the organization monitoring time usage o k response
func (o *OrganizationMonitoringTimeUsageOK) Code() int {
	return 200
}

func (o *OrganizationMonitoringTimeUsageOK) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/resources/organizationUsageHistory][%d] organizationMonitoringTimeUsageOK  %+v", 200, o.Payload)
}

func (o *OrganizationMonitoringTimeUsageOK) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/resources/organizationUsageHistory][%d] organizationMonitoringTimeUsageOK  %+v", 200, o.Payload)
}

func (o *OrganizationMonitoringTimeUsageOK) GetPayload() *models.MonitoringTimeAndPackageHistoryDTO {
	return o.Payload
}

func (o *OrganizationMonitoringTimeUsageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MonitoringTimeAndPackageHistoryDTO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrganizationMonitoringTimeUsageBadRequest creates a OrganizationMonitoringTimeUsageBadRequest with default headers values
func NewOrganizationMonitoringTimeUsageBadRequest() *OrganizationMonitoringTimeUsageBadRequest {
	return &OrganizationMonitoringTimeUsageBadRequest{}
}

/*
OrganizationMonitoringTimeUsageBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type OrganizationMonitoringTimeUsageBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this organization monitoring time usage bad request response has a 2xx status code
func (o *OrganizationMonitoringTimeUsageBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this organization monitoring time usage bad request response has a 3xx status code
func (o *OrganizationMonitoringTimeUsageBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this organization monitoring time usage bad request response has a 4xx status code
func (o *OrganizationMonitoringTimeUsageBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this organization monitoring time usage bad request response has a 5xx status code
func (o *OrganizationMonitoringTimeUsageBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this organization monitoring time usage bad request response a status code equal to that given
func (o *OrganizationMonitoringTimeUsageBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the organization monitoring time usage bad request response
func (o *OrganizationMonitoringTimeUsageBadRequest) Code() int {
	return 400
}

func (o *OrganizationMonitoringTimeUsageBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/resources/organizationUsageHistory][%d] organizationMonitoringTimeUsageBadRequest  %+v", 400, o.Payload)
}

func (o *OrganizationMonitoringTimeUsageBadRequest) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/resources/organizationUsageHistory][%d] organizationMonitoringTimeUsageBadRequest  %+v", 400, o.Payload)
}

func (o *OrganizationMonitoringTimeUsageBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *OrganizationMonitoringTimeUsageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrganizationMonitoringTimeUsageUnauthorized creates a OrganizationMonitoringTimeUsageUnauthorized with default headers values
func NewOrganizationMonitoringTimeUsageUnauthorized() *OrganizationMonitoringTimeUsageUnauthorized {
	return &OrganizationMonitoringTimeUsageUnauthorized{}
}

/*
OrganizationMonitoringTimeUsageUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type OrganizationMonitoringTimeUsageUnauthorized struct {
}

// IsSuccess returns true when this organization monitoring time usage unauthorized response has a 2xx status code
func (o *OrganizationMonitoringTimeUsageUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this organization monitoring time usage unauthorized response has a 3xx status code
func (o *OrganizationMonitoringTimeUsageUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this organization monitoring time usage unauthorized response has a 4xx status code
func (o *OrganizationMonitoringTimeUsageUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this organization monitoring time usage unauthorized response has a 5xx status code
func (o *OrganizationMonitoringTimeUsageUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this organization monitoring time usage unauthorized response a status code equal to that given
func (o *OrganizationMonitoringTimeUsageUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the organization monitoring time usage unauthorized response
func (o *OrganizationMonitoringTimeUsageUnauthorized) Code() int {
	return 401
}

func (o *OrganizationMonitoringTimeUsageUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/resources/organizationUsageHistory][%d] organizationMonitoringTimeUsageUnauthorized ", 401)
}

func (o *OrganizationMonitoringTimeUsageUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/resources/organizationUsageHistory][%d] organizationMonitoringTimeUsageUnauthorized ", 401)
}

func (o *OrganizationMonitoringTimeUsageUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewOrganizationMonitoringTimeUsageForbidden creates a OrganizationMonitoringTimeUsageForbidden with default headers values
func NewOrganizationMonitoringTimeUsageForbidden() *OrganizationMonitoringTimeUsageForbidden {
	return &OrganizationMonitoringTimeUsageForbidden{}
}

/*
OrganizationMonitoringTimeUsageForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type OrganizationMonitoringTimeUsageForbidden struct {
}

// IsSuccess returns true when this organization monitoring time usage forbidden response has a 2xx status code
func (o *OrganizationMonitoringTimeUsageForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this organization monitoring time usage forbidden response has a 3xx status code
func (o *OrganizationMonitoringTimeUsageForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this organization monitoring time usage forbidden response has a 4xx status code
func (o *OrganizationMonitoringTimeUsageForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this organization monitoring time usage forbidden response has a 5xx status code
func (o *OrganizationMonitoringTimeUsageForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this organization monitoring time usage forbidden response a status code equal to that given
func (o *OrganizationMonitoringTimeUsageForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the organization monitoring time usage forbidden response
func (o *OrganizationMonitoringTimeUsageForbidden) Code() int {
	return 403
}

func (o *OrganizationMonitoringTimeUsageForbidden) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/resources/organizationUsageHistory][%d] organizationMonitoringTimeUsageForbidden ", 403)
}

func (o *OrganizationMonitoringTimeUsageForbidden) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/resources/organizationUsageHistory][%d] organizationMonitoringTimeUsageForbidden ", 403)
}

func (o *OrganizationMonitoringTimeUsageForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewOrganizationMonitoringTimeUsageNotFound creates a OrganizationMonitoringTimeUsageNotFound with default headers values
func NewOrganizationMonitoringTimeUsageNotFound() *OrganizationMonitoringTimeUsageNotFound {
	return &OrganizationMonitoringTimeUsageNotFound{}
}

/*
OrganizationMonitoringTimeUsageNotFound describes a response with status code 404, with default header values.

Not Found
*/
type OrganizationMonitoringTimeUsageNotFound struct {
}

// IsSuccess returns true when this organization monitoring time usage not found response has a 2xx status code
func (o *OrganizationMonitoringTimeUsageNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this organization monitoring time usage not found response has a 3xx status code
func (o *OrganizationMonitoringTimeUsageNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this organization monitoring time usage not found response has a 4xx status code
func (o *OrganizationMonitoringTimeUsageNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this organization monitoring time usage not found response has a 5xx status code
func (o *OrganizationMonitoringTimeUsageNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this organization monitoring time usage not found response a status code equal to that given
func (o *OrganizationMonitoringTimeUsageNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the organization monitoring time usage not found response
func (o *OrganizationMonitoringTimeUsageNotFound) Code() int {
	return 404
}

func (o *OrganizationMonitoringTimeUsageNotFound) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/resources/organizationUsageHistory][%d] organizationMonitoringTimeUsageNotFound ", 404)
}

func (o *OrganizationMonitoringTimeUsageNotFound) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/resources/organizationUsageHistory][%d] organizationMonitoringTimeUsageNotFound ", 404)
}

func (o *OrganizationMonitoringTimeUsageNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewOrganizationMonitoringTimeUsageRequestTimeout creates a OrganizationMonitoringTimeUsageRequestTimeout with default headers values
func NewOrganizationMonitoringTimeUsageRequestTimeout() *OrganizationMonitoringTimeUsageRequestTimeout {
	return &OrganizationMonitoringTimeUsageRequestTimeout{}
}

/*
OrganizationMonitoringTimeUsageRequestTimeout describes a response with status code 408, with default header values.

Request Timeout
*/
type OrganizationMonitoringTimeUsageRequestTimeout struct {
}

// IsSuccess returns true when this organization monitoring time usage request timeout response has a 2xx status code
func (o *OrganizationMonitoringTimeUsageRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this organization monitoring time usage request timeout response has a 3xx status code
func (o *OrganizationMonitoringTimeUsageRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this organization monitoring time usage request timeout response has a 4xx status code
func (o *OrganizationMonitoringTimeUsageRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this organization monitoring time usage request timeout response has a 5xx status code
func (o *OrganizationMonitoringTimeUsageRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this organization monitoring time usage request timeout response a status code equal to that given
func (o *OrganizationMonitoringTimeUsageRequestTimeout) IsCode(code int) bool {
	return code == 408
}

// Code gets the status code for the organization monitoring time usage request timeout response
func (o *OrganizationMonitoringTimeUsageRequestTimeout) Code() int {
	return 408
}

func (o *OrganizationMonitoringTimeUsageRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/resources/organizationUsageHistory][%d] organizationMonitoringTimeUsageRequestTimeout ", 408)
}

func (o *OrganizationMonitoringTimeUsageRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/resources/organizationUsageHistory][%d] organizationMonitoringTimeUsageRequestTimeout ", 408)
}

func (o *OrganizationMonitoringTimeUsageRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewOrganizationMonitoringTimeUsageUnprocessableEntity creates a OrganizationMonitoringTimeUsageUnprocessableEntity with default headers values
func NewOrganizationMonitoringTimeUsageUnprocessableEntity() *OrganizationMonitoringTimeUsageUnprocessableEntity {
	return &OrganizationMonitoringTimeUsageUnprocessableEntity{}
}

/*
OrganizationMonitoringTimeUsageUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type OrganizationMonitoringTimeUsageUnprocessableEntity struct {
}

// IsSuccess returns true when this organization monitoring time usage unprocessable entity response has a 2xx status code
func (o *OrganizationMonitoringTimeUsageUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this organization monitoring time usage unprocessable entity response has a 3xx status code
func (o *OrganizationMonitoringTimeUsageUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this organization monitoring time usage unprocessable entity response has a 4xx status code
func (o *OrganizationMonitoringTimeUsageUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this organization monitoring time usage unprocessable entity response has a 5xx status code
func (o *OrganizationMonitoringTimeUsageUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this organization monitoring time usage unprocessable entity response a status code equal to that given
func (o *OrganizationMonitoringTimeUsageUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the organization monitoring time usage unprocessable entity response
func (o *OrganizationMonitoringTimeUsageUnprocessableEntity) Code() int {
	return 422
}

func (o *OrganizationMonitoringTimeUsageUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/resources/organizationUsageHistory][%d] organizationMonitoringTimeUsageUnprocessableEntity ", 422)
}

func (o *OrganizationMonitoringTimeUsageUnprocessableEntity) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/resources/organizationUsageHistory][%d] organizationMonitoringTimeUsageUnprocessableEntity ", 422)
}

func (o *OrganizationMonitoringTimeUsageUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
