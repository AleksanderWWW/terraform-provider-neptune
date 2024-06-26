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

// DeleteProjectReader is a Reader for the DeleteProject structure.
type DeleteProjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteProjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteProjectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteProjectBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteProjectUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteProjectForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteProjectNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewDeleteProjectRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewDeleteProjectUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /api/backend/v1/projects] deleteProject", response, response.Code())
	}
}

// NewDeleteProjectOK creates a DeleteProjectOK with default headers values
func NewDeleteProjectOK() *DeleteProjectOK {
	return &DeleteProjectOK{}
}

/*
DeleteProjectOK describes a response with status code 200, with default header values.

No response
*/
type DeleteProjectOK struct {
}

// IsSuccess returns true when this delete project o k response has a 2xx status code
func (o *DeleteProjectOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete project o k response has a 3xx status code
func (o *DeleteProjectOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete project o k response has a 4xx status code
func (o *DeleteProjectOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete project o k response has a 5xx status code
func (o *DeleteProjectOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete project o k response a status code equal to that given
func (o *DeleteProjectOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete project o k response
func (o *DeleteProjectOK) Code() int {
	return 200
}

func (o *DeleteProjectOK) Error() string {
	return fmt.Sprintf("[DELETE /api/backend/v1/projects][%d] deleteProjectOK ", 200)
}

func (o *DeleteProjectOK) String() string {
	return fmt.Sprintf("[DELETE /api/backend/v1/projects][%d] deleteProjectOK ", 200)
}

func (o *DeleteProjectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteProjectBadRequest creates a DeleteProjectBadRequest with default headers values
func NewDeleteProjectBadRequest() *DeleteProjectBadRequest {
	return &DeleteProjectBadRequest{}
}

/*
DeleteProjectBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DeleteProjectBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this delete project bad request response has a 2xx status code
func (o *DeleteProjectBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete project bad request response has a 3xx status code
func (o *DeleteProjectBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete project bad request response has a 4xx status code
func (o *DeleteProjectBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete project bad request response has a 5xx status code
func (o *DeleteProjectBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete project bad request response a status code equal to that given
func (o *DeleteProjectBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the delete project bad request response
func (o *DeleteProjectBadRequest) Code() int {
	return 400
}

func (o *DeleteProjectBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/backend/v1/projects][%d] deleteProjectBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteProjectBadRequest) String() string {
	return fmt.Sprintf("[DELETE /api/backend/v1/projects][%d] deleteProjectBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteProjectBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteProjectBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteProjectUnauthorized creates a DeleteProjectUnauthorized with default headers values
func NewDeleteProjectUnauthorized() *DeleteProjectUnauthorized {
	return &DeleteProjectUnauthorized{}
}

/*
DeleteProjectUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DeleteProjectUnauthorized struct {
}

// IsSuccess returns true when this delete project unauthorized response has a 2xx status code
func (o *DeleteProjectUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete project unauthorized response has a 3xx status code
func (o *DeleteProjectUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete project unauthorized response has a 4xx status code
func (o *DeleteProjectUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete project unauthorized response has a 5xx status code
func (o *DeleteProjectUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete project unauthorized response a status code equal to that given
func (o *DeleteProjectUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the delete project unauthorized response
func (o *DeleteProjectUnauthorized) Code() int {
	return 401
}

func (o *DeleteProjectUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/backend/v1/projects][%d] deleteProjectUnauthorized ", 401)
}

func (o *DeleteProjectUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /api/backend/v1/projects][%d] deleteProjectUnauthorized ", 401)
}

func (o *DeleteProjectUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteProjectForbidden creates a DeleteProjectForbidden with default headers values
func NewDeleteProjectForbidden() *DeleteProjectForbidden {
	return &DeleteProjectForbidden{}
}

/*
DeleteProjectForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteProjectForbidden struct {
}

// IsSuccess returns true when this delete project forbidden response has a 2xx status code
func (o *DeleteProjectForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete project forbidden response has a 3xx status code
func (o *DeleteProjectForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete project forbidden response has a 4xx status code
func (o *DeleteProjectForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete project forbidden response has a 5xx status code
func (o *DeleteProjectForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete project forbidden response a status code equal to that given
func (o *DeleteProjectForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete project forbidden response
func (o *DeleteProjectForbidden) Code() int {
	return 403
}

func (o *DeleteProjectForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/backend/v1/projects][%d] deleteProjectForbidden ", 403)
}

func (o *DeleteProjectForbidden) String() string {
	return fmt.Sprintf("[DELETE /api/backend/v1/projects][%d] deleteProjectForbidden ", 403)
}

func (o *DeleteProjectForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteProjectNotFound creates a DeleteProjectNotFound with default headers values
func NewDeleteProjectNotFound() *DeleteProjectNotFound {
	return &DeleteProjectNotFound{}
}

/*
DeleteProjectNotFound describes a response with status code 404, with default header values.

Not Found
*/
type DeleteProjectNotFound struct {
}

// IsSuccess returns true when this delete project not found response has a 2xx status code
func (o *DeleteProjectNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete project not found response has a 3xx status code
func (o *DeleteProjectNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete project not found response has a 4xx status code
func (o *DeleteProjectNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete project not found response has a 5xx status code
func (o *DeleteProjectNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete project not found response a status code equal to that given
func (o *DeleteProjectNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete project not found response
func (o *DeleteProjectNotFound) Code() int {
	return 404
}

func (o *DeleteProjectNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/backend/v1/projects][%d] deleteProjectNotFound ", 404)
}

func (o *DeleteProjectNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/backend/v1/projects][%d] deleteProjectNotFound ", 404)
}

func (o *DeleteProjectNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteProjectRequestTimeout creates a DeleteProjectRequestTimeout with default headers values
func NewDeleteProjectRequestTimeout() *DeleteProjectRequestTimeout {
	return &DeleteProjectRequestTimeout{}
}

/*
DeleteProjectRequestTimeout describes a response with status code 408, with default header values.

Request Timeout
*/
type DeleteProjectRequestTimeout struct {
}

// IsSuccess returns true when this delete project request timeout response has a 2xx status code
func (o *DeleteProjectRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete project request timeout response has a 3xx status code
func (o *DeleteProjectRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete project request timeout response has a 4xx status code
func (o *DeleteProjectRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete project request timeout response has a 5xx status code
func (o *DeleteProjectRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this delete project request timeout response a status code equal to that given
func (o *DeleteProjectRequestTimeout) IsCode(code int) bool {
	return code == 408
}

// Code gets the status code for the delete project request timeout response
func (o *DeleteProjectRequestTimeout) Code() int {
	return 408
}

func (o *DeleteProjectRequestTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/backend/v1/projects][%d] deleteProjectRequestTimeout ", 408)
}

func (o *DeleteProjectRequestTimeout) String() string {
	return fmt.Sprintf("[DELETE /api/backend/v1/projects][%d] deleteProjectRequestTimeout ", 408)
}

func (o *DeleteProjectRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteProjectUnprocessableEntity creates a DeleteProjectUnprocessableEntity with default headers values
func NewDeleteProjectUnprocessableEntity() *DeleteProjectUnprocessableEntity {
	return &DeleteProjectUnprocessableEntity{}
}

/*
DeleteProjectUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type DeleteProjectUnprocessableEntity struct {
}

// IsSuccess returns true when this delete project unprocessable entity response has a 2xx status code
func (o *DeleteProjectUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete project unprocessable entity response has a 3xx status code
func (o *DeleteProjectUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete project unprocessable entity response has a 4xx status code
func (o *DeleteProjectUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete project unprocessable entity response has a 5xx status code
func (o *DeleteProjectUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this delete project unprocessable entity response a status code equal to that given
func (o *DeleteProjectUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the delete project unprocessable entity response
func (o *DeleteProjectUnprocessableEntity) Code() int {
	return 422
}

func (o *DeleteProjectUnprocessableEntity) Error() string {
	return fmt.Sprintf("[DELETE /api/backend/v1/projects][%d] deleteProjectUnprocessableEntity ", 422)
}

func (o *DeleteProjectUnprocessableEntity) String() string {
	return fmt.Sprintf("[DELETE /api/backend/v1/projects][%d] deleteProjectUnprocessableEntity ", 422)
}

func (o *DeleteProjectUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
