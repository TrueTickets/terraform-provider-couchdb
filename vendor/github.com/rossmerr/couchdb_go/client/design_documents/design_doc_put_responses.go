// Code generated by go-swagger; DO NOT EDIT.

package design_documents

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/rossmerr/couchdb_go/models"
)

// DesignDocPutReader is a Reader for the DesignDocPut structure.
type DesignDocPutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DesignDocPutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDesignDocPutCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewDesignDocPutAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDesignDocPutBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDesignDocPutUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDesignDocPutNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDesignDocPutConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDesignDocPutCreated creates a DesignDocPutCreated with default headers values
func NewDesignDocPutCreated() *DesignDocPutCreated {
	return &DesignDocPutCreated{}
}

/*DesignDocPutCreated handles this case with default header values.

Document created and stored on disk
*/
type DesignDocPutCreated struct {
	/*Double quoted document’s revision token
	 */
	ETag string
	/*Document URI
	 */
	Location strfmt.URI

	Payload *models.DocumentOK
}

func (o *DesignDocPutCreated) Error() string {
	return fmt.Sprintf("[PUT /{db}/_design/{ddoc}][%d] designDocPutCreated  %+v", 201, o.Payload)
}

func (o *DesignDocPutCreated) GetPayload() *models.DocumentOK {
	return o.Payload
}

func (o *DesignDocPutCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header ETag
	o.ETag = response.GetHeader("ETag")

	// response header Location

	location, err := formats.Parse("uri", response.GetHeader("Location"))
	if err != nil {
		return errors.InvalidType("Location", "header", "strfmt.URI", response.GetHeader("Location"))
	}
	o.Location = *(location.(*strfmt.URI))

	o.Payload = new(models.DocumentOK)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDesignDocPutAccepted creates a DesignDocPutAccepted with default headers values
func NewDesignDocPutAccepted() *DesignDocPutAccepted {
	return &DesignDocPutAccepted{}
}

/*DesignDocPutAccepted handles this case with default header values.

Document data accepted, but not yet stored on disk
*/
type DesignDocPutAccepted struct {
	/*Double quoted document’s revision token
	 */
	ETag string
	/*Document URI
	 */
	Location strfmt.URI

	Payload *models.DocumentOK
}

func (o *DesignDocPutAccepted) Error() string {
	return fmt.Sprintf("[PUT /{db}/_design/{ddoc}][%d] designDocPutAccepted  %+v", 202, o.Payload)
}

func (o *DesignDocPutAccepted) GetPayload() *models.DocumentOK {
	return o.Payload
}

func (o *DesignDocPutAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header ETag
	o.ETag = response.GetHeader("ETag")

	// response header Location

	location, err := formats.Parse("uri", response.GetHeader("Location"))
	if err != nil {
		return errors.InvalidType("Location", "header", "strfmt.URI", response.GetHeader("Location"))
	}
	o.Location = *(location.(*strfmt.URI))

	o.Payload = new(models.DocumentOK)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDesignDocPutBadRequest creates a DesignDocPutBadRequest with default headers values
func NewDesignDocPutBadRequest() *DesignDocPutBadRequest {
	return &DesignDocPutBadRequest{}
}

/*DesignDocPutBadRequest handles this case with default header values.

Invalid request body or parameters
*/
type DesignDocPutBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *DesignDocPutBadRequest) Error() string {
	return fmt.Sprintf("[PUT /{db}/_design/{ddoc}][%d] designDocPutBadRequest  %+v", 400, o.Payload)
}

func (o *DesignDocPutBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DesignDocPutBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDesignDocPutUnauthorized creates a DesignDocPutUnauthorized with default headers values
func NewDesignDocPutUnauthorized() *DesignDocPutUnauthorized {
	return &DesignDocPutUnauthorized{}
}

/*DesignDocPutUnauthorized handles this case with default header values.

Write privileges required
*/
type DesignDocPutUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *DesignDocPutUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /{db}/_design/{ddoc}][%d] designDocPutUnauthorized  %+v", 401, o.Payload)
}

func (o *DesignDocPutUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DesignDocPutUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDesignDocPutNotFound creates a DesignDocPutNotFound with default headers values
func NewDesignDocPutNotFound() *DesignDocPutNotFound {
	return &DesignDocPutNotFound{}
}

/*DesignDocPutNotFound handles this case with default header values.

Specified database or document ID doesn’t exists
*/
type DesignDocPutNotFound struct {
	Payload *models.ErrorResponse
}

func (o *DesignDocPutNotFound) Error() string {
	return fmt.Sprintf("[PUT /{db}/_design/{ddoc}][%d] designDocPutNotFound  %+v", 404, o.Payload)
}

func (o *DesignDocPutNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DesignDocPutNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDesignDocPutConflict creates a DesignDocPutConflict with default headers values
func NewDesignDocPutConflict() *DesignDocPutConflict {
	return &DesignDocPutConflict{}
}

/*DesignDocPutConflict handles this case with default header values.

Document with the specified ID already exists or specified revision is not latest for target document
*/
type DesignDocPutConflict struct {
	Payload *models.ErrorResponse
}

func (o *DesignDocPutConflict) Error() string {
	return fmt.Sprintf("[PUT /{db}/_design/{ddoc}][%d] designDocPutConflict  %+v", 409, o.Payload)
}

func (o *DesignDocPutConflict) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DesignDocPutConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
