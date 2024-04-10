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

// GetOrganizationAvatarReader is a Reader for the GetOrganizationAvatar structure.
type GetOrganizationAvatarReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrganizationAvatarReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrganizationAvatarOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetOrganizationAvatarBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetOrganizationAvatarUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetOrganizationAvatarForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetOrganizationAvatarNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetOrganizationAvatarRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGetOrganizationAvatarUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /api/backend/v1/organizations2/{organizationName}/avatar] getOrganizationAvatar", response, response.Code())
	}
}

// NewGetOrganizationAvatarOK creates a GetOrganizationAvatarOK with default headers values
func NewGetOrganizationAvatarOK() *GetOrganizationAvatarOK {
	return &GetOrganizationAvatarOK{}
}

/*
GetOrganizationAvatarOK describes a response with status code 200, with default header values.

No response
*/
type GetOrganizationAvatarOK struct {
}

// IsSuccess returns true when this get organization avatar o k response has a 2xx status code
func (o *GetOrganizationAvatarOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get organization avatar o k response has a 3xx status code
func (o *GetOrganizationAvatarOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organization avatar o k response has a 4xx status code
func (o *GetOrganizationAvatarOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get organization avatar o k response has a 5xx status code
func (o *GetOrganizationAvatarOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get organization avatar o k response a status code equal to that given
func (o *GetOrganizationAvatarOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get organization avatar o k response
func (o *GetOrganizationAvatarOK) Code() int {
	return 200
}

func (o *GetOrganizationAvatarOK) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/organizations2/{organizationName}/avatar][%d] getOrganizationAvatarOK ", 200)
}

func (o *GetOrganizationAvatarOK) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/organizations2/{organizationName}/avatar][%d] getOrganizationAvatarOK ", 200)
}

func (o *GetOrganizationAvatarOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetOrganizationAvatarBadRequest creates a GetOrganizationAvatarBadRequest with default headers values
func NewGetOrganizationAvatarBadRequest() *GetOrganizationAvatarBadRequest {
	return &GetOrganizationAvatarBadRequest{}
}

/*
GetOrganizationAvatarBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetOrganizationAvatarBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this get organization avatar bad request response has a 2xx status code
func (o *GetOrganizationAvatarBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get organization avatar bad request response has a 3xx status code
func (o *GetOrganizationAvatarBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organization avatar bad request response has a 4xx status code
func (o *GetOrganizationAvatarBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get organization avatar bad request response has a 5xx status code
func (o *GetOrganizationAvatarBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get organization avatar bad request response a status code equal to that given
func (o *GetOrganizationAvatarBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get organization avatar bad request response
func (o *GetOrganizationAvatarBadRequest) Code() int {
	return 400
}

func (o *GetOrganizationAvatarBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/organizations2/{organizationName}/avatar][%d] getOrganizationAvatarBadRequest  %+v", 400, o.Payload)
}

func (o *GetOrganizationAvatarBadRequest) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/organizations2/{organizationName}/avatar][%d] getOrganizationAvatarBadRequest  %+v", 400, o.Payload)
}

func (o *GetOrganizationAvatarBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetOrganizationAvatarBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationAvatarUnauthorized creates a GetOrganizationAvatarUnauthorized with default headers values
func NewGetOrganizationAvatarUnauthorized() *GetOrganizationAvatarUnauthorized {
	return &GetOrganizationAvatarUnauthorized{}
}

/*
GetOrganizationAvatarUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetOrganizationAvatarUnauthorized struct {
}

// IsSuccess returns true when this get organization avatar unauthorized response has a 2xx status code
func (o *GetOrganizationAvatarUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get organization avatar unauthorized response has a 3xx status code
func (o *GetOrganizationAvatarUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organization avatar unauthorized response has a 4xx status code
func (o *GetOrganizationAvatarUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get organization avatar unauthorized response has a 5xx status code
func (o *GetOrganizationAvatarUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get organization avatar unauthorized response a status code equal to that given
func (o *GetOrganizationAvatarUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get organization avatar unauthorized response
func (o *GetOrganizationAvatarUnauthorized) Code() int {
	return 401
}

func (o *GetOrganizationAvatarUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/organizations2/{organizationName}/avatar][%d] getOrganizationAvatarUnauthorized ", 401)
}

func (o *GetOrganizationAvatarUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/organizations2/{organizationName}/avatar][%d] getOrganizationAvatarUnauthorized ", 401)
}

func (o *GetOrganizationAvatarUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetOrganizationAvatarForbidden creates a GetOrganizationAvatarForbidden with default headers values
func NewGetOrganizationAvatarForbidden() *GetOrganizationAvatarForbidden {
	return &GetOrganizationAvatarForbidden{}
}

/*
GetOrganizationAvatarForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetOrganizationAvatarForbidden struct {
}

// IsSuccess returns true when this get organization avatar forbidden response has a 2xx status code
func (o *GetOrganizationAvatarForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get organization avatar forbidden response has a 3xx status code
func (o *GetOrganizationAvatarForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organization avatar forbidden response has a 4xx status code
func (o *GetOrganizationAvatarForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get organization avatar forbidden response has a 5xx status code
func (o *GetOrganizationAvatarForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get organization avatar forbidden response a status code equal to that given
func (o *GetOrganizationAvatarForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get organization avatar forbidden response
func (o *GetOrganizationAvatarForbidden) Code() int {
	return 403
}

func (o *GetOrganizationAvatarForbidden) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/organizations2/{organizationName}/avatar][%d] getOrganizationAvatarForbidden ", 403)
}

func (o *GetOrganizationAvatarForbidden) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/organizations2/{organizationName}/avatar][%d] getOrganizationAvatarForbidden ", 403)
}

func (o *GetOrganizationAvatarForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetOrganizationAvatarNotFound creates a GetOrganizationAvatarNotFound with default headers values
func NewGetOrganizationAvatarNotFound() *GetOrganizationAvatarNotFound {
	return &GetOrganizationAvatarNotFound{}
}

/*
GetOrganizationAvatarNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetOrganizationAvatarNotFound struct {
}

// IsSuccess returns true when this get organization avatar not found response has a 2xx status code
func (o *GetOrganizationAvatarNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get organization avatar not found response has a 3xx status code
func (o *GetOrganizationAvatarNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organization avatar not found response has a 4xx status code
func (o *GetOrganizationAvatarNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get organization avatar not found response has a 5xx status code
func (o *GetOrganizationAvatarNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get organization avatar not found response a status code equal to that given
func (o *GetOrganizationAvatarNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get organization avatar not found response
func (o *GetOrganizationAvatarNotFound) Code() int {
	return 404
}

func (o *GetOrganizationAvatarNotFound) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/organizations2/{organizationName}/avatar][%d] getOrganizationAvatarNotFound ", 404)
}

func (o *GetOrganizationAvatarNotFound) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/organizations2/{organizationName}/avatar][%d] getOrganizationAvatarNotFound ", 404)
}

func (o *GetOrganizationAvatarNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetOrganizationAvatarRequestTimeout creates a GetOrganizationAvatarRequestTimeout with default headers values
func NewGetOrganizationAvatarRequestTimeout() *GetOrganizationAvatarRequestTimeout {
	return &GetOrganizationAvatarRequestTimeout{}
}

/*
GetOrganizationAvatarRequestTimeout describes a response with status code 408, with default header values.

Request Timeout
*/
type GetOrganizationAvatarRequestTimeout struct {
}

// IsSuccess returns true when this get organization avatar request timeout response has a 2xx status code
func (o *GetOrganizationAvatarRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get organization avatar request timeout response has a 3xx status code
func (o *GetOrganizationAvatarRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organization avatar request timeout response has a 4xx status code
func (o *GetOrganizationAvatarRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get organization avatar request timeout response has a 5xx status code
func (o *GetOrganizationAvatarRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get organization avatar request timeout response a status code equal to that given
func (o *GetOrganizationAvatarRequestTimeout) IsCode(code int) bool {
	return code == 408
}

// Code gets the status code for the get organization avatar request timeout response
func (o *GetOrganizationAvatarRequestTimeout) Code() int {
	return 408
}

func (o *GetOrganizationAvatarRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/organizations2/{organizationName}/avatar][%d] getOrganizationAvatarRequestTimeout ", 408)
}

func (o *GetOrganizationAvatarRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/organizations2/{organizationName}/avatar][%d] getOrganizationAvatarRequestTimeout ", 408)
}

func (o *GetOrganizationAvatarRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetOrganizationAvatarUnprocessableEntity creates a GetOrganizationAvatarUnprocessableEntity with default headers values
func NewGetOrganizationAvatarUnprocessableEntity() *GetOrganizationAvatarUnprocessableEntity {
	return &GetOrganizationAvatarUnprocessableEntity{}
}

/*
GetOrganizationAvatarUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type GetOrganizationAvatarUnprocessableEntity struct {
}

// IsSuccess returns true when this get organization avatar unprocessable entity response has a 2xx status code
func (o *GetOrganizationAvatarUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get organization avatar unprocessable entity response has a 3xx status code
func (o *GetOrganizationAvatarUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organization avatar unprocessable entity response has a 4xx status code
func (o *GetOrganizationAvatarUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this get organization avatar unprocessable entity response has a 5xx status code
func (o *GetOrganizationAvatarUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this get organization avatar unprocessable entity response a status code equal to that given
func (o *GetOrganizationAvatarUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the get organization avatar unprocessable entity response
func (o *GetOrganizationAvatarUnprocessableEntity) Code() int {
	return 422
}

func (o *GetOrganizationAvatarUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/organizations2/{organizationName}/avatar][%d] getOrganizationAvatarUnprocessableEntity ", 422)
}

func (o *GetOrganizationAvatarUnprocessableEntity) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/organizations2/{organizationName}/avatar][%d] getOrganizationAvatarUnprocessableEntity ", 422)
}

func (o *GetOrganizationAvatarUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
