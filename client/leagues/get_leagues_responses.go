// Code generated by go-swagger; DO NOT EDIT.

package leagues

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetLeaguesReader is a Reader for the GetLeagues structure.
type GetLeaguesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLeaguesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLeaguesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLeaguesOK creates a GetLeaguesOK with default headers values
func NewGetLeaguesOK() *GetLeaguesOK {
	return &GetLeaguesOK{}
}

/* GetLeaguesOK describes a response with status code 200, with default header values.

Success
*/
type GetLeaguesOK struct {
	Payload []*GetLeaguesOKBodyItems0
}

func (o *GetLeaguesOK) Error() string {
	return fmt.Sprintf("[GET /leagues][%d] getLeaguesOK  %+v", 200, o.Payload)
}
func (o *GetLeaguesOK) GetPayload() []*GetLeaguesOKBodyItems0 {
	return o.Payload
}

func (o *GetLeaguesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetLeaguesOKBodyItems0 get leagues o k body items0
swagger:model GetLeaguesOKBodyItems0
*/
type GetLeaguesOKBodyItems0 struct {

	// banner
	Banner string `json:"banner,omitempty"`

	// leagueid
	Leagueid int64 `json:"leagueid,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// ticket
	Ticket string `json:"ticket,omitempty"`

	// tier
	Tier string `json:"tier,omitempty"`
}

// Validate validates this get leagues o k body items0
func (o *GetLeaguesOKBodyItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get leagues o k body items0 based on context it is used
func (o *GetLeaguesOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetLeaguesOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetLeaguesOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetLeaguesOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
