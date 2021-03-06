// Code generated by go-swagger; DO NOT EDIT.

package request

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetRequestJobIDReader is a Reader for the GetRequestJobID structure.
type GetRequestJobIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRequestJobIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRequestJobIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRequestJobIDOK creates a GetRequestJobIDOK with default headers values
func NewGetRequestJobIDOK() *GetRequestJobIDOK {
	return &GetRequestJobIDOK{}
}

/* GetRequestJobIDOK describes a response with status code 200, with default header values.

Success
*/
type GetRequestJobIDOK struct {
	Payload interface{}
}

func (o *GetRequestJobIDOK) Error() string {
	return fmt.Sprintf("[GET /request/{jobId}][%d] getRequestJobIdOK  %+v", 200, o.Payload)
}
func (o *GetRequestJobIDOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetRequestJobIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
