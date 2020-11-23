// Code generated by go-swagger; DO NOT EDIT.

package database

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/rossmerr/couchdb_go/models"
)

// DocGetAllReader is a Reader for the DocGetAll structure.
type DocGetAllReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DocGetAllReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDocGetAllOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDocGetAllNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDocGetAllOK creates a DocGetAllOK with default headers values
func NewDocGetAllOK() *DocGetAllOK {
	return &DocGetAllOK{}
}

/*DocGetAllOK handles this case with default header values.

Request completed successfully
*/
type DocGetAllOK struct {
	Payload *models.Pagination
}

func (o *DocGetAllOK) Error() string {
	return fmt.Sprintf("[GET /{db}/_all_docs][%d] docGetAllOK  %+v", 200, o.Payload)
}

func (o *DocGetAllOK) GetPayload() *models.Pagination {
	return o.Payload
}

func (o *DocGetAllOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Pagination)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDocGetAllNotFound creates a DocGetAllNotFound with default headers values
func NewDocGetAllNotFound() *DocGetAllNotFound {
	return &DocGetAllNotFound{}
}

/*DocGetAllNotFound handles this case with default header values.

Requested database not found
*/
type DocGetAllNotFound struct {
	Payload *models.ErrorResponse
}

func (o *DocGetAllNotFound) Error() string {
	return fmt.Sprintf("[GET /{db}/_all_docs][%d] docGetAllNotFound  %+v", 404, o.Payload)
}

func (o *DocGetAllNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DocGetAllNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
