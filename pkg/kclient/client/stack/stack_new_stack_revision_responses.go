// Code generated by go-swagger; DO NOT EDIT.

package stack

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/koyeb/koyeb-cli/pkg/kclient/models"
)

// StackNewStackRevisionReader is a Reader for the StackNewStackRevision structure.
type StackNewStackRevisionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StackNewStackRevisionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStackNewStackRevisionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewStackNewStackRevisionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewStackNewStackRevisionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewStackNewStackRevisionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewStackNewStackRevisionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStackNewStackRevisionOK creates a StackNewStackRevisionOK with default headers values
func NewStackNewStackRevisionOK() *StackNewStackRevisionOK {
	return &StackNewStackRevisionOK{}
}

/*StackNewStackRevisionOK handles this case with default header values.

A successful response.
*/
type StackNewStackRevisionOK struct {
	Payload *models.StorageGetStackRevisionReply
}

func (o *StackNewStackRevisionOK) Error() string {
	return fmt.Sprintf("[POST /v1/stacks/{stack_id}/revisions][%d] stackNewStackRevisionOK  %+v", 200, o.Payload)
}

func (o *StackNewStackRevisionOK) GetPayload() *models.StorageGetStackRevisionReply {
	return o.Payload
}

func (o *StackNewStackRevisionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StorageGetStackRevisionReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStackNewStackRevisionBadRequest creates a StackNewStackRevisionBadRequest with default headers values
func NewStackNewStackRevisionBadRequest() *StackNewStackRevisionBadRequest {
	return &StackNewStackRevisionBadRequest{}
}

/*StackNewStackRevisionBadRequest handles this case with default header values.

Validation error
*/
type StackNewStackRevisionBadRequest struct {
	Payload *models.CommonErrorWithFields
}

func (o *StackNewStackRevisionBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/stacks/{stack_id}/revisions][%d] stackNewStackRevisionBadRequest  %+v", 400, o.Payload)
}

func (o *StackNewStackRevisionBadRequest) GetPayload() *models.CommonErrorWithFields {
	return o.Payload
}

func (o *StackNewStackRevisionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonErrorWithFields)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStackNewStackRevisionForbidden creates a StackNewStackRevisionForbidden with default headers values
func NewStackNewStackRevisionForbidden() *StackNewStackRevisionForbidden {
	return &StackNewStackRevisionForbidden{}
}

/*StackNewStackRevisionForbidden handles this case with default header values.

Returned when the user does not have permission to access the resource.
*/
type StackNewStackRevisionForbidden struct {
	Payload *models.CommonError
}

func (o *StackNewStackRevisionForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/stacks/{stack_id}/revisions][%d] stackNewStackRevisionForbidden  %+v", 403, o.Payload)
}

func (o *StackNewStackRevisionForbidden) GetPayload() *models.CommonError {
	return o.Payload
}

func (o *StackNewStackRevisionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStackNewStackRevisionNotFound creates a StackNewStackRevisionNotFound with default headers values
func NewStackNewStackRevisionNotFound() *StackNewStackRevisionNotFound {
	return &StackNewStackRevisionNotFound{}
}

/*StackNewStackRevisionNotFound handles this case with default header values.

Returned when the resource does not exist.
*/
type StackNewStackRevisionNotFound struct {
	Payload *models.CommonError
}

func (o *StackNewStackRevisionNotFound) Error() string {
	return fmt.Sprintf("[POST /v1/stacks/{stack_id}/revisions][%d] stackNewStackRevisionNotFound  %+v", 404, o.Payload)
}

func (o *StackNewStackRevisionNotFound) GetPayload() *models.CommonError {
	return o.Payload
}

func (o *StackNewStackRevisionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStackNewStackRevisionDefault creates a StackNewStackRevisionDefault with default headers values
func NewStackNewStackRevisionDefault(code int) *StackNewStackRevisionDefault {
	return &StackNewStackRevisionDefault{
		_statusCode: code,
	}
}

/*StackNewStackRevisionDefault handles this case with default header values.

An unexpected error response
*/
type StackNewStackRevisionDefault struct {
	_statusCode int

	Payload *models.GatewayruntimeError
}

// Code gets the status code for the stack new stack revision default response
func (o *StackNewStackRevisionDefault) Code() int {
	return o._statusCode
}

func (o *StackNewStackRevisionDefault) Error() string {
	return fmt.Sprintf("[POST /v1/stacks/{stack_id}/revisions][%d] stack_NewStackRevision default  %+v", o._statusCode, o.Payload)
}

func (o *StackNewStackRevisionDefault) GetPayload() *models.GatewayruntimeError {
	return o.Payload
}

func (o *StackNewStackRevisionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayruntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
