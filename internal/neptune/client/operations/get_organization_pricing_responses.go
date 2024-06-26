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

// GetOrganizationPricingReader is a Reader for the GetOrganizationPricing structure.
type GetOrganizationPricingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrganizationPricingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrganizationPricingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetOrganizationPricingBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetOrganizationPricingUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetOrganizationPricingForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetOrganizationPricingNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetOrganizationPricingRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGetOrganizationPricingUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /api/backend/v1/payments/{organizationIdentifier}] getOrganizationPricing", response, response.Code())
	}
}

// NewGetOrganizationPricingOK creates a GetOrganizationPricingOK with default headers values
func NewGetOrganizationPricingOK() *GetOrganizationPricingOK {
	return &GetOrganizationPricingOK{}
}

/*
GetOrganizationPricingOK describes a response with status code 200, with default header values.

OK
*/
type GetOrganizationPricingOK struct {
	Payload *models.OrganizationPricingDTO
}

// IsSuccess returns true when this get organization pricing o k response has a 2xx status code
func (o *GetOrganizationPricingOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get organization pricing o k response has a 3xx status code
func (o *GetOrganizationPricingOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organization pricing o k response has a 4xx status code
func (o *GetOrganizationPricingOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get organization pricing o k response has a 5xx status code
func (o *GetOrganizationPricingOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get organization pricing o k response a status code equal to that given
func (o *GetOrganizationPricingOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get organization pricing o k response
func (o *GetOrganizationPricingOK) Code() int {
	return 200
}

func (o *GetOrganizationPricingOK) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/payments/{organizationIdentifier}][%d] getOrganizationPricingOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationPricingOK) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/payments/{organizationIdentifier}][%d] getOrganizationPricingOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationPricingOK) GetPayload() *models.OrganizationPricingDTO {
	return o.Payload
}

func (o *GetOrganizationPricingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OrganizationPricingDTO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationPricingBadRequest creates a GetOrganizationPricingBadRequest with default headers values
func NewGetOrganizationPricingBadRequest() *GetOrganizationPricingBadRequest {
	return &GetOrganizationPricingBadRequest{}
}

/*
GetOrganizationPricingBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetOrganizationPricingBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this get organization pricing bad request response has a 2xx status code
func (o *GetOrganizationPricingBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get organization pricing bad request response has a 3xx status code
func (o *GetOrganizationPricingBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organization pricing bad request response has a 4xx status code
func (o *GetOrganizationPricingBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get organization pricing bad request response has a 5xx status code
func (o *GetOrganizationPricingBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get organization pricing bad request response a status code equal to that given
func (o *GetOrganizationPricingBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get organization pricing bad request response
func (o *GetOrganizationPricingBadRequest) Code() int {
	return 400
}

func (o *GetOrganizationPricingBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/payments/{organizationIdentifier}][%d] getOrganizationPricingBadRequest  %+v", 400, o.Payload)
}

func (o *GetOrganizationPricingBadRequest) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/payments/{organizationIdentifier}][%d] getOrganizationPricingBadRequest  %+v", 400, o.Payload)
}

func (o *GetOrganizationPricingBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetOrganizationPricingBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationPricingUnauthorized creates a GetOrganizationPricingUnauthorized with default headers values
func NewGetOrganizationPricingUnauthorized() *GetOrganizationPricingUnauthorized {
	return &GetOrganizationPricingUnauthorized{}
}

/*
GetOrganizationPricingUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetOrganizationPricingUnauthorized struct {
}

// IsSuccess returns true when this get organization pricing unauthorized response has a 2xx status code
func (o *GetOrganizationPricingUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get organization pricing unauthorized response has a 3xx status code
func (o *GetOrganizationPricingUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organization pricing unauthorized response has a 4xx status code
func (o *GetOrganizationPricingUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get organization pricing unauthorized response has a 5xx status code
func (o *GetOrganizationPricingUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get organization pricing unauthorized response a status code equal to that given
func (o *GetOrganizationPricingUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get organization pricing unauthorized response
func (o *GetOrganizationPricingUnauthorized) Code() int {
	return 401
}

func (o *GetOrganizationPricingUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/payments/{organizationIdentifier}][%d] getOrganizationPricingUnauthorized ", 401)
}

func (o *GetOrganizationPricingUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/payments/{organizationIdentifier}][%d] getOrganizationPricingUnauthorized ", 401)
}

func (o *GetOrganizationPricingUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetOrganizationPricingForbidden creates a GetOrganizationPricingForbidden with default headers values
func NewGetOrganizationPricingForbidden() *GetOrganizationPricingForbidden {
	return &GetOrganizationPricingForbidden{}
}

/*
GetOrganizationPricingForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetOrganizationPricingForbidden struct {
}

// IsSuccess returns true when this get organization pricing forbidden response has a 2xx status code
func (o *GetOrganizationPricingForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get organization pricing forbidden response has a 3xx status code
func (o *GetOrganizationPricingForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organization pricing forbidden response has a 4xx status code
func (o *GetOrganizationPricingForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get organization pricing forbidden response has a 5xx status code
func (o *GetOrganizationPricingForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get organization pricing forbidden response a status code equal to that given
func (o *GetOrganizationPricingForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get organization pricing forbidden response
func (o *GetOrganizationPricingForbidden) Code() int {
	return 403
}

func (o *GetOrganizationPricingForbidden) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/payments/{organizationIdentifier}][%d] getOrganizationPricingForbidden ", 403)
}

func (o *GetOrganizationPricingForbidden) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/payments/{organizationIdentifier}][%d] getOrganizationPricingForbidden ", 403)
}

func (o *GetOrganizationPricingForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetOrganizationPricingNotFound creates a GetOrganizationPricingNotFound with default headers values
func NewGetOrganizationPricingNotFound() *GetOrganizationPricingNotFound {
	return &GetOrganizationPricingNotFound{}
}

/*
GetOrganizationPricingNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetOrganizationPricingNotFound struct {
}

// IsSuccess returns true when this get organization pricing not found response has a 2xx status code
func (o *GetOrganizationPricingNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get organization pricing not found response has a 3xx status code
func (o *GetOrganizationPricingNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organization pricing not found response has a 4xx status code
func (o *GetOrganizationPricingNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get organization pricing not found response has a 5xx status code
func (o *GetOrganizationPricingNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get organization pricing not found response a status code equal to that given
func (o *GetOrganizationPricingNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get organization pricing not found response
func (o *GetOrganizationPricingNotFound) Code() int {
	return 404
}

func (o *GetOrganizationPricingNotFound) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/payments/{organizationIdentifier}][%d] getOrganizationPricingNotFound ", 404)
}

func (o *GetOrganizationPricingNotFound) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/payments/{organizationIdentifier}][%d] getOrganizationPricingNotFound ", 404)
}

func (o *GetOrganizationPricingNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetOrganizationPricingRequestTimeout creates a GetOrganizationPricingRequestTimeout with default headers values
func NewGetOrganizationPricingRequestTimeout() *GetOrganizationPricingRequestTimeout {
	return &GetOrganizationPricingRequestTimeout{}
}

/*
GetOrganizationPricingRequestTimeout describes a response with status code 408, with default header values.

Request Timeout
*/
type GetOrganizationPricingRequestTimeout struct {
}

// IsSuccess returns true when this get organization pricing request timeout response has a 2xx status code
func (o *GetOrganizationPricingRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get organization pricing request timeout response has a 3xx status code
func (o *GetOrganizationPricingRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organization pricing request timeout response has a 4xx status code
func (o *GetOrganizationPricingRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get organization pricing request timeout response has a 5xx status code
func (o *GetOrganizationPricingRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get organization pricing request timeout response a status code equal to that given
func (o *GetOrganizationPricingRequestTimeout) IsCode(code int) bool {
	return code == 408
}

// Code gets the status code for the get organization pricing request timeout response
func (o *GetOrganizationPricingRequestTimeout) Code() int {
	return 408
}

func (o *GetOrganizationPricingRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/payments/{organizationIdentifier}][%d] getOrganizationPricingRequestTimeout ", 408)
}

func (o *GetOrganizationPricingRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/payments/{organizationIdentifier}][%d] getOrganizationPricingRequestTimeout ", 408)
}

func (o *GetOrganizationPricingRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetOrganizationPricingUnprocessableEntity creates a GetOrganizationPricingUnprocessableEntity with default headers values
func NewGetOrganizationPricingUnprocessableEntity() *GetOrganizationPricingUnprocessableEntity {
	return &GetOrganizationPricingUnprocessableEntity{}
}

/*
GetOrganizationPricingUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type GetOrganizationPricingUnprocessableEntity struct {
}

// IsSuccess returns true when this get organization pricing unprocessable entity response has a 2xx status code
func (o *GetOrganizationPricingUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get organization pricing unprocessable entity response has a 3xx status code
func (o *GetOrganizationPricingUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organization pricing unprocessable entity response has a 4xx status code
func (o *GetOrganizationPricingUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this get organization pricing unprocessable entity response has a 5xx status code
func (o *GetOrganizationPricingUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this get organization pricing unprocessable entity response a status code equal to that given
func (o *GetOrganizationPricingUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the get organization pricing unprocessable entity response
func (o *GetOrganizationPricingUnprocessableEntity) Code() int {
	return 422
}

func (o *GetOrganizationPricingUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/payments/{organizationIdentifier}][%d] getOrganizationPricingUnprocessableEntity ", 422)
}

func (o *GetOrganizationPricingUnprocessableEntity) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/payments/{organizationIdentifier}][%d] getOrganizationPricingUnprocessableEntity ", 422)
}

func (o *GetOrganizationPricingUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
