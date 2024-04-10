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

// DeleteStorageLimitReader is a Reader for the DeleteStorageLimit structure.
type DeleteStorageLimitReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteStorageLimitReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteStorageLimitOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteStorageLimitBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteStorageLimitUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteStorageLimitForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteStorageLimitNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewDeleteStorageLimitRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewDeleteStorageLimitUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /api/backend/v1/projects/storage-limit] deleteStorageLimit", response, response.Code())
	}
}

// NewDeleteStorageLimitOK creates a DeleteStorageLimitOK with default headers values
func NewDeleteStorageLimitOK() *DeleteStorageLimitOK {
	return &DeleteStorageLimitOK{}
}

/*
DeleteStorageLimitOK describes a response with status code 200, with default header values.

No response
*/
type DeleteStorageLimitOK struct {
}

// IsSuccess returns true when this delete storage limit o k response has a 2xx status code
func (o *DeleteStorageLimitOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete storage limit o k response has a 3xx status code
func (o *DeleteStorageLimitOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete storage limit o k response has a 4xx status code
func (o *DeleteStorageLimitOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete storage limit o k response has a 5xx status code
func (o *DeleteStorageLimitOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete storage limit o k response a status code equal to that given
func (o *DeleteStorageLimitOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete storage limit o k response
func (o *DeleteStorageLimitOK) Code() int {
	return 200
}

func (o *DeleteStorageLimitOK) Error() string {
	return fmt.Sprintf("[DELETE /api/backend/v1/projects/storage-limit][%d] deleteStorageLimitOK ", 200)
}

func (o *DeleteStorageLimitOK) String() string {
	return fmt.Sprintf("[DELETE /api/backend/v1/projects/storage-limit][%d] deleteStorageLimitOK ", 200)
}

func (o *DeleteStorageLimitOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteStorageLimitBadRequest creates a DeleteStorageLimitBadRequest with default headers values
func NewDeleteStorageLimitBadRequest() *DeleteStorageLimitBadRequest {
	return &DeleteStorageLimitBadRequest{}
}

/*
DeleteStorageLimitBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DeleteStorageLimitBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this delete storage limit bad request response has a 2xx status code
func (o *DeleteStorageLimitBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete storage limit bad request response has a 3xx status code
func (o *DeleteStorageLimitBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete storage limit bad request response has a 4xx status code
func (o *DeleteStorageLimitBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete storage limit bad request response has a 5xx status code
func (o *DeleteStorageLimitBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete storage limit bad request response a status code equal to that given
func (o *DeleteStorageLimitBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the delete storage limit bad request response
func (o *DeleteStorageLimitBadRequest) Code() int {
	return 400
}

func (o *DeleteStorageLimitBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/backend/v1/projects/storage-limit][%d] deleteStorageLimitBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteStorageLimitBadRequest) String() string {
	return fmt.Sprintf("[DELETE /api/backend/v1/projects/storage-limit][%d] deleteStorageLimitBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteStorageLimitBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteStorageLimitBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteStorageLimitUnauthorized creates a DeleteStorageLimitUnauthorized with default headers values
func NewDeleteStorageLimitUnauthorized() *DeleteStorageLimitUnauthorized {
	return &DeleteStorageLimitUnauthorized{}
}

/*
DeleteStorageLimitUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DeleteStorageLimitUnauthorized struct {
}

// IsSuccess returns true when this delete storage limit unauthorized response has a 2xx status code
func (o *DeleteStorageLimitUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete storage limit unauthorized response has a 3xx status code
func (o *DeleteStorageLimitUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete storage limit unauthorized response has a 4xx status code
func (o *DeleteStorageLimitUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete storage limit unauthorized response has a 5xx status code
func (o *DeleteStorageLimitUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete storage limit unauthorized response a status code equal to that given
func (o *DeleteStorageLimitUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the delete storage limit unauthorized response
func (o *DeleteStorageLimitUnauthorized) Code() int {
	return 401
}

func (o *DeleteStorageLimitUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/backend/v1/projects/storage-limit][%d] deleteStorageLimitUnauthorized ", 401)
}

func (o *DeleteStorageLimitUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /api/backend/v1/projects/storage-limit][%d] deleteStorageLimitUnauthorized ", 401)
}

func (o *DeleteStorageLimitUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteStorageLimitForbidden creates a DeleteStorageLimitForbidden with default headers values
func NewDeleteStorageLimitForbidden() *DeleteStorageLimitForbidden {
	return &DeleteStorageLimitForbidden{}
}

/*
DeleteStorageLimitForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteStorageLimitForbidden struct {
}

// IsSuccess returns true when this delete storage limit forbidden response has a 2xx status code
func (o *DeleteStorageLimitForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete storage limit forbidden response has a 3xx status code
func (o *DeleteStorageLimitForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete storage limit forbidden response has a 4xx status code
func (o *DeleteStorageLimitForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete storage limit forbidden response has a 5xx status code
func (o *DeleteStorageLimitForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete storage limit forbidden response a status code equal to that given
func (o *DeleteStorageLimitForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete storage limit forbidden response
func (o *DeleteStorageLimitForbidden) Code() int {
	return 403
}

func (o *DeleteStorageLimitForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/backend/v1/projects/storage-limit][%d] deleteStorageLimitForbidden ", 403)
}

func (o *DeleteStorageLimitForbidden) String() string {
	return fmt.Sprintf("[DELETE /api/backend/v1/projects/storage-limit][%d] deleteStorageLimitForbidden ", 403)
}

func (o *DeleteStorageLimitForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteStorageLimitNotFound creates a DeleteStorageLimitNotFound with default headers values
func NewDeleteStorageLimitNotFound() *DeleteStorageLimitNotFound {
	return &DeleteStorageLimitNotFound{}
}

/*
DeleteStorageLimitNotFound describes a response with status code 404, with default header values.

Not Found
*/
type DeleteStorageLimitNotFound struct {
}

// IsSuccess returns true when this delete storage limit not found response has a 2xx status code
func (o *DeleteStorageLimitNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete storage limit not found response has a 3xx status code
func (o *DeleteStorageLimitNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete storage limit not found response has a 4xx status code
func (o *DeleteStorageLimitNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete storage limit not found response has a 5xx status code
func (o *DeleteStorageLimitNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete storage limit not found response a status code equal to that given
func (o *DeleteStorageLimitNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete storage limit not found response
func (o *DeleteStorageLimitNotFound) Code() int {
	return 404
}

func (o *DeleteStorageLimitNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/backend/v1/projects/storage-limit][%d] deleteStorageLimitNotFound ", 404)
}

func (o *DeleteStorageLimitNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/backend/v1/projects/storage-limit][%d] deleteStorageLimitNotFound ", 404)
}

func (o *DeleteStorageLimitNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteStorageLimitRequestTimeout creates a DeleteStorageLimitRequestTimeout with default headers values
func NewDeleteStorageLimitRequestTimeout() *DeleteStorageLimitRequestTimeout {
	return &DeleteStorageLimitRequestTimeout{}
}

/*
DeleteStorageLimitRequestTimeout describes a response with status code 408, with default header values.

Request Timeout
*/
type DeleteStorageLimitRequestTimeout struct {
}

// IsSuccess returns true when this delete storage limit request timeout response has a 2xx status code
func (o *DeleteStorageLimitRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete storage limit request timeout response has a 3xx status code
func (o *DeleteStorageLimitRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete storage limit request timeout response has a 4xx status code
func (o *DeleteStorageLimitRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete storage limit request timeout response has a 5xx status code
func (o *DeleteStorageLimitRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this delete storage limit request timeout response a status code equal to that given
func (o *DeleteStorageLimitRequestTimeout) IsCode(code int) bool {
	return code == 408
}

// Code gets the status code for the delete storage limit request timeout response
func (o *DeleteStorageLimitRequestTimeout) Code() int {
	return 408
}

func (o *DeleteStorageLimitRequestTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/backend/v1/projects/storage-limit][%d] deleteStorageLimitRequestTimeout ", 408)
}

func (o *DeleteStorageLimitRequestTimeout) String() string {
	return fmt.Sprintf("[DELETE /api/backend/v1/projects/storage-limit][%d] deleteStorageLimitRequestTimeout ", 408)
}

func (o *DeleteStorageLimitRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteStorageLimitUnprocessableEntity creates a DeleteStorageLimitUnprocessableEntity with default headers values
func NewDeleteStorageLimitUnprocessableEntity() *DeleteStorageLimitUnprocessableEntity {
	return &DeleteStorageLimitUnprocessableEntity{}
}

/*
DeleteStorageLimitUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type DeleteStorageLimitUnprocessableEntity struct {
}

// IsSuccess returns true when this delete storage limit unprocessable entity response has a 2xx status code
func (o *DeleteStorageLimitUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete storage limit unprocessable entity response has a 3xx status code
func (o *DeleteStorageLimitUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete storage limit unprocessable entity response has a 4xx status code
func (o *DeleteStorageLimitUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete storage limit unprocessable entity response has a 5xx status code
func (o *DeleteStorageLimitUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this delete storage limit unprocessable entity response a status code equal to that given
func (o *DeleteStorageLimitUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the delete storage limit unprocessable entity response
func (o *DeleteStorageLimitUnprocessableEntity) Code() int {
	return 422
}

func (o *DeleteStorageLimitUnprocessableEntity) Error() string {
	return fmt.Sprintf("[DELETE /api/backend/v1/projects/storage-limit][%d] deleteStorageLimitUnprocessableEntity ", 422)
}

func (o *DeleteStorageLimitUnprocessableEntity) String() string {
	return fmt.Sprintf("[DELETE /api/backend/v1/projects/storage-limit][%d] deleteStorageLimitUnprocessableEntity ", 422)
}

func (o *DeleteStorageLimitUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
