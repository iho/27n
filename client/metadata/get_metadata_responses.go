// Code generated by go-swagger; DO NOT EDIT.

package metadata

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetMetadataReader is a Reader for the GetMetadata structure.
type GetMetadataReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMetadataReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetMetadataOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetMetadataOK creates a GetMetadataOK with default headers values
func NewGetMetadataOK() *GetMetadataOK {
	return &GetMetadataOK{}
}

/* GetMetadataOK describes a response with status code 200, with default header values.

Success
*/
type GetMetadataOK struct {
	Payload *GetMetadataOKBody
}

func (o *GetMetadataOK) Error() string {
	return fmt.Sprintf("[GET /metadata][%d] getMetadataOK  %+v", 200, o.Payload)
}
func (o *GetMetadataOK) GetPayload() *GetMetadataOKBody {
	return o.Payload
}

func (o *GetMetadataOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetMetadataOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetMetadataOKBody get metadata o k body
swagger:model GetMetadataOKBody
*/
type GetMetadataOKBody struct {

	// banner
	Banner interface{} `json:"banner,omitempty"`

	// cheese
	Cheese *GetMetadataOKBodyCheese `json:"cheese,omitempty"`
}

// Validate validates this get metadata o k body
func (o *GetMetadataOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCheese(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetMetadataOKBody) validateCheese(formats strfmt.Registry) error {
	if swag.IsZero(o.Cheese) { // not required
		return nil
	}

	if o.Cheese != nil {
		if err := o.Cheese.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getMetadataOK" + "." + "cheese")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get metadata o k body based on the context it is used
func (o *GetMetadataOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateCheese(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetMetadataOKBody) contextValidateCheese(ctx context.Context, formats strfmt.Registry) error {

	if o.Cheese != nil {
		if err := o.Cheese.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getMetadataOK" + "." + "cheese")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetMetadataOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetMetadataOKBody) UnmarshalBinary(b []byte) error {
	var res GetMetadataOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetMetadataOKBodyCheese cheese
swagger:model GetMetadataOKBodyCheese
*/
type GetMetadataOKBodyCheese struct {

	// cheese
	Cheese string `json:"cheese,omitempty"`

	// goal
	Goal string `json:"goal,omitempty"`
}

// Validate validates this get metadata o k body cheese
func (o *GetMetadataOKBodyCheese) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get metadata o k body cheese based on context it is used
func (o *GetMetadataOKBodyCheese) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetMetadataOKBodyCheese) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetMetadataOKBodyCheese) UnmarshalBinary(b []byte) error {
	var res GetMetadataOKBodyCheese
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
