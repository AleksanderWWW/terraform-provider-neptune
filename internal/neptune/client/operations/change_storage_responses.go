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

// ChangeStorageReader is a Reader for the ChangeStorage structure.
type ChangeStorageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChangeStorageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewChangeStorageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewChangeStorageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewChangeStorageUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewChangeStorageForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewChangeStorageNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewChangeStorageRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewChangeStorageUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /api/backend/v1/payments/limits/changeStorage] changeStorage", response, response.Code())
	}
}

// NewChangeStorageOK creates a ChangeStorageOK with default headers values
func NewChangeStorageOK() *ChangeStorageOK {
	return &ChangeStorageOK{}
}

/*
ChangeStorageOK describes a response with status code 200, with default header values.

OK
*/
type ChangeStorageOK struct {
	Payload *models.PurchaseResultDTO
}

// IsSuccess returns true when this change storage o k response has a 2xx status code
func (o *ChangeStorageOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this change storage o k response has a 3xx status code
func (o *ChangeStorageOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this change storage o k response has a 4xx status code
func (o *ChangeStorageOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this change storage o k response has a 5xx status code
func (o *ChangeStorageOK) IsServerError() bool {
	return false
}

// IsCode returns true when this change storage o k response a status code equal to that given
func (o *ChangeStorageOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the change storage o k response
func (o *ChangeStorageOK) Code() int {
	return 200
}

func (o *ChangeStorageOK) Error() string {
	return fmt.Sprintf("[POST /api/backend/v1/payments/limits/changeStorage][%d] changeStorageOK  %+v", 200, o.Payload)
}

func (o *ChangeStorageOK) String() string {
	return fmt.Sprintf("[POST /api/backend/v1/payments/limits/changeStorage][%d] changeStorageOK  %+v", 200, o.Payload)
}

func (o *ChangeStorageOK) GetPayload() *models.PurchaseResultDTO {
	return o.Payload
}

func (o *ChangeStorageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PurchaseResultDTO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewChangeStorageBadRequest creates a ChangeStorageBadRequest with default headers values
func NewChangeStorageBadRequest() *ChangeStorageBadRequest {
	return &ChangeStorageBadRequest{}
}

/*
ChangeStorageBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ChangeStorageBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this change storage bad request response has a 2xx status code
func (o *ChangeStorageBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this change storage bad request response has a 3xx status code
func (o *ChangeStorageBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this change storage bad request response has a 4xx status code
func (o *ChangeStorageBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this change storage bad request response has a 5xx status code
func (o *ChangeStorageBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this change storage bad request response a status code equal to that given
func (o *ChangeStorageBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the change storage bad request response
func (o *ChangeStorageBadRequest) Code() int {
	return 400
}

func (o *ChangeStorageBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/backend/v1/payments/limits/changeStorage][%d] changeStorageBadRequest  %+v", 400, o.Payload)
}

func (o *ChangeStorageBadRequest) String() string {
	return fmt.Sprintf("[POST /api/backend/v1/payments/limits/changeStorage][%d] changeStorageBadRequest  %+v", 400, o.Payload)
}

func (o *ChangeStorageBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ChangeStorageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewChangeStorageUnauthorized creates a ChangeStorageUnauthorized with default headers values
func NewChangeStorageUnauthorized() *ChangeStorageUnauthorized {
	return &ChangeStorageUnauthorized{}
}

/*
ChangeStorageUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ChangeStorageUnauthorized struct {
}

// IsSuccess returns true when this change storage unauthorized response has a 2xx status code
func (o *ChangeStorageUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this change storage unauthorized response has a 3xx status code
func (o *ChangeStorageUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this change storage unauthorized response has a 4xx status code
func (o *ChangeStorageUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this change storage unauthorized response has a 5xx status code
func (o *ChangeStorageUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this change storage unauthorized response a status code equal to that given
func (o *ChangeStorageUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the change storage unauthorized response
func (o *ChangeStorageUnauthorized) Code() int {
	return 401
}

func (o *ChangeStorageUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/backend/v1/payments/limits/changeStorage][%d] changeStorageUnauthorized ", 401)
}

func (o *ChangeStorageUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/backend/v1/payments/limits/changeStorage][%d] changeStorageUnauthorized ", 401)
}

func (o *ChangeStorageUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewChangeStorageForbidden creates a ChangeStorageForbidden with default headers values
func NewChangeStorageForbidden() *ChangeStorageForbidden {
	return &ChangeStorageForbidden{}
}

/*
ChangeStorageForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ChangeStorageForbidden struct {
}

// IsSuccess returns true when this change storage forbidden response has a 2xx status code
func (o *ChangeStorageForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this change storage forbidden response has a 3xx status code
func (o *ChangeStorageForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this change storage forbidden response has a 4xx status code
func (o *ChangeStorageForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this change storage forbidden response has a 5xx status code
func (o *ChangeStorageForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this change storage forbidden response a status code equal to that given
func (o *ChangeStorageForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the change storage forbidden response
func (o *ChangeStorageForbidden) Code() int {
	return 403
}

func (o *ChangeStorageForbidden) Error() string {
	return fmt.Sprintf("[POST /api/backend/v1/payments/limits/changeStorage][%d] changeStorageForbidden ", 403)
}

func (o *ChangeStorageForbidden) String() string {
	return fmt.Sprintf("[POST /api/backend/v1/payments/limits/changeStorage][%d] changeStorageForbidden ", 403)
}

func (o *ChangeStorageForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewChangeStorageNotFound creates a ChangeStorageNotFound with default headers values
func NewChangeStorageNotFound() *ChangeStorageNotFound {
	return &ChangeStorageNotFound{}
}

/*
ChangeStorageNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ChangeStorageNotFound struct {
}

// IsSuccess returns true when this change storage not found response has a 2xx status code
func (o *ChangeStorageNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this change storage not found response has a 3xx status code
func (o *ChangeStorageNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this change storage not found response has a 4xx status code
func (o *ChangeStorageNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this change storage not found response has a 5xx status code
func (o *ChangeStorageNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this change storage not found response a status code equal to that given
func (o *ChangeStorageNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the change storage not found response
func (o *ChangeStorageNotFound) Code() int {
	return 404
}

func (o *ChangeStorageNotFound) Error() string {
	return fmt.Sprintf("[POST /api/backend/v1/payments/limits/changeStorage][%d] changeStorageNotFound ", 404)
}

func (o *ChangeStorageNotFound) String() string {
	return fmt.Sprintf("[POST /api/backend/v1/payments/limits/changeStorage][%d] changeStorageNotFound ", 404)
}

func (o *ChangeStorageNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewChangeStorageRequestTimeout creates a ChangeStorageRequestTimeout with default headers values
func NewChangeStorageRequestTimeout() *ChangeStorageRequestTimeout {
	return &ChangeStorageRequestTimeout{}
}

/*
ChangeStorageRequestTimeout describes a response with status code 408, with default header values.

Request Timeout
*/
type ChangeStorageRequestTimeout struct {
}

// IsSuccess returns true when this change storage request timeout response has a 2xx status code
func (o *ChangeStorageRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this change storage request timeout response has a 3xx status code
func (o *ChangeStorageRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this change storage request timeout response has a 4xx status code
func (o *ChangeStorageRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this change storage request timeout response has a 5xx status code
func (o *ChangeStorageRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this change storage request timeout response a status code equal to that given
func (o *ChangeStorageRequestTimeout) IsCode(code int) bool {
	return code == 408
}

// Code gets the status code for the change storage request timeout response
func (o *ChangeStorageRequestTimeout) Code() int {
	return 408
}

func (o *ChangeStorageRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/backend/v1/payments/limits/changeStorage][%d] changeStorageRequestTimeout ", 408)
}

func (o *ChangeStorageRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/backend/v1/payments/limits/changeStorage][%d] changeStorageRequestTimeout ", 408)
}

func (o *ChangeStorageRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewChangeStorageUnprocessableEntity creates a ChangeStorageUnprocessableEntity with default headers values
func NewChangeStorageUnprocessableEntity() *ChangeStorageUnprocessableEntity {
	return &ChangeStorageUnprocessableEntity{}
}

/*
ChangeStorageUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type ChangeStorageUnprocessableEntity struct {
}

// IsSuccess returns true when this change storage unprocessable entity response has a 2xx status code
func (o *ChangeStorageUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this change storage unprocessable entity response has a 3xx status code
func (o *ChangeStorageUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this change storage unprocessable entity response has a 4xx status code
func (o *ChangeStorageUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this change storage unprocessable entity response has a 5xx status code
func (o *ChangeStorageUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this change storage unprocessable entity response a status code equal to that given
func (o *ChangeStorageUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the change storage unprocessable entity response
func (o *ChangeStorageUnprocessableEntity) Code() int {
	return 422
}

func (o *ChangeStorageUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /api/backend/v1/payments/limits/changeStorage][%d] changeStorageUnprocessableEntity ", 422)
}

func (o *ChangeStorageUnprocessableEntity) String() string {
	return fmt.Sprintf("[POST /api/backend/v1/payments/limits/changeStorage][%d] changeStorageUnprocessableEntity ", 422)
}

func (o *ChangeStorageUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}