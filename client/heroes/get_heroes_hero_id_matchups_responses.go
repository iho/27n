// Code generated by go-swagger; DO NOT EDIT.

package heroes

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

// GetHeroesHeroIDMatchupsReader is a Reader for the GetHeroesHeroIDMatchups structure.
type GetHeroesHeroIDMatchupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetHeroesHeroIDMatchupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetHeroesHeroIDMatchupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetHeroesHeroIDMatchupsOK creates a GetHeroesHeroIDMatchupsOK with default headers values
func NewGetHeroesHeroIDMatchupsOK() *GetHeroesHeroIDMatchupsOK {
	return &GetHeroesHeroIDMatchupsOK{}
}

/* GetHeroesHeroIDMatchupsOK describes a response with status code 200, with default header values.

Success
*/
type GetHeroesHeroIDMatchupsOK struct {
	Payload []*GetHeroesHeroIDMatchupsOKBodyItems0
}

func (o *GetHeroesHeroIDMatchupsOK) Error() string {
	return fmt.Sprintf("[GET /heroes/{hero_id}/matchups][%d] getHeroesHeroIdMatchupsOK  %+v", 200, o.Payload)
}
func (o *GetHeroesHeroIDMatchupsOK) GetPayload() []*GetHeroesHeroIDMatchupsOKBodyItems0 {
	return o.Payload
}

func (o *GetHeroesHeroIDMatchupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetHeroesHeroIDMatchupsOKBodyItems0 get heroes hero ID matchups o k body items0
swagger:model GetHeroesHeroIDMatchupsOKBodyItems0
*/
type GetHeroesHeroIDMatchupsOKBodyItems0 struct {

	// Number of games played
	GamesPlayed int64 `json:"games_played,omitempty"`

	// Numeric identifier for the hero object
	HeroID int64 `json:"hero_id,omitempty"`

	// Number of games won
	Wins int64 `json:"wins,omitempty"`
}

// Validate validates this get heroes hero ID matchups o k body items0
func (o *GetHeroesHeroIDMatchupsOKBodyItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get heroes hero ID matchups o k body items0 based on context it is used
func (o *GetHeroesHeroIDMatchupsOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetHeroesHeroIDMatchupsOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetHeroesHeroIDMatchupsOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetHeroesHeroIDMatchupsOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
