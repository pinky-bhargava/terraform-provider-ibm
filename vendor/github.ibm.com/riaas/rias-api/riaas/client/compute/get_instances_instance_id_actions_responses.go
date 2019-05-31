// Code generated by go-swagger; DO NOT EDIT.

package compute

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.ibm.com/riaas/rias-api/riaas/models"
)

// GetInstancesInstanceIDActionsReader is a Reader for the GetInstancesInstanceIDActions structure.
type GetInstancesInstanceIDActionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInstancesInstanceIDActionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetInstancesInstanceIDActionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetInstancesInstanceIDActionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetInstancesInstanceIDActionsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		
		return nil, result
	}
}

// NewGetInstancesInstanceIDActionsOK creates a GetInstancesInstanceIDActionsOK with default headers values
func NewGetInstancesInstanceIDActionsOK() *GetInstancesInstanceIDActionsOK {
	return &GetInstancesInstanceIDActionsOK{}
}

/*GetInstancesInstanceIDActionsOK handles this case with default header values.

dummy
*/
type GetInstancesInstanceIDActionsOK struct {
	Payload *models.GetInstancesInstanceIDActionsOKBody
}

func (o *GetInstancesInstanceIDActionsOK) Error() string {
	return fmt.Sprintf("[GET /instances/{instance_id}/actions][%d] getInstancesInstanceIdActionsOK  %+v", 200, o.Payload)
}

func (o *GetInstancesInstanceIDActionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetInstancesInstanceIDActionsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstancesInstanceIDActionsNotFound creates a GetInstancesInstanceIDActionsNotFound with default headers values
func NewGetInstancesInstanceIDActionsNotFound() *GetInstancesInstanceIDActionsNotFound {
	return &GetInstancesInstanceIDActionsNotFound{}
}

/*GetInstancesInstanceIDActionsNotFound handles this case with default header values.

error
*/
type GetInstancesInstanceIDActionsNotFound struct {
	Payload *models.Riaaserror
}

func (o *GetInstancesInstanceIDActionsNotFound) Error() string {
	return fmt.Sprintf("[GET /instances/{instance_id}/actions][%d] getInstancesInstanceIdActionsNotFound  %+v", 404, o.Payload)
}

func (o *GetInstancesInstanceIDActionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstancesInstanceIDActionsDefault creates a GetInstancesInstanceIDActionsDefault with default headers values
func NewGetInstancesInstanceIDActionsDefault(code int) *GetInstancesInstanceIDActionsDefault {
	return &GetInstancesInstanceIDActionsDefault{
		_statusCode: code,
	}
}

/*GetInstancesInstanceIDActionsDefault handles this case with default header values.

unexpectederror
*/
type GetInstancesInstanceIDActionsDefault struct {
	_statusCode int

	Payload *models.Riaaserror
}

// Code gets the status code for the get instances instance ID actions default response
func (o *GetInstancesInstanceIDActionsDefault) Code() int {
	return o._statusCode
}

func (o *GetInstancesInstanceIDActionsDefault) Error() string {
	return fmt.Sprintf("[GET /instances/{instance_id}/actions][%d] GetInstancesInstanceIDActions default  %+v", o._statusCode, o.Payload)
}

func (o *GetInstancesInstanceIDActionsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}