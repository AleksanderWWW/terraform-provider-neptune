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

// DetachCreditCardReader is a Reader for the DetachCreditCard structure.
type DetachCreditCardReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DetachCreditCardReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDetachCreditCardOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDetachCreditCardBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDetachCreditCardUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDetachCreditCardForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDetachCreditCardNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewDetachCreditCardRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewDetachCreditCardUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /api/backend/v1/payments/{organizationIdentifier}/creditCard] detachCreditCard", response, response.Code())
	}
}

// NewDetachCreditCardOK creates a DetachCreditCardOK with default headers values
func NewDetachCreditCardOK() *DetachCreditCardOK {
	return &DetachCreditCardOK{}
}

/*
DetachCreditCardOK describes a response with status code 200, with default header values.

No response
*/
type DetachCreditCardOK struct {
}

// IsSuccess returns true when this detach credit card o k response has a 2xx status code
func (o *DetachCreditCardOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this detach credit card o k response has a 3xx status code
func (o *DetachCreditCardOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this detach credit card o k response has a 4xx status code
func (o *DetachCreditCardOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this detach credit card o k response has a 5xx status code
func (o *DetachCreditCardOK) IsServerError() bool {
	return false
}

// IsCode returns true when this detach credit card o k response a status code equal to that given
func (o *DetachCreditCardOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the detach credit card o k response
func (o *DetachCreditCardOK) Code() int {
	return 200
}

func (o *DetachCreditCardOK) Error() string {
	return fmt.Sprintf("[DELETE /api/backend/v1/payments/{organizationIdentifier}/creditCard][%d] detachCreditCardOK ", 200)
}

func (o *DetachCreditCardOK) String() string {
	return fmt.Sprintf("[DELETE /api/backend/v1/payments/{organizationIdentifier}/creditCard][%d] detachCreditCardOK ", 200)
}

func (o *DetachCreditCardOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDetachCreditCardBadRequest creates a DetachCreditCardBadRequest with default headers values
func NewDetachCreditCardBadRequest() *DetachCreditCardBadRequest {
	return &DetachCreditCardBadRequest{}
}

/*
DetachCreditCardBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DetachCreditCardBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this detach credit card bad request response has a 2xx status code
func (o *DetachCreditCardBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this detach credit card bad request response has a 3xx status code
func (o *DetachCreditCardBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this detach credit card bad request response has a 4xx status code
func (o *DetachCreditCardBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this detach credit card bad request response has a 5xx status code
func (o *DetachCreditCardBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this detach credit card bad request response a status code equal to that given
func (o *DetachCreditCardBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the detach credit card bad request response
func (o *DetachCreditCardBadRequest) Code() int {
	return 400
}

func (o *DetachCreditCardBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/backend/v1/payments/{organizationIdentifier}/creditCard][%d] detachCreditCardBadRequest  %+v", 400, o.Payload)
}

func (o *DetachCreditCardBadRequest) String() string {
	return fmt.Sprintf("[DELETE /api/backend/v1/payments/{organizationIdentifier}/creditCard][%d] detachCreditCardBadRequest  %+v", 400, o.Payload)
}

func (o *DetachCreditCardBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *DetachCreditCardBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDetachCreditCardUnauthorized creates a DetachCreditCardUnauthorized with default headers values
func NewDetachCreditCardUnauthorized() *DetachCreditCardUnauthorized {
	return &DetachCreditCardUnauthorized{}
}

/*
DetachCreditCardUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DetachCreditCardUnauthorized struct {
}

// IsSuccess returns true when this detach credit card unauthorized response has a 2xx status code
func (o *DetachCreditCardUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this detach credit card unauthorized response has a 3xx status code
func (o *DetachCreditCardUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this detach credit card unauthorized response has a 4xx status code
func (o *DetachCreditCardUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this detach credit card unauthorized response has a 5xx status code
func (o *DetachCreditCardUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this detach credit card unauthorized response a status code equal to that given
func (o *DetachCreditCardUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the detach credit card unauthorized response
func (o *DetachCreditCardUnauthorized) Code() int {
	return 401
}

func (o *DetachCreditCardUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/backend/v1/payments/{organizationIdentifier}/creditCard][%d] detachCreditCardUnauthorized ", 401)
}

func (o *DetachCreditCardUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /api/backend/v1/payments/{organizationIdentifier}/creditCard][%d] detachCreditCardUnauthorized ", 401)
}

func (o *DetachCreditCardUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDetachCreditCardForbidden creates a DetachCreditCardForbidden with default headers values
func NewDetachCreditCardForbidden() *DetachCreditCardForbidden {
	return &DetachCreditCardForbidden{}
}

/*
DetachCreditCardForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DetachCreditCardForbidden struct {
}

// IsSuccess returns true when this detach credit card forbidden response has a 2xx status code
func (o *DetachCreditCardForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this detach credit card forbidden response has a 3xx status code
func (o *DetachCreditCardForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this detach credit card forbidden response has a 4xx status code
func (o *DetachCreditCardForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this detach credit card forbidden response has a 5xx status code
func (o *DetachCreditCardForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this detach credit card forbidden response a status code equal to that given
func (o *DetachCreditCardForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the detach credit card forbidden response
func (o *DetachCreditCardForbidden) Code() int {
	return 403
}

func (o *DetachCreditCardForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/backend/v1/payments/{organizationIdentifier}/creditCard][%d] detachCreditCardForbidden ", 403)
}

func (o *DetachCreditCardForbidden) String() string {
	return fmt.Sprintf("[DELETE /api/backend/v1/payments/{organizationIdentifier}/creditCard][%d] detachCreditCardForbidden ", 403)
}

func (o *DetachCreditCardForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDetachCreditCardNotFound creates a DetachCreditCardNotFound with default headers values
func NewDetachCreditCardNotFound() *DetachCreditCardNotFound {
	return &DetachCreditCardNotFound{}
}

/*
DetachCreditCardNotFound describes a response with status code 404, with default header values.

Not Found
*/
type DetachCreditCardNotFound struct {
}

// IsSuccess returns true when this detach credit card not found response has a 2xx status code
func (o *DetachCreditCardNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this detach credit card not found response has a 3xx status code
func (o *DetachCreditCardNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this detach credit card not found response has a 4xx status code
func (o *DetachCreditCardNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this detach credit card not found response has a 5xx status code
func (o *DetachCreditCardNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this detach credit card not found response a status code equal to that given
func (o *DetachCreditCardNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the detach credit card not found response
func (o *DetachCreditCardNotFound) Code() int {
	return 404
}

func (o *DetachCreditCardNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/backend/v1/payments/{organizationIdentifier}/creditCard][%d] detachCreditCardNotFound ", 404)
}

func (o *DetachCreditCardNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/backend/v1/payments/{organizationIdentifier}/creditCard][%d] detachCreditCardNotFound ", 404)
}

func (o *DetachCreditCardNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDetachCreditCardRequestTimeout creates a DetachCreditCardRequestTimeout with default headers values
func NewDetachCreditCardRequestTimeout() *DetachCreditCardRequestTimeout {
	return &DetachCreditCardRequestTimeout{}
}

/*
DetachCreditCardRequestTimeout describes a response with status code 408, with default header values.

Request Timeout
*/
type DetachCreditCardRequestTimeout struct {
}

// IsSuccess returns true when this detach credit card request timeout response has a 2xx status code
func (o *DetachCreditCardRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this detach credit card request timeout response has a 3xx status code
func (o *DetachCreditCardRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this detach credit card request timeout response has a 4xx status code
func (o *DetachCreditCardRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this detach credit card request timeout response has a 5xx status code
func (o *DetachCreditCardRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this detach credit card request timeout response a status code equal to that given
func (o *DetachCreditCardRequestTimeout) IsCode(code int) bool {
	return code == 408
}

// Code gets the status code for the detach credit card request timeout response
func (o *DetachCreditCardRequestTimeout) Code() int {
	return 408
}

func (o *DetachCreditCardRequestTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/backend/v1/payments/{organizationIdentifier}/creditCard][%d] detachCreditCardRequestTimeout ", 408)
}

func (o *DetachCreditCardRequestTimeout) String() string {
	return fmt.Sprintf("[DELETE /api/backend/v1/payments/{organizationIdentifier}/creditCard][%d] detachCreditCardRequestTimeout ", 408)
}

func (o *DetachCreditCardRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDetachCreditCardUnprocessableEntity creates a DetachCreditCardUnprocessableEntity with default headers values
func NewDetachCreditCardUnprocessableEntity() *DetachCreditCardUnprocessableEntity {
	return &DetachCreditCardUnprocessableEntity{}
}

/*
DetachCreditCardUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type DetachCreditCardUnprocessableEntity struct {
}

// IsSuccess returns true when this detach credit card unprocessable entity response has a 2xx status code
func (o *DetachCreditCardUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this detach credit card unprocessable entity response has a 3xx status code
func (o *DetachCreditCardUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this detach credit card unprocessable entity response has a 4xx status code
func (o *DetachCreditCardUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this detach credit card unprocessable entity response has a 5xx status code
func (o *DetachCreditCardUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this detach credit card unprocessable entity response a status code equal to that given
func (o *DetachCreditCardUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the detach credit card unprocessable entity response
func (o *DetachCreditCardUnprocessableEntity) Code() int {
	return 422
}

func (o *DetachCreditCardUnprocessableEntity) Error() string {
	return fmt.Sprintf("[DELETE /api/backend/v1/payments/{organizationIdentifier}/creditCard][%d] detachCreditCardUnprocessableEntity ", 422)
}

func (o *DetachCreditCardUnprocessableEntity) String() string {
	return fmt.Sprintf("[DELETE /api/backend/v1/payments/{organizationIdentifier}/creditCard][%d] detachCreditCardUnprocessableEntity ", 422)
}

func (o *DetachCreditCardUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
