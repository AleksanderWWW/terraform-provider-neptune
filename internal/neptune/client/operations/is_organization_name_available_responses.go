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

// IsOrganizationNameAvailableReader is a Reader for the IsOrganizationNameAvailable structure.
type IsOrganizationNameAvailableReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IsOrganizationNameAvailableReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIsOrganizationNameAvailableOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewIsOrganizationNameAvailableBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewIsOrganizationNameAvailableUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewIsOrganizationNameAvailableForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewIsOrganizationNameAvailableNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewIsOrganizationNameAvailableRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewIsOrganizationNameAvailableUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /api/backend/v1/organizations2/{organizationName}/available] isOrganizationNameAvailable", response, response.Code())
	}
}

// NewIsOrganizationNameAvailableOK creates a IsOrganizationNameAvailableOK with default headers values
func NewIsOrganizationNameAvailableOK() *IsOrganizationNameAvailableOK {
	return &IsOrganizationNameAvailableOK{}
}

/*
IsOrganizationNameAvailableOK describes a response with status code 200, with default header values.

OK
*/
type IsOrganizationNameAvailableOK struct {
	Payload *models.OrganizationNameAvailableDTO
}

// IsSuccess returns true when this is organization name available o k response has a 2xx status code
func (o *IsOrganizationNameAvailableOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this is organization name available o k response has a 3xx status code
func (o *IsOrganizationNameAvailableOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this is organization name available o k response has a 4xx status code
func (o *IsOrganizationNameAvailableOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this is organization name available o k response has a 5xx status code
func (o *IsOrganizationNameAvailableOK) IsServerError() bool {
	return false
}

// IsCode returns true when this is organization name available o k response a status code equal to that given
func (o *IsOrganizationNameAvailableOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the is organization name available o k response
func (o *IsOrganizationNameAvailableOK) Code() int {
	return 200
}

func (o *IsOrganizationNameAvailableOK) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/organizations2/{organizationName}/available][%d] isOrganizationNameAvailableOK  %+v", 200, o.Payload)
}

func (o *IsOrganizationNameAvailableOK) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/organizations2/{organizationName}/available][%d] isOrganizationNameAvailableOK  %+v", 200, o.Payload)
}

func (o *IsOrganizationNameAvailableOK) GetPayload() *models.OrganizationNameAvailableDTO {
	return o.Payload
}

func (o *IsOrganizationNameAvailableOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OrganizationNameAvailableDTO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIsOrganizationNameAvailableBadRequest creates a IsOrganizationNameAvailableBadRequest with default headers values
func NewIsOrganizationNameAvailableBadRequest() *IsOrganizationNameAvailableBadRequest {
	return &IsOrganizationNameAvailableBadRequest{}
}

/*
IsOrganizationNameAvailableBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type IsOrganizationNameAvailableBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this is organization name available bad request response has a 2xx status code
func (o *IsOrganizationNameAvailableBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this is organization name available bad request response has a 3xx status code
func (o *IsOrganizationNameAvailableBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this is organization name available bad request response has a 4xx status code
func (o *IsOrganizationNameAvailableBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this is organization name available bad request response has a 5xx status code
func (o *IsOrganizationNameAvailableBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this is organization name available bad request response a status code equal to that given
func (o *IsOrganizationNameAvailableBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the is organization name available bad request response
func (o *IsOrganizationNameAvailableBadRequest) Code() int {
	return 400
}

func (o *IsOrganizationNameAvailableBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/organizations2/{organizationName}/available][%d] isOrganizationNameAvailableBadRequest  %+v", 400, o.Payload)
}

func (o *IsOrganizationNameAvailableBadRequest) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/organizations2/{organizationName}/available][%d] isOrganizationNameAvailableBadRequest  %+v", 400, o.Payload)
}

func (o *IsOrganizationNameAvailableBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *IsOrganizationNameAvailableBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIsOrganizationNameAvailableUnauthorized creates a IsOrganizationNameAvailableUnauthorized with default headers values
func NewIsOrganizationNameAvailableUnauthorized() *IsOrganizationNameAvailableUnauthorized {
	return &IsOrganizationNameAvailableUnauthorized{}
}

/*
IsOrganizationNameAvailableUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type IsOrganizationNameAvailableUnauthorized struct {
}

// IsSuccess returns true when this is organization name available unauthorized response has a 2xx status code
func (o *IsOrganizationNameAvailableUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this is organization name available unauthorized response has a 3xx status code
func (o *IsOrganizationNameAvailableUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this is organization name available unauthorized response has a 4xx status code
func (o *IsOrganizationNameAvailableUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this is organization name available unauthorized response has a 5xx status code
func (o *IsOrganizationNameAvailableUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this is organization name available unauthorized response a status code equal to that given
func (o *IsOrganizationNameAvailableUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the is organization name available unauthorized response
func (o *IsOrganizationNameAvailableUnauthorized) Code() int {
	return 401
}

func (o *IsOrganizationNameAvailableUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/organizations2/{organizationName}/available][%d] isOrganizationNameAvailableUnauthorized ", 401)
}

func (o *IsOrganizationNameAvailableUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/organizations2/{organizationName}/available][%d] isOrganizationNameAvailableUnauthorized ", 401)
}

func (o *IsOrganizationNameAvailableUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewIsOrganizationNameAvailableForbidden creates a IsOrganizationNameAvailableForbidden with default headers values
func NewIsOrganizationNameAvailableForbidden() *IsOrganizationNameAvailableForbidden {
	return &IsOrganizationNameAvailableForbidden{}
}

/*
IsOrganizationNameAvailableForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type IsOrganizationNameAvailableForbidden struct {
}

// IsSuccess returns true when this is organization name available forbidden response has a 2xx status code
func (o *IsOrganizationNameAvailableForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this is organization name available forbidden response has a 3xx status code
func (o *IsOrganizationNameAvailableForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this is organization name available forbidden response has a 4xx status code
func (o *IsOrganizationNameAvailableForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this is organization name available forbidden response has a 5xx status code
func (o *IsOrganizationNameAvailableForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this is organization name available forbidden response a status code equal to that given
func (o *IsOrganizationNameAvailableForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the is organization name available forbidden response
func (o *IsOrganizationNameAvailableForbidden) Code() int {
	return 403
}

func (o *IsOrganizationNameAvailableForbidden) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/organizations2/{organizationName}/available][%d] isOrganizationNameAvailableForbidden ", 403)
}

func (o *IsOrganizationNameAvailableForbidden) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/organizations2/{organizationName}/available][%d] isOrganizationNameAvailableForbidden ", 403)
}

func (o *IsOrganizationNameAvailableForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewIsOrganizationNameAvailableNotFound creates a IsOrganizationNameAvailableNotFound with default headers values
func NewIsOrganizationNameAvailableNotFound() *IsOrganizationNameAvailableNotFound {
	return &IsOrganizationNameAvailableNotFound{}
}

/*
IsOrganizationNameAvailableNotFound describes a response with status code 404, with default header values.

Not Found
*/
type IsOrganizationNameAvailableNotFound struct {
}

// IsSuccess returns true when this is organization name available not found response has a 2xx status code
func (o *IsOrganizationNameAvailableNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this is organization name available not found response has a 3xx status code
func (o *IsOrganizationNameAvailableNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this is organization name available not found response has a 4xx status code
func (o *IsOrganizationNameAvailableNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this is organization name available not found response has a 5xx status code
func (o *IsOrganizationNameAvailableNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this is organization name available not found response a status code equal to that given
func (o *IsOrganizationNameAvailableNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the is organization name available not found response
func (o *IsOrganizationNameAvailableNotFound) Code() int {
	return 404
}

func (o *IsOrganizationNameAvailableNotFound) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/organizations2/{organizationName}/available][%d] isOrganizationNameAvailableNotFound ", 404)
}

func (o *IsOrganizationNameAvailableNotFound) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/organizations2/{organizationName}/available][%d] isOrganizationNameAvailableNotFound ", 404)
}

func (o *IsOrganizationNameAvailableNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewIsOrganizationNameAvailableRequestTimeout creates a IsOrganizationNameAvailableRequestTimeout with default headers values
func NewIsOrganizationNameAvailableRequestTimeout() *IsOrganizationNameAvailableRequestTimeout {
	return &IsOrganizationNameAvailableRequestTimeout{}
}

/*
IsOrganizationNameAvailableRequestTimeout describes a response with status code 408, with default header values.

Request Timeout
*/
type IsOrganizationNameAvailableRequestTimeout struct {
}

// IsSuccess returns true when this is organization name available request timeout response has a 2xx status code
func (o *IsOrganizationNameAvailableRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this is organization name available request timeout response has a 3xx status code
func (o *IsOrganizationNameAvailableRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this is organization name available request timeout response has a 4xx status code
func (o *IsOrganizationNameAvailableRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this is organization name available request timeout response has a 5xx status code
func (o *IsOrganizationNameAvailableRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this is organization name available request timeout response a status code equal to that given
func (o *IsOrganizationNameAvailableRequestTimeout) IsCode(code int) bool {
	return code == 408
}

// Code gets the status code for the is organization name available request timeout response
func (o *IsOrganizationNameAvailableRequestTimeout) Code() int {
	return 408
}

func (o *IsOrganizationNameAvailableRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/organizations2/{organizationName}/available][%d] isOrganizationNameAvailableRequestTimeout ", 408)
}

func (o *IsOrganizationNameAvailableRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/organizations2/{organizationName}/available][%d] isOrganizationNameAvailableRequestTimeout ", 408)
}

func (o *IsOrganizationNameAvailableRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewIsOrganizationNameAvailableUnprocessableEntity creates a IsOrganizationNameAvailableUnprocessableEntity with default headers values
func NewIsOrganizationNameAvailableUnprocessableEntity() *IsOrganizationNameAvailableUnprocessableEntity {
	return &IsOrganizationNameAvailableUnprocessableEntity{}
}

/*
IsOrganizationNameAvailableUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type IsOrganizationNameAvailableUnprocessableEntity struct {
}

// IsSuccess returns true when this is organization name available unprocessable entity response has a 2xx status code
func (o *IsOrganizationNameAvailableUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this is organization name available unprocessable entity response has a 3xx status code
func (o *IsOrganizationNameAvailableUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this is organization name available unprocessable entity response has a 4xx status code
func (o *IsOrganizationNameAvailableUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this is organization name available unprocessable entity response has a 5xx status code
func (o *IsOrganizationNameAvailableUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this is organization name available unprocessable entity response a status code equal to that given
func (o *IsOrganizationNameAvailableUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the is organization name available unprocessable entity response
func (o *IsOrganizationNameAvailableUnprocessableEntity) Code() int {
	return 422
}

func (o *IsOrganizationNameAvailableUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/organizations2/{organizationName}/available][%d] isOrganizationNameAvailableUnprocessableEntity ", 422)
}

func (o *IsOrganizationNameAvailableUnprocessableEntity) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/organizations2/{organizationName}/available][%d] isOrganizationNameAvailableUnprocessableEntity ", 422)
}

func (o *IsOrganizationNameAvailableUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
