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

// GetLeaguesLeagueIDTeamsReader is a Reader for the GetLeaguesLeagueIDTeams structure.
type GetLeaguesLeagueIDTeamsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLeaguesLeagueIDTeamsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLeaguesLeagueIDTeamsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLeaguesLeagueIDTeamsOK creates a GetLeaguesLeagueIDTeamsOK with default headers values
func NewGetLeaguesLeagueIDTeamsOK() *GetLeaguesLeagueIDTeamsOK {
	return &GetLeaguesLeagueIDTeamsOK{}
}

/* GetLeaguesLeagueIDTeamsOK describes a response with status code 200, with default header values.

Success
*/
type GetLeaguesLeagueIDTeamsOK struct {
	Payload *GetLeaguesLeagueIDTeamsOKBody
}

func (o *GetLeaguesLeagueIDTeamsOK) Error() string {
	return fmt.Sprintf("[GET /leagues/{league_id}/teams][%d] getLeaguesLeagueIdTeamsOK  %+v", 200, o.Payload)
}
func (o *GetLeaguesLeagueIDTeamsOK) GetPayload() *GetLeaguesLeagueIDTeamsOKBody {
	return o.Payload
}

func (o *GetLeaguesLeagueIDTeamsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetLeaguesLeagueIDTeamsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetLeaguesLeagueIDTeamsOKBody get leagues league ID teams o k body
swagger:model GetLeaguesLeagueIDTeamsOKBody
*/
type GetLeaguesLeagueIDTeamsOKBody struct {

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

// Validate validates this get leagues league ID teams o k body
func (o *GetLeaguesLeagueIDTeamsOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get leagues league ID teams o k body based on context it is used
func (o *GetLeaguesLeagueIDTeamsOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetLeaguesLeagueIDTeamsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetLeaguesLeagueIDTeamsOKBody) UnmarshalBinary(b []byte) error {
	var res GetLeaguesLeagueIDTeamsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
