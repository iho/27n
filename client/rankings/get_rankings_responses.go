// Code generated by go-swagger; DO NOT EDIT.

package rankings

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
	"github.com/go-openapi/validate"
)

// GetRankingsReader is a Reader for the GetRankings structure.
type GetRankingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRankingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRankingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRankingsOK creates a GetRankingsOK with default headers values
func NewGetRankingsOK() *GetRankingsOK {
	return &GetRankingsOK{}
}

/* GetRankingsOK describes a response with status code 200, with default header values.

Success
*/
type GetRankingsOK struct {
	Payload *GetRankingsOKBody
}

func (o *GetRankingsOK) Error() string {
	return fmt.Sprintf("[GET /rankings][%d] getRankingsOK  %+v", 200, o.Payload)
}
func (o *GetRankingsOK) GetPayload() *GetRankingsOKBody {
	return o.Payload
}

func (o *GetRankingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetRankingsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetRankingsOKBody get rankings o k body
swagger:model GetRankingsOKBody
*/
type GetRankingsOKBody struct {

	// The ID value of the hero played
	HeroID int64 `json:"hero_id,omitempty"`

	// rankings
	Rankings *GetRankingsOKBodyRankings `json:"rankings,omitempty"`
}

// Validate validates this get rankings o k body
func (o *GetRankingsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateRankings(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetRankingsOKBody) validateRankings(formats strfmt.Registry) error {
	if swag.IsZero(o.Rankings) { // not required
		return nil
	}

	if o.Rankings != nil {
		if err := o.Rankings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getRankingsOK" + "." + "rankings")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get rankings o k body based on the context it is used
func (o *GetRankingsOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateRankings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetRankingsOKBody) contextValidateRankings(ctx context.Context, formats strfmt.Registry) error {

	if o.Rankings != nil {
		if err := o.Rankings.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getRankingsOK" + "." + "rankings")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetRankingsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetRankingsOKBody) UnmarshalBinary(b []byte) error {
	var res GetRankingsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetRankingsOKBodyRankings rankings
swagger:model GetRankingsOKBodyRankings
*/
type GetRankingsOKBodyRankings struct {

	// account_id
	AccountID int64 `json:"account_id,omitempty"`

	// avatar
	Avatar string `json:"avatar,omitempty"`

	// avatarfull
	Avatarfull string `json:"avatarfull,omitempty"`

	// avatarmedium
	Avatarmedium string `json:"avatarmedium,omitempty"`

	// cheese
	Cheese int64 `json:"cheese,omitempty"`

	// fh_unavailable
	FhUnavailable bool `json:"fh_unavailable,omitempty"`

	// full_history_time
	// Format: date-time
	FullHistoryTime strfmt.DateTime `json:"full_history_time,omitempty"`

	// last_login
	// Format: date-time
	LastLogin strfmt.DateTime `json:"last_login,omitempty"`

	// loccountrycode
	Loccountrycode string `json:"loccountrycode,omitempty"`

	// personaname
	Personaname string `json:"personaname,omitempty"`

	// profileurl
	Profileurl string `json:"profileurl,omitempty"`

	// rank_tier
	RankTier int64 `json:"rank_tier,omitempty"`

	// score
	Score string `json:"score,omitempty"`

	// steamid
	Steamid string `json:"steamid,omitempty"`
}

// Validate validates this get rankings o k body rankings
func (o *GetRankingsOKBodyRankings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateFullHistoryTime(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateLastLogin(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetRankingsOKBodyRankings) validateFullHistoryTime(formats strfmt.Registry) error {
	if swag.IsZero(o.FullHistoryTime) { // not required
		return nil
	}

	if err := validate.FormatOf("getRankingsOK"+"."+"rankings"+"."+"full_history_time", "body", "date-time", o.FullHistoryTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *GetRankingsOKBodyRankings) validateLastLogin(formats strfmt.Registry) error {
	if swag.IsZero(o.LastLogin) { // not required
		return nil
	}

	if err := validate.FormatOf("getRankingsOK"+"."+"rankings"+"."+"last_login", "body", "date-time", o.LastLogin.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get rankings o k body rankings based on context it is used
func (o *GetRankingsOKBodyRankings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetRankingsOKBodyRankings) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetRankingsOKBodyRankings) UnmarshalBinary(b []byte) error {
	var res GetRankingsOKBodyRankings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
