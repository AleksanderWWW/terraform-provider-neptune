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

// RevokeServiceAccountAPITokenReader is a Reader for the RevokeServiceAccountAPIToken structure.
type RevokeServiceAccountAPITokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RevokeServiceAccountAPITokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRevokeServiceAccountAPITokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRevokeServiceAccountAPITokenBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewRevokeServiceAccountAPITokenUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRevokeServiceAccountAPITokenForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRevokeServiceAccountAPITokenNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewRevokeServiceAccountAPITokenRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewRevokeServiceAccountAPITokenUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /api/backend/v1/authorization/serviceAccounts/api-token] revokeServiceAccountApiToken", response, response.Code())
	}
}

// NewRevokeServiceAccountAPITokenOK creates a RevokeServiceAccountAPITokenOK with default headers values
func NewRevokeServiceAccountAPITokenOK() *RevokeServiceAccountAPITokenOK {
	return &RevokeServiceAccountAPITokenOK{}
}

/*
RevokeServiceAccountAPITokenOK describes a response with status code 200, with default header values.

No response
*/
type RevokeServiceAccountAPITokenOK struct {
}

// IsSuccess returns true when this revoke service account Api token o k response has a 2xx status code
func (o *RevokeServiceAccountAPITokenOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this revoke service account Api token o k response has a 3xx status code
func (o *RevokeServiceAccountAPITokenOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this revoke service account Api token o k response has a 4xx status code
func (o *RevokeServiceAccountAPITokenOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this revoke service account Api token o k response has a 5xx status code
func (o *RevokeServiceAccountAPITokenOK) IsServerError() bool {
	return false
}

// IsCode returns true when this revoke service account Api token o k response a status code equal to that given
func (o *RevokeServiceAccountAPITokenOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the revoke service account Api token o k response
func (o *RevokeServiceAccountAPITokenOK) Code() int {
	return 200
}

func (o *RevokeServiceAccountAPITokenOK) Error() string {
	return fmt.Sprintf("[DELETE /api/backend/v1/authorization/serviceAccounts/api-token][%d] revokeServiceAccountApiTokenOK ", 200)
}

func (o *RevokeServiceAccountAPITokenOK) String() string {
	return fmt.Sprintf("[DELETE /api/backend/v1/authorization/serviceAccounts/api-token][%d] revokeServiceAccountApiTokenOK ", 200)
}

func (o *RevokeServiceAccountAPITokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRevokeServiceAccountAPITokenBadRequest creates a RevokeServiceAccountAPITokenBadRequest with default headers values
func NewRevokeServiceAccountAPITokenBadRequest() *RevokeServiceAccountAPITokenBadRequest {
	return &RevokeServiceAccountAPITokenBadRequest{}
}

/*
RevokeServiceAccountAPITokenBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type RevokeServiceAccountAPITokenBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this revoke service account Api token bad request response has a 2xx status code
func (o *RevokeServiceAccountAPITokenBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this revoke service account Api token bad request response has a 3xx status code
func (o *RevokeServiceAccountAPITokenBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this revoke service account Api token bad request response has a 4xx status code
func (o *RevokeServiceAccountAPITokenBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this revoke service account Api token bad request response has a 5xx status code
func (o *RevokeServiceAccountAPITokenBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this revoke service account Api token bad request response a status code equal to that given
func (o *RevokeServiceAccountAPITokenBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the revoke service account Api token bad request response
func (o *RevokeServiceAccountAPITokenBadRequest) Code() int {
	return 400
}

func (o *RevokeServiceAccountAPITokenBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/backend/v1/authorization/serviceAccounts/api-token][%d] revokeServiceAccountApiTokenBadRequest  %+v", 400, o.Payload)
}

func (o *RevokeServiceAccountAPITokenBadRequest) String() string {
	return fmt.Sprintf("[DELETE /api/backend/v1/authorization/serviceAccounts/api-token][%d] revokeServiceAccountApiTokenBadRequest  %+v", 400, o.Payload)
}

func (o *RevokeServiceAccountAPITokenBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *RevokeServiceAccountAPITokenBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRevokeServiceAccountAPITokenUnauthorized creates a RevokeServiceAccountAPITokenUnauthorized with default headers values
func NewRevokeServiceAccountAPITokenUnauthorized() *RevokeServiceAccountAPITokenUnauthorized {
	return &RevokeServiceAccountAPITokenUnauthorized{}
}

/*
RevokeServiceAccountAPITokenUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type RevokeServiceAccountAPITokenUnauthorized struct {
}

// IsSuccess returns true when this revoke service account Api token unauthorized response has a 2xx status code
func (o *RevokeServiceAccountAPITokenUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this revoke service account Api token unauthorized response has a 3xx status code
func (o *RevokeServiceAccountAPITokenUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this revoke service account Api token unauthorized response has a 4xx status code
func (o *RevokeServiceAccountAPITokenUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this revoke service account Api token unauthorized response has a 5xx status code
func (o *RevokeServiceAccountAPITokenUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this revoke service account Api token unauthorized response a status code equal to that given
func (o *RevokeServiceAccountAPITokenUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the revoke service account Api token unauthorized response
func (o *RevokeServiceAccountAPITokenUnauthorized) Code() int {
	return 401
}

func (o *RevokeServiceAccountAPITokenUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/backend/v1/authorization/serviceAccounts/api-token][%d] revokeServiceAccountApiTokenUnauthorized ", 401)
}

func (o *RevokeServiceAccountAPITokenUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /api/backend/v1/authorization/serviceAccounts/api-token][%d] revokeServiceAccountApiTokenUnauthorized ", 401)
}

func (o *RevokeServiceAccountAPITokenUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRevokeServiceAccountAPITokenForbidden creates a RevokeServiceAccountAPITokenForbidden with default headers values
func NewRevokeServiceAccountAPITokenForbidden() *RevokeServiceAccountAPITokenForbidden {
	return &RevokeServiceAccountAPITokenForbidden{}
}

/*
RevokeServiceAccountAPITokenForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type RevokeServiceAccountAPITokenForbidden struct {
}

// IsSuccess returns true when this revoke service account Api token forbidden response has a 2xx status code
func (o *RevokeServiceAccountAPITokenForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this revoke service account Api token forbidden response has a 3xx status code
func (o *RevokeServiceAccountAPITokenForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this revoke service account Api token forbidden response has a 4xx status code
func (o *RevokeServiceAccountAPITokenForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this revoke service account Api token forbidden response has a 5xx status code
func (o *RevokeServiceAccountAPITokenForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this revoke service account Api token forbidden response a status code equal to that given
func (o *RevokeServiceAccountAPITokenForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the revoke service account Api token forbidden response
func (o *RevokeServiceAccountAPITokenForbidden) Code() int {
	return 403
}

func (o *RevokeServiceAccountAPITokenForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/backend/v1/authorization/serviceAccounts/api-token][%d] revokeServiceAccountApiTokenForbidden ", 403)
}

func (o *RevokeServiceAccountAPITokenForbidden) String() string {
	return fmt.Sprintf("[DELETE /api/backend/v1/authorization/serviceAccounts/api-token][%d] revokeServiceAccountApiTokenForbidden ", 403)
}

func (o *RevokeServiceAccountAPITokenForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRevokeServiceAccountAPITokenNotFound creates a RevokeServiceAccountAPITokenNotFound with default headers values
func NewRevokeServiceAccountAPITokenNotFound() *RevokeServiceAccountAPITokenNotFound {
	return &RevokeServiceAccountAPITokenNotFound{}
}

/*
RevokeServiceAccountAPITokenNotFound describes a response with status code 404, with default header values.

Not Found
*/
type RevokeServiceAccountAPITokenNotFound struct {
}

// IsSuccess returns true when this revoke service account Api token not found response has a 2xx status code
func (o *RevokeServiceAccountAPITokenNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this revoke service account Api token not found response has a 3xx status code
func (o *RevokeServiceAccountAPITokenNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this revoke service account Api token not found response has a 4xx status code
func (o *RevokeServiceAccountAPITokenNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this revoke service account Api token not found response has a 5xx status code
func (o *RevokeServiceAccountAPITokenNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this revoke service account Api token not found response a status code equal to that given
func (o *RevokeServiceAccountAPITokenNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the revoke service account Api token not found response
func (o *RevokeServiceAccountAPITokenNotFound) Code() int {
	return 404
}

func (o *RevokeServiceAccountAPITokenNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/backend/v1/authorization/serviceAccounts/api-token][%d] revokeServiceAccountApiTokenNotFound ", 404)
}

func (o *RevokeServiceAccountAPITokenNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/backend/v1/authorization/serviceAccounts/api-token][%d] revokeServiceAccountApiTokenNotFound ", 404)
}

func (o *RevokeServiceAccountAPITokenNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRevokeServiceAccountAPITokenRequestTimeout creates a RevokeServiceAccountAPITokenRequestTimeout with default headers values
func NewRevokeServiceAccountAPITokenRequestTimeout() *RevokeServiceAccountAPITokenRequestTimeout {
	return &RevokeServiceAccountAPITokenRequestTimeout{}
}

/*
RevokeServiceAccountAPITokenRequestTimeout describes a response with status code 408, with default header values.

Request Timeout
*/
type RevokeServiceAccountAPITokenRequestTimeout struct {
}

// IsSuccess returns true when this revoke service account Api token request timeout response has a 2xx status code
func (o *RevokeServiceAccountAPITokenRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this revoke service account Api token request timeout response has a 3xx status code
func (o *RevokeServiceAccountAPITokenRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this revoke service account Api token request timeout response has a 4xx status code
func (o *RevokeServiceAccountAPITokenRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this revoke service account Api token request timeout response has a 5xx status code
func (o *RevokeServiceAccountAPITokenRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this revoke service account Api token request timeout response a status code equal to that given
func (o *RevokeServiceAccountAPITokenRequestTimeout) IsCode(code int) bool {
	return code == 408
}

// Code gets the status code for the revoke service account Api token request timeout response
func (o *RevokeServiceAccountAPITokenRequestTimeout) Code() int {
	return 408
}

func (o *RevokeServiceAccountAPITokenRequestTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/backend/v1/authorization/serviceAccounts/api-token][%d] revokeServiceAccountApiTokenRequestTimeout ", 408)
}

func (o *RevokeServiceAccountAPITokenRequestTimeout) String() string {
	return fmt.Sprintf("[DELETE /api/backend/v1/authorization/serviceAccounts/api-token][%d] revokeServiceAccountApiTokenRequestTimeout ", 408)
}

func (o *RevokeServiceAccountAPITokenRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRevokeServiceAccountAPITokenUnprocessableEntity creates a RevokeServiceAccountAPITokenUnprocessableEntity with default headers values
func NewRevokeServiceAccountAPITokenUnprocessableEntity() *RevokeServiceAccountAPITokenUnprocessableEntity {
	return &RevokeServiceAccountAPITokenUnprocessableEntity{}
}

/*
RevokeServiceAccountAPITokenUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type RevokeServiceAccountAPITokenUnprocessableEntity struct {
}

// IsSuccess returns true when this revoke service account Api token unprocessable entity response has a 2xx status code
func (o *RevokeServiceAccountAPITokenUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this revoke service account Api token unprocessable entity response has a 3xx status code
func (o *RevokeServiceAccountAPITokenUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this revoke service account Api token unprocessable entity response has a 4xx status code
func (o *RevokeServiceAccountAPITokenUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this revoke service account Api token unprocessable entity response has a 5xx status code
func (o *RevokeServiceAccountAPITokenUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this revoke service account Api token unprocessable entity response a status code equal to that given
func (o *RevokeServiceAccountAPITokenUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the revoke service account Api token unprocessable entity response
func (o *RevokeServiceAccountAPITokenUnprocessableEntity) Code() int {
	return 422
}

func (o *RevokeServiceAccountAPITokenUnprocessableEntity) Error() string {
	return fmt.Sprintf("[DELETE /api/backend/v1/authorization/serviceAccounts/api-token][%d] revokeServiceAccountApiTokenUnprocessableEntity ", 422)
}

func (o *RevokeServiceAccountAPITokenUnprocessableEntity) String() string {
	return fmt.Sprintf("[DELETE /api/backend/v1/authorization/serviceAccounts/api-token][%d] revokeServiceAccountApiTokenUnprocessableEntity ", 422)
}

func (o *RevokeServiceAccountAPITokenUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
