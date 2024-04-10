// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"client/models"
)

// ListProjectsMembersPostReader is a Reader for the ListProjectsMembersPost structure.
type ListProjectsMembersPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListProjectsMembersPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListProjectsMembersPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListProjectsMembersPostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewListProjectsMembersPostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListProjectsMembersPostForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListProjectsMembersPostNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewListProjectsMembersPostRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewListProjectsMembersPostUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /api/backend/v1/projects/membersList] listProjectsMembersPost", response, response.Code())
	}
}

// NewListProjectsMembersPostOK creates a ListProjectsMembersPostOK with default headers values
func NewListProjectsMembersPostOK() *ListProjectsMembersPostOK {
	return &ListProjectsMembersPostOK{}
}

/*
ListProjectsMembersPostOK describes a response with status code 200, with default header values.

OK
*/
type ListProjectsMembersPostOK struct {
	Payload []*models.ProjectMembersDTO
}

// IsSuccess returns true when this list projects members post o k response has a 2xx status code
func (o *ListProjectsMembersPostOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list projects members post o k response has a 3xx status code
func (o *ListProjectsMembersPostOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list projects members post o k response has a 4xx status code
func (o *ListProjectsMembersPostOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list projects members post o k response has a 5xx status code
func (o *ListProjectsMembersPostOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list projects members post o k response a status code equal to that given
func (o *ListProjectsMembersPostOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list projects members post o k response
func (o *ListProjectsMembersPostOK) Code() int {
	return 200
}

func (o *ListProjectsMembersPostOK) Error() string {
	return fmt.Sprintf("[POST /api/backend/v1/projects/membersList][%d] listProjectsMembersPostOK  %+v", 200, o.Payload)
}

func (o *ListProjectsMembersPostOK) String() string {
	return fmt.Sprintf("[POST /api/backend/v1/projects/membersList][%d] listProjectsMembersPostOK  %+v", 200, o.Payload)
}

func (o *ListProjectsMembersPostOK) GetPayload() []*models.ProjectMembersDTO {
	return o.Payload
}

func (o *ListProjectsMembersPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListProjectsMembersPostBadRequest creates a ListProjectsMembersPostBadRequest with default headers values
func NewListProjectsMembersPostBadRequest() *ListProjectsMembersPostBadRequest {
	return &ListProjectsMembersPostBadRequest{}
}

/*
ListProjectsMembersPostBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ListProjectsMembersPostBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this list projects members post bad request response has a 2xx status code
func (o *ListProjectsMembersPostBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list projects members post bad request response has a 3xx status code
func (o *ListProjectsMembersPostBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list projects members post bad request response has a 4xx status code
func (o *ListProjectsMembersPostBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this list projects members post bad request response has a 5xx status code
func (o *ListProjectsMembersPostBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this list projects members post bad request response a status code equal to that given
func (o *ListProjectsMembersPostBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the list projects members post bad request response
func (o *ListProjectsMembersPostBadRequest) Code() int {
	return 400
}

func (o *ListProjectsMembersPostBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/backend/v1/projects/membersList][%d] listProjectsMembersPostBadRequest  %+v", 400, o.Payload)
}

func (o *ListProjectsMembersPostBadRequest) String() string {
	return fmt.Sprintf("[POST /api/backend/v1/projects/membersList][%d] listProjectsMembersPostBadRequest  %+v", 400, o.Payload)
}

func (o *ListProjectsMembersPostBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListProjectsMembersPostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListProjectsMembersPostUnauthorized creates a ListProjectsMembersPostUnauthorized with default headers values
func NewListProjectsMembersPostUnauthorized() *ListProjectsMembersPostUnauthorized {
	return &ListProjectsMembersPostUnauthorized{}
}

/*
ListProjectsMembersPostUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ListProjectsMembersPostUnauthorized struct {
}

// IsSuccess returns true when this list projects members post unauthorized response has a 2xx status code
func (o *ListProjectsMembersPostUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list projects members post unauthorized response has a 3xx status code
func (o *ListProjectsMembersPostUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list projects members post unauthorized response has a 4xx status code
func (o *ListProjectsMembersPostUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this list projects members post unauthorized response has a 5xx status code
func (o *ListProjectsMembersPostUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this list projects members post unauthorized response a status code equal to that given
func (o *ListProjectsMembersPostUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the list projects members post unauthorized response
func (o *ListProjectsMembersPostUnauthorized) Code() int {
	return 401
}

func (o *ListProjectsMembersPostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/backend/v1/projects/membersList][%d] listProjectsMembersPostUnauthorized ", 401)
}

func (o *ListProjectsMembersPostUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/backend/v1/projects/membersList][%d] listProjectsMembersPostUnauthorized ", 401)
}

func (o *ListProjectsMembersPostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListProjectsMembersPostForbidden creates a ListProjectsMembersPostForbidden with default headers values
func NewListProjectsMembersPostForbidden() *ListProjectsMembersPostForbidden {
	return &ListProjectsMembersPostForbidden{}
}

/*
ListProjectsMembersPostForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ListProjectsMembersPostForbidden struct {
}

// IsSuccess returns true when this list projects members post forbidden response has a 2xx status code
func (o *ListProjectsMembersPostForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list projects members post forbidden response has a 3xx status code
func (o *ListProjectsMembersPostForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list projects members post forbidden response has a 4xx status code
func (o *ListProjectsMembersPostForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list projects members post forbidden response has a 5xx status code
func (o *ListProjectsMembersPostForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list projects members post forbidden response a status code equal to that given
func (o *ListProjectsMembersPostForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the list projects members post forbidden response
func (o *ListProjectsMembersPostForbidden) Code() int {
	return 403
}

func (o *ListProjectsMembersPostForbidden) Error() string {
	return fmt.Sprintf("[POST /api/backend/v1/projects/membersList][%d] listProjectsMembersPostForbidden ", 403)
}

func (o *ListProjectsMembersPostForbidden) String() string {
	return fmt.Sprintf("[POST /api/backend/v1/projects/membersList][%d] listProjectsMembersPostForbidden ", 403)
}

func (o *ListProjectsMembersPostForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListProjectsMembersPostNotFound creates a ListProjectsMembersPostNotFound with default headers values
func NewListProjectsMembersPostNotFound() *ListProjectsMembersPostNotFound {
	return &ListProjectsMembersPostNotFound{}
}

/*
ListProjectsMembersPostNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ListProjectsMembersPostNotFound struct {
}

// IsSuccess returns true when this list projects members post not found response has a 2xx status code
func (o *ListProjectsMembersPostNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list projects members post not found response has a 3xx status code
func (o *ListProjectsMembersPostNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list projects members post not found response has a 4xx status code
func (o *ListProjectsMembersPostNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this list projects members post not found response has a 5xx status code
func (o *ListProjectsMembersPostNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this list projects members post not found response a status code equal to that given
func (o *ListProjectsMembersPostNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the list projects members post not found response
func (o *ListProjectsMembersPostNotFound) Code() int {
	return 404
}

func (o *ListProjectsMembersPostNotFound) Error() string {
	return fmt.Sprintf("[POST /api/backend/v1/projects/membersList][%d] listProjectsMembersPostNotFound ", 404)
}

func (o *ListProjectsMembersPostNotFound) String() string {
	return fmt.Sprintf("[POST /api/backend/v1/projects/membersList][%d] listProjectsMembersPostNotFound ", 404)
}

func (o *ListProjectsMembersPostNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListProjectsMembersPostRequestTimeout creates a ListProjectsMembersPostRequestTimeout with default headers values
func NewListProjectsMembersPostRequestTimeout() *ListProjectsMembersPostRequestTimeout {
	return &ListProjectsMembersPostRequestTimeout{}
}

/*
ListProjectsMembersPostRequestTimeout describes a response with status code 408, with default header values.

Request Timeout
*/
type ListProjectsMembersPostRequestTimeout struct {
}

// IsSuccess returns true when this list projects members post request timeout response has a 2xx status code
func (o *ListProjectsMembersPostRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list projects members post request timeout response has a 3xx status code
func (o *ListProjectsMembersPostRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list projects members post request timeout response has a 4xx status code
func (o *ListProjectsMembersPostRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this list projects members post request timeout response has a 5xx status code
func (o *ListProjectsMembersPostRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this list projects members post request timeout response a status code equal to that given
func (o *ListProjectsMembersPostRequestTimeout) IsCode(code int) bool {
	return code == 408
}

// Code gets the status code for the list projects members post request timeout response
func (o *ListProjectsMembersPostRequestTimeout) Code() int {
	return 408
}

func (o *ListProjectsMembersPostRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/backend/v1/projects/membersList][%d] listProjectsMembersPostRequestTimeout ", 408)
}

func (o *ListProjectsMembersPostRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/backend/v1/projects/membersList][%d] listProjectsMembersPostRequestTimeout ", 408)
}

func (o *ListProjectsMembersPostRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListProjectsMembersPostUnprocessableEntity creates a ListProjectsMembersPostUnprocessableEntity with default headers values
func NewListProjectsMembersPostUnprocessableEntity() *ListProjectsMembersPostUnprocessableEntity {
	return &ListProjectsMembersPostUnprocessableEntity{}
}

/*
ListProjectsMembersPostUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type ListProjectsMembersPostUnprocessableEntity struct {
}

// IsSuccess returns true when this list projects members post unprocessable entity response has a 2xx status code
func (o *ListProjectsMembersPostUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list projects members post unprocessable entity response has a 3xx status code
func (o *ListProjectsMembersPostUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list projects members post unprocessable entity response has a 4xx status code
func (o *ListProjectsMembersPostUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this list projects members post unprocessable entity response has a 5xx status code
func (o *ListProjectsMembersPostUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this list projects members post unprocessable entity response a status code equal to that given
func (o *ListProjectsMembersPostUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the list projects members post unprocessable entity response
func (o *ListProjectsMembersPostUnprocessableEntity) Code() int {
	return 422
}

func (o *ListProjectsMembersPostUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /api/backend/v1/projects/membersList][%d] listProjectsMembersPostUnprocessableEntity ", 422)
}

func (o *ListProjectsMembersPostUnprocessableEntity) String() string {
	return fmt.Sprintf("[POST /api/backend/v1/projects/membersList][%d] listProjectsMembersPostUnprocessableEntity ", 422)
}

func (o *ListProjectsMembersPostUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
