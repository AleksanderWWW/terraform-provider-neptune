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

// AddProjectMemberReader is a Reader for the AddProjectMember structure.
type AddProjectMemberReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddProjectMemberReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddProjectMemberOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAddProjectMemberBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAddProjectMemberUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAddProjectMemberForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAddProjectMemberNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewAddProjectMemberRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewAddProjectMemberConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewAddProjectMemberUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /api/backend/v1/projects/members] addProjectMember", response, response.Code())
	}
}

// NewAddProjectMemberOK creates a AddProjectMemberOK with default headers values
func NewAddProjectMemberOK() *AddProjectMemberOK {
	return &AddProjectMemberOK{}
}

/*
AddProjectMemberOK describes a response with status code 200, with default header values.

OK
*/
type AddProjectMemberOK struct {
	Payload *models.ProjectMemberDTO
}

// IsSuccess returns true when this add project member o k response has a 2xx status code
func (o *AddProjectMemberOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this add project member o k response has a 3xx status code
func (o *AddProjectMemberOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add project member o k response has a 4xx status code
func (o *AddProjectMemberOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this add project member o k response has a 5xx status code
func (o *AddProjectMemberOK) IsServerError() bool {
	return false
}

// IsCode returns true when this add project member o k response a status code equal to that given
func (o *AddProjectMemberOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the add project member o k response
func (o *AddProjectMemberOK) Code() int {
	return 200
}

func (o *AddProjectMemberOK) Error() string {
	return fmt.Sprintf("[POST /api/backend/v1/projects/members][%d] addProjectMemberOK  %+v", 200, o.Payload)
}

func (o *AddProjectMemberOK) String() string {
	return fmt.Sprintf("[POST /api/backend/v1/projects/members][%d] addProjectMemberOK  %+v", 200, o.Payload)
}

func (o *AddProjectMemberOK) GetPayload() *models.ProjectMemberDTO {
	return o.Payload
}

func (o *AddProjectMemberOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProjectMemberDTO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddProjectMemberBadRequest creates a AddProjectMemberBadRequest with default headers values
func NewAddProjectMemberBadRequest() *AddProjectMemberBadRequest {
	return &AddProjectMemberBadRequest{}
}

/*
AddProjectMemberBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AddProjectMemberBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this add project member bad request response has a 2xx status code
func (o *AddProjectMemberBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add project member bad request response has a 3xx status code
func (o *AddProjectMemberBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add project member bad request response has a 4xx status code
func (o *AddProjectMemberBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this add project member bad request response has a 5xx status code
func (o *AddProjectMemberBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this add project member bad request response a status code equal to that given
func (o *AddProjectMemberBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the add project member bad request response
func (o *AddProjectMemberBadRequest) Code() int {
	return 400
}

func (o *AddProjectMemberBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/backend/v1/projects/members][%d] addProjectMemberBadRequest  %+v", 400, o.Payload)
}

func (o *AddProjectMemberBadRequest) String() string {
	return fmt.Sprintf("[POST /api/backend/v1/projects/members][%d] addProjectMemberBadRequest  %+v", 400, o.Payload)
}

func (o *AddProjectMemberBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *AddProjectMemberBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddProjectMemberUnauthorized creates a AddProjectMemberUnauthorized with default headers values
func NewAddProjectMemberUnauthorized() *AddProjectMemberUnauthorized {
	return &AddProjectMemberUnauthorized{}
}

/*
AddProjectMemberUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type AddProjectMemberUnauthorized struct {
}

// IsSuccess returns true when this add project member unauthorized response has a 2xx status code
func (o *AddProjectMemberUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add project member unauthorized response has a 3xx status code
func (o *AddProjectMemberUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add project member unauthorized response has a 4xx status code
func (o *AddProjectMemberUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this add project member unauthorized response has a 5xx status code
func (o *AddProjectMemberUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this add project member unauthorized response a status code equal to that given
func (o *AddProjectMemberUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the add project member unauthorized response
func (o *AddProjectMemberUnauthorized) Code() int {
	return 401
}

func (o *AddProjectMemberUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/backend/v1/projects/members][%d] addProjectMemberUnauthorized ", 401)
}

func (o *AddProjectMemberUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/backend/v1/projects/members][%d] addProjectMemberUnauthorized ", 401)
}

func (o *AddProjectMemberUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddProjectMemberForbidden creates a AddProjectMemberForbidden with default headers values
func NewAddProjectMemberForbidden() *AddProjectMemberForbidden {
	return &AddProjectMemberForbidden{}
}

/*
AddProjectMemberForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type AddProjectMemberForbidden struct {
}

// IsSuccess returns true when this add project member forbidden response has a 2xx status code
func (o *AddProjectMemberForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add project member forbidden response has a 3xx status code
func (o *AddProjectMemberForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add project member forbidden response has a 4xx status code
func (o *AddProjectMemberForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this add project member forbidden response has a 5xx status code
func (o *AddProjectMemberForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this add project member forbidden response a status code equal to that given
func (o *AddProjectMemberForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the add project member forbidden response
func (o *AddProjectMemberForbidden) Code() int {
	return 403
}

func (o *AddProjectMemberForbidden) Error() string {
	return fmt.Sprintf("[POST /api/backend/v1/projects/members][%d] addProjectMemberForbidden ", 403)
}

func (o *AddProjectMemberForbidden) String() string {
	return fmt.Sprintf("[POST /api/backend/v1/projects/members][%d] addProjectMemberForbidden ", 403)
}

func (o *AddProjectMemberForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddProjectMemberNotFound creates a AddProjectMemberNotFound with default headers values
func NewAddProjectMemberNotFound() *AddProjectMemberNotFound {
	return &AddProjectMemberNotFound{}
}

/*
AddProjectMemberNotFound describes a response with status code 404, with default header values.

Not Found
*/
type AddProjectMemberNotFound struct {
}

// IsSuccess returns true when this add project member not found response has a 2xx status code
func (o *AddProjectMemberNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add project member not found response has a 3xx status code
func (o *AddProjectMemberNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add project member not found response has a 4xx status code
func (o *AddProjectMemberNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this add project member not found response has a 5xx status code
func (o *AddProjectMemberNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this add project member not found response a status code equal to that given
func (o *AddProjectMemberNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the add project member not found response
func (o *AddProjectMemberNotFound) Code() int {
	return 404
}

func (o *AddProjectMemberNotFound) Error() string {
	return fmt.Sprintf("[POST /api/backend/v1/projects/members][%d] addProjectMemberNotFound ", 404)
}

func (o *AddProjectMemberNotFound) String() string {
	return fmt.Sprintf("[POST /api/backend/v1/projects/members][%d] addProjectMemberNotFound ", 404)
}

func (o *AddProjectMemberNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddProjectMemberRequestTimeout creates a AddProjectMemberRequestTimeout with default headers values
func NewAddProjectMemberRequestTimeout() *AddProjectMemberRequestTimeout {
	return &AddProjectMemberRequestTimeout{}
}

/*
AddProjectMemberRequestTimeout describes a response with status code 408, with default header values.

Request Timeout
*/
type AddProjectMemberRequestTimeout struct {
}

// IsSuccess returns true when this add project member request timeout response has a 2xx status code
func (o *AddProjectMemberRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add project member request timeout response has a 3xx status code
func (o *AddProjectMemberRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add project member request timeout response has a 4xx status code
func (o *AddProjectMemberRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this add project member request timeout response has a 5xx status code
func (o *AddProjectMemberRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this add project member request timeout response a status code equal to that given
func (o *AddProjectMemberRequestTimeout) IsCode(code int) bool {
	return code == 408
}

// Code gets the status code for the add project member request timeout response
func (o *AddProjectMemberRequestTimeout) Code() int {
	return 408
}

func (o *AddProjectMemberRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/backend/v1/projects/members][%d] addProjectMemberRequestTimeout ", 408)
}

func (o *AddProjectMemberRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/backend/v1/projects/members][%d] addProjectMemberRequestTimeout ", 408)
}

func (o *AddProjectMemberRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddProjectMemberConflict creates a AddProjectMemberConflict with default headers values
func NewAddProjectMemberConflict() *AddProjectMemberConflict {
	return &AddProjectMemberConflict{}
}

/*
AddProjectMemberConflict describes a response with status code 409, with default header values.

Conflict
*/
type AddProjectMemberConflict struct {
}

// IsSuccess returns true when this add project member conflict response has a 2xx status code
func (o *AddProjectMemberConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add project member conflict response has a 3xx status code
func (o *AddProjectMemberConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add project member conflict response has a 4xx status code
func (o *AddProjectMemberConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this add project member conflict response has a 5xx status code
func (o *AddProjectMemberConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this add project member conflict response a status code equal to that given
func (o *AddProjectMemberConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the add project member conflict response
func (o *AddProjectMemberConflict) Code() int {
	return 409
}

func (o *AddProjectMemberConflict) Error() string {
	return fmt.Sprintf("[POST /api/backend/v1/projects/members][%d] addProjectMemberConflict ", 409)
}

func (o *AddProjectMemberConflict) String() string {
	return fmt.Sprintf("[POST /api/backend/v1/projects/members][%d] addProjectMemberConflict ", 409)
}

func (o *AddProjectMemberConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddProjectMemberUnprocessableEntity creates a AddProjectMemberUnprocessableEntity with default headers values
func NewAddProjectMemberUnprocessableEntity() *AddProjectMemberUnprocessableEntity {
	return &AddProjectMemberUnprocessableEntity{}
}

/*
AddProjectMemberUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type AddProjectMemberUnprocessableEntity struct {
}

// IsSuccess returns true when this add project member unprocessable entity response has a 2xx status code
func (o *AddProjectMemberUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add project member unprocessable entity response has a 3xx status code
func (o *AddProjectMemberUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add project member unprocessable entity response has a 4xx status code
func (o *AddProjectMemberUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this add project member unprocessable entity response has a 5xx status code
func (o *AddProjectMemberUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this add project member unprocessable entity response a status code equal to that given
func (o *AddProjectMemberUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the add project member unprocessable entity response
func (o *AddProjectMemberUnprocessableEntity) Code() int {
	return 422
}

func (o *AddProjectMemberUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /api/backend/v1/projects/members][%d] addProjectMemberUnprocessableEntity ", 422)
}

func (o *AddProjectMemberUnprocessableEntity) String() string {
	return fmt.Sprintf("[POST /api/backend/v1/projects/members][%d] addProjectMemberUnprocessableEntity ", 422)
}

func (o *AddProjectMemberUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}