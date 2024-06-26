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

// GetProjectBackgroundReader is a Reader for the GetProjectBackground structure.
type GetProjectBackgroundReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProjectBackgroundReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetProjectBackgroundOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetProjectBackgroundBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetProjectBackgroundUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetProjectBackgroundForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetProjectBackgroundNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetProjectBackgroundRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGetProjectBackgroundUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /api/backend/v1/organizations/{organizationName}/projects/{projectName}/background] getProjectBackground", response, response.Code())
	}
}

// NewGetProjectBackgroundOK creates a GetProjectBackgroundOK with default headers values
func NewGetProjectBackgroundOK() *GetProjectBackgroundOK {
	return &GetProjectBackgroundOK{}
}

/*
GetProjectBackgroundOK describes a response with status code 200, with default header values.

No response
*/
type GetProjectBackgroundOK struct {
}

// IsSuccess returns true when this get project background o k response has a 2xx status code
func (o *GetProjectBackgroundOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get project background o k response has a 3xx status code
func (o *GetProjectBackgroundOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get project background o k response has a 4xx status code
func (o *GetProjectBackgroundOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get project background o k response has a 5xx status code
func (o *GetProjectBackgroundOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get project background o k response a status code equal to that given
func (o *GetProjectBackgroundOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get project background o k response
func (o *GetProjectBackgroundOK) Code() int {
	return 200
}

func (o *GetProjectBackgroundOK) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/organizations/{organizationName}/projects/{projectName}/background][%d] getProjectBackgroundOK ", 200)
}

func (o *GetProjectBackgroundOK) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/organizations/{organizationName}/projects/{projectName}/background][%d] getProjectBackgroundOK ", 200)
}

func (o *GetProjectBackgroundOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProjectBackgroundBadRequest creates a GetProjectBackgroundBadRequest with default headers values
func NewGetProjectBackgroundBadRequest() *GetProjectBackgroundBadRequest {
	return &GetProjectBackgroundBadRequest{}
}

/*
GetProjectBackgroundBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetProjectBackgroundBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this get project background bad request response has a 2xx status code
func (o *GetProjectBackgroundBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get project background bad request response has a 3xx status code
func (o *GetProjectBackgroundBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get project background bad request response has a 4xx status code
func (o *GetProjectBackgroundBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get project background bad request response has a 5xx status code
func (o *GetProjectBackgroundBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get project background bad request response a status code equal to that given
func (o *GetProjectBackgroundBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get project background bad request response
func (o *GetProjectBackgroundBadRequest) Code() int {
	return 400
}

func (o *GetProjectBackgroundBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/organizations/{organizationName}/projects/{projectName}/background][%d] getProjectBackgroundBadRequest  %+v", 400, o.Payload)
}

func (o *GetProjectBackgroundBadRequest) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/organizations/{organizationName}/projects/{projectName}/background][%d] getProjectBackgroundBadRequest  %+v", 400, o.Payload)
}

func (o *GetProjectBackgroundBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetProjectBackgroundBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProjectBackgroundUnauthorized creates a GetProjectBackgroundUnauthorized with default headers values
func NewGetProjectBackgroundUnauthorized() *GetProjectBackgroundUnauthorized {
	return &GetProjectBackgroundUnauthorized{}
}

/*
GetProjectBackgroundUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetProjectBackgroundUnauthorized struct {
}

// IsSuccess returns true when this get project background unauthorized response has a 2xx status code
func (o *GetProjectBackgroundUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get project background unauthorized response has a 3xx status code
func (o *GetProjectBackgroundUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get project background unauthorized response has a 4xx status code
func (o *GetProjectBackgroundUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get project background unauthorized response has a 5xx status code
func (o *GetProjectBackgroundUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get project background unauthorized response a status code equal to that given
func (o *GetProjectBackgroundUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get project background unauthorized response
func (o *GetProjectBackgroundUnauthorized) Code() int {
	return 401
}

func (o *GetProjectBackgroundUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/organizations/{organizationName}/projects/{projectName}/background][%d] getProjectBackgroundUnauthorized ", 401)
}

func (o *GetProjectBackgroundUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/organizations/{organizationName}/projects/{projectName}/background][%d] getProjectBackgroundUnauthorized ", 401)
}

func (o *GetProjectBackgroundUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProjectBackgroundForbidden creates a GetProjectBackgroundForbidden with default headers values
func NewGetProjectBackgroundForbidden() *GetProjectBackgroundForbidden {
	return &GetProjectBackgroundForbidden{}
}

/*
GetProjectBackgroundForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetProjectBackgroundForbidden struct {
}

// IsSuccess returns true when this get project background forbidden response has a 2xx status code
func (o *GetProjectBackgroundForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get project background forbidden response has a 3xx status code
func (o *GetProjectBackgroundForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get project background forbidden response has a 4xx status code
func (o *GetProjectBackgroundForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get project background forbidden response has a 5xx status code
func (o *GetProjectBackgroundForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get project background forbidden response a status code equal to that given
func (o *GetProjectBackgroundForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get project background forbidden response
func (o *GetProjectBackgroundForbidden) Code() int {
	return 403
}

func (o *GetProjectBackgroundForbidden) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/organizations/{organizationName}/projects/{projectName}/background][%d] getProjectBackgroundForbidden ", 403)
}

func (o *GetProjectBackgroundForbidden) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/organizations/{organizationName}/projects/{projectName}/background][%d] getProjectBackgroundForbidden ", 403)
}

func (o *GetProjectBackgroundForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProjectBackgroundNotFound creates a GetProjectBackgroundNotFound with default headers values
func NewGetProjectBackgroundNotFound() *GetProjectBackgroundNotFound {
	return &GetProjectBackgroundNotFound{}
}

/*
GetProjectBackgroundNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetProjectBackgroundNotFound struct {
}

// IsSuccess returns true when this get project background not found response has a 2xx status code
func (o *GetProjectBackgroundNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get project background not found response has a 3xx status code
func (o *GetProjectBackgroundNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get project background not found response has a 4xx status code
func (o *GetProjectBackgroundNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get project background not found response has a 5xx status code
func (o *GetProjectBackgroundNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get project background not found response a status code equal to that given
func (o *GetProjectBackgroundNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get project background not found response
func (o *GetProjectBackgroundNotFound) Code() int {
	return 404
}

func (o *GetProjectBackgroundNotFound) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/organizations/{organizationName}/projects/{projectName}/background][%d] getProjectBackgroundNotFound ", 404)
}

func (o *GetProjectBackgroundNotFound) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/organizations/{organizationName}/projects/{projectName}/background][%d] getProjectBackgroundNotFound ", 404)
}

func (o *GetProjectBackgroundNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProjectBackgroundRequestTimeout creates a GetProjectBackgroundRequestTimeout with default headers values
func NewGetProjectBackgroundRequestTimeout() *GetProjectBackgroundRequestTimeout {
	return &GetProjectBackgroundRequestTimeout{}
}

/*
GetProjectBackgroundRequestTimeout describes a response with status code 408, with default header values.

Request Timeout
*/
type GetProjectBackgroundRequestTimeout struct {
}

// IsSuccess returns true when this get project background request timeout response has a 2xx status code
func (o *GetProjectBackgroundRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get project background request timeout response has a 3xx status code
func (o *GetProjectBackgroundRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get project background request timeout response has a 4xx status code
func (o *GetProjectBackgroundRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get project background request timeout response has a 5xx status code
func (o *GetProjectBackgroundRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get project background request timeout response a status code equal to that given
func (o *GetProjectBackgroundRequestTimeout) IsCode(code int) bool {
	return code == 408
}

// Code gets the status code for the get project background request timeout response
func (o *GetProjectBackgroundRequestTimeout) Code() int {
	return 408
}

func (o *GetProjectBackgroundRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/organizations/{organizationName}/projects/{projectName}/background][%d] getProjectBackgroundRequestTimeout ", 408)
}

func (o *GetProjectBackgroundRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/organizations/{organizationName}/projects/{projectName}/background][%d] getProjectBackgroundRequestTimeout ", 408)
}

func (o *GetProjectBackgroundRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProjectBackgroundUnprocessableEntity creates a GetProjectBackgroundUnprocessableEntity with default headers values
func NewGetProjectBackgroundUnprocessableEntity() *GetProjectBackgroundUnprocessableEntity {
	return &GetProjectBackgroundUnprocessableEntity{}
}

/*
GetProjectBackgroundUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type GetProjectBackgroundUnprocessableEntity struct {
}

// IsSuccess returns true when this get project background unprocessable entity response has a 2xx status code
func (o *GetProjectBackgroundUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get project background unprocessable entity response has a 3xx status code
func (o *GetProjectBackgroundUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get project background unprocessable entity response has a 4xx status code
func (o *GetProjectBackgroundUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this get project background unprocessable entity response has a 5xx status code
func (o *GetProjectBackgroundUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this get project background unprocessable entity response a status code equal to that given
func (o *GetProjectBackgroundUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the get project background unprocessable entity response
func (o *GetProjectBackgroundUnprocessableEntity) Code() int {
	return 422
}

func (o *GetProjectBackgroundUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/organizations/{organizationName}/projects/{projectName}/background][%d] getProjectBackgroundUnprocessableEntity ", 422)
}

func (o *GetProjectBackgroundUnprocessableEntity) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/organizations/{organizationName}/projects/{projectName}/background][%d] getProjectBackgroundUnprocessableEntity ", 422)
}

func (o *GetProjectBackgroundUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
