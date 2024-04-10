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

// ListMembersReader is a Reader for the ListMembers structure.
type ListMembersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListMembersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListMembersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListMembersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewListMembersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListMembersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListMembersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewListMembersRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewListMembersUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /api/backend/v1/projects/users] listMembers", response, response.Code())
	}
}

// NewListMembersOK creates a ListMembersOK with default headers values
func NewListMembersOK() *ListMembersOK {
	return &ListMembersOK{}
}

/*
ListMembersOK describes a response with status code 200, with default header values.

OK
*/
type ListMembersOK struct {
	Payload *models.UserListDTO
}

// IsSuccess returns true when this list members o k response has a 2xx status code
func (o *ListMembersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list members o k response has a 3xx status code
func (o *ListMembersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list members o k response has a 4xx status code
func (o *ListMembersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list members o k response has a 5xx status code
func (o *ListMembersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list members o k response a status code equal to that given
func (o *ListMembersOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list members o k response
func (o *ListMembersOK) Code() int {
	return 200
}

func (o *ListMembersOK) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/projects/users][%d] listMembersOK  %+v", 200, o.Payload)
}

func (o *ListMembersOK) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/projects/users][%d] listMembersOK  %+v", 200, o.Payload)
}

func (o *ListMembersOK) GetPayload() *models.UserListDTO {
	return o.Payload
}

func (o *ListMembersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserListDTO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListMembersBadRequest creates a ListMembersBadRequest with default headers values
func NewListMembersBadRequest() *ListMembersBadRequest {
	return &ListMembersBadRequest{}
}

/*
ListMembersBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ListMembersBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this list members bad request response has a 2xx status code
func (o *ListMembersBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list members bad request response has a 3xx status code
func (o *ListMembersBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list members bad request response has a 4xx status code
func (o *ListMembersBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this list members bad request response has a 5xx status code
func (o *ListMembersBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this list members bad request response a status code equal to that given
func (o *ListMembersBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the list members bad request response
func (o *ListMembersBadRequest) Code() int {
	return 400
}

func (o *ListMembersBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/projects/users][%d] listMembersBadRequest  %+v", 400, o.Payload)
}

func (o *ListMembersBadRequest) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/projects/users][%d] listMembersBadRequest  %+v", 400, o.Payload)
}

func (o *ListMembersBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListMembersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListMembersUnauthorized creates a ListMembersUnauthorized with default headers values
func NewListMembersUnauthorized() *ListMembersUnauthorized {
	return &ListMembersUnauthorized{}
}

/*
ListMembersUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ListMembersUnauthorized struct {
}

// IsSuccess returns true when this list members unauthorized response has a 2xx status code
func (o *ListMembersUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list members unauthorized response has a 3xx status code
func (o *ListMembersUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list members unauthorized response has a 4xx status code
func (o *ListMembersUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this list members unauthorized response has a 5xx status code
func (o *ListMembersUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this list members unauthorized response a status code equal to that given
func (o *ListMembersUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the list members unauthorized response
func (o *ListMembersUnauthorized) Code() int {
	return 401
}

func (o *ListMembersUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/projects/users][%d] listMembersUnauthorized ", 401)
}

func (o *ListMembersUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/projects/users][%d] listMembersUnauthorized ", 401)
}

func (o *ListMembersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListMembersForbidden creates a ListMembersForbidden with default headers values
func NewListMembersForbidden() *ListMembersForbidden {
	return &ListMembersForbidden{}
}

/*
ListMembersForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ListMembersForbidden struct {
}

// IsSuccess returns true when this list members forbidden response has a 2xx status code
func (o *ListMembersForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list members forbidden response has a 3xx status code
func (o *ListMembersForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list members forbidden response has a 4xx status code
func (o *ListMembersForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list members forbidden response has a 5xx status code
func (o *ListMembersForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list members forbidden response a status code equal to that given
func (o *ListMembersForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the list members forbidden response
func (o *ListMembersForbidden) Code() int {
	return 403
}

func (o *ListMembersForbidden) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/projects/users][%d] listMembersForbidden ", 403)
}

func (o *ListMembersForbidden) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/projects/users][%d] listMembersForbidden ", 403)
}

func (o *ListMembersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListMembersNotFound creates a ListMembersNotFound with default headers values
func NewListMembersNotFound() *ListMembersNotFound {
	return &ListMembersNotFound{}
}

/*
ListMembersNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ListMembersNotFound struct {
}

// IsSuccess returns true when this list members not found response has a 2xx status code
func (o *ListMembersNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list members not found response has a 3xx status code
func (o *ListMembersNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list members not found response has a 4xx status code
func (o *ListMembersNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this list members not found response has a 5xx status code
func (o *ListMembersNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this list members not found response a status code equal to that given
func (o *ListMembersNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the list members not found response
func (o *ListMembersNotFound) Code() int {
	return 404
}

func (o *ListMembersNotFound) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/projects/users][%d] listMembersNotFound ", 404)
}

func (o *ListMembersNotFound) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/projects/users][%d] listMembersNotFound ", 404)
}

func (o *ListMembersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListMembersRequestTimeout creates a ListMembersRequestTimeout with default headers values
func NewListMembersRequestTimeout() *ListMembersRequestTimeout {
	return &ListMembersRequestTimeout{}
}

/*
ListMembersRequestTimeout describes a response with status code 408, with default header values.

Request Timeout
*/
type ListMembersRequestTimeout struct {
}

// IsSuccess returns true when this list members request timeout response has a 2xx status code
func (o *ListMembersRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list members request timeout response has a 3xx status code
func (o *ListMembersRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list members request timeout response has a 4xx status code
func (o *ListMembersRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this list members request timeout response has a 5xx status code
func (o *ListMembersRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this list members request timeout response a status code equal to that given
func (o *ListMembersRequestTimeout) IsCode(code int) bool {
	return code == 408
}

// Code gets the status code for the list members request timeout response
func (o *ListMembersRequestTimeout) Code() int {
	return 408
}

func (o *ListMembersRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/projects/users][%d] listMembersRequestTimeout ", 408)
}

func (o *ListMembersRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/projects/users][%d] listMembersRequestTimeout ", 408)
}

func (o *ListMembersRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListMembersUnprocessableEntity creates a ListMembersUnprocessableEntity with default headers values
func NewListMembersUnprocessableEntity() *ListMembersUnprocessableEntity {
	return &ListMembersUnprocessableEntity{}
}

/*
ListMembersUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type ListMembersUnprocessableEntity struct {
}

// IsSuccess returns true when this list members unprocessable entity response has a 2xx status code
func (o *ListMembersUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list members unprocessable entity response has a 3xx status code
func (o *ListMembersUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list members unprocessable entity response has a 4xx status code
func (o *ListMembersUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this list members unprocessable entity response has a 5xx status code
func (o *ListMembersUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this list members unprocessable entity response a status code equal to that given
func (o *ListMembersUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the list members unprocessable entity response
func (o *ListMembersUnprocessableEntity) Code() int {
	return 422
}

func (o *ListMembersUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/projects/users][%d] listMembersUnprocessableEntity ", 422)
}

func (o *ListMembersUnprocessableEntity) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/projects/users][%d] listMembersUnprocessableEntity ", 422)
}

func (o *ListMembersUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
