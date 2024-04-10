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

// GetProjectAvatarReader is a Reader for the GetProjectAvatar structure.
type GetProjectAvatarReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProjectAvatarReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetProjectAvatarOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetProjectAvatarBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetProjectAvatarUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetProjectAvatarForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetProjectAvatarNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetProjectAvatarRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGetProjectAvatarUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /api/backend/v1/organizations/{organizationName}/projects/{projectName}/avatar] getProjectAvatar", response, response.Code())
	}
}

// NewGetProjectAvatarOK creates a GetProjectAvatarOK with default headers values
func NewGetProjectAvatarOK() *GetProjectAvatarOK {
	return &GetProjectAvatarOK{}
}

/*
GetProjectAvatarOK describes a response with status code 200, with default header values.

No response
*/
type GetProjectAvatarOK struct {
}

// IsSuccess returns true when this get project avatar o k response has a 2xx status code
func (o *GetProjectAvatarOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get project avatar o k response has a 3xx status code
func (o *GetProjectAvatarOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get project avatar o k response has a 4xx status code
func (o *GetProjectAvatarOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get project avatar o k response has a 5xx status code
func (o *GetProjectAvatarOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get project avatar o k response a status code equal to that given
func (o *GetProjectAvatarOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get project avatar o k response
func (o *GetProjectAvatarOK) Code() int {
	return 200
}

func (o *GetProjectAvatarOK) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/organizations/{organizationName}/projects/{projectName}/avatar][%d] getProjectAvatarOK ", 200)
}

func (o *GetProjectAvatarOK) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/organizations/{organizationName}/projects/{projectName}/avatar][%d] getProjectAvatarOK ", 200)
}

func (o *GetProjectAvatarOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProjectAvatarBadRequest creates a GetProjectAvatarBadRequest with default headers values
func NewGetProjectAvatarBadRequest() *GetProjectAvatarBadRequest {
	return &GetProjectAvatarBadRequest{}
}

/*
GetProjectAvatarBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetProjectAvatarBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this get project avatar bad request response has a 2xx status code
func (o *GetProjectAvatarBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get project avatar bad request response has a 3xx status code
func (o *GetProjectAvatarBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get project avatar bad request response has a 4xx status code
func (o *GetProjectAvatarBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get project avatar bad request response has a 5xx status code
func (o *GetProjectAvatarBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get project avatar bad request response a status code equal to that given
func (o *GetProjectAvatarBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get project avatar bad request response
func (o *GetProjectAvatarBadRequest) Code() int {
	return 400
}

func (o *GetProjectAvatarBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/organizations/{organizationName}/projects/{projectName}/avatar][%d] getProjectAvatarBadRequest  %+v", 400, o.Payload)
}

func (o *GetProjectAvatarBadRequest) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/organizations/{organizationName}/projects/{projectName}/avatar][%d] getProjectAvatarBadRequest  %+v", 400, o.Payload)
}

func (o *GetProjectAvatarBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetProjectAvatarBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProjectAvatarUnauthorized creates a GetProjectAvatarUnauthorized with default headers values
func NewGetProjectAvatarUnauthorized() *GetProjectAvatarUnauthorized {
	return &GetProjectAvatarUnauthorized{}
}

/*
GetProjectAvatarUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetProjectAvatarUnauthorized struct {
}

// IsSuccess returns true when this get project avatar unauthorized response has a 2xx status code
func (o *GetProjectAvatarUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get project avatar unauthorized response has a 3xx status code
func (o *GetProjectAvatarUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get project avatar unauthorized response has a 4xx status code
func (o *GetProjectAvatarUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get project avatar unauthorized response has a 5xx status code
func (o *GetProjectAvatarUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get project avatar unauthorized response a status code equal to that given
func (o *GetProjectAvatarUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get project avatar unauthorized response
func (o *GetProjectAvatarUnauthorized) Code() int {
	return 401
}

func (o *GetProjectAvatarUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/organizations/{organizationName}/projects/{projectName}/avatar][%d] getProjectAvatarUnauthorized ", 401)
}

func (o *GetProjectAvatarUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/organizations/{organizationName}/projects/{projectName}/avatar][%d] getProjectAvatarUnauthorized ", 401)
}

func (o *GetProjectAvatarUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProjectAvatarForbidden creates a GetProjectAvatarForbidden with default headers values
func NewGetProjectAvatarForbidden() *GetProjectAvatarForbidden {
	return &GetProjectAvatarForbidden{}
}

/*
GetProjectAvatarForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetProjectAvatarForbidden struct {
}

// IsSuccess returns true when this get project avatar forbidden response has a 2xx status code
func (o *GetProjectAvatarForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get project avatar forbidden response has a 3xx status code
func (o *GetProjectAvatarForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get project avatar forbidden response has a 4xx status code
func (o *GetProjectAvatarForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get project avatar forbidden response has a 5xx status code
func (o *GetProjectAvatarForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get project avatar forbidden response a status code equal to that given
func (o *GetProjectAvatarForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get project avatar forbidden response
func (o *GetProjectAvatarForbidden) Code() int {
	return 403
}

func (o *GetProjectAvatarForbidden) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/organizations/{organizationName}/projects/{projectName}/avatar][%d] getProjectAvatarForbidden ", 403)
}

func (o *GetProjectAvatarForbidden) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/organizations/{organizationName}/projects/{projectName}/avatar][%d] getProjectAvatarForbidden ", 403)
}

func (o *GetProjectAvatarForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProjectAvatarNotFound creates a GetProjectAvatarNotFound with default headers values
func NewGetProjectAvatarNotFound() *GetProjectAvatarNotFound {
	return &GetProjectAvatarNotFound{}
}

/*
GetProjectAvatarNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetProjectAvatarNotFound struct {
}

// IsSuccess returns true when this get project avatar not found response has a 2xx status code
func (o *GetProjectAvatarNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get project avatar not found response has a 3xx status code
func (o *GetProjectAvatarNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get project avatar not found response has a 4xx status code
func (o *GetProjectAvatarNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get project avatar not found response has a 5xx status code
func (o *GetProjectAvatarNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get project avatar not found response a status code equal to that given
func (o *GetProjectAvatarNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get project avatar not found response
func (o *GetProjectAvatarNotFound) Code() int {
	return 404
}

func (o *GetProjectAvatarNotFound) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/organizations/{organizationName}/projects/{projectName}/avatar][%d] getProjectAvatarNotFound ", 404)
}

func (o *GetProjectAvatarNotFound) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/organizations/{organizationName}/projects/{projectName}/avatar][%d] getProjectAvatarNotFound ", 404)
}

func (o *GetProjectAvatarNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProjectAvatarRequestTimeout creates a GetProjectAvatarRequestTimeout with default headers values
func NewGetProjectAvatarRequestTimeout() *GetProjectAvatarRequestTimeout {
	return &GetProjectAvatarRequestTimeout{}
}

/*
GetProjectAvatarRequestTimeout describes a response with status code 408, with default header values.

Request Timeout
*/
type GetProjectAvatarRequestTimeout struct {
}

// IsSuccess returns true when this get project avatar request timeout response has a 2xx status code
func (o *GetProjectAvatarRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get project avatar request timeout response has a 3xx status code
func (o *GetProjectAvatarRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get project avatar request timeout response has a 4xx status code
func (o *GetProjectAvatarRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get project avatar request timeout response has a 5xx status code
func (o *GetProjectAvatarRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get project avatar request timeout response a status code equal to that given
func (o *GetProjectAvatarRequestTimeout) IsCode(code int) bool {
	return code == 408
}

// Code gets the status code for the get project avatar request timeout response
func (o *GetProjectAvatarRequestTimeout) Code() int {
	return 408
}

func (o *GetProjectAvatarRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/organizations/{organizationName}/projects/{projectName}/avatar][%d] getProjectAvatarRequestTimeout ", 408)
}

func (o *GetProjectAvatarRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/organizations/{organizationName}/projects/{projectName}/avatar][%d] getProjectAvatarRequestTimeout ", 408)
}

func (o *GetProjectAvatarRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProjectAvatarUnprocessableEntity creates a GetProjectAvatarUnprocessableEntity with default headers values
func NewGetProjectAvatarUnprocessableEntity() *GetProjectAvatarUnprocessableEntity {
	return &GetProjectAvatarUnprocessableEntity{}
}

/*
GetProjectAvatarUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type GetProjectAvatarUnprocessableEntity struct {
}

// IsSuccess returns true when this get project avatar unprocessable entity response has a 2xx status code
func (o *GetProjectAvatarUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get project avatar unprocessable entity response has a 3xx status code
func (o *GetProjectAvatarUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get project avatar unprocessable entity response has a 4xx status code
func (o *GetProjectAvatarUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this get project avatar unprocessable entity response has a 5xx status code
func (o *GetProjectAvatarUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this get project avatar unprocessable entity response a status code equal to that given
func (o *GetProjectAvatarUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the get project avatar unprocessable entity response
func (o *GetProjectAvatarUnprocessableEntity) Code() int {
	return 422
}

func (o *GetProjectAvatarUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/organizations/{organizationName}/projects/{projectName}/avatar][%d] getProjectAvatarUnprocessableEntity ", 422)
}

func (o *GetProjectAvatarUnprocessableEntity) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/organizations/{organizationName}/projects/{projectName}/avatar][%d] getProjectAvatarUnprocessableEntity ", 422)
}

func (o *GetProjectAvatarUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
