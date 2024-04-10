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

// GetServiceAccountAPITokenReader is a Reader for the GetServiceAccountAPIToken structure.
type GetServiceAccountAPITokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetServiceAccountAPITokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetServiceAccountAPITokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetServiceAccountAPITokenBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetServiceAccountAPITokenUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetServiceAccountAPITokenForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetServiceAccountAPITokenNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetServiceAccountAPITokenRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGetServiceAccountAPITokenUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /api/backend/v1/authorization/serviceAccounts/api-token] getServiceAccountApiToken", response, response.Code())
	}
}

// NewGetServiceAccountAPITokenOK creates a GetServiceAccountAPITokenOK with default headers values
func NewGetServiceAccountAPITokenOK() *GetServiceAccountAPITokenOK {
	return &GetServiceAccountAPITokenOK{}
}

/*
GetServiceAccountAPITokenOK describes a response with status code 200, with default header values.

OK
*/
type GetServiceAccountAPITokenOK struct {
	Payload *models.NeptuneAPIToken
}

// IsSuccess returns true when this get service account Api token o k response has a 2xx status code
func (o *GetServiceAccountAPITokenOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get service account Api token o k response has a 3xx status code
func (o *GetServiceAccountAPITokenOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get service account Api token o k response has a 4xx status code
func (o *GetServiceAccountAPITokenOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get service account Api token o k response has a 5xx status code
func (o *GetServiceAccountAPITokenOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get service account Api token o k response a status code equal to that given
func (o *GetServiceAccountAPITokenOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get service account Api token o k response
func (o *GetServiceAccountAPITokenOK) Code() int {
	return 200
}

func (o *GetServiceAccountAPITokenOK) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/authorization/serviceAccounts/api-token][%d] getServiceAccountApiTokenOK  %+v", 200, o.Payload)
}

func (o *GetServiceAccountAPITokenOK) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/authorization/serviceAccounts/api-token][%d] getServiceAccountApiTokenOK  %+v", 200, o.Payload)
}

func (o *GetServiceAccountAPITokenOK) GetPayload() *models.NeptuneAPIToken {
	return o.Payload
}

func (o *GetServiceAccountAPITokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NeptuneAPIToken)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetServiceAccountAPITokenBadRequest creates a GetServiceAccountAPITokenBadRequest with default headers values
func NewGetServiceAccountAPITokenBadRequest() *GetServiceAccountAPITokenBadRequest {
	return &GetServiceAccountAPITokenBadRequest{}
}

/*
GetServiceAccountAPITokenBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetServiceAccountAPITokenBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this get service account Api token bad request response has a 2xx status code
func (o *GetServiceAccountAPITokenBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get service account Api token bad request response has a 3xx status code
func (o *GetServiceAccountAPITokenBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get service account Api token bad request response has a 4xx status code
func (o *GetServiceAccountAPITokenBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get service account Api token bad request response has a 5xx status code
func (o *GetServiceAccountAPITokenBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get service account Api token bad request response a status code equal to that given
func (o *GetServiceAccountAPITokenBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get service account Api token bad request response
func (o *GetServiceAccountAPITokenBadRequest) Code() int {
	return 400
}

func (o *GetServiceAccountAPITokenBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/authorization/serviceAccounts/api-token][%d] getServiceAccountApiTokenBadRequest  %+v", 400, o.Payload)
}

func (o *GetServiceAccountAPITokenBadRequest) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/authorization/serviceAccounts/api-token][%d] getServiceAccountApiTokenBadRequest  %+v", 400, o.Payload)
}

func (o *GetServiceAccountAPITokenBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetServiceAccountAPITokenBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetServiceAccountAPITokenUnauthorized creates a GetServiceAccountAPITokenUnauthorized with default headers values
func NewGetServiceAccountAPITokenUnauthorized() *GetServiceAccountAPITokenUnauthorized {
	return &GetServiceAccountAPITokenUnauthorized{}
}

/*
GetServiceAccountAPITokenUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetServiceAccountAPITokenUnauthorized struct {
}

// IsSuccess returns true when this get service account Api token unauthorized response has a 2xx status code
func (o *GetServiceAccountAPITokenUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get service account Api token unauthorized response has a 3xx status code
func (o *GetServiceAccountAPITokenUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get service account Api token unauthorized response has a 4xx status code
func (o *GetServiceAccountAPITokenUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get service account Api token unauthorized response has a 5xx status code
func (o *GetServiceAccountAPITokenUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get service account Api token unauthorized response a status code equal to that given
func (o *GetServiceAccountAPITokenUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get service account Api token unauthorized response
func (o *GetServiceAccountAPITokenUnauthorized) Code() int {
	return 401
}

func (o *GetServiceAccountAPITokenUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/authorization/serviceAccounts/api-token][%d] getServiceAccountApiTokenUnauthorized ", 401)
}

func (o *GetServiceAccountAPITokenUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/authorization/serviceAccounts/api-token][%d] getServiceAccountApiTokenUnauthorized ", 401)
}

func (o *GetServiceAccountAPITokenUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetServiceAccountAPITokenForbidden creates a GetServiceAccountAPITokenForbidden with default headers values
func NewGetServiceAccountAPITokenForbidden() *GetServiceAccountAPITokenForbidden {
	return &GetServiceAccountAPITokenForbidden{}
}

/*
GetServiceAccountAPITokenForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetServiceAccountAPITokenForbidden struct {
}

// IsSuccess returns true when this get service account Api token forbidden response has a 2xx status code
func (o *GetServiceAccountAPITokenForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get service account Api token forbidden response has a 3xx status code
func (o *GetServiceAccountAPITokenForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get service account Api token forbidden response has a 4xx status code
func (o *GetServiceAccountAPITokenForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get service account Api token forbidden response has a 5xx status code
func (o *GetServiceAccountAPITokenForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get service account Api token forbidden response a status code equal to that given
func (o *GetServiceAccountAPITokenForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get service account Api token forbidden response
func (o *GetServiceAccountAPITokenForbidden) Code() int {
	return 403
}

func (o *GetServiceAccountAPITokenForbidden) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/authorization/serviceAccounts/api-token][%d] getServiceAccountApiTokenForbidden ", 403)
}

func (o *GetServiceAccountAPITokenForbidden) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/authorization/serviceAccounts/api-token][%d] getServiceAccountApiTokenForbidden ", 403)
}

func (o *GetServiceAccountAPITokenForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetServiceAccountAPITokenNotFound creates a GetServiceAccountAPITokenNotFound with default headers values
func NewGetServiceAccountAPITokenNotFound() *GetServiceAccountAPITokenNotFound {
	return &GetServiceAccountAPITokenNotFound{}
}

/*
GetServiceAccountAPITokenNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetServiceAccountAPITokenNotFound struct {
}

// IsSuccess returns true when this get service account Api token not found response has a 2xx status code
func (o *GetServiceAccountAPITokenNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get service account Api token not found response has a 3xx status code
func (o *GetServiceAccountAPITokenNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get service account Api token not found response has a 4xx status code
func (o *GetServiceAccountAPITokenNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get service account Api token not found response has a 5xx status code
func (o *GetServiceAccountAPITokenNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get service account Api token not found response a status code equal to that given
func (o *GetServiceAccountAPITokenNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get service account Api token not found response
func (o *GetServiceAccountAPITokenNotFound) Code() int {
	return 404
}

func (o *GetServiceAccountAPITokenNotFound) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/authorization/serviceAccounts/api-token][%d] getServiceAccountApiTokenNotFound ", 404)
}

func (o *GetServiceAccountAPITokenNotFound) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/authorization/serviceAccounts/api-token][%d] getServiceAccountApiTokenNotFound ", 404)
}

func (o *GetServiceAccountAPITokenNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetServiceAccountAPITokenRequestTimeout creates a GetServiceAccountAPITokenRequestTimeout with default headers values
func NewGetServiceAccountAPITokenRequestTimeout() *GetServiceAccountAPITokenRequestTimeout {
	return &GetServiceAccountAPITokenRequestTimeout{}
}

/*
GetServiceAccountAPITokenRequestTimeout describes a response with status code 408, with default header values.

Request Timeout
*/
type GetServiceAccountAPITokenRequestTimeout struct {
}

// IsSuccess returns true when this get service account Api token request timeout response has a 2xx status code
func (o *GetServiceAccountAPITokenRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get service account Api token request timeout response has a 3xx status code
func (o *GetServiceAccountAPITokenRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get service account Api token request timeout response has a 4xx status code
func (o *GetServiceAccountAPITokenRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get service account Api token request timeout response has a 5xx status code
func (o *GetServiceAccountAPITokenRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get service account Api token request timeout response a status code equal to that given
func (o *GetServiceAccountAPITokenRequestTimeout) IsCode(code int) bool {
	return code == 408
}

// Code gets the status code for the get service account Api token request timeout response
func (o *GetServiceAccountAPITokenRequestTimeout) Code() int {
	return 408
}

func (o *GetServiceAccountAPITokenRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/authorization/serviceAccounts/api-token][%d] getServiceAccountApiTokenRequestTimeout ", 408)
}

func (o *GetServiceAccountAPITokenRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/authorization/serviceAccounts/api-token][%d] getServiceAccountApiTokenRequestTimeout ", 408)
}

func (o *GetServiceAccountAPITokenRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetServiceAccountAPITokenUnprocessableEntity creates a GetServiceAccountAPITokenUnprocessableEntity with default headers values
func NewGetServiceAccountAPITokenUnprocessableEntity() *GetServiceAccountAPITokenUnprocessableEntity {
	return &GetServiceAccountAPITokenUnprocessableEntity{}
}

/*
GetServiceAccountAPITokenUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type GetServiceAccountAPITokenUnprocessableEntity struct {
}

// IsSuccess returns true when this get service account Api token unprocessable entity response has a 2xx status code
func (o *GetServiceAccountAPITokenUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get service account Api token unprocessable entity response has a 3xx status code
func (o *GetServiceAccountAPITokenUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get service account Api token unprocessable entity response has a 4xx status code
func (o *GetServiceAccountAPITokenUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this get service account Api token unprocessable entity response has a 5xx status code
func (o *GetServiceAccountAPITokenUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this get service account Api token unprocessable entity response a status code equal to that given
func (o *GetServiceAccountAPITokenUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the get service account Api token unprocessable entity response
func (o *GetServiceAccountAPITokenUnprocessableEntity) Code() int {
	return 422
}

func (o *GetServiceAccountAPITokenUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/authorization/serviceAccounts/api-token][%d] getServiceAccountApiTokenUnprocessableEntity ", 422)
}

func (o *GetServiceAccountAPITokenUnprocessableEntity) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/authorization/serviceAccounts/api-token][%d] getServiceAccountApiTokenUnprocessableEntity ", 422)
}

func (o *GetServiceAccountAPITokenUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
