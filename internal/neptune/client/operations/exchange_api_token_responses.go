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

// ExchangeAPITokenReader is a Reader for the ExchangeAPIToken structure.
type ExchangeAPITokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExchangeAPITokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExchangeAPITokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewExchangeAPITokenBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewExchangeAPITokenUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewExchangeAPITokenForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewExchangeAPITokenNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewExchangeAPITokenRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewExchangeAPITokenUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /api/backend/v1/authorization/oauth-token] exchangeApiToken", response, response.Code())
	}
}

// NewExchangeAPITokenOK creates a ExchangeAPITokenOK with default headers values
func NewExchangeAPITokenOK() *ExchangeAPITokenOK {
	return &ExchangeAPITokenOK{}
}

/*
ExchangeAPITokenOK describes a response with status code 200, with default header values.

OK
*/
type ExchangeAPITokenOK struct {
	Payload *models.NeptuneOauthToken
}

// IsSuccess returns true when this exchange Api token o k response has a 2xx status code
func (o *ExchangeAPITokenOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this exchange Api token o k response has a 3xx status code
func (o *ExchangeAPITokenOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this exchange Api token o k response has a 4xx status code
func (o *ExchangeAPITokenOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this exchange Api token o k response has a 5xx status code
func (o *ExchangeAPITokenOK) IsServerError() bool {
	return false
}

// IsCode returns true when this exchange Api token o k response a status code equal to that given
func (o *ExchangeAPITokenOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the exchange Api token o k response
func (o *ExchangeAPITokenOK) Code() int {
	return 200
}

func (o *ExchangeAPITokenOK) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/authorization/oauth-token][%d] exchangeApiTokenOK  %+v", 200, o.Payload)
}

func (o *ExchangeAPITokenOK) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/authorization/oauth-token][%d] exchangeApiTokenOK  %+v", 200, o.Payload)
}

func (o *ExchangeAPITokenOK) GetPayload() *models.NeptuneOauthToken {
	return o.Payload
}

func (o *ExchangeAPITokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NeptuneOauthToken)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExchangeAPITokenBadRequest creates a ExchangeAPITokenBadRequest with default headers values
func NewExchangeAPITokenBadRequest() *ExchangeAPITokenBadRequest {
	return &ExchangeAPITokenBadRequest{}
}

/*
ExchangeAPITokenBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ExchangeAPITokenBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this exchange Api token bad request response has a 2xx status code
func (o *ExchangeAPITokenBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this exchange Api token bad request response has a 3xx status code
func (o *ExchangeAPITokenBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this exchange Api token bad request response has a 4xx status code
func (o *ExchangeAPITokenBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this exchange Api token bad request response has a 5xx status code
func (o *ExchangeAPITokenBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this exchange Api token bad request response a status code equal to that given
func (o *ExchangeAPITokenBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the exchange Api token bad request response
func (o *ExchangeAPITokenBadRequest) Code() int {
	return 400
}

func (o *ExchangeAPITokenBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/authorization/oauth-token][%d] exchangeApiTokenBadRequest  %+v", 400, o.Payload)
}

func (o *ExchangeAPITokenBadRequest) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/authorization/oauth-token][%d] exchangeApiTokenBadRequest  %+v", 400, o.Payload)
}

func (o *ExchangeAPITokenBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ExchangeAPITokenBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExchangeAPITokenUnauthorized creates a ExchangeAPITokenUnauthorized with default headers values
func NewExchangeAPITokenUnauthorized() *ExchangeAPITokenUnauthorized {
	return &ExchangeAPITokenUnauthorized{}
}

/*
ExchangeAPITokenUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ExchangeAPITokenUnauthorized struct {
}

// IsSuccess returns true when this exchange Api token unauthorized response has a 2xx status code
func (o *ExchangeAPITokenUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this exchange Api token unauthorized response has a 3xx status code
func (o *ExchangeAPITokenUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this exchange Api token unauthorized response has a 4xx status code
func (o *ExchangeAPITokenUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this exchange Api token unauthorized response has a 5xx status code
func (o *ExchangeAPITokenUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this exchange Api token unauthorized response a status code equal to that given
func (o *ExchangeAPITokenUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the exchange Api token unauthorized response
func (o *ExchangeAPITokenUnauthorized) Code() int {
	return 401
}

func (o *ExchangeAPITokenUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/authorization/oauth-token][%d] exchangeApiTokenUnauthorized ", 401)
}

func (o *ExchangeAPITokenUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/authorization/oauth-token][%d] exchangeApiTokenUnauthorized ", 401)
}

func (o *ExchangeAPITokenUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewExchangeAPITokenForbidden creates a ExchangeAPITokenForbidden with default headers values
func NewExchangeAPITokenForbidden() *ExchangeAPITokenForbidden {
	return &ExchangeAPITokenForbidden{}
}

/*
ExchangeAPITokenForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ExchangeAPITokenForbidden struct {
}

// IsSuccess returns true when this exchange Api token forbidden response has a 2xx status code
func (o *ExchangeAPITokenForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this exchange Api token forbidden response has a 3xx status code
func (o *ExchangeAPITokenForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this exchange Api token forbidden response has a 4xx status code
func (o *ExchangeAPITokenForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this exchange Api token forbidden response has a 5xx status code
func (o *ExchangeAPITokenForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this exchange Api token forbidden response a status code equal to that given
func (o *ExchangeAPITokenForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the exchange Api token forbidden response
func (o *ExchangeAPITokenForbidden) Code() int {
	return 403
}

func (o *ExchangeAPITokenForbidden) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/authorization/oauth-token][%d] exchangeApiTokenForbidden ", 403)
}

func (o *ExchangeAPITokenForbidden) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/authorization/oauth-token][%d] exchangeApiTokenForbidden ", 403)
}

func (o *ExchangeAPITokenForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewExchangeAPITokenNotFound creates a ExchangeAPITokenNotFound with default headers values
func NewExchangeAPITokenNotFound() *ExchangeAPITokenNotFound {
	return &ExchangeAPITokenNotFound{}
}

/*
ExchangeAPITokenNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ExchangeAPITokenNotFound struct {
}

// IsSuccess returns true when this exchange Api token not found response has a 2xx status code
func (o *ExchangeAPITokenNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this exchange Api token not found response has a 3xx status code
func (o *ExchangeAPITokenNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this exchange Api token not found response has a 4xx status code
func (o *ExchangeAPITokenNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this exchange Api token not found response has a 5xx status code
func (o *ExchangeAPITokenNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this exchange Api token not found response a status code equal to that given
func (o *ExchangeAPITokenNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the exchange Api token not found response
func (o *ExchangeAPITokenNotFound) Code() int {
	return 404
}

func (o *ExchangeAPITokenNotFound) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/authorization/oauth-token][%d] exchangeApiTokenNotFound ", 404)
}

func (o *ExchangeAPITokenNotFound) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/authorization/oauth-token][%d] exchangeApiTokenNotFound ", 404)
}

func (o *ExchangeAPITokenNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewExchangeAPITokenRequestTimeout creates a ExchangeAPITokenRequestTimeout with default headers values
func NewExchangeAPITokenRequestTimeout() *ExchangeAPITokenRequestTimeout {
	return &ExchangeAPITokenRequestTimeout{}
}

/*
ExchangeAPITokenRequestTimeout describes a response with status code 408, with default header values.

Request Timeout
*/
type ExchangeAPITokenRequestTimeout struct {
}

// IsSuccess returns true when this exchange Api token request timeout response has a 2xx status code
func (o *ExchangeAPITokenRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this exchange Api token request timeout response has a 3xx status code
func (o *ExchangeAPITokenRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this exchange Api token request timeout response has a 4xx status code
func (o *ExchangeAPITokenRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this exchange Api token request timeout response has a 5xx status code
func (o *ExchangeAPITokenRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this exchange Api token request timeout response a status code equal to that given
func (o *ExchangeAPITokenRequestTimeout) IsCode(code int) bool {
	return code == 408
}

// Code gets the status code for the exchange Api token request timeout response
func (o *ExchangeAPITokenRequestTimeout) Code() int {
	return 408
}

func (o *ExchangeAPITokenRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/authorization/oauth-token][%d] exchangeApiTokenRequestTimeout ", 408)
}

func (o *ExchangeAPITokenRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/authorization/oauth-token][%d] exchangeApiTokenRequestTimeout ", 408)
}

func (o *ExchangeAPITokenRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewExchangeAPITokenUnprocessableEntity creates a ExchangeAPITokenUnprocessableEntity with default headers values
func NewExchangeAPITokenUnprocessableEntity() *ExchangeAPITokenUnprocessableEntity {
	return &ExchangeAPITokenUnprocessableEntity{}
}

/*
ExchangeAPITokenUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type ExchangeAPITokenUnprocessableEntity struct {
}

// IsSuccess returns true when this exchange Api token unprocessable entity response has a 2xx status code
func (o *ExchangeAPITokenUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this exchange Api token unprocessable entity response has a 3xx status code
func (o *ExchangeAPITokenUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this exchange Api token unprocessable entity response has a 4xx status code
func (o *ExchangeAPITokenUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this exchange Api token unprocessable entity response has a 5xx status code
func (o *ExchangeAPITokenUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this exchange Api token unprocessable entity response a status code equal to that given
func (o *ExchangeAPITokenUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the exchange Api token unprocessable entity response
func (o *ExchangeAPITokenUnprocessableEntity) Code() int {
	return 422
}

func (o *ExchangeAPITokenUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/authorization/oauth-token][%d] exchangeApiTokenUnprocessableEntity ", 422)
}

func (o *ExchangeAPITokenUnprocessableEntity) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/authorization/oauth-token][%d] exchangeApiTokenUnprocessableEntity ", 422)
}

func (o *ExchangeAPITokenUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}