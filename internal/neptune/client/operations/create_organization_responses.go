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

// CreateOrganizationReader is a Reader for the CreateOrganization structure.
type CreateOrganizationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateOrganizationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateOrganizationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateOrganizationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateOrganizationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateOrganizationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateOrganizationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewCreateOrganizationRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewCreateOrganizationUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /api/backend/v1/organizations2] createOrganization", response, response.Code())
	}
}

// NewCreateOrganizationOK creates a CreateOrganizationOK with default headers values
func NewCreateOrganizationOK() *CreateOrganizationOK {
	return &CreateOrganizationOK{}
}

/*
CreateOrganizationOK describes a response with status code 200, with default header values.

OK
*/
type CreateOrganizationOK struct {
	Payload *models.OrganizationDTO
}

// IsSuccess returns true when this create organization o k response has a 2xx status code
func (o *CreateOrganizationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create organization o k response has a 3xx status code
func (o *CreateOrganizationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create organization o k response has a 4xx status code
func (o *CreateOrganizationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create organization o k response has a 5xx status code
func (o *CreateOrganizationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create organization o k response a status code equal to that given
func (o *CreateOrganizationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create organization o k response
func (o *CreateOrganizationOK) Code() int {
	return 200
}

func (o *CreateOrganizationOK) Error() string {
	return fmt.Sprintf("[POST /api/backend/v1/organizations2][%d] createOrganizationOK  %+v", 200, o.Payload)
}

func (o *CreateOrganizationOK) String() string {
	return fmt.Sprintf("[POST /api/backend/v1/organizations2][%d] createOrganizationOK  %+v", 200, o.Payload)
}

func (o *CreateOrganizationOK) GetPayload() *models.OrganizationDTO {
	return o.Payload
}

func (o *CreateOrganizationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OrganizationDTO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateOrganizationBadRequest creates a CreateOrganizationBadRequest with default headers values
func NewCreateOrganizationBadRequest() *CreateOrganizationBadRequest {
	return &CreateOrganizationBadRequest{}
}

/*
CreateOrganizationBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CreateOrganizationBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this create organization bad request response has a 2xx status code
func (o *CreateOrganizationBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create organization bad request response has a 3xx status code
func (o *CreateOrganizationBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create organization bad request response has a 4xx status code
func (o *CreateOrganizationBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create organization bad request response has a 5xx status code
func (o *CreateOrganizationBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create organization bad request response a status code equal to that given
func (o *CreateOrganizationBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create organization bad request response
func (o *CreateOrganizationBadRequest) Code() int {
	return 400
}

func (o *CreateOrganizationBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/backend/v1/organizations2][%d] createOrganizationBadRequest  %+v", 400, o.Payload)
}

func (o *CreateOrganizationBadRequest) String() string {
	return fmt.Sprintf("[POST /api/backend/v1/organizations2][%d] createOrganizationBadRequest  %+v", 400, o.Payload)
}

func (o *CreateOrganizationBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateOrganizationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateOrganizationUnauthorized creates a CreateOrganizationUnauthorized with default headers values
func NewCreateOrganizationUnauthorized() *CreateOrganizationUnauthorized {
	return &CreateOrganizationUnauthorized{}
}

/*
CreateOrganizationUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CreateOrganizationUnauthorized struct {
}

// IsSuccess returns true when this create organization unauthorized response has a 2xx status code
func (o *CreateOrganizationUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create organization unauthorized response has a 3xx status code
func (o *CreateOrganizationUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create organization unauthorized response has a 4xx status code
func (o *CreateOrganizationUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create organization unauthorized response has a 5xx status code
func (o *CreateOrganizationUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create organization unauthorized response a status code equal to that given
func (o *CreateOrganizationUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the create organization unauthorized response
func (o *CreateOrganizationUnauthorized) Code() int {
	return 401
}

func (o *CreateOrganizationUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/backend/v1/organizations2][%d] createOrganizationUnauthorized ", 401)
}

func (o *CreateOrganizationUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/backend/v1/organizations2][%d] createOrganizationUnauthorized ", 401)
}

func (o *CreateOrganizationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateOrganizationForbidden creates a CreateOrganizationForbidden with default headers values
func NewCreateOrganizationForbidden() *CreateOrganizationForbidden {
	return &CreateOrganizationForbidden{}
}

/*
CreateOrganizationForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CreateOrganizationForbidden struct {
}

// IsSuccess returns true when this create organization forbidden response has a 2xx status code
func (o *CreateOrganizationForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create organization forbidden response has a 3xx status code
func (o *CreateOrganizationForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create organization forbidden response has a 4xx status code
func (o *CreateOrganizationForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create organization forbidden response has a 5xx status code
func (o *CreateOrganizationForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create organization forbidden response a status code equal to that given
func (o *CreateOrganizationForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the create organization forbidden response
func (o *CreateOrganizationForbidden) Code() int {
	return 403
}

func (o *CreateOrganizationForbidden) Error() string {
	return fmt.Sprintf("[POST /api/backend/v1/organizations2][%d] createOrganizationForbidden ", 403)
}

func (o *CreateOrganizationForbidden) String() string {
	return fmt.Sprintf("[POST /api/backend/v1/organizations2][%d] createOrganizationForbidden ", 403)
}

func (o *CreateOrganizationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateOrganizationNotFound creates a CreateOrganizationNotFound with default headers values
func NewCreateOrganizationNotFound() *CreateOrganizationNotFound {
	return &CreateOrganizationNotFound{}
}

/*
CreateOrganizationNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CreateOrganizationNotFound struct {
}

// IsSuccess returns true when this create organization not found response has a 2xx status code
func (o *CreateOrganizationNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create organization not found response has a 3xx status code
func (o *CreateOrganizationNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create organization not found response has a 4xx status code
func (o *CreateOrganizationNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this create organization not found response has a 5xx status code
func (o *CreateOrganizationNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this create organization not found response a status code equal to that given
func (o *CreateOrganizationNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the create organization not found response
func (o *CreateOrganizationNotFound) Code() int {
	return 404
}

func (o *CreateOrganizationNotFound) Error() string {
	return fmt.Sprintf("[POST /api/backend/v1/organizations2][%d] createOrganizationNotFound ", 404)
}

func (o *CreateOrganizationNotFound) String() string {
	return fmt.Sprintf("[POST /api/backend/v1/organizations2][%d] createOrganizationNotFound ", 404)
}

func (o *CreateOrganizationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateOrganizationRequestTimeout creates a CreateOrganizationRequestTimeout with default headers values
func NewCreateOrganizationRequestTimeout() *CreateOrganizationRequestTimeout {
	return &CreateOrganizationRequestTimeout{}
}

/*
CreateOrganizationRequestTimeout describes a response with status code 408, with default header values.

Request Timeout
*/
type CreateOrganizationRequestTimeout struct {
}

// IsSuccess returns true when this create organization request timeout response has a 2xx status code
func (o *CreateOrganizationRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create organization request timeout response has a 3xx status code
func (o *CreateOrganizationRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create organization request timeout response has a 4xx status code
func (o *CreateOrganizationRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this create organization request timeout response has a 5xx status code
func (o *CreateOrganizationRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this create organization request timeout response a status code equal to that given
func (o *CreateOrganizationRequestTimeout) IsCode(code int) bool {
	return code == 408
}

// Code gets the status code for the create organization request timeout response
func (o *CreateOrganizationRequestTimeout) Code() int {
	return 408
}

func (o *CreateOrganizationRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/backend/v1/organizations2][%d] createOrganizationRequestTimeout ", 408)
}

func (o *CreateOrganizationRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/backend/v1/organizations2][%d] createOrganizationRequestTimeout ", 408)
}

func (o *CreateOrganizationRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateOrganizationUnprocessableEntity creates a CreateOrganizationUnprocessableEntity with default headers values
func NewCreateOrganizationUnprocessableEntity() *CreateOrganizationUnprocessableEntity {
	return &CreateOrganizationUnprocessableEntity{}
}

/*
CreateOrganizationUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type CreateOrganizationUnprocessableEntity struct {
}

// IsSuccess returns true when this create organization unprocessable entity response has a 2xx status code
func (o *CreateOrganizationUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create organization unprocessable entity response has a 3xx status code
func (o *CreateOrganizationUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create organization unprocessable entity response has a 4xx status code
func (o *CreateOrganizationUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this create organization unprocessable entity response has a 5xx status code
func (o *CreateOrganizationUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this create organization unprocessable entity response a status code equal to that given
func (o *CreateOrganizationUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the create organization unprocessable entity response
func (o *CreateOrganizationUnprocessableEntity) Code() int {
	return 422
}

func (o *CreateOrganizationUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /api/backend/v1/organizations2][%d] createOrganizationUnprocessableEntity ", 422)
}

func (o *CreateOrganizationUnprocessableEntity) String() string {
	return fmt.Sprintf("[POST /api/backend/v1/organizations2][%d] createOrganizationUnprocessableEntity ", 422)
}

func (o *CreateOrganizationUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
