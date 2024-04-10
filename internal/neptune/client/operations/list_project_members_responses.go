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

// ListProjectMembersReader is a Reader for the ListProjectMembers structure.
type ListProjectMembersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListProjectMembersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListProjectMembersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListProjectMembersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewListProjectMembersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListProjectMembersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListProjectMembersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewListProjectMembersRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewListProjectMembersUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /api/backend/v1/projects/membersOf] listProjectMembers", response, response.Code())
	}
}

// NewListProjectMembersOK creates a ListProjectMembersOK with default headers values
func NewListProjectMembersOK() *ListProjectMembersOK {
	return &ListProjectMembersOK{}
}

/*
ListProjectMembersOK describes a response with status code 200, with default header values.

OK
*/
type ListProjectMembersOK struct {
	Payload []*models.ProjectMemberDTO
}

// IsSuccess returns true when this list project members o k response has a 2xx status code
func (o *ListProjectMembersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list project members o k response has a 3xx status code
func (o *ListProjectMembersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list project members o k response has a 4xx status code
func (o *ListProjectMembersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list project members o k response has a 5xx status code
func (o *ListProjectMembersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list project members o k response a status code equal to that given
func (o *ListProjectMembersOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list project members o k response
func (o *ListProjectMembersOK) Code() int {
	return 200
}

func (o *ListProjectMembersOK) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/projects/membersOf][%d] listProjectMembersOK  %+v", 200, o.Payload)
}

func (o *ListProjectMembersOK) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/projects/membersOf][%d] listProjectMembersOK  %+v", 200, o.Payload)
}

func (o *ListProjectMembersOK) GetPayload() []*models.ProjectMemberDTO {
	return o.Payload
}

func (o *ListProjectMembersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListProjectMembersBadRequest creates a ListProjectMembersBadRequest with default headers values
func NewListProjectMembersBadRequest() *ListProjectMembersBadRequest {
	return &ListProjectMembersBadRequest{}
}

/*
ListProjectMembersBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ListProjectMembersBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this list project members bad request response has a 2xx status code
func (o *ListProjectMembersBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list project members bad request response has a 3xx status code
func (o *ListProjectMembersBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list project members bad request response has a 4xx status code
func (o *ListProjectMembersBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this list project members bad request response has a 5xx status code
func (o *ListProjectMembersBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this list project members bad request response a status code equal to that given
func (o *ListProjectMembersBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the list project members bad request response
func (o *ListProjectMembersBadRequest) Code() int {
	return 400
}

func (o *ListProjectMembersBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/projects/membersOf][%d] listProjectMembersBadRequest  %+v", 400, o.Payload)
}

func (o *ListProjectMembersBadRequest) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/projects/membersOf][%d] listProjectMembersBadRequest  %+v", 400, o.Payload)
}

func (o *ListProjectMembersBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListProjectMembersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListProjectMembersUnauthorized creates a ListProjectMembersUnauthorized with default headers values
func NewListProjectMembersUnauthorized() *ListProjectMembersUnauthorized {
	return &ListProjectMembersUnauthorized{}
}

/*
ListProjectMembersUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ListProjectMembersUnauthorized struct {
}

// IsSuccess returns true when this list project members unauthorized response has a 2xx status code
func (o *ListProjectMembersUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list project members unauthorized response has a 3xx status code
func (o *ListProjectMembersUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list project members unauthorized response has a 4xx status code
func (o *ListProjectMembersUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this list project members unauthorized response has a 5xx status code
func (o *ListProjectMembersUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this list project members unauthorized response a status code equal to that given
func (o *ListProjectMembersUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the list project members unauthorized response
func (o *ListProjectMembersUnauthorized) Code() int {
	return 401
}

func (o *ListProjectMembersUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/projects/membersOf][%d] listProjectMembersUnauthorized ", 401)
}

func (o *ListProjectMembersUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/projects/membersOf][%d] listProjectMembersUnauthorized ", 401)
}

func (o *ListProjectMembersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListProjectMembersForbidden creates a ListProjectMembersForbidden with default headers values
func NewListProjectMembersForbidden() *ListProjectMembersForbidden {
	return &ListProjectMembersForbidden{}
}

/*
ListProjectMembersForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ListProjectMembersForbidden struct {
}

// IsSuccess returns true when this list project members forbidden response has a 2xx status code
func (o *ListProjectMembersForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list project members forbidden response has a 3xx status code
func (o *ListProjectMembersForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list project members forbidden response has a 4xx status code
func (o *ListProjectMembersForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list project members forbidden response has a 5xx status code
func (o *ListProjectMembersForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list project members forbidden response a status code equal to that given
func (o *ListProjectMembersForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the list project members forbidden response
func (o *ListProjectMembersForbidden) Code() int {
	return 403
}

func (o *ListProjectMembersForbidden) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/projects/membersOf][%d] listProjectMembersForbidden ", 403)
}

func (o *ListProjectMembersForbidden) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/projects/membersOf][%d] listProjectMembersForbidden ", 403)
}

func (o *ListProjectMembersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListProjectMembersNotFound creates a ListProjectMembersNotFound with default headers values
func NewListProjectMembersNotFound() *ListProjectMembersNotFound {
	return &ListProjectMembersNotFound{}
}

/*
ListProjectMembersNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ListProjectMembersNotFound struct {
}

// IsSuccess returns true when this list project members not found response has a 2xx status code
func (o *ListProjectMembersNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list project members not found response has a 3xx status code
func (o *ListProjectMembersNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list project members not found response has a 4xx status code
func (o *ListProjectMembersNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this list project members not found response has a 5xx status code
func (o *ListProjectMembersNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this list project members not found response a status code equal to that given
func (o *ListProjectMembersNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the list project members not found response
func (o *ListProjectMembersNotFound) Code() int {
	return 404
}

func (o *ListProjectMembersNotFound) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/projects/membersOf][%d] listProjectMembersNotFound ", 404)
}

func (o *ListProjectMembersNotFound) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/projects/membersOf][%d] listProjectMembersNotFound ", 404)
}

func (o *ListProjectMembersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListProjectMembersRequestTimeout creates a ListProjectMembersRequestTimeout with default headers values
func NewListProjectMembersRequestTimeout() *ListProjectMembersRequestTimeout {
	return &ListProjectMembersRequestTimeout{}
}

/*
ListProjectMembersRequestTimeout describes a response with status code 408, with default header values.

Request Timeout
*/
type ListProjectMembersRequestTimeout struct {
}

// IsSuccess returns true when this list project members request timeout response has a 2xx status code
func (o *ListProjectMembersRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list project members request timeout response has a 3xx status code
func (o *ListProjectMembersRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list project members request timeout response has a 4xx status code
func (o *ListProjectMembersRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this list project members request timeout response has a 5xx status code
func (o *ListProjectMembersRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this list project members request timeout response a status code equal to that given
func (o *ListProjectMembersRequestTimeout) IsCode(code int) bool {
	return code == 408
}

// Code gets the status code for the list project members request timeout response
func (o *ListProjectMembersRequestTimeout) Code() int {
	return 408
}

func (o *ListProjectMembersRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/projects/membersOf][%d] listProjectMembersRequestTimeout ", 408)
}

func (o *ListProjectMembersRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/projects/membersOf][%d] listProjectMembersRequestTimeout ", 408)
}

func (o *ListProjectMembersRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListProjectMembersUnprocessableEntity creates a ListProjectMembersUnprocessableEntity with default headers values
func NewListProjectMembersUnprocessableEntity() *ListProjectMembersUnprocessableEntity {
	return &ListProjectMembersUnprocessableEntity{}
}

/*
ListProjectMembersUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type ListProjectMembersUnprocessableEntity struct {
}

// IsSuccess returns true when this list project members unprocessable entity response has a 2xx status code
func (o *ListProjectMembersUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list project members unprocessable entity response has a 3xx status code
func (o *ListProjectMembersUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list project members unprocessable entity response has a 4xx status code
func (o *ListProjectMembersUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this list project members unprocessable entity response has a 5xx status code
func (o *ListProjectMembersUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this list project members unprocessable entity response a status code equal to that given
func (o *ListProjectMembersUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the list project members unprocessable entity response
func (o *ListProjectMembersUnprocessableEntity) Code() int {
	return 422
}

func (o *ListProjectMembersUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/projects/membersOf][%d] listProjectMembersUnprocessableEntity ", 422)
}

func (o *ListProjectMembersUnprocessableEntity) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/projects/membersOf][%d] listProjectMembersUnprocessableEntity ", 422)
}

func (o *ListProjectMembersUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
