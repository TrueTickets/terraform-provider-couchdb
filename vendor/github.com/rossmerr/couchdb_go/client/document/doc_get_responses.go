// Code generated by go-swagger; DO NOT EDIT.

package document

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/rossmerr/couchdb_go/models"
)

// DocGetReader is a Reader for the DocGet structure.
type DocGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DocGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDocGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewDocGetNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewDocGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDocGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDocGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDocGetOK creates a DocGetOK with default headers values
func NewDocGetOK() *DocGetOK {
	return &DocGetOK{}
}

/*DocGetOK handles this case with default header values.

Request completed successfully
*/
type DocGetOK struct {
	/*Double quoted document’s revision token
	 */
	ETag string
	/*chunked. Available if requested with query parameter open_revs
	 */
	TransferEncoding string

	Payload models.Document
}

func (o *DocGetOK) Error() string {
	return fmt.Sprintf("[GET /{db}/{docid}][%d] docGetOK  %+v", 200, o.Payload)
}

func (o *DocGetOK) GetPayload() models.Document {
	return o.Payload
}

func (o *DocGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header ETag
	o.ETag = response.GetHeader("ETag")

	// response header Transfer-Encoding
	o.TransferEncoding = response.GetHeader("Transfer-Encoding")

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDocGetNotModified creates a DocGetNotModified with default headers values
func NewDocGetNotModified() *DocGetNotModified {
	return &DocGetNotModified{}
}

/*DocGetNotModified handles this case with default header values.

Document wasn’t modified since specified revision
*/
type DocGetNotModified struct {
}

func (o *DocGetNotModified) Error() string {
	return fmt.Sprintf("[GET /{db}/{docid}][%d] docGetNotModified ", 304)
}

func (o *DocGetNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDocGetBadRequest creates a DocGetBadRequest with default headers values
func NewDocGetBadRequest() *DocGetBadRequest {
	return &DocGetBadRequest{}
}

/*DocGetBadRequest handles this case with default header values.

The format of the request or revision was invalid
*/
type DocGetBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *DocGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /{db}/{docid}][%d] docGetBadRequest  %+v", 400, o.Payload)
}

func (o *DocGetBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DocGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDocGetUnauthorized creates a DocGetUnauthorized with default headers values
func NewDocGetUnauthorized() *DocGetUnauthorized {
	return &DocGetUnauthorized{}
}

/*DocGetUnauthorized handles this case with default header values.

Read privilege required
*/
type DocGetUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *DocGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /{db}/{docid}][%d] docGetUnauthorized  %+v", 401, o.Payload)
}

func (o *DocGetUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DocGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDocGetNotFound creates a DocGetNotFound with default headers values
func NewDocGetNotFound() *DocGetNotFound {
	return &DocGetNotFound{}
}

/*DocGetNotFound handles this case with default header values.

Document not found
*/
type DocGetNotFound struct {
	Payload *models.ErrorResponse
}

func (o *DocGetNotFound) Error() string {
	return fmt.Sprintf("[GET /{db}/{docid}][%d] docGetNotFound  %+v", 404, o.Payload)
}

func (o *DocGetNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DocGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
