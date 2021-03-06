// Code generated by go-swagger; DO NOT EDIT.

package teams

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

// GetTeamsReader is a Reader for the GetTeams structure.
type GetTeamsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTeamsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTeamsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTeamsOK creates a GetTeamsOK with default headers values
func NewGetTeamsOK() *GetTeamsOK {
	return &GetTeamsOK{}
}

/* GetTeamsOK describes a response with status code 200, with default header values.

Success
*/
type GetTeamsOK struct {
	Payload []*GetTeamsOKBodyItems0
}

func (o *GetTeamsOK) Error() string {
	return fmt.Sprintf("[GET /teams][%d] getTeamsOK  %+v", 200, o.Payload)
}
func (o *GetTeamsOK) GetPayload() []*GetTeamsOKBodyItems0 {
	return o.Payload
}

func (o *GetTeamsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetTeamsOKBodyItems0 get teams o k body items0
swagger:model GetTeamsOKBodyItems0
*/
type GetTeamsOKBodyItems0 struct {

	// The Unix timestamp of the last match played by this team
	LastMatchTime int64 `json:"last_match_time,omitempty"`

	// The number of losses by this team
	Losses int64 `json:"losses,omitempty"`

	// Team name, eg. 'Newbee'
	Name string `json:"name,omitempty"`

	// The Elo rating of the team
	Rating float64 `json:"rating,omitempty"`

	// The team tag/abbreviation
	Tag string `json:"tag,omitempty"`

	// Team's identifier
	TeamID int64 `json:"team_id,omitempty"`

	// The number of games won by this team
	Wins int64 `json:"wins,omitempty"`
}

// Validate validates this get teams o k body items0
func (o *GetTeamsOKBodyItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get teams o k body items0 based on context it is used
func (o *GetTeamsOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetTeamsOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetTeamsOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetTeamsOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
