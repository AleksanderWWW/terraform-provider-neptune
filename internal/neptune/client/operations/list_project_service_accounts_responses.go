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

// ListProjectServiceAccountsReader is a Reader for the ListProjectServiceAccounts structure.
type ListProjectServiceAccountsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListProjectServiceAccountsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListProjectServiceAccountsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListProjectServiceAccountsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewListProjectServiceAccountsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListProjectServiceAccountsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListProjectServiceAccountsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewListProjectServiceAccountsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewListProjectServiceAccountsUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /api/backend/v1/projects/service_accounts] listProjectServiceAccounts", response, response.Code())
	}
}

// NewListProjectServiceAccountsOK creates a ListProjectServiceAccountsOK with default headers values
func NewListProjectServiceAccountsOK() *ListProjectServiceAccountsOK {
	return &ListProjectServiceAccountsOK{}
}

/*
ListProjectServiceAccountsOK describes a response with status code 200, with default header values.

OK
*/
type ListProjectServiceAccountsOK struct {
	Payload []*models.ServiceAccountWithRoleDTO
}

// IsSuccess returns true when this list project service accounts o k response has a 2xx status code
func (o *ListProjectServiceAccountsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list project service accounts o k response has a 3xx status code
func (o *ListProjectServiceAccountsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list project service accounts o k response has a 4xx status code
func (o *ListProjectServiceAccountsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list project service accounts o k response has a 5xx status code
func (o *ListProjectServiceAccountsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list project service accounts o k response a status code equal to that given
func (o *ListProjectServiceAccountsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list project service accounts o k response
func (o *ListProjectServiceAccountsOK) Code() int {
	return 200
}

func (o *ListProjectServiceAccountsOK) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/projects/service_accounts][%d] listProjectServiceAccountsOK  %+v", 200, o.Payload)
}

func (o *ListProjectServiceAccountsOK) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/projects/service_accounts][%d] listProjectServiceAccountsOK  %+v", 200, o.Payload)
}

func (o *ListProjectServiceAccountsOK) GetPayload() []*models.ServiceAccountWithRoleDTO {
	return o.Payload
}

func (o *ListProjectServiceAccountsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListProjectServiceAccountsBadRequest creates a ListProjectServiceAccountsBadRequest with default headers values
func NewListProjectServiceAccountsBadRequest() *ListProjectServiceAccountsBadRequest {
	return &ListProjectServiceAccountsBadRequest{}
}

/*
ListProjectServiceAccountsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ListProjectServiceAccountsBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this list project service accounts bad request response has a 2xx status code
func (o *ListProjectServiceAccountsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list project service accounts bad request response has a 3xx status code
func (o *ListProjectServiceAccountsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list project service accounts bad request response has a 4xx status code
func (o *ListProjectServiceAccountsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this list project service accounts bad request response has a 5xx status code
func (o *ListProjectServiceAccountsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this list project service accounts bad request response a status code equal to that given
func (o *ListProjectServiceAccountsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the list project service accounts bad request response
func (o *ListProjectServiceAccountsBadRequest) Code() int {
	return 400
}

func (o *ListProjectServiceAccountsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/projects/service_accounts][%d] listProjectServiceAccountsBadRequest  %+v", 400, o.Payload)
}

func (o *ListProjectServiceAccountsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/projects/service_accounts][%d] listProjectServiceAccountsBadRequest  %+v", 400, o.Payload)
}

func (o *ListProjectServiceAccountsBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListProjectServiceAccountsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListProjectServiceAccountsUnauthorized creates a ListProjectServiceAccountsUnauthorized with default headers values
func NewListProjectServiceAccountsUnauthorized() *ListProjectServiceAccountsUnauthorized {
	return &ListProjectServiceAccountsUnauthorized{}
}

/*
ListProjectServiceAccountsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ListProjectServiceAccountsUnauthorized struct {
}

// IsSuccess returns true when this list project service accounts unauthorized response has a 2xx status code
func (o *ListProjectServiceAccountsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list project service accounts unauthorized response has a 3xx status code
func (o *ListProjectServiceAccountsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list project service accounts unauthorized response has a 4xx status code
func (o *ListProjectServiceAccountsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this list project service accounts unauthorized response has a 5xx status code
func (o *ListProjectServiceAccountsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this list project service accounts unauthorized response a status code equal to that given
func (o *ListProjectServiceAccountsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the list project service accounts unauthorized response
func (o *ListProjectServiceAccountsUnauthorized) Code() int {
	return 401
}

func (o *ListProjectServiceAccountsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/projects/service_accounts][%d] listProjectServiceAccountsUnauthorized ", 401)
}

func (o *ListProjectServiceAccountsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/projects/service_accounts][%d] listProjectServiceAccountsUnauthorized ", 401)
}

func (o *ListProjectServiceAccountsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListProjectServiceAccountsForbidden creates a ListProjectServiceAccountsForbidden with default headers values
func NewListProjectServiceAccountsForbidden() *ListProjectServiceAccountsForbidden {
	return &ListProjectServiceAccountsForbidden{}
}

/*
ListProjectServiceAccountsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ListProjectServiceAccountsForbidden struct {
}

// IsSuccess returns true when this list project service accounts forbidden response has a 2xx status code
func (o *ListProjectServiceAccountsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list project service accounts forbidden response has a 3xx status code
func (o *ListProjectServiceAccountsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list project service accounts forbidden response has a 4xx status code
func (o *ListProjectServiceAccountsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list project service accounts forbidden response has a 5xx status code
func (o *ListProjectServiceAccountsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list project service accounts forbidden response a status code equal to that given
func (o *ListProjectServiceAccountsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the list project service accounts forbidden response
func (o *ListProjectServiceAccountsForbidden) Code() int {
	return 403
}

func (o *ListProjectServiceAccountsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/projects/service_accounts][%d] listProjectServiceAccountsForbidden ", 403)
}

func (o *ListProjectServiceAccountsForbidden) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/projects/service_accounts][%d] listProjectServiceAccountsForbidden ", 403)
}

func (o *ListProjectServiceAccountsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListProjectServiceAccountsNotFound creates a ListProjectServiceAccountsNotFound with default headers values
func NewListProjectServiceAccountsNotFound() *ListProjectServiceAccountsNotFound {
	return &ListProjectServiceAccountsNotFound{}
}

/*
ListProjectServiceAccountsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ListProjectServiceAccountsNotFound struct {
}

// IsSuccess returns true when this list project service accounts not found response has a 2xx status code
func (o *ListProjectServiceAccountsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list project service accounts not found response has a 3xx status code
func (o *ListProjectServiceAccountsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list project service accounts not found response has a 4xx status code
func (o *ListProjectServiceAccountsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this list project service accounts not found response has a 5xx status code
func (o *ListProjectServiceAccountsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this list project service accounts not found response a status code equal to that given
func (o *ListProjectServiceAccountsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the list project service accounts not found response
func (o *ListProjectServiceAccountsNotFound) Code() int {
	return 404
}

func (o *ListProjectServiceAccountsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/projects/service_accounts][%d] listProjectServiceAccountsNotFound ", 404)
}

func (o *ListProjectServiceAccountsNotFound) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/projects/service_accounts][%d] listProjectServiceAccountsNotFound ", 404)
}

func (o *ListProjectServiceAccountsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListProjectServiceAccountsRequestTimeout creates a ListProjectServiceAccountsRequestTimeout with default headers values
func NewListProjectServiceAccountsRequestTimeout() *ListProjectServiceAccountsRequestTimeout {
	return &ListProjectServiceAccountsRequestTimeout{}
}

/*
ListProjectServiceAccountsRequestTimeout describes a response with status code 408, with default header values.

Request Timeout
*/
type ListProjectServiceAccountsRequestTimeout struct {
}

// IsSuccess returns true when this list project service accounts request timeout response has a 2xx status code
func (o *ListProjectServiceAccountsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list project service accounts request timeout response has a 3xx status code
func (o *ListProjectServiceAccountsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list project service accounts request timeout response has a 4xx status code
func (o *ListProjectServiceAccountsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this list project service accounts request timeout response has a 5xx status code
func (o *ListProjectServiceAccountsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this list project service accounts request timeout response a status code equal to that given
func (o *ListProjectServiceAccountsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

// Code gets the status code for the list project service accounts request timeout response
func (o *ListProjectServiceAccountsRequestTimeout) Code() int {
	return 408
}

func (o *ListProjectServiceAccountsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/projects/service_accounts][%d] listProjectServiceAccountsRequestTimeout ", 408)
}

func (o *ListProjectServiceAccountsRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/projects/service_accounts][%d] listProjectServiceAccountsRequestTimeout ", 408)
}

func (o *ListProjectServiceAccountsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListProjectServiceAccountsUnprocessableEntity creates a ListProjectServiceAccountsUnprocessableEntity with default headers values
func NewListProjectServiceAccountsUnprocessableEntity() *ListProjectServiceAccountsUnprocessableEntity {
	return &ListProjectServiceAccountsUnprocessableEntity{}
}

/*
ListProjectServiceAccountsUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type ListProjectServiceAccountsUnprocessableEntity struct {
}

// IsSuccess returns true when this list project service accounts unprocessable entity response has a 2xx status code
func (o *ListProjectServiceAccountsUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list project service accounts unprocessable entity response has a 3xx status code
func (o *ListProjectServiceAccountsUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list project service accounts unprocessable entity response has a 4xx status code
func (o *ListProjectServiceAccountsUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this list project service accounts unprocessable entity response has a 5xx status code
func (o *ListProjectServiceAccountsUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this list project service accounts unprocessable entity response a status code equal to that given
func (o *ListProjectServiceAccountsUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the list project service accounts unprocessable entity response
func (o *ListProjectServiceAccountsUnprocessableEntity) Code() int {
	return 422
}

func (o *ListProjectServiceAccountsUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/projects/service_accounts][%d] listProjectServiceAccountsUnprocessableEntity ", 422)
}

func (o *ListProjectServiceAccountsUnprocessableEntity) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/projects/service_accounts][%d] listProjectServiceAccountsUnprocessableEntity ", 422)
}

func (o *ListProjectServiceAccountsUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
