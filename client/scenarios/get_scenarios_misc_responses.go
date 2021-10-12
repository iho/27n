// Code generated by go-swagger; DO NOT EDIT.

package scenarios

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

// GetScenariosMiscReader is a Reader for the GetScenariosMisc structure.
type GetScenariosMiscReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetScenariosMiscReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetScenariosMiscOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetScenariosMiscOK creates a GetScenariosMiscOK with default headers values
func NewGetScenariosMiscOK() *GetScenariosMiscOK {
	return &GetScenariosMiscOK{}
}

/* GetScenariosMiscOK describes a response with status code 200, with default header values.

Success
*/
type GetScenariosMiscOK struct {
	Payload []*GetScenariosMiscOKBodyItems0
}

func (o *GetScenariosMiscOK) Error() string {
	return fmt.Sprintf("[GET /scenarios/misc][%d] getScenariosMiscOK  %+v", 200, o.Payload)
}
func (o *GetScenariosMiscOK) GetPayload() []*GetScenariosMiscOKBodyItems0 {
	return o.Payload
}

func (o *GetScenariosMiscOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetScenariosMiscOKBodyItems0 get scenarios misc o k body items0
swagger:model GetScenariosMiscOKBodyItems0
*/
type GetScenariosMiscOKBodyItems0 struct {

	// The number of games where this scenario occurred
	Games string `json:"games,omitempty"`

	// Boolean indicating whether Radiant executed this scenario
	IsRadiant bool `json:"is_radiant,omitempty"`

	// Region the game was played in
	Region int64 `json:"region,omitempty"`

	// The scenario's name or description
	Scenario string `json:"scenario,omitempty"`

	// The number of games won where this scenario occured
	Wins string `json:"wins,omitempty"`
}

// Validate validates this get scenarios misc o k body items0
func (o *GetScenariosMiscOKBodyItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get scenarios misc o k body items0 based on context it is used
func (o *GetScenariosMiscOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetScenariosMiscOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetScenariosMiscOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetScenariosMiscOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
