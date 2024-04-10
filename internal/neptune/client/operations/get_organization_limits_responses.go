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

// GetOrganizationLimitsReader is a Reader for the GetOrganizationLimits structure.
type GetOrganizationLimitsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrganizationLimitsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrganizationLimitsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetOrganizationLimitsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetOrganizationLimitsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetOrganizationLimitsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetOrganizationLimitsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetOrganizationLimitsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGetOrganizationLimitsUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /api/backend/v1/organizations2/{organizationIdentifier}/limits] getOrganizationLimits", response, response.Code())
	}
}

// NewGetOrganizationLimitsOK creates a GetOrganizationLimitsOK with default headers values
func NewGetOrganizationLimitsOK() *GetOrganizationLimitsOK {
	return &GetOrganizationLimitsOK{}
}

/*
GetOrganizationLimitsOK describes a response with status code 200, with default header values.

OK
*/
type GetOrganizationLimitsOK struct {
	Payload *models.OrganizationLimitsDTO
}

// IsSuccess returns true when this get organization limits o k response has a 2xx status code
func (o *GetOrganizationLimitsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get organization limits o k response has a 3xx status code
func (o *GetOrganizationLimitsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organization limits o k response has a 4xx status code
func (o *GetOrganizationLimitsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get organization limits o k response has a 5xx status code
func (o *GetOrganizationLimitsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get organization limits o k response a status code equal to that given
func (o *GetOrganizationLimitsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get organization limits o k response
func (o *GetOrganizationLimitsOK) Code() int {
	return 200
}

func (o *GetOrganizationLimitsOK) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/organizations2/{organizationIdentifier}/limits][%d] getOrganizationLimitsOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationLimitsOK) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/organizations2/{organizationIdentifier}/limits][%d] getOrganizationLimitsOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationLimitsOK) GetPayload() *models.OrganizationLimitsDTO {
	return o.Payload
}

func (o *GetOrganizationLimitsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OrganizationLimitsDTO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationLimitsBadRequest creates a GetOrganizationLimitsBadRequest with default headers values
func NewGetOrganizationLimitsBadRequest() *GetOrganizationLimitsBadRequest {
	return &GetOrganizationLimitsBadRequest{}
}

/*
GetOrganizationLimitsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetOrganizationLimitsBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this get organization limits bad request response has a 2xx status code
func (o *GetOrganizationLimitsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get organization limits bad request response has a 3xx status code
func (o *GetOrganizationLimitsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organization limits bad request response has a 4xx status code
func (o *GetOrganizationLimitsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get organization limits bad request response has a 5xx status code
func (o *GetOrganizationLimitsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get organization limits bad request response a status code equal to that given
func (o *GetOrganizationLimitsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get organization limits bad request response
func (o *GetOrganizationLimitsBadRequest) Code() int {
	return 400
}

func (o *GetOrganizationLimitsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/organizations2/{organizationIdentifier}/limits][%d] getOrganizationLimitsBadRequest  %+v", 400, o.Payload)
}

func (o *GetOrganizationLimitsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/organizations2/{organizationIdentifier}/limits][%d] getOrganizationLimitsBadRequest  %+v", 400, o.Payload)
}

func (o *GetOrganizationLimitsBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetOrganizationLimitsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationLimitsUnauthorized creates a GetOrganizationLimitsUnauthorized with default headers values
func NewGetOrganizationLimitsUnauthorized() *GetOrganizationLimitsUnauthorized {
	return &GetOrganizationLimitsUnauthorized{}
}

/*
GetOrganizationLimitsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetOrganizationLimitsUnauthorized struct {
}

// IsSuccess returns true when this get organization limits unauthorized response has a 2xx status code
func (o *GetOrganizationLimitsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get organization limits unauthorized response has a 3xx status code
func (o *GetOrganizationLimitsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organization limits unauthorized response has a 4xx status code
func (o *GetOrganizationLimitsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get organization limits unauthorized response has a 5xx status code
func (o *GetOrganizationLimitsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get organization limits unauthorized response a status code equal to that given
func (o *GetOrganizationLimitsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get organization limits unauthorized response
func (o *GetOrganizationLimitsUnauthorized) Code() int {
	return 401
}

func (o *GetOrganizationLimitsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/organizations2/{organizationIdentifier}/limits][%d] getOrganizationLimitsUnauthorized ", 401)
}

func (o *GetOrganizationLimitsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/organizations2/{organizationIdentifier}/limits][%d] getOrganizationLimitsUnauthorized ", 401)
}

func (o *GetOrganizationLimitsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetOrganizationLimitsForbidden creates a GetOrganizationLimitsForbidden with default headers values
func NewGetOrganizationLimitsForbidden() *GetOrganizationLimitsForbidden {
	return &GetOrganizationLimitsForbidden{}
}

/*
GetOrganizationLimitsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetOrganizationLimitsForbidden struct {
}

// IsSuccess returns true when this get organization limits forbidden response has a 2xx status code
func (o *GetOrganizationLimitsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get organization limits forbidden response has a 3xx status code
func (o *GetOrganizationLimitsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organization limits forbidden response has a 4xx status code
func (o *GetOrganizationLimitsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get organization limits forbidden response has a 5xx status code
func (o *GetOrganizationLimitsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get organization limits forbidden response a status code equal to that given
func (o *GetOrganizationLimitsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get organization limits forbidden response
func (o *GetOrganizationLimitsForbidden) Code() int {
	return 403
}

func (o *GetOrganizationLimitsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/organizations2/{organizationIdentifier}/limits][%d] getOrganizationLimitsForbidden ", 403)
}

func (o *GetOrganizationLimitsForbidden) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/organizations2/{organizationIdentifier}/limits][%d] getOrganizationLimitsForbidden ", 403)
}

func (o *GetOrganizationLimitsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetOrganizationLimitsNotFound creates a GetOrganizationLimitsNotFound with default headers values
func NewGetOrganizationLimitsNotFound() *GetOrganizationLimitsNotFound {
	return &GetOrganizationLimitsNotFound{}
}

/*
GetOrganizationLimitsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetOrganizationLimitsNotFound struct {
}

// IsSuccess returns true when this get organization limits not found response has a 2xx status code
func (o *GetOrganizationLimitsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get organization limits not found response has a 3xx status code
func (o *GetOrganizationLimitsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organization limits not found response has a 4xx status code
func (o *GetOrganizationLimitsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get organization limits not found response has a 5xx status code
func (o *GetOrganizationLimitsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get organization limits not found response a status code equal to that given
func (o *GetOrganizationLimitsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get organization limits not found response
func (o *GetOrganizationLimitsNotFound) Code() int {
	return 404
}

func (o *GetOrganizationLimitsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/organizations2/{organizationIdentifier}/limits][%d] getOrganizationLimitsNotFound ", 404)
}

func (o *GetOrganizationLimitsNotFound) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/organizations2/{organizationIdentifier}/limits][%d] getOrganizationLimitsNotFound ", 404)
}

func (o *GetOrganizationLimitsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetOrganizationLimitsRequestTimeout creates a GetOrganizationLimitsRequestTimeout with default headers values
func NewGetOrganizationLimitsRequestTimeout() *GetOrganizationLimitsRequestTimeout {
	return &GetOrganizationLimitsRequestTimeout{}
}

/*
GetOrganizationLimitsRequestTimeout describes a response with status code 408, with default header values.

Request Timeout
*/
type GetOrganizationLimitsRequestTimeout struct {
}

// IsSuccess returns true when this get organization limits request timeout response has a 2xx status code
func (o *GetOrganizationLimitsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get organization limits request timeout response has a 3xx status code
func (o *GetOrganizationLimitsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organization limits request timeout response has a 4xx status code
func (o *GetOrganizationLimitsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get organization limits request timeout response has a 5xx status code
func (o *GetOrganizationLimitsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get organization limits request timeout response a status code equal to that given
func (o *GetOrganizationLimitsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

// Code gets the status code for the get organization limits request timeout response
func (o *GetOrganizationLimitsRequestTimeout) Code() int {
	return 408
}

func (o *GetOrganizationLimitsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/organizations2/{organizationIdentifier}/limits][%d] getOrganizationLimitsRequestTimeout ", 408)
}

func (o *GetOrganizationLimitsRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/organizations2/{organizationIdentifier}/limits][%d] getOrganizationLimitsRequestTimeout ", 408)
}

func (o *GetOrganizationLimitsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetOrganizationLimitsUnprocessableEntity creates a GetOrganizationLimitsUnprocessableEntity with default headers values
func NewGetOrganizationLimitsUnprocessableEntity() *GetOrganizationLimitsUnprocessableEntity {
	return &GetOrganizationLimitsUnprocessableEntity{}
}

/*
GetOrganizationLimitsUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type GetOrganizationLimitsUnprocessableEntity struct {
}

// IsSuccess returns true when this get organization limits unprocessable entity response has a 2xx status code
func (o *GetOrganizationLimitsUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get organization limits unprocessable entity response has a 3xx status code
func (o *GetOrganizationLimitsUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organization limits unprocessable entity response has a 4xx status code
func (o *GetOrganizationLimitsUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this get organization limits unprocessable entity response has a 5xx status code
func (o *GetOrganizationLimitsUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this get organization limits unprocessable entity response a status code equal to that given
func (o *GetOrganizationLimitsUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the get organization limits unprocessable entity response
func (o *GetOrganizationLimitsUnprocessableEntity) Code() int {
	return 422
}

func (o *GetOrganizationLimitsUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/organizations2/{organizationIdentifier}/limits][%d] getOrganizationLimitsUnprocessableEntity ", 422)
}

func (o *GetOrganizationLimitsUnprocessableEntity) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/organizations2/{organizationIdentifier}/limits][%d] getOrganizationLimitsUnprocessableEntity ", 422)
}

func (o *GetOrganizationLimitsUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}