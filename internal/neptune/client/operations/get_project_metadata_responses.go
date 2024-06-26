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

// GetProjectMetadataReader is a Reader for the GetProjectMetadata structure.
type GetProjectMetadataReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProjectMetadataReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetProjectMetadataOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetProjectMetadataBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetProjectMetadataUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetProjectMetadataForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetProjectMetadataNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetProjectMetadataRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGetProjectMetadataUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /api/backend/v1/projects/metadata] getProjectMetadata", response, response.Code())
	}
}

// NewGetProjectMetadataOK creates a GetProjectMetadataOK with default headers values
func NewGetProjectMetadataOK() *GetProjectMetadataOK {
	return &GetProjectMetadataOK{}
}

/*
GetProjectMetadataOK describes a response with status code 200, with default header values.

OK
*/
type GetProjectMetadataOK struct {
	Payload *models.ProjectMetadataDTO
}

// IsSuccess returns true when this get project metadata o k response has a 2xx status code
func (o *GetProjectMetadataOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get project metadata o k response has a 3xx status code
func (o *GetProjectMetadataOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get project metadata o k response has a 4xx status code
func (o *GetProjectMetadataOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get project metadata o k response has a 5xx status code
func (o *GetProjectMetadataOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get project metadata o k response a status code equal to that given
func (o *GetProjectMetadataOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get project metadata o k response
func (o *GetProjectMetadataOK) Code() int {
	return 200
}

func (o *GetProjectMetadataOK) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/projects/metadata][%d] getProjectMetadataOK  %+v", 200, o.Payload)
}

func (o *GetProjectMetadataOK) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/projects/metadata][%d] getProjectMetadataOK  %+v", 200, o.Payload)
}

func (o *GetProjectMetadataOK) GetPayload() *models.ProjectMetadataDTO {
	return o.Payload
}

func (o *GetProjectMetadataOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProjectMetadataDTO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProjectMetadataBadRequest creates a GetProjectMetadataBadRequest with default headers values
func NewGetProjectMetadataBadRequest() *GetProjectMetadataBadRequest {
	return &GetProjectMetadataBadRequest{}
}

/*
GetProjectMetadataBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetProjectMetadataBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this get project metadata bad request response has a 2xx status code
func (o *GetProjectMetadataBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get project metadata bad request response has a 3xx status code
func (o *GetProjectMetadataBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get project metadata bad request response has a 4xx status code
func (o *GetProjectMetadataBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get project metadata bad request response has a 5xx status code
func (o *GetProjectMetadataBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get project metadata bad request response a status code equal to that given
func (o *GetProjectMetadataBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get project metadata bad request response
func (o *GetProjectMetadataBadRequest) Code() int {
	return 400
}

func (o *GetProjectMetadataBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/projects/metadata][%d] getProjectMetadataBadRequest  %+v", 400, o.Payload)
}

func (o *GetProjectMetadataBadRequest) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/projects/metadata][%d] getProjectMetadataBadRequest  %+v", 400, o.Payload)
}

func (o *GetProjectMetadataBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetProjectMetadataBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProjectMetadataUnauthorized creates a GetProjectMetadataUnauthorized with default headers values
func NewGetProjectMetadataUnauthorized() *GetProjectMetadataUnauthorized {
	return &GetProjectMetadataUnauthorized{}
}

/*
GetProjectMetadataUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetProjectMetadataUnauthorized struct {
}

// IsSuccess returns true when this get project metadata unauthorized response has a 2xx status code
func (o *GetProjectMetadataUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get project metadata unauthorized response has a 3xx status code
func (o *GetProjectMetadataUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get project metadata unauthorized response has a 4xx status code
func (o *GetProjectMetadataUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get project metadata unauthorized response has a 5xx status code
func (o *GetProjectMetadataUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get project metadata unauthorized response a status code equal to that given
func (o *GetProjectMetadataUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get project metadata unauthorized response
func (o *GetProjectMetadataUnauthorized) Code() int {
	return 401
}

func (o *GetProjectMetadataUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/projects/metadata][%d] getProjectMetadataUnauthorized ", 401)
}

func (o *GetProjectMetadataUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/projects/metadata][%d] getProjectMetadataUnauthorized ", 401)
}

func (o *GetProjectMetadataUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProjectMetadataForbidden creates a GetProjectMetadataForbidden with default headers values
func NewGetProjectMetadataForbidden() *GetProjectMetadataForbidden {
	return &GetProjectMetadataForbidden{}
}

/*
GetProjectMetadataForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetProjectMetadataForbidden struct {
}

// IsSuccess returns true when this get project metadata forbidden response has a 2xx status code
func (o *GetProjectMetadataForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get project metadata forbidden response has a 3xx status code
func (o *GetProjectMetadataForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get project metadata forbidden response has a 4xx status code
func (o *GetProjectMetadataForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get project metadata forbidden response has a 5xx status code
func (o *GetProjectMetadataForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get project metadata forbidden response a status code equal to that given
func (o *GetProjectMetadataForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get project metadata forbidden response
func (o *GetProjectMetadataForbidden) Code() int {
	return 403
}

func (o *GetProjectMetadataForbidden) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/projects/metadata][%d] getProjectMetadataForbidden ", 403)
}

func (o *GetProjectMetadataForbidden) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/projects/metadata][%d] getProjectMetadataForbidden ", 403)
}

func (o *GetProjectMetadataForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProjectMetadataNotFound creates a GetProjectMetadataNotFound with default headers values
func NewGetProjectMetadataNotFound() *GetProjectMetadataNotFound {
	return &GetProjectMetadataNotFound{}
}

/*
GetProjectMetadataNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetProjectMetadataNotFound struct {
}

// IsSuccess returns true when this get project metadata not found response has a 2xx status code
func (o *GetProjectMetadataNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get project metadata not found response has a 3xx status code
func (o *GetProjectMetadataNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get project metadata not found response has a 4xx status code
func (o *GetProjectMetadataNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get project metadata not found response has a 5xx status code
func (o *GetProjectMetadataNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get project metadata not found response a status code equal to that given
func (o *GetProjectMetadataNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get project metadata not found response
func (o *GetProjectMetadataNotFound) Code() int {
	return 404
}

func (o *GetProjectMetadataNotFound) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/projects/metadata][%d] getProjectMetadataNotFound ", 404)
}

func (o *GetProjectMetadataNotFound) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/projects/metadata][%d] getProjectMetadataNotFound ", 404)
}

func (o *GetProjectMetadataNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProjectMetadataRequestTimeout creates a GetProjectMetadataRequestTimeout with default headers values
func NewGetProjectMetadataRequestTimeout() *GetProjectMetadataRequestTimeout {
	return &GetProjectMetadataRequestTimeout{}
}

/*
GetProjectMetadataRequestTimeout describes a response with status code 408, with default header values.

Request Timeout
*/
type GetProjectMetadataRequestTimeout struct {
}

// IsSuccess returns true when this get project metadata request timeout response has a 2xx status code
func (o *GetProjectMetadataRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get project metadata request timeout response has a 3xx status code
func (o *GetProjectMetadataRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get project metadata request timeout response has a 4xx status code
func (o *GetProjectMetadataRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get project metadata request timeout response has a 5xx status code
func (o *GetProjectMetadataRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get project metadata request timeout response a status code equal to that given
func (o *GetProjectMetadataRequestTimeout) IsCode(code int) bool {
	return code == 408
}

// Code gets the status code for the get project metadata request timeout response
func (o *GetProjectMetadataRequestTimeout) Code() int {
	return 408
}

func (o *GetProjectMetadataRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/projects/metadata][%d] getProjectMetadataRequestTimeout ", 408)
}

func (o *GetProjectMetadataRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/projects/metadata][%d] getProjectMetadataRequestTimeout ", 408)
}

func (o *GetProjectMetadataRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProjectMetadataUnprocessableEntity creates a GetProjectMetadataUnprocessableEntity with default headers values
func NewGetProjectMetadataUnprocessableEntity() *GetProjectMetadataUnprocessableEntity {
	return &GetProjectMetadataUnprocessableEntity{}
}

/*
GetProjectMetadataUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type GetProjectMetadataUnprocessableEntity struct {
}

// IsSuccess returns true when this get project metadata unprocessable entity response has a 2xx status code
func (o *GetProjectMetadataUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get project metadata unprocessable entity response has a 3xx status code
func (o *GetProjectMetadataUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get project metadata unprocessable entity response has a 4xx status code
func (o *GetProjectMetadataUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this get project metadata unprocessable entity response has a 5xx status code
func (o *GetProjectMetadataUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this get project metadata unprocessable entity response a status code equal to that given
func (o *GetProjectMetadataUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the get project metadata unprocessable entity response
func (o *GetProjectMetadataUnprocessableEntity) Code() int {
	return 422
}

func (o *GetProjectMetadataUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /api/backend/v1/projects/metadata][%d] getProjectMetadataUnprocessableEntity ", 422)
}

func (o *GetProjectMetadataUnprocessableEntity) String() string {
	return fmt.Sprintf("[GET /api/backend/v1/projects/metadata][%d] getProjectMetadataUnprocessableEntity ", 422)
}

func (o *GetProjectMetadataUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
