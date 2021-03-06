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

// GetHeroesHeroIDDurationsReader is a Reader for the GetHeroesHeroIDDurations structure.
type GetHeroesHeroIDDurationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetHeroesHeroIDDurationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetHeroesHeroIDDurationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetHeroesHeroIDDurationsOK creates a GetHeroesHeroIDDurationsOK with default headers values
func NewGetHeroesHeroIDDurationsOK() *GetHeroesHeroIDDurationsOK {
	return &GetHeroesHeroIDDurationsOK{}
}

/* GetHeroesHeroIDDurationsOK describes a response with status code 200, with default header values.

Success
*/
type GetHeroesHeroIDDurationsOK struct {
	Payload []*GetHeroesHeroIDDurationsOKBodyItems0
}

func (o *GetHeroesHeroIDDurationsOK) Error() string {
	return fmt.Sprintf("[GET /heroes/{hero_id}/durations][%d] getHeroesHeroIdDurationsOK  %+v", 200, o.Payload)
}
func (o *GetHeroesHeroIDDurationsOK) GetPayload() []*GetHeroesHeroIDDurationsOKBodyItems0 {
	return o.Payload
}

func (o *GetHeroesHeroIDDurationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetHeroesHeroIDDurationsOKBodyItems0 get heroes hero ID durations o k body items0
swagger:model GetHeroesHeroIDDurationsOKBodyItems0
*/
type GetHeroesHeroIDDurationsOKBodyItems0 struct {

	// Lower bound of number of seconds the match lasted
	DurationBin string `json:"duration_bin,omitempty"`

	// Number of games played
	GamesPlayed int64 `json:"games_played,omitempty"`

	// Number of wins
	Wins int64 `json:"wins,omitempty"`
}

// Validate validates this get heroes hero ID durations o k body items0
func (o *GetHeroesHeroIDDurationsOKBodyItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get heroes hero ID durations o k body items0 based on context it is used
func (o *GetHeroesHeroIDDurationsOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetHeroesHeroIDDurationsOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetHeroesHeroIDDurationsOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetHeroesHeroIDDurationsOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
