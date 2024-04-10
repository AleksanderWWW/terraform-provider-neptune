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

// ResendOrganizationInvitationReader is a Reader for the ResendOrganizationInvitation structure.
type ResendOrganizationInvitationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ResendOrganizationInvitationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewResendOrganizationInvitationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewResendOrganizationInvitationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewResendOrganizationInvitationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewResendOrganizationInvitationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewResendOrganizationInvitationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewResendOrganizationInvitationRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewResendOrganizationInvitationUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /api/backend/v1/invitations/organization/resend] resendOrganizationInvitation", response, response.Code())
	}
}

// NewResendOrganizationInvitationOK creates a ResendOrganizationInvitationOK with default headers values
func NewResendOrganizationInvitationOK() *ResendOrganizationInvitationOK {
	return &ResendOrganizationInvitationOK{}
}

/*
ResendOrganizationInvitationOK describes a response with status code 200, with default header values.

No response
*/
type ResendOrganizationInvitationOK struct {
}

// IsSuccess returns true when this resend organization invitation o k response has a 2xx status code
func (o *ResendOrganizationInvitationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this resend organization invitation o k response has a 3xx status code
func (o *ResendOrganizationInvitationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this resend organization invitation o k response has a 4xx status code
func (o *ResendOrganizationInvitationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this resend organization invitation o k response has a 5xx status code
func (o *ResendOrganizationInvitationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this resend organization invitation o k response a status code equal to that given
func (o *ResendOrganizationInvitationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the resend organization invitation o k response
func (o *ResendOrganizationInvitationOK) Code() int {
	return 200
}

func (o *ResendOrganizationInvitationOK) Error() string {
	return fmt.Sprintf("[POST /api/backend/v1/invitations/organization/resend][%d] resendOrganizationInvitationOK ", 200)
}

func (o *ResendOrganizationInvitationOK) String() string {
	return fmt.Sprintf("[POST /api/backend/v1/invitations/organization/resend][%d] resendOrganizationInvitationOK ", 200)
}

func (o *ResendOrganizationInvitationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewResendOrganizationInvitationBadRequest creates a ResendOrganizationInvitationBadRequest with default headers values
func NewResendOrganizationInvitationBadRequest() *ResendOrganizationInvitationBadRequest {
	return &ResendOrganizationInvitationBadRequest{}
}

/*
ResendOrganizationInvitationBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ResendOrganizationInvitationBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this resend organization invitation bad request response has a 2xx status code
func (o *ResendOrganizationInvitationBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this resend organization invitation bad request response has a 3xx status code
func (o *ResendOrganizationInvitationBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this resend organization invitation bad request response has a 4xx status code
func (o *ResendOrganizationInvitationBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this resend organization invitation bad request response has a 5xx status code
func (o *ResendOrganizationInvitationBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this resend organization invitation bad request response a status code equal to that given
func (o *ResendOrganizationInvitationBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the resend organization invitation bad request response
func (o *ResendOrganizationInvitationBadRequest) Code() int {
	return 400
}

func (o *ResendOrganizationInvitationBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/backend/v1/invitations/organization/resend][%d] resendOrganizationInvitationBadRequest  %+v", 400, o.Payload)
}

func (o *ResendOrganizationInvitationBadRequest) String() string {
	return fmt.Sprintf("[POST /api/backend/v1/invitations/organization/resend][%d] resendOrganizationInvitationBadRequest  %+v", 400, o.Payload)
}

func (o *ResendOrganizationInvitationBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ResendOrganizationInvitationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResendOrganizationInvitationUnauthorized creates a ResendOrganizationInvitationUnauthorized with default headers values
func NewResendOrganizationInvitationUnauthorized() *ResendOrganizationInvitationUnauthorized {
	return &ResendOrganizationInvitationUnauthorized{}
}

/*
ResendOrganizationInvitationUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ResendOrganizationInvitationUnauthorized struct {
}

// IsSuccess returns true when this resend organization invitation unauthorized response has a 2xx status code
func (o *ResendOrganizationInvitationUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this resend organization invitation unauthorized response has a 3xx status code
func (o *ResendOrganizationInvitationUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this resend organization invitation unauthorized response has a 4xx status code
func (o *ResendOrganizationInvitationUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this resend organization invitation unauthorized response has a 5xx status code
func (o *ResendOrganizationInvitationUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this resend organization invitation unauthorized response a status code equal to that given
func (o *ResendOrganizationInvitationUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the resend organization invitation unauthorized response
func (o *ResendOrganizationInvitationUnauthorized) Code() int {
	return 401
}

func (o *ResendOrganizationInvitationUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/backend/v1/invitations/organization/resend][%d] resendOrganizationInvitationUnauthorized ", 401)
}

func (o *ResendOrganizationInvitationUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/backend/v1/invitations/organization/resend][%d] resendOrganizationInvitationUnauthorized ", 401)
}

func (o *ResendOrganizationInvitationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewResendOrganizationInvitationForbidden creates a ResendOrganizationInvitationForbidden with default headers values
func NewResendOrganizationInvitationForbidden() *ResendOrganizationInvitationForbidden {
	return &ResendOrganizationInvitationForbidden{}
}

/*
ResendOrganizationInvitationForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ResendOrganizationInvitationForbidden struct {
}

// IsSuccess returns true when this resend organization invitation forbidden response has a 2xx status code
func (o *ResendOrganizationInvitationForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this resend organization invitation forbidden response has a 3xx status code
func (o *ResendOrganizationInvitationForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this resend organization invitation forbidden response has a 4xx status code
func (o *ResendOrganizationInvitationForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this resend organization invitation forbidden response has a 5xx status code
func (o *ResendOrganizationInvitationForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this resend organization invitation forbidden response a status code equal to that given
func (o *ResendOrganizationInvitationForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the resend organization invitation forbidden response
func (o *ResendOrganizationInvitationForbidden) Code() int {
	return 403
}

func (o *ResendOrganizationInvitationForbidden) Error() string {
	return fmt.Sprintf("[POST /api/backend/v1/invitations/organization/resend][%d] resendOrganizationInvitationForbidden ", 403)
}

func (o *ResendOrganizationInvitationForbidden) String() string {
	return fmt.Sprintf("[POST /api/backend/v1/invitations/organization/resend][%d] resendOrganizationInvitationForbidden ", 403)
}

func (o *ResendOrganizationInvitationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewResendOrganizationInvitationNotFound creates a ResendOrganizationInvitationNotFound with default headers values
func NewResendOrganizationInvitationNotFound() *ResendOrganizationInvitationNotFound {
	return &ResendOrganizationInvitationNotFound{}
}

/*
ResendOrganizationInvitationNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ResendOrganizationInvitationNotFound struct {
}

// IsSuccess returns true when this resend organization invitation not found response has a 2xx status code
func (o *ResendOrganizationInvitationNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this resend organization invitation not found response has a 3xx status code
func (o *ResendOrganizationInvitationNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this resend organization invitation not found response has a 4xx status code
func (o *ResendOrganizationInvitationNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this resend organization invitation not found response has a 5xx status code
func (o *ResendOrganizationInvitationNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this resend organization invitation not found response a status code equal to that given
func (o *ResendOrganizationInvitationNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the resend organization invitation not found response
func (o *ResendOrganizationInvitationNotFound) Code() int {
	return 404
}

func (o *ResendOrganizationInvitationNotFound) Error() string {
	return fmt.Sprintf("[POST /api/backend/v1/invitations/organization/resend][%d] resendOrganizationInvitationNotFound ", 404)
}

func (o *ResendOrganizationInvitationNotFound) String() string {
	return fmt.Sprintf("[POST /api/backend/v1/invitations/organization/resend][%d] resendOrganizationInvitationNotFound ", 404)
}

func (o *ResendOrganizationInvitationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewResendOrganizationInvitationRequestTimeout creates a ResendOrganizationInvitationRequestTimeout with default headers values
func NewResendOrganizationInvitationRequestTimeout() *ResendOrganizationInvitationRequestTimeout {
	return &ResendOrganizationInvitationRequestTimeout{}
}

/*
ResendOrganizationInvitationRequestTimeout describes a response with status code 408, with default header values.

Request Timeout
*/
type ResendOrganizationInvitationRequestTimeout struct {
}

// IsSuccess returns true when this resend organization invitation request timeout response has a 2xx status code
func (o *ResendOrganizationInvitationRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this resend organization invitation request timeout response has a 3xx status code
func (o *ResendOrganizationInvitationRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this resend organization invitation request timeout response has a 4xx status code
func (o *ResendOrganizationInvitationRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this resend organization invitation request timeout response has a 5xx status code
func (o *ResendOrganizationInvitationRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this resend organization invitation request timeout response a status code equal to that given
func (o *ResendOrganizationInvitationRequestTimeout) IsCode(code int) bool {
	return code == 408
}

// Code gets the status code for the resend organization invitation request timeout response
func (o *ResendOrganizationInvitationRequestTimeout) Code() int {
	return 408
}

func (o *ResendOrganizationInvitationRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/backend/v1/invitations/organization/resend][%d] resendOrganizationInvitationRequestTimeout ", 408)
}

func (o *ResendOrganizationInvitationRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/backend/v1/invitations/organization/resend][%d] resendOrganizationInvitationRequestTimeout ", 408)
}

func (o *ResendOrganizationInvitationRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewResendOrganizationInvitationUnprocessableEntity creates a ResendOrganizationInvitationUnprocessableEntity with default headers values
func NewResendOrganizationInvitationUnprocessableEntity() *ResendOrganizationInvitationUnprocessableEntity {
	return &ResendOrganizationInvitationUnprocessableEntity{}
}

/*
ResendOrganizationInvitationUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type ResendOrganizationInvitationUnprocessableEntity struct {
}

// IsSuccess returns true when this resend organization invitation unprocessable entity response has a 2xx status code
func (o *ResendOrganizationInvitationUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this resend organization invitation unprocessable entity response has a 3xx status code
func (o *ResendOrganizationInvitationUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this resend organization invitation unprocessable entity response has a 4xx status code
func (o *ResendOrganizationInvitationUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this resend organization invitation unprocessable entity response has a 5xx status code
func (o *ResendOrganizationInvitationUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this resend organization invitation unprocessable entity response a status code equal to that given
func (o *ResendOrganizationInvitationUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the resend organization invitation unprocessable entity response
func (o *ResendOrganizationInvitationUnprocessableEntity) Code() int {
	return 422
}

func (o *ResendOrganizationInvitationUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /api/backend/v1/invitations/organization/resend][%d] resendOrganizationInvitationUnprocessableEntity ", 422)
}

func (o *ResendOrganizationInvitationUnprocessableEntity) String() string {
	return fmt.Sprintf("[POST /api/backend/v1/invitations/organization/resend][%d] resendOrganizationInvitationUnprocessableEntity ", 422)
}

func (o *ResendOrganizationInvitationUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
