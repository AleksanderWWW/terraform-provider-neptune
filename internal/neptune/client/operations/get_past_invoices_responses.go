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

// GetPastInvoicesReader is a Reader for the GetPastInvoices structure.
type GetPastInvoicesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPastInvoicesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPastInvoicesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetPastInvoicesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetPastInvoicesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetPastInvoicesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetPastInvoicesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetPastInvoicesRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGetPastInvoicesUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /api/backend/v1/payments/{organizationIdentifier}/invoices/past] getPastInvoices", response, response.Code())
	}
}

// NewGetPastInvoicesOK creates a GetPastInvoicesOK with default headers values
func NewGetPastInvoicesOK() *GetPastInvoicesOK {
	return &GetPastInvoicesOK{}
}

/*
GetPastInvoicesOK describes a response with status code 200, with default header values.

OK
*/
type GetPastInvoicesOK struct {
	Payload string
}

// IsSuccess returns true when this get past invoices o k response has a 2xx status code
func (o *GetPastInvoicesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get past invoices o k response has a 3xx status code
func (o *GetPastInvoicesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get past invoices o k response has a 4xx status code
func (o *GetPastInvoicesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get past invoices o k response has a 5xx status code
func (o *GetPastInvoicesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get past invoices o k response a status code equal to that given
func (o *GetPastInvoicesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get past invoices o k response
func (o *GetPastInvoicesOK) Code() int {
	return 200
}

func (o *GetPastInvoicesOK) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/payments/{organizationIdentifier}/invoices/past][%d] getPastInvoicesOK  %+v", 200, o.Payload)
}

func (o *GetPastInvoicesOK) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/payments/{organizationIdentifier}/invoices/past][%d] getPastInvoicesOK  %+v", 200, o.Payload)
}

func (o *GetPastInvoicesOK) GetPayload() string {
	return o.Payload
}

func (o *GetPastInvoicesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPastInvoicesBadRequest creates a GetPastInvoicesBadRequest with default headers values
func NewGetPastInvoicesBadRequest() *GetPastInvoicesBadRequest {
	return &GetPastInvoicesBadRequest{}
}

/*
GetPastInvoicesBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetPastInvoicesBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this get past invoices bad request response has a 2xx status code
func (o *GetPastInvoicesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get past invoices bad request response has a 3xx status code
func (o *GetPastInvoicesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get past invoices bad request response has a 4xx status code
func (o *GetPastInvoicesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get past invoices bad request response has a 5xx status code
func (o *GetPastInvoicesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get past invoices bad request response a status code equal to that given
func (o *GetPastInvoicesBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get past invoices bad request response
func (o *GetPastInvoicesBadRequest) Code() int {
	return 400
}

func (o *GetPastInvoicesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/payments/{organizationIdentifier}/invoices/past][%d] getPastInvoicesBadRequest  %+v", 400, o.Payload)
}

func (o *GetPastInvoicesBadRequest) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/payments/{organizationIdentifier}/invoices/past][%d] getPastInvoicesBadRequest  %+v", 400, o.Payload)
}

func (o *GetPastInvoicesBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetPastInvoicesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPastInvoicesUnauthorized creates a GetPastInvoicesUnauthorized with default headers values
func NewGetPastInvoicesUnauthorized() *GetPastInvoicesUnauthorized {
	return &GetPastInvoicesUnauthorized{}
}

/*
GetPastInvoicesUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetPastInvoicesUnauthorized struct {
}

// IsSuccess returns true when this get past invoices unauthorized response has a 2xx status code
func (o *GetPastInvoicesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get past invoices unauthorized response has a 3xx status code
func (o *GetPastInvoicesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get past invoices unauthorized response has a 4xx status code
func (o *GetPastInvoicesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get past invoices unauthorized response has a 5xx status code
func (o *GetPastInvoicesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get past invoices unauthorized response a status code equal to that given
func (o *GetPastInvoicesUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get past invoices unauthorized response
func (o *GetPastInvoicesUnauthorized) Code() int {
	return 401
}

func (o *GetPastInvoicesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/payments/{organizationIdentifier}/invoices/past][%d] getPastInvoicesUnauthorized ", 401)
}

func (o *GetPastInvoicesUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/payments/{organizationIdentifier}/invoices/past][%d] getPastInvoicesUnauthorized ", 401)
}

func (o *GetPastInvoicesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPastInvoicesForbidden creates a GetPastInvoicesForbidden with default headers values
func NewGetPastInvoicesForbidden() *GetPastInvoicesForbidden {
	return &GetPastInvoicesForbidden{}
}

/*
GetPastInvoicesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetPastInvoicesForbidden struct {
}

// IsSuccess returns true when this get past invoices forbidden response has a 2xx status code
func (o *GetPastInvoicesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get past invoices forbidden response has a 3xx status code
func (o *GetPastInvoicesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get past invoices forbidden response has a 4xx status code
func (o *GetPastInvoicesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get past invoices forbidden response has a 5xx status code
func (o *GetPastInvoicesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get past invoices forbidden response a status code equal to that given
func (o *GetPastInvoicesForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get past invoices forbidden response
func (o *GetPastInvoicesForbidden) Code() int {
	return 403
}

func (o *GetPastInvoicesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/payments/{organizationIdentifier}/invoices/past][%d] getPastInvoicesForbidden ", 403)
}

func (o *GetPastInvoicesForbidden) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/payments/{organizationIdentifier}/invoices/past][%d] getPastInvoicesForbidden ", 403)
}

func (o *GetPastInvoicesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPastInvoicesNotFound creates a GetPastInvoicesNotFound with default headers values
func NewGetPastInvoicesNotFound() *GetPastInvoicesNotFound {
	return &GetPastInvoicesNotFound{}
}

/*
GetPastInvoicesNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetPastInvoicesNotFound struct {
}

// IsSuccess returns true when this get past invoices not found response has a 2xx status code
func (o *GetPastInvoicesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get past invoices not found response has a 3xx status code
func (o *GetPastInvoicesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get past invoices not found response has a 4xx status code
func (o *GetPastInvoicesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get past invoices not found response has a 5xx status code
func (o *GetPastInvoicesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get past invoices not found response a status code equal to that given
func (o *GetPastInvoicesNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get past invoices not found response
func (o *GetPastInvoicesNotFound) Code() int {
	return 404
}

func (o *GetPastInvoicesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/payments/{organizationIdentifier}/invoices/past][%d] getPastInvoicesNotFound ", 404)
}

func (o *GetPastInvoicesNotFound) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/payments/{organizationIdentifier}/invoices/past][%d] getPastInvoicesNotFound ", 404)
}

func (o *GetPastInvoicesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPastInvoicesRequestTimeout creates a GetPastInvoicesRequestTimeout with default headers values
func NewGetPastInvoicesRequestTimeout() *GetPastInvoicesRequestTimeout {
	return &GetPastInvoicesRequestTimeout{}
}

/*
GetPastInvoicesRequestTimeout describes a response with status code 408, with default header values.

Request Timeout
*/
type GetPastInvoicesRequestTimeout struct {
}

// IsSuccess returns true when this get past invoices request timeout response has a 2xx status code
func (o *GetPastInvoicesRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get past invoices request timeout response has a 3xx status code
func (o *GetPastInvoicesRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get past invoices request timeout response has a 4xx status code
func (o *GetPastInvoicesRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get past invoices request timeout response has a 5xx status code
func (o *GetPastInvoicesRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get past invoices request timeout response a status code equal to that given
func (o *GetPastInvoicesRequestTimeout) IsCode(code int) bool {
	return code == 408
}

// Code gets the status code for the get past invoices request timeout response
func (o *GetPastInvoicesRequestTimeout) Code() int {
	return 408
}

func (o *GetPastInvoicesRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/payments/{organizationIdentifier}/invoices/past][%d] getPastInvoicesRequestTimeout ", 408)
}

func (o *GetPastInvoicesRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/payments/{organizationIdentifier}/invoices/past][%d] getPastInvoicesRequestTimeout ", 408)
}

func (o *GetPastInvoicesRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPastInvoicesUnprocessableEntity creates a GetPastInvoicesUnprocessableEntity with default headers values
func NewGetPastInvoicesUnprocessableEntity() *GetPastInvoicesUnprocessableEntity {
	return &GetPastInvoicesUnprocessableEntity{}
}

/*
GetPastInvoicesUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type GetPastInvoicesUnprocessableEntity struct {
}

// IsSuccess returns true when this get past invoices unprocessable entity response has a 2xx status code
func (o *GetPastInvoicesUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get past invoices unprocessable entity response has a 3xx status code
func (o *GetPastInvoicesUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get past invoices unprocessable entity response has a 4xx status code
func (o *GetPastInvoicesUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this get past invoices unprocessable entity response has a 5xx status code
func (o *GetPastInvoicesUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this get past invoices unprocessable entity response a status code equal to that given
func (o *GetPastInvoicesUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the get past invoices unprocessable entity response
func (o *GetPastInvoicesUnprocessableEntity) Code() int {
	return 422
}

func (o *GetPastInvoicesUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/payments/{organizationIdentifier}/invoices/past][%d] getPastInvoicesUnprocessableEntity ", 422)
}

func (o *GetPastInvoicesUnprocessableEntity) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/payments/{organizationIdentifier}/invoices/past][%d] getPastInvoicesUnprocessableEntity ", 422)
}

func (o *GetPastInvoicesUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}