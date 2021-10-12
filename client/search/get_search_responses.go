// Code generated by go-swagger; DO NOT EDIT.

package search

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

// GetSearchReader is a Reader for the GetSearch structure.
type GetSearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSearchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSearchOK creates a GetSearchOK with default headers values
func NewGetSearchOK() *GetSearchOK {
	return &GetSearchOK{}
}

/* GetSearchOK describes a response with status code 200, with default header values.

Success
*/
type GetSearchOK struct {
	Payload []*GetSearchOKBodyItems0
}

func (o *GetSearchOK) Error() string {
	return fmt.Sprintf("[GET /search][%d] getSearchOK  %+v", 200, o.Payload)
}
func (o *GetSearchOK) GetPayload() []*GetSearchOKBodyItems0 {
	return o.Payload
}

func (o *GetSearchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetSearchOKBodyItems0 get search o k body items0
swagger:model GetSearchOKBodyItems0
*/
type GetSearchOKBodyItems0 struct {

	// account_id
	AccountID int64 `json:"account_id,omitempty"`

	// avatarfull
	Avatarfull string `json:"avatarfull,omitempty"`

	// last_match_time. May not be present or null.
	LastMatchTime string `json:"last_match_time,omitempty"`

	// personaname
	Personaname string `json:"personaname,omitempty"`

	// similarity
	Similarity float64 `json:"similarity,omitempty"`
}

// Validate validates this get search o k body items0
func (o *GetSearchOKBodyItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get search o k body items0 based on context it is used
func (o *GetSearchOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetSearchOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetSearchOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetSearchOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}