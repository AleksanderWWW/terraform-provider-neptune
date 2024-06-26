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

// GetAPITokenReader is a Reader for the GetAPIToken structure.
type GetAPITokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPITokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAPITokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAPITokenBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetAPITokenUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAPITokenForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAPITokenNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetAPITokenRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGetAPITokenUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /api/backend/v1/authorization/api-token] getApiToken", response, response.Code())
	}
}

// NewGetAPITokenOK creates a GetAPITokenOK with default headers values
func NewGetAPITokenOK() *GetAPITokenOK {
	return &GetAPITokenOK{}
}

/*
GetAPITokenOK describes a response with status code 200, with default header values.

OK
*/
type GetAPITokenOK struct {
	Payload *models.NeptuneAPIToken
}

// IsSuccess returns true when this get Api token o k response has a 2xx status code
func (o *GetAPITokenOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get Api token o k response has a 3xx status code
func (o *GetAPITokenOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Api token o k response has a 4xx status code
func (o *GetAPITokenOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get Api token o k response has a 5xx status code
func (o *GetAPITokenOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get Api token o k response a status code equal to that given
func (o *GetAPITokenOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get Api token o k response
func (o *GetAPITokenOK) Code() int {
	return 200
}

func (o *GetAPITokenOK) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/authorization/api-token][%d] getApiTokenOK  %+v", 200, o.Payload)
}

func (o *GetAPITokenOK) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/authorization/api-token][%d] getApiTokenOK  %+v", 200, o.Payload)
}

func (o *GetAPITokenOK) GetPayload() *models.NeptuneAPIToken {
	return o.Payload
}

func (o *GetAPITokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NeptuneAPIToken)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPITokenBadRequest creates a GetAPITokenBadRequest with default headers values
func NewGetAPITokenBadRequest() *GetAPITokenBadRequest {
	return &GetAPITokenBadRequest{}
}

/*
GetAPITokenBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetAPITokenBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this get Api token bad request response has a 2xx status code
func (o *GetAPITokenBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get Api token bad request response has a 3xx status code
func (o *GetAPITokenBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Api token bad request response has a 4xx status code
func (o *GetAPITokenBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get Api token bad request response has a 5xx status code
func (o *GetAPITokenBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get Api token bad request response a status code equal to that given
func (o *GetAPITokenBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get Api token bad request response
func (o *GetAPITokenBadRequest) Code() int {
	return 400
}

func (o *GetAPITokenBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/authorization/api-token][%d] getApiTokenBadRequest  %+v", 400, o.Payload)
}

func (o *GetAPITokenBadRequest) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/authorization/api-token][%d] getApiTokenBadRequest  %+v", 400, o.Payload)
}

func (o *GetAPITokenBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAPITokenBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPITokenUnauthorized creates a GetAPITokenUnauthorized with default headers values
func NewGetAPITokenUnauthorized() *GetAPITokenUnauthorized {
	return &GetAPITokenUnauthorized{}
}

/*
GetAPITokenUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetAPITokenUnauthorized struct {
}

// IsSuccess returns true when this get Api token unauthorized response has a 2xx status code
func (o *GetAPITokenUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get Api token unauthorized response has a 3xx status code
func (o *GetAPITokenUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Api token unauthorized response has a 4xx status code
func (o *GetAPITokenUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get Api token unauthorized response has a 5xx status code
func (o *GetAPITokenUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get Api token unauthorized response a status code equal to that given
func (o *GetAPITokenUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get Api token unauthorized response
func (o *GetAPITokenUnauthorized) Code() int {
	return 401
}

func (o *GetAPITokenUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/authorization/api-token][%d] getApiTokenUnauthorized ", 401)
}

func (o *GetAPITokenUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/authorization/api-token][%d] getApiTokenUnauthorized ", 401)
}

func (o *GetAPITokenUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAPITokenForbidden creates a GetAPITokenForbidden with default headers values
func NewGetAPITokenForbidden() *GetAPITokenForbidden {
	return &GetAPITokenForbidden{}
}

/*
GetAPITokenForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetAPITokenForbidden struct {
}

// IsSuccess returns true when this get Api token forbidden response has a 2xx status code
func (o *GetAPITokenForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get Api token forbidden response has a 3xx status code
func (o *GetAPITokenForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Api token forbidden response has a 4xx status code
func (o *GetAPITokenForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get Api token forbidden response has a 5xx status code
func (o *GetAPITokenForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get Api token forbidden response a status code equal to that given
func (o *GetAPITokenForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get Api token forbidden response
func (o *GetAPITokenForbidden) Code() int {
	return 403
}

func (o *GetAPITokenForbidden) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/authorization/api-token][%d] getApiTokenForbidden ", 403)
}

func (o *GetAPITokenForbidden) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/authorization/api-token][%d] getApiTokenForbidden ", 403)
}

func (o *GetAPITokenForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAPITokenNotFound creates a GetAPITokenNotFound with default headers values
func NewGetAPITokenNotFound() *GetAPITokenNotFound {
	return &GetAPITokenNotFound{}
}

/*
GetAPITokenNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetAPITokenNotFound struct {
}

// IsSuccess returns true when this get Api token not found response has a 2xx status code
func (o *GetAPITokenNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get Api token not found response has a 3xx status code
func (o *GetAPITokenNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Api token not found response has a 4xx status code
func (o *GetAPITokenNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get Api token not found response has a 5xx status code
func (o *GetAPITokenNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get Api token not found response a status code equal to that given
func (o *GetAPITokenNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get Api token not found response
func (o *GetAPITokenNotFound) Code() int {
	return 404
}

func (o *GetAPITokenNotFound) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/authorization/api-token][%d] getApiTokenNotFound ", 404)
}

func (o *GetAPITokenNotFound) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/authorization/api-token][%d] getApiTokenNotFound ", 404)
}

func (o *GetAPITokenNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAPITokenRequestTimeout creates a GetAPITokenRequestTimeout with default headers values
func NewGetAPITokenRequestTimeout() *GetAPITokenRequestTimeout {
	return &GetAPITokenRequestTimeout{}
}

/*
GetAPITokenRequestTimeout describes a response with status code 408, with default header values.

Request Timeout
*/
type GetAPITokenRequestTimeout struct {
}

// IsSuccess returns true when this get Api token request timeout response has a 2xx status code
func (o *GetAPITokenRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get Api token request timeout response has a 3xx status code
func (o *GetAPITokenRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Api token request timeout response has a 4xx status code
func (o *GetAPITokenRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get Api token request timeout response has a 5xx status code
func (o *GetAPITokenRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get Api token request timeout response a status code equal to that given
func (o *GetAPITokenRequestTimeout) IsCode(code int) bool {
	return code == 408
}

// Code gets the status code for the get Api token request timeout response
func (o *GetAPITokenRequestTimeout) Code() int {
	return 408
}

func (o *GetAPITokenRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/authorization/api-token][%d] getApiTokenRequestTimeout ", 408)
}

func (o *GetAPITokenRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/authorization/api-token][%d] getApiTokenRequestTimeout ", 408)
}

func (o *GetAPITokenRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAPITokenUnprocessableEntity creates a GetAPITokenUnprocessableEntity with default headers values
func NewGetAPITokenUnprocessableEntity() *GetAPITokenUnprocessableEntity {
	return &GetAPITokenUnprocessableEntity{}
}

/*
GetAPITokenUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type GetAPITokenUnprocessableEntity struct {
}

// IsSuccess returns true when this get Api token unprocessable entity response has a 2xx status code
func (o *GetAPITokenUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get Api token unprocessable entity response has a 3xx status code
func (o *GetAPITokenUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Api token unprocessable entity response has a 4xx status code
func (o *GetAPITokenUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this get Api token unprocessable entity response has a 5xx status code
func (o *GetAPITokenUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this get Api token unprocessable entity response a status code equal to that given
func (o *GetAPITokenUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the get Api token unprocessable entity response
func (o *GetAPITokenUnprocessableEntity) Code() int {
	return 422
}

func (o *GetAPITokenUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/authorization/api-token][%d] getApiTokenUnprocessableEntity ", 422)
}

func (o *GetAPITokenUnprocessableEntity) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/authorization/api-token][%d] getApiTokenUnprocessableEntity ", 422)
}

func (o *GetAPITokenUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
