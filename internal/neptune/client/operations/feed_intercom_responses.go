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

// FeedIntercomReader is a Reader for the FeedIntercom structure.
type FeedIntercomReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FeedIntercomReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFeedIntercomOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewFeedIntercomBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewFeedIntercomUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewFeedIntercomForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewFeedIntercomNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewFeedIntercomRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewFeedIntercomUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /api/backend/v1/communicationData/submit] feedIntercom", response, response.Code())
	}
}

// NewFeedIntercomOK creates a FeedIntercomOK with default headers values
func NewFeedIntercomOK() *FeedIntercomOK {
	return &FeedIntercomOK{}
}

/*
FeedIntercomOK describes a response with status code 200, with default header values.

No response
*/
type FeedIntercomOK struct {
}

// IsSuccess returns true when this feed intercom o k response has a 2xx status code
func (o *FeedIntercomOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this feed intercom o k response has a 3xx status code
func (o *FeedIntercomOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this feed intercom o k response has a 4xx status code
func (o *FeedIntercomOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this feed intercom o k response has a 5xx status code
func (o *FeedIntercomOK) IsServerError() bool {
	return false
}

// IsCode returns true when this feed intercom o k response a status code equal to that given
func (o *FeedIntercomOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the feed intercom o k response
func (o *FeedIntercomOK) Code() int {
	return 200
}

func (o *FeedIntercomOK) Error() string {
	return fmt.Sprintf("[POST /api/backend/v1/communicationData/submit][%d] feedIntercomOK ", 200)
}

func (o *FeedIntercomOK) String() string {
	return fmt.Sprintf("[POST /api/backend/v1/communicationData/submit][%d] feedIntercomOK ", 200)
}

func (o *FeedIntercomOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFeedIntercomBadRequest creates a FeedIntercomBadRequest with default headers values
func NewFeedIntercomBadRequest() *FeedIntercomBadRequest {
	return &FeedIntercomBadRequest{}
}

/*
FeedIntercomBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type FeedIntercomBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this feed intercom bad request response has a 2xx status code
func (o *FeedIntercomBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this feed intercom bad request response has a 3xx status code
func (o *FeedIntercomBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this feed intercom bad request response has a 4xx status code
func (o *FeedIntercomBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this feed intercom bad request response has a 5xx status code
func (o *FeedIntercomBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this feed intercom bad request response a status code equal to that given
func (o *FeedIntercomBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the feed intercom bad request response
func (o *FeedIntercomBadRequest) Code() int {
	return 400
}

func (o *FeedIntercomBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/backend/v1/communicationData/submit][%d] feedIntercomBadRequest  %+v", 400, o.Payload)
}

func (o *FeedIntercomBadRequest) String() string {
	return fmt.Sprintf("[POST /api/backend/v1/communicationData/submit][%d] feedIntercomBadRequest  %+v", 400, o.Payload)
}

func (o *FeedIntercomBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *FeedIntercomBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFeedIntercomUnauthorized creates a FeedIntercomUnauthorized with default headers values
func NewFeedIntercomUnauthorized() *FeedIntercomUnauthorized {
	return &FeedIntercomUnauthorized{}
}

/*
FeedIntercomUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type FeedIntercomUnauthorized struct {
}

// IsSuccess returns true when this feed intercom unauthorized response has a 2xx status code
func (o *FeedIntercomUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this feed intercom unauthorized response has a 3xx status code
func (o *FeedIntercomUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this feed intercom unauthorized response has a 4xx status code
func (o *FeedIntercomUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this feed intercom unauthorized response has a 5xx status code
func (o *FeedIntercomUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this feed intercom unauthorized response a status code equal to that given
func (o *FeedIntercomUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the feed intercom unauthorized response
func (o *FeedIntercomUnauthorized) Code() int {
	return 401
}

func (o *FeedIntercomUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/backend/v1/communicationData/submit][%d] feedIntercomUnauthorized ", 401)
}

func (o *FeedIntercomUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/backend/v1/communicationData/submit][%d] feedIntercomUnauthorized ", 401)
}

func (o *FeedIntercomUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFeedIntercomForbidden creates a FeedIntercomForbidden with default headers values
func NewFeedIntercomForbidden() *FeedIntercomForbidden {
	return &FeedIntercomForbidden{}
}

/*
FeedIntercomForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type FeedIntercomForbidden struct {
}

// IsSuccess returns true when this feed intercom forbidden response has a 2xx status code
func (o *FeedIntercomForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this feed intercom forbidden response has a 3xx status code
func (o *FeedIntercomForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this feed intercom forbidden response has a 4xx status code
func (o *FeedIntercomForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this feed intercom forbidden response has a 5xx status code
func (o *FeedIntercomForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this feed intercom forbidden response a status code equal to that given
func (o *FeedIntercomForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the feed intercom forbidden response
func (o *FeedIntercomForbidden) Code() int {
	return 403
}

func (o *FeedIntercomForbidden) Error() string {
	return fmt.Sprintf("[POST /api/backend/v1/communicationData/submit][%d] feedIntercomForbidden ", 403)
}

func (o *FeedIntercomForbidden) String() string {
	return fmt.Sprintf("[POST /api/backend/v1/communicationData/submit][%d] feedIntercomForbidden ", 403)
}

func (o *FeedIntercomForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFeedIntercomNotFound creates a FeedIntercomNotFound with default headers values
func NewFeedIntercomNotFound() *FeedIntercomNotFound {
	return &FeedIntercomNotFound{}
}

/*
FeedIntercomNotFound describes a response with status code 404, with default header values.

Not Found
*/
type FeedIntercomNotFound struct {
}

// IsSuccess returns true when this feed intercom not found response has a 2xx status code
func (o *FeedIntercomNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this feed intercom not found response has a 3xx status code
func (o *FeedIntercomNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this feed intercom not found response has a 4xx status code
func (o *FeedIntercomNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this feed intercom not found response has a 5xx status code
func (o *FeedIntercomNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this feed intercom not found response a status code equal to that given
func (o *FeedIntercomNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the feed intercom not found response
func (o *FeedIntercomNotFound) Code() int {
	return 404
}

func (o *FeedIntercomNotFound) Error() string {
	return fmt.Sprintf("[POST /api/backend/v1/communicationData/submit][%d] feedIntercomNotFound ", 404)
}

func (o *FeedIntercomNotFound) String() string {
	return fmt.Sprintf("[POST /api/backend/v1/communicationData/submit][%d] feedIntercomNotFound ", 404)
}

func (o *FeedIntercomNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFeedIntercomRequestTimeout creates a FeedIntercomRequestTimeout with default headers values
func NewFeedIntercomRequestTimeout() *FeedIntercomRequestTimeout {
	return &FeedIntercomRequestTimeout{}
}

/*
FeedIntercomRequestTimeout describes a response with status code 408, with default header values.

Request Timeout
*/
type FeedIntercomRequestTimeout struct {
}

// IsSuccess returns true when this feed intercom request timeout response has a 2xx status code
func (o *FeedIntercomRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this feed intercom request timeout response has a 3xx status code
func (o *FeedIntercomRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this feed intercom request timeout response has a 4xx status code
func (o *FeedIntercomRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this feed intercom request timeout response has a 5xx status code
func (o *FeedIntercomRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this feed intercom request timeout response a status code equal to that given
func (o *FeedIntercomRequestTimeout) IsCode(code int) bool {
	return code == 408
}

// Code gets the status code for the feed intercom request timeout response
func (o *FeedIntercomRequestTimeout) Code() int {
	return 408
}

func (o *FeedIntercomRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/backend/v1/communicationData/submit][%d] feedIntercomRequestTimeout ", 408)
}

func (o *FeedIntercomRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/backend/v1/communicationData/submit][%d] feedIntercomRequestTimeout ", 408)
}

func (o *FeedIntercomRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFeedIntercomUnprocessableEntity creates a FeedIntercomUnprocessableEntity with default headers values
func NewFeedIntercomUnprocessableEntity() *FeedIntercomUnprocessableEntity {
	return &FeedIntercomUnprocessableEntity{}
}

/*
FeedIntercomUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type FeedIntercomUnprocessableEntity struct {
}

// IsSuccess returns true when this feed intercom unprocessable entity response has a 2xx status code
func (o *FeedIntercomUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this feed intercom unprocessable entity response has a 3xx status code
func (o *FeedIntercomUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this feed intercom unprocessable entity response has a 4xx status code
func (o *FeedIntercomUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this feed intercom unprocessable entity response has a 5xx status code
func (o *FeedIntercomUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this feed intercom unprocessable entity response a status code equal to that given
func (o *FeedIntercomUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the feed intercom unprocessable entity response
func (o *FeedIntercomUnprocessableEntity) Code() int {
	return 422
}

func (o *FeedIntercomUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /api/backend/v1/communicationData/submit][%d] feedIntercomUnprocessableEntity ", 422)
}

func (o *FeedIntercomUnprocessableEntity) String() string {
	return fmt.Sprintf("[POST /api/backend/v1/communicationData/submit][%d] feedIntercomUnprocessableEntity ", 422)
}

func (o *FeedIntercomUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
