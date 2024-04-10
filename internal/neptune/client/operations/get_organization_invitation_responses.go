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

// GetOrganizationInvitationReader is a Reader for the GetOrganizationInvitation structure.
type GetOrganizationInvitationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrganizationInvitationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrganizationInvitationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetOrganizationInvitationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetOrganizationInvitationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetOrganizationInvitationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetOrganizationInvitationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetOrganizationInvitationRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGetOrganizationInvitationUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /api/backend/v1/invitations/organization/{invitationId}] getOrganizationInvitation", response, response.Code())
	}
}

// NewGetOrganizationInvitationOK creates a GetOrganizationInvitationOK with default headers values
func NewGetOrganizationInvitationOK() *GetOrganizationInvitationOK {
	return &GetOrganizationInvitationOK{}
}

/*
GetOrganizationInvitationOK describes a response with status code 200, with default header values.

OK
*/
type GetOrganizationInvitationOK struct {
	Payload *models.OrganizationInvitationDTO
}

// IsSuccess returns true when this get organization invitation o k response has a 2xx status code
func (o *GetOrganizationInvitationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get organization invitation o k response has a 3xx status code
func (o *GetOrganizationInvitationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organization invitation o k response has a 4xx status code
func (o *GetOrganizationInvitationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get organization invitation o k response has a 5xx status code
func (o *GetOrganizationInvitationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get organization invitation o k response a status code equal to that given
func (o *GetOrganizationInvitationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get organization invitation o k response
func (o *GetOrganizationInvitationOK) Code() int {
	return 200
}

func (o *GetOrganizationInvitationOK) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/invitations/organization/{invitationId}][%d] getOrganizationInvitationOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationInvitationOK) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/invitations/organization/{invitationId}][%d] getOrganizationInvitationOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationInvitationOK) GetPayload() *models.OrganizationInvitationDTO {
	return o.Payload
}

func (o *GetOrganizationInvitationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OrganizationInvitationDTO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationInvitationBadRequest creates a GetOrganizationInvitationBadRequest with default headers values
func NewGetOrganizationInvitationBadRequest() *GetOrganizationInvitationBadRequest {
	return &GetOrganizationInvitationBadRequest{}
}

/*
GetOrganizationInvitationBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetOrganizationInvitationBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this get organization invitation bad request response has a 2xx status code
func (o *GetOrganizationInvitationBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get organization invitation bad request response has a 3xx status code
func (o *GetOrganizationInvitationBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organization invitation bad request response has a 4xx status code
func (o *GetOrganizationInvitationBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get organization invitation bad request response has a 5xx status code
func (o *GetOrganizationInvitationBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get organization invitation bad request response a status code equal to that given
func (o *GetOrganizationInvitationBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get organization invitation bad request response
func (o *GetOrganizationInvitationBadRequest) Code() int {
	return 400
}

func (o *GetOrganizationInvitationBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/invitations/organization/{invitationId}][%d] getOrganizationInvitationBadRequest  %+v", 400, o.Payload)
}

func (o *GetOrganizationInvitationBadRequest) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/invitations/organization/{invitationId}][%d] getOrganizationInvitationBadRequest  %+v", 400, o.Payload)
}

func (o *GetOrganizationInvitationBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetOrganizationInvitationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationInvitationUnauthorized creates a GetOrganizationInvitationUnauthorized with default headers values
func NewGetOrganizationInvitationUnauthorized() *GetOrganizationInvitationUnauthorized {
	return &GetOrganizationInvitationUnauthorized{}
}

/*
GetOrganizationInvitationUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetOrganizationInvitationUnauthorized struct {
}

// IsSuccess returns true when this get organization invitation unauthorized response has a 2xx status code
func (o *GetOrganizationInvitationUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get organization invitation unauthorized response has a 3xx status code
func (o *GetOrganizationInvitationUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organization invitation unauthorized response has a 4xx status code
func (o *GetOrganizationInvitationUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get organization invitation unauthorized response has a 5xx status code
func (o *GetOrganizationInvitationUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get organization invitation unauthorized response a status code equal to that given
func (o *GetOrganizationInvitationUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get organization invitation unauthorized response
func (o *GetOrganizationInvitationUnauthorized) Code() int {
	return 401
}

func (o *GetOrganizationInvitationUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/invitations/organization/{invitationId}][%d] getOrganizationInvitationUnauthorized ", 401)
}

func (o *GetOrganizationInvitationUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/invitations/organization/{invitationId}][%d] getOrganizationInvitationUnauthorized ", 401)
}

func (o *GetOrganizationInvitationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetOrganizationInvitationForbidden creates a GetOrganizationInvitationForbidden with default headers values
func NewGetOrganizationInvitationForbidden() *GetOrganizationInvitationForbidden {
	return &GetOrganizationInvitationForbidden{}
}

/*
GetOrganizationInvitationForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetOrganizationInvitationForbidden struct {
}

// IsSuccess returns true when this get organization invitation forbidden response has a 2xx status code
func (o *GetOrganizationInvitationForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get organization invitation forbidden response has a 3xx status code
func (o *GetOrganizationInvitationForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organization invitation forbidden response has a 4xx status code
func (o *GetOrganizationInvitationForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get organization invitation forbidden response has a 5xx status code
func (o *GetOrganizationInvitationForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get organization invitation forbidden response a status code equal to that given
func (o *GetOrganizationInvitationForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get organization invitation forbidden response
func (o *GetOrganizationInvitationForbidden) Code() int {
	return 403
}

func (o *GetOrganizationInvitationForbidden) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/invitations/organization/{invitationId}][%d] getOrganizationInvitationForbidden ", 403)
}

func (o *GetOrganizationInvitationForbidden) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/invitations/organization/{invitationId}][%d] getOrganizationInvitationForbidden ", 403)
}

func (o *GetOrganizationInvitationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetOrganizationInvitationNotFound creates a GetOrganizationInvitationNotFound with default headers values
func NewGetOrganizationInvitationNotFound() *GetOrganizationInvitationNotFound {
	return &GetOrganizationInvitationNotFound{}
}

/*
GetOrganizationInvitationNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetOrganizationInvitationNotFound struct {
}

// IsSuccess returns true when this get organization invitation not found response has a 2xx status code
func (o *GetOrganizationInvitationNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get organization invitation not found response has a 3xx status code
func (o *GetOrganizationInvitationNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organization invitation not found response has a 4xx status code
func (o *GetOrganizationInvitationNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get organization invitation not found response has a 5xx status code
func (o *GetOrganizationInvitationNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get organization invitation not found response a status code equal to that given
func (o *GetOrganizationInvitationNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get organization invitation not found response
func (o *GetOrganizationInvitationNotFound) Code() int {
	return 404
}

func (o *GetOrganizationInvitationNotFound) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/invitations/organization/{invitationId}][%d] getOrganizationInvitationNotFound ", 404)
}

func (o *GetOrganizationInvitationNotFound) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/invitations/organization/{invitationId}][%d] getOrganizationInvitationNotFound ", 404)
}

func (o *GetOrganizationInvitationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetOrganizationInvitationRequestTimeout creates a GetOrganizationInvitationRequestTimeout with default headers values
func NewGetOrganizationInvitationRequestTimeout() *GetOrganizationInvitationRequestTimeout {
	return &GetOrganizationInvitationRequestTimeout{}
}

/*
GetOrganizationInvitationRequestTimeout describes a response with status code 408, with default header values.

Request Timeout
*/
type GetOrganizationInvitationRequestTimeout struct {
}

// IsSuccess returns true when this get organization invitation request timeout response has a 2xx status code
func (o *GetOrganizationInvitationRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get organization invitation request timeout response has a 3xx status code
func (o *GetOrganizationInvitationRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organization invitation request timeout response has a 4xx status code
func (o *GetOrganizationInvitationRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get organization invitation request timeout response has a 5xx status code
func (o *GetOrganizationInvitationRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get organization invitation request timeout response a status code equal to that given
func (o *GetOrganizationInvitationRequestTimeout) IsCode(code int) bool {
	return code == 408
}

// Code gets the status code for the get organization invitation request timeout response
func (o *GetOrganizationInvitationRequestTimeout) Code() int {
	return 408
}

func (o *GetOrganizationInvitationRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/invitations/organization/{invitationId}][%d] getOrganizationInvitationRequestTimeout ", 408)
}

func (o *GetOrganizationInvitationRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/invitations/organization/{invitationId}][%d] getOrganizationInvitationRequestTimeout ", 408)
}

func (o *GetOrganizationInvitationRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetOrganizationInvitationUnprocessableEntity creates a GetOrganizationInvitationUnprocessableEntity with default headers values
func NewGetOrganizationInvitationUnprocessableEntity() *GetOrganizationInvitationUnprocessableEntity {
	return &GetOrganizationInvitationUnprocessableEntity{}
}

/*
GetOrganizationInvitationUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type GetOrganizationInvitationUnprocessableEntity struct {
}

// IsSuccess returns true when this get organization invitation unprocessable entity response has a 2xx status code
func (o *GetOrganizationInvitationUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get organization invitation unprocessable entity response has a 3xx status code
func (o *GetOrganizationInvitationUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organization invitation unprocessable entity response has a 4xx status code
func (o *GetOrganizationInvitationUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this get organization invitation unprocessable entity response has a 5xx status code
func (o *GetOrganizationInvitationUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this get organization invitation unprocessable entity response a status code equal to that given
func (o *GetOrganizationInvitationUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the get organization invitation unprocessable entity response
func (o *GetOrganizationInvitationUnprocessableEntity) Code() int {
	return 422
}

func (o *GetOrganizationInvitationUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/invitations/organization/{invitationId}][%d] getOrganizationInvitationUnprocessableEntity ", 422)
}

func (o *GetOrganizationInvitationUnprocessableEntity) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/invitations/organization/{invitationId}][%d] getOrganizationInvitationUnprocessableEntity ", 422)
}

func (o *GetOrganizationInvitationUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
