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

// UpdateOrganizationInvitationReader is a Reader for the UpdateOrganizationInvitation structure.
type UpdateOrganizationInvitationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateOrganizationInvitationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateOrganizationInvitationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateOrganizationInvitationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateOrganizationInvitationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateOrganizationInvitationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateOrganizationInvitationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewUpdateOrganizationInvitationRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewUpdateOrganizationInvitationUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /api/backend/v1/invitations/organization/{invitationId}] updateOrganizationInvitation", response, response.Code())
	}
}

// NewUpdateOrganizationInvitationOK creates a UpdateOrganizationInvitationOK with default headers values
func NewUpdateOrganizationInvitationOK() *UpdateOrganizationInvitationOK {
	return &UpdateOrganizationInvitationOK{}
}

/*
UpdateOrganizationInvitationOK describes a response with status code 200, with default header values.

OK
*/
type UpdateOrganizationInvitationOK struct {
	Payload *models.OrganizationInvitationDTO
}

// IsSuccess returns true when this update organization invitation o k response has a 2xx status code
func (o *UpdateOrganizationInvitationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update organization invitation o k response has a 3xx status code
func (o *UpdateOrganizationInvitationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update organization invitation o k response has a 4xx status code
func (o *UpdateOrganizationInvitationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update organization invitation o k response has a 5xx status code
func (o *UpdateOrganizationInvitationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update organization invitation o k response a status code equal to that given
func (o *UpdateOrganizationInvitationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update organization invitation o k response
func (o *UpdateOrganizationInvitationOK) Code() int {
	return 200
}

func (o *UpdateOrganizationInvitationOK) Error() string {
	return fmt.Sprintf("[PUT /api/backend/v1/invitations/organization/{invitationId}][%d] updateOrganizationInvitationOK  %+v", 200, o.Payload)
}

func (o *UpdateOrganizationInvitationOK) String() string {
	return fmt.Sprintf("[PUT /api/backend/v1/invitations/organization/{invitationId}][%d] updateOrganizationInvitationOK  %+v", 200, o.Payload)
}

func (o *UpdateOrganizationInvitationOK) GetPayload() *models.OrganizationInvitationDTO {
	return o.Payload
}

func (o *UpdateOrganizationInvitationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OrganizationInvitationDTO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateOrganizationInvitationBadRequest creates a UpdateOrganizationInvitationBadRequest with default headers values
func NewUpdateOrganizationInvitationBadRequest() *UpdateOrganizationInvitationBadRequest {
	return &UpdateOrganizationInvitationBadRequest{}
}

/*
UpdateOrganizationInvitationBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UpdateOrganizationInvitationBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this update organization invitation bad request response has a 2xx status code
func (o *UpdateOrganizationInvitationBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update organization invitation bad request response has a 3xx status code
func (o *UpdateOrganizationInvitationBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update organization invitation bad request response has a 4xx status code
func (o *UpdateOrganizationInvitationBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update organization invitation bad request response has a 5xx status code
func (o *UpdateOrganizationInvitationBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update organization invitation bad request response a status code equal to that given
func (o *UpdateOrganizationInvitationBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update organization invitation bad request response
func (o *UpdateOrganizationInvitationBadRequest) Code() int {
	return 400
}

func (o *UpdateOrganizationInvitationBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/backend/v1/invitations/organization/{invitationId}][%d] updateOrganizationInvitationBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateOrganizationInvitationBadRequest) String() string {
	return fmt.Sprintf("[PUT /api/backend/v1/invitations/organization/{invitationId}][%d] updateOrganizationInvitationBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateOrganizationInvitationBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateOrganizationInvitationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateOrganizationInvitationUnauthorized creates a UpdateOrganizationInvitationUnauthorized with default headers values
func NewUpdateOrganizationInvitationUnauthorized() *UpdateOrganizationInvitationUnauthorized {
	return &UpdateOrganizationInvitationUnauthorized{}
}

/*
UpdateOrganizationInvitationUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type UpdateOrganizationInvitationUnauthorized struct {
}

// IsSuccess returns true when this update organization invitation unauthorized response has a 2xx status code
func (o *UpdateOrganizationInvitationUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update organization invitation unauthorized response has a 3xx status code
func (o *UpdateOrganizationInvitationUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update organization invitation unauthorized response has a 4xx status code
func (o *UpdateOrganizationInvitationUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update organization invitation unauthorized response has a 5xx status code
func (o *UpdateOrganizationInvitationUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update organization invitation unauthorized response a status code equal to that given
func (o *UpdateOrganizationInvitationUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the update organization invitation unauthorized response
func (o *UpdateOrganizationInvitationUnauthorized) Code() int {
	return 401
}

func (o *UpdateOrganizationInvitationUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/backend/v1/invitations/organization/{invitationId}][%d] updateOrganizationInvitationUnauthorized ", 401)
}

func (o *UpdateOrganizationInvitationUnauthorized) String() string {
	return fmt.Sprintf("[PUT /api/backend/v1/invitations/organization/{invitationId}][%d] updateOrganizationInvitationUnauthorized ", 401)
}

func (o *UpdateOrganizationInvitationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateOrganizationInvitationForbidden creates a UpdateOrganizationInvitationForbidden with default headers values
func NewUpdateOrganizationInvitationForbidden() *UpdateOrganizationInvitationForbidden {
	return &UpdateOrganizationInvitationForbidden{}
}

/*
UpdateOrganizationInvitationForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UpdateOrganizationInvitationForbidden struct {
}

// IsSuccess returns true when this update organization invitation forbidden response has a 2xx status code
func (o *UpdateOrganizationInvitationForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update organization invitation forbidden response has a 3xx status code
func (o *UpdateOrganizationInvitationForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update organization invitation forbidden response has a 4xx status code
func (o *UpdateOrganizationInvitationForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update organization invitation forbidden response has a 5xx status code
func (o *UpdateOrganizationInvitationForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update organization invitation forbidden response a status code equal to that given
func (o *UpdateOrganizationInvitationForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the update organization invitation forbidden response
func (o *UpdateOrganizationInvitationForbidden) Code() int {
	return 403
}

func (o *UpdateOrganizationInvitationForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/backend/v1/invitations/organization/{invitationId}][%d] updateOrganizationInvitationForbidden ", 403)
}

func (o *UpdateOrganizationInvitationForbidden) String() string {
	return fmt.Sprintf("[PUT /api/backend/v1/invitations/organization/{invitationId}][%d] updateOrganizationInvitationForbidden ", 403)
}

func (o *UpdateOrganizationInvitationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateOrganizationInvitationNotFound creates a UpdateOrganizationInvitationNotFound with default headers values
func NewUpdateOrganizationInvitationNotFound() *UpdateOrganizationInvitationNotFound {
	return &UpdateOrganizationInvitationNotFound{}
}

/*
UpdateOrganizationInvitationNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UpdateOrganizationInvitationNotFound struct {
}

// IsSuccess returns true when this update organization invitation not found response has a 2xx status code
func (o *UpdateOrganizationInvitationNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update organization invitation not found response has a 3xx status code
func (o *UpdateOrganizationInvitationNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update organization invitation not found response has a 4xx status code
func (o *UpdateOrganizationInvitationNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update organization invitation not found response has a 5xx status code
func (o *UpdateOrganizationInvitationNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update organization invitation not found response a status code equal to that given
func (o *UpdateOrganizationInvitationNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update organization invitation not found response
func (o *UpdateOrganizationInvitationNotFound) Code() int {
	return 404
}

func (o *UpdateOrganizationInvitationNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/backend/v1/invitations/organization/{invitationId}][%d] updateOrganizationInvitationNotFound ", 404)
}

func (o *UpdateOrganizationInvitationNotFound) String() string {
	return fmt.Sprintf("[PUT /api/backend/v1/invitations/organization/{invitationId}][%d] updateOrganizationInvitationNotFound ", 404)
}

func (o *UpdateOrganizationInvitationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateOrganizationInvitationRequestTimeout creates a UpdateOrganizationInvitationRequestTimeout with default headers values
func NewUpdateOrganizationInvitationRequestTimeout() *UpdateOrganizationInvitationRequestTimeout {
	return &UpdateOrganizationInvitationRequestTimeout{}
}

/*
UpdateOrganizationInvitationRequestTimeout describes a response with status code 408, with default header values.

Request Timeout
*/
type UpdateOrganizationInvitationRequestTimeout struct {
}

// IsSuccess returns true when this update organization invitation request timeout response has a 2xx status code
func (o *UpdateOrganizationInvitationRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update organization invitation request timeout response has a 3xx status code
func (o *UpdateOrganizationInvitationRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update organization invitation request timeout response has a 4xx status code
func (o *UpdateOrganizationInvitationRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this update organization invitation request timeout response has a 5xx status code
func (o *UpdateOrganizationInvitationRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this update organization invitation request timeout response a status code equal to that given
func (o *UpdateOrganizationInvitationRequestTimeout) IsCode(code int) bool {
	return code == 408
}

// Code gets the status code for the update organization invitation request timeout response
func (o *UpdateOrganizationInvitationRequestTimeout) Code() int {
	return 408
}

func (o *UpdateOrganizationInvitationRequestTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/backend/v1/invitations/organization/{invitationId}][%d] updateOrganizationInvitationRequestTimeout ", 408)
}

func (o *UpdateOrganizationInvitationRequestTimeout) String() string {
	return fmt.Sprintf("[PUT /api/backend/v1/invitations/organization/{invitationId}][%d] updateOrganizationInvitationRequestTimeout ", 408)
}

func (o *UpdateOrganizationInvitationRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateOrganizationInvitationUnprocessableEntity creates a UpdateOrganizationInvitationUnprocessableEntity with default headers values
func NewUpdateOrganizationInvitationUnprocessableEntity() *UpdateOrganizationInvitationUnprocessableEntity {
	return &UpdateOrganizationInvitationUnprocessableEntity{}
}

/*
UpdateOrganizationInvitationUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type UpdateOrganizationInvitationUnprocessableEntity struct {
}

// IsSuccess returns true when this update organization invitation unprocessable entity response has a 2xx status code
func (o *UpdateOrganizationInvitationUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update organization invitation unprocessable entity response has a 3xx status code
func (o *UpdateOrganizationInvitationUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update organization invitation unprocessable entity response has a 4xx status code
func (o *UpdateOrganizationInvitationUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this update organization invitation unprocessable entity response has a 5xx status code
func (o *UpdateOrganizationInvitationUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this update organization invitation unprocessable entity response a status code equal to that given
func (o *UpdateOrganizationInvitationUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the update organization invitation unprocessable entity response
func (o *UpdateOrganizationInvitationUnprocessableEntity) Code() int {
	return 422
}

func (o *UpdateOrganizationInvitationUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /api/backend/v1/invitations/organization/{invitationId}][%d] updateOrganizationInvitationUnprocessableEntity ", 422)
}

func (o *UpdateOrganizationInvitationUnprocessableEntity) String() string {
	return fmt.Sprintf("[PUT /api/backend/v1/invitations/organization/{invitationId}][%d] updateOrganizationInvitationUnprocessableEntity ", 422)
}

func (o *UpdateOrganizationInvitationUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
