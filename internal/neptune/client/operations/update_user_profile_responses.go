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

// UpdateUserProfileReader is a Reader for the UpdateUserProfile structure.
type UpdateUserProfileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateUserProfileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateUserProfileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateUserProfileBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateUserProfileUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateUserProfileForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateUserProfileNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewUpdateUserProfileRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewUpdateUserProfileUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PATCH /api/backend/v1/userProfile] updateUserProfile", response, response.Code())
	}
}

// NewUpdateUserProfileOK creates a UpdateUserProfileOK with default headers values
func NewUpdateUserProfileOK() *UpdateUserProfileOK {
	return &UpdateUserProfileOK{}
}

/*
UpdateUserProfileOK describes a response with status code 200, with default header values.

OK
*/
type UpdateUserProfileOK struct {
	Payload *models.UserProfileDTO
}

// IsSuccess returns true when this update user profile o k response has a 2xx status code
func (o *UpdateUserProfileOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update user profile o k response has a 3xx status code
func (o *UpdateUserProfileOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update user profile o k response has a 4xx status code
func (o *UpdateUserProfileOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update user profile o k response has a 5xx status code
func (o *UpdateUserProfileOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update user profile o k response a status code equal to that given
func (o *UpdateUserProfileOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update user profile o k response
func (o *UpdateUserProfileOK) Code() int {
	return 200
}

func (o *UpdateUserProfileOK) Error() string {
	return fmt.Sprintf("[PATCH /api/backend/v1/userProfile][%d] updateUserProfileOK  %+v", 200, o.Payload)
}

func (o *UpdateUserProfileOK) String() string {
	return fmt.Sprintf("[PATCH /api/backend/v1/userProfile][%d] updateUserProfileOK  %+v", 200, o.Payload)
}

func (o *UpdateUserProfileOK) GetPayload() *models.UserProfileDTO {
	return o.Payload
}

func (o *UpdateUserProfileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserProfileDTO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateUserProfileBadRequest creates a UpdateUserProfileBadRequest with default headers values
func NewUpdateUserProfileBadRequest() *UpdateUserProfileBadRequest {
	return &UpdateUserProfileBadRequest{}
}

/*
UpdateUserProfileBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UpdateUserProfileBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this update user profile bad request response has a 2xx status code
func (o *UpdateUserProfileBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update user profile bad request response has a 3xx status code
func (o *UpdateUserProfileBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update user profile bad request response has a 4xx status code
func (o *UpdateUserProfileBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update user profile bad request response has a 5xx status code
func (o *UpdateUserProfileBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update user profile bad request response a status code equal to that given
func (o *UpdateUserProfileBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update user profile bad request response
func (o *UpdateUserProfileBadRequest) Code() int {
	return 400
}

func (o *UpdateUserProfileBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/backend/v1/userProfile][%d] updateUserProfileBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateUserProfileBadRequest) String() string {
	return fmt.Sprintf("[PATCH /api/backend/v1/userProfile][%d] updateUserProfileBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateUserProfileBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateUserProfileBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateUserProfileUnauthorized creates a UpdateUserProfileUnauthorized with default headers values
func NewUpdateUserProfileUnauthorized() *UpdateUserProfileUnauthorized {
	return &UpdateUserProfileUnauthorized{}
}

/*
UpdateUserProfileUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type UpdateUserProfileUnauthorized struct {
}

// IsSuccess returns true when this update user profile unauthorized response has a 2xx status code
func (o *UpdateUserProfileUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update user profile unauthorized response has a 3xx status code
func (o *UpdateUserProfileUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update user profile unauthorized response has a 4xx status code
func (o *UpdateUserProfileUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update user profile unauthorized response has a 5xx status code
func (o *UpdateUserProfileUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update user profile unauthorized response a status code equal to that given
func (o *UpdateUserProfileUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the update user profile unauthorized response
func (o *UpdateUserProfileUnauthorized) Code() int {
	return 401
}

func (o *UpdateUserProfileUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/backend/v1/userProfile][%d] updateUserProfileUnauthorized ", 401)
}

func (o *UpdateUserProfileUnauthorized) String() string {
	return fmt.Sprintf("[PATCH /api/backend/v1/userProfile][%d] updateUserProfileUnauthorized ", 401)
}

func (o *UpdateUserProfileUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateUserProfileForbidden creates a UpdateUserProfileForbidden with default headers values
func NewUpdateUserProfileForbidden() *UpdateUserProfileForbidden {
	return &UpdateUserProfileForbidden{}
}

/*
UpdateUserProfileForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UpdateUserProfileForbidden struct {
}

// IsSuccess returns true when this update user profile forbidden response has a 2xx status code
func (o *UpdateUserProfileForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update user profile forbidden response has a 3xx status code
func (o *UpdateUserProfileForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update user profile forbidden response has a 4xx status code
func (o *UpdateUserProfileForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update user profile forbidden response has a 5xx status code
func (o *UpdateUserProfileForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update user profile forbidden response a status code equal to that given
func (o *UpdateUserProfileForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the update user profile forbidden response
func (o *UpdateUserProfileForbidden) Code() int {
	return 403
}

func (o *UpdateUserProfileForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/backend/v1/userProfile][%d] updateUserProfileForbidden ", 403)
}

func (o *UpdateUserProfileForbidden) String() string {
	return fmt.Sprintf("[PATCH /api/backend/v1/userProfile][%d] updateUserProfileForbidden ", 403)
}

func (o *UpdateUserProfileForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateUserProfileNotFound creates a UpdateUserProfileNotFound with default headers values
func NewUpdateUserProfileNotFound() *UpdateUserProfileNotFound {
	return &UpdateUserProfileNotFound{}
}

/*
UpdateUserProfileNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UpdateUserProfileNotFound struct {
}

// IsSuccess returns true when this update user profile not found response has a 2xx status code
func (o *UpdateUserProfileNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update user profile not found response has a 3xx status code
func (o *UpdateUserProfileNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update user profile not found response has a 4xx status code
func (o *UpdateUserProfileNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update user profile not found response has a 5xx status code
func (o *UpdateUserProfileNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update user profile not found response a status code equal to that given
func (o *UpdateUserProfileNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update user profile not found response
func (o *UpdateUserProfileNotFound) Code() int {
	return 404
}

func (o *UpdateUserProfileNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/backend/v1/userProfile][%d] updateUserProfileNotFound ", 404)
}

func (o *UpdateUserProfileNotFound) String() string {
	return fmt.Sprintf("[PATCH /api/backend/v1/userProfile][%d] updateUserProfileNotFound ", 404)
}

func (o *UpdateUserProfileNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateUserProfileRequestTimeout creates a UpdateUserProfileRequestTimeout with default headers values
func NewUpdateUserProfileRequestTimeout() *UpdateUserProfileRequestTimeout {
	return &UpdateUserProfileRequestTimeout{}
}

/*
UpdateUserProfileRequestTimeout describes a response with status code 408, with default header values.

Request Timeout
*/
type UpdateUserProfileRequestTimeout struct {
}

// IsSuccess returns true when this update user profile request timeout response has a 2xx status code
func (o *UpdateUserProfileRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update user profile request timeout response has a 3xx status code
func (o *UpdateUserProfileRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update user profile request timeout response has a 4xx status code
func (o *UpdateUserProfileRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this update user profile request timeout response has a 5xx status code
func (o *UpdateUserProfileRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this update user profile request timeout response a status code equal to that given
func (o *UpdateUserProfileRequestTimeout) IsCode(code int) bool {
	return code == 408
}

// Code gets the status code for the update user profile request timeout response
func (o *UpdateUserProfileRequestTimeout) Code() int {
	return 408
}

func (o *UpdateUserProfileRequestTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/backend/v1/userProfile][%d] updateUserProfileRequestTimeout ", 408)
}

func (o *UpdateUserProfileRequestTimeout) String() string {
	return fmt.Sprintf("[PATCH /api/backend/v1/userProfile][%d] updateUserProfileRequestTimeout ", 408)
}

func (o *UpdateUserProfileRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateUserProfileUnprocessableEntity creates a UpdateUserProfileUnprocessableEntity with default headers values
func NewUpdateUserProfileUnprocessableEntity() *UpdateUserProfileUnprocessableEntity {
	return &UpdateUserProfileUnprocessableEntity{}
}

/*
UpdateUserProfileUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type UpdateUserProfileUnprocessableEntity struct {
}

// IsSuccess returns true when this update user profile unprocessable entity response has a 2xx status code
func (o *UpdateUserProfileUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update user profile unprocessable entity response has a 3xx status code
func (o *UpdateUserProfileUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update user profile unprocessable entity response has a 4xx status code
func (o *UpdateUserProfileUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this update user profile unprocessable entity response has a 5xx status code
func (o *UpdateUserProfileUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this update user profile unprocessable entity response a status code equal to that given
func (o *UpdateUserProfileUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the update user profile unprocessable entity response
func (o *UpdateUserProfileUnprocessableEntity) Code() int {
	return 422
}

func (o *UpdateUserProfileUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PATCH /api/backend/v1/userProfile][%d] updateUserProfileUnprocessableEntity ", 422)
}

func (o *UpdateUserProfileUnprocessableEntity) String() string {
	return fmt.Sprintf("[PATCH /api/backend/v1/userProfile][%d] updateUserProfileUnprocessableEntity ", 422)
}

func (o *UpdateUserProfileUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}