// Code generated by go-swagger; DO NOT EDIT.

package distributions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetDistributionsReader is a Reader for the GetDistributions structure.
type GetDistributionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDistributionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDistributionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDistributionsOK creates a GetDistributionsOK with default headers values
func NewGetDistributionsOK() *GetDistributionsOK {
	return &GetDistributionsOK{}
}

/* GetDistributionsOK describes a response with status code 200, with default header values.

Success
*/
type GetDistributionsOK struct {
	Payload *GetDistributionsOKBody
}

func (o *GetDistributionsOK) Error() string {
	return fmt.Sprintf("[GET /distributions][%d] getDistributionsOK  %+v", 200, o.Payload)
}
func (o *GetDistributionsOK) GetPayload() *GetDistributionsOKBody {
	return o.Payload
}

func (o *GetDistributionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetDistributionsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetDistributionsOKBody get distributions o k body
swagger:model GetDistributionsOKBody
*/
type GetDistributionsOKBody struct {

	// country mmr
	CountryMmr *GetDistributionsOKBodyCountryMmr `json:"country_mmr,omitempty"`

	// mmr
	Mmr *GetDistributionsOKBodyMmr `json:"mmr,omitempty"`

	// ranks
	Ranks *GetDistributionsOKBodyRanks `json:"ranks,omitempty"`
}

// Validate validates this get distributions o k body
func (o *GetDistributionsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCountryMmr(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMmr(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRanks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetDistributionsOKBody) validateCountryMmr(formats strfmt.Registry) error {
	if swag.IsZero(o.CountryMmr) { // not required
		return nil
	}

	if o.CountryMmr != nil {
		if err := o.CountryMmr.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getDistributionsOK" + "." + "country_mmr")
			}
			return err
		}
	}

	return nil
}

func (o *GetDistributionsOKBody) validateMmr(formats strfmt.Registry) error {
	if swag.IsZero(o.Mmr) { // not required
		return nil
	}

	if o.Mmr != nil {
		if err := o.Mmr.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getDistributionsOK" + "." + "mmr")
			}
			return err
		}
	}

	return nil
}

func (o *GetDistributionsOKBody) validateRanks(formats strfmt.Registry) error {
	if swag.IsZero(o.Ranks) { // not required
		return nil
	}

	if o.Ranks != nil {
		if err := o.Ranks.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getDistributionsOK" + "." + "ranks")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get distributions o k body based on the context it is used
func (o *GetDistributionsOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateCountryMmr(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateMmr(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateRanks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetDistributionsOKBody) contextValidateCountryMmr(ctx context.Context, formats strfmt.Registry) error {

	if o.CountryMmr != nil {
		if err := o.CountryMmr.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getDistributionsOK" + "." + "country_mmr")
			}
			return err
		}
	}

	return nil
}

func (o *GetDistributionsOKBody) contextValidateMmr(ctx context.Context, formats strfmt.Registry) error {

	if o.Mmr != nil {
		if err := o.Mmr.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getDistributionsOK" + "." + "mmr")
			}
			return err
		}
	}

	return nil
}

func (o *GetDistributionsOKBody) contextValidateRanks(ctx context.Context, formats strfmt.Registry) error {

	if o.Ranks != nil {
		if err := o.Ranks.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getDistributionsOK" + "." + "ranks")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetDistributionsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetDistributionsOKBody) UnmarshalBinary(b []byte) error {
	var res GetDistributionsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetDistributionsOKBodyCountryMmr country_mmr
swagger:model GetDistributionsOKBodyCountryMmr
*/
type GetDistributionsOKBodyCountryMmr struct {

	// command
	Commmand string `json:"commmand,omitempty"`

	// fields
	Fields []*GetDistributionsOKBodyCountryMmrFieldsItems0 `json:"fields"`

	// rowAsArray
	RowAsArray bool `json:"rowAsArray,omitempty"`

	// rowCount
	RowCount int64 `json:"rowCount,omitempty"`

	// rows
	Rows []*GetDistributionsOKBodyCountryMmrRowsItems0 `json:"rows"`
}

// Validate validates this get distributions o k body country mmr
func (o *GetDistributionsOKBodyCountryMmr) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateFields(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRows(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetDistributionsOKBodyCountryMmr) validateFields(formats strfmt.Registry) error {
	if swag.IsZero(o.Fields) { // not required
		return nil
	}

	for i := 0; i < len(o.Fields); i++ {
		if swag.IsZero(o.Fields[i]) { // not required
			continue
		}

		if o.Fields[i] != nil {
			if err := o.Fields[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getDistributionsOK" + "." + "country_mmr" + "." + "fields" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetDistributionsOKBodyCountryMmr) validateRows(formats strfmt.Registry) error {
	if swag.IsZero(o.Rows) { // not required
		return nil
	}

	for i := 0; i < len(o.Rows); i++ {
		if swag.IsZero(o.Rows[i]) { // not required
			continue
		}

		if o.Rows[i] != nil {
			if err := o.Rows[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getDistributionsOK" + "." + "country_mmr" + "." + "rows" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get distributions o k body country mmr based on the context it is used
func (o *GetDistributionsOKBodyCountryMmr) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateFields(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateRows(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetDistributionsOKBodyCountryMmr) contextValidateFields(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Fields); i++ {

		if o.Fields[i] != nil {
			if err := o.Fields[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getDistributionsOK" + "." + "country_mmr" + "." + "fields" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetDistributionsOKBodyCountryMmr) contextValidateRows(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Rows); i++ {

		if o.Rows[i] != nil {
			if err := o.Rows[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getDistributionsOK" + "." + "country_mmr" + "." + "rows" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetDistributionsOKBodyCountryMmr) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetDistributionsOKBodyCountryMmr) UnmarshalBinary(b []byte) error {
	var res GetDistributionsOKBodyCountryMmr
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetDistributionsOKBodyCountryMmrFieldsItems0 get distributions o k body country mmr fields items0
swagger:model GetDistributionsOKBodyCountryMmrFieldsItems0
*/
type GetDistributionsOKBodyCountryMmrFieldsItems0 struct {

	// columnID
	ColumnID int64 `json:"columnID,omitempty"`

	// dataTypeID
	DataTypeID int64 `json:"dataTypeID,omitempty"`

	// dataTypeModifier
	DataTypeModifier int64 `json:"dataTypeModifier,omitempty"`

	// dataTypeSize
	DataTypeSize int64 `json:"dataTypeSize,omitempty"`

	// format
	Format string `json:"format,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// tableID
	TableID int64 `json:"tableID,omitempty"`
}

// Validate validates this get distributions o k body country mmr fields items0
func (o *GetDistributionsOKBodyCountryMmrFieldsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get distributions o k body country mmr fields items0 based on context it is used
func (o *GetDistributionsOKBodyCountryMmrFieldsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetDistributionsOKBodyCountryMmrFieldsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetDistributionsOKBodyCountryMmrFieldsItems0) UnmarshalBinary(b []byte) error {
	var res GetDistributionsOKBodyCountryMmrFieldsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetDistributionsOKBodyCountryMmrRowsItems0 get distributions o k body country mmr rows items0
swagger:model GetDistributionsOKBodyCountryMmrRowsItems0
*/
type GetDistributionsOKBodyCountryMmrRowsItems0 struct {

	// avg
	Avg string `json:"avg,omitempty"`

	// common
	Common string `json:"common,omitempty"`

	// count
	Count int64 `json:"count,omitempty"`

	// loccountrycode
	Loccountrycode string `json:"loccountrycode,omitempty"`
}

// Validate validates this get distributions o k body country mmr rows items0
func (o *GetDistributionsOKBodyCountryMmrRowsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get distributions o k body country mmr rows items0 based on context it is used
func (o *GetDistributionsOKBodyCountryMmrRowsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetDistributionsOKBodyCountryMmrRowsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetDistributionsOKBodyCountryMmrRowsItems0) UnmarshalBinary(b []byte) error {
	var res GetDistributionsOKBodyCountryMmrRowsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetDistributionsOKBodyMmr mmr
swagger:model GetDistributionsOKBodyMmr
*/
type GetDistributionsOKBodyMmr struct {

	// command
	Commmand string `json:"commmand,omitempty"`

	// fields
	Fields []*GetDistributionsOKBodyMmrFieldsItems0 `json:"fields"`

	// rowAsArray
	RowAsArray bool `json:"rowAsArray,omitempty"`

	// rowCount
	RowCount int64 `json:"rowCount,omitempty"`

	// rows
	Rows []*GetDistributionsOKBodyMmrRowsItems0 `json:"rows"`

	// sum
	Sum *GetDistributionsOKBodyMmrSum `json:"sum,omitempty"`
}

// Validate validates this get distributions o k body mmr
func (o *GetDistributionsOKBodyMmr) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateFields(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRows(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSum(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetDistributionsOKBodyMmr) validateFields(formats strfmt.Registry) error {
	if swag.IsZero(o.Fields) { // not required
		return nil
	}

	for i := 0; i < len(o.Fields); i++ {
		if swag.IsZero(o.Fields[i]) { // not required
			continue
		}

		if o.Fields[i] != nil {
			if err := o.Fields[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getDistributionsOK" + "." + "mmr" + "." + "fields" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetDistributionsOKBodyMmr) validateRows(formats strfmt.Registry) error {
	if swag.IsZero(o.Rows) { // not required
		return nil
	}

	for i := 0; i < len(o.Rows); i++ {
		if swag.IsZero(o.Rows[i]) { // not required
			continue
		}

		if o.Rows[i] != nil {
			if err := o.Rows[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getDistributionsOK" + "." + "mmr" + "." + "rows" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetDistributionsOKBodyMmr) validateSum(formats strfmt.Registry) error {
	if swag.IsZero(o.Sum) { // not required
		return nil
	}

	if o.Sum != nil {
		if err := o.Sum.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getDistributionsOK" + "." + "mmr" + "." + "sum")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get distributions o k body mmr based on the context it is used
func (o *GetDistributionsOKBodyMmr) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateFields(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateRows(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateSum(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetDistributionsOKBodyMmr) contextValidateFields(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Fields); i++ {

		if o.Fields[i] != nil {
			if err := o.Fields[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getDistributionsOK" + "." + "mmr" + "." + "fields" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetDistributionsOKBodyMmr) contextValidateRows(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Rows); i++ {

		if o.Rows[i] != nil {
			if err := o.Rows[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getDistributionsOK" + "." + "mmr" + "." + "rows" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetDistributionsOKBodyMmr) contextValidateSum(ctx context.Context, formats strfmt.Registry) error {

	if o.Sum != nil {
		if err := o.Sum.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getDistributionsOK" + "." + "mmr" + "." + "sum")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetDistributionsOKBodyMmr) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetDistributionsOKBodyMmr) UnmarshalBinary(b []byte) error {
	var res GetDistributionsOKBodyMmr
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetDistributionsOKBodyMmrFieldsItems0 get distributions o k body mmr fields items0
swagger:model GetDistributionsOKBodyMmrFieldsItems0
*/
type GetDistributionsOKBodyMmrFieldsItems0 struct {

	// columnID
	ColumnID int64 `json:"columnID,omitempty"`

	// dataTypeID
	DataTypeID int64 `json:"dataTypeID,omitempty"`

	// dataTypeModifier
	DataTypeModifier string `json:"dataTypeModifier,omitempty"`

	// dataTypeSize
	DataTypeSize int64 `json:"dataTypeSize,omitempty"`

	// format
	Format string `json:"format,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// tableID
	TableID int64 `json:"tableID,omitempty"`
}

// Validate validates this get distributions o k body mmr fields items0
func (o *GetDistributionsOKBodyMmrFieldsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get distributions o k body mmr fields items0 based on context it is used
func (o *GetDistributionsOKBodyMmrFieldsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetDistributionsOKBodyMmrFieldsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetDistributionsOKBodyMmrFieldsItems0) UnmarshalBinary(b []byte) error {
	var res GetDistributionsOKBodyMmrFieldsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetDistributionsOKBodyMmrRowsItems0 get distributions o k body mmr rows items0
swagger:model GetDistributionsOKBodyMmrRowsItems0
*/
type GetDistributionsOKBodyMmrRowsItems0 struct {

	// bin
	Bin int64 `json:"bin,omitempty"`

	// bin_name
	BinName int64 `json:"bin_name,omitempty"`

	// count
	Count int64 `json:"count,omitempty"`

	// cumulative_sum
	CumulativeSum int64 `json:"cumulative_sum,omitempty"`
}

// Validate validates this get distributions o k body mmr rows items0
func (o *GetDistributionsOKBodyMmrRowsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get distributions o k body mmr rows items0 based on context it is used
func (o *GetDistributionsOKBodyMmrRowsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetDistributionsOKBodyMmrRowsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetDistributionsOKBodyMmrRowsItems0) UnmarshalBinary(b []byte) error {
	var res GetDistributionsOKBodyMmrRowsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetDistributionsOKBodyMmrSum sum
swagger:model GetDistributionsOKBodyMmrSum
*/
type GetDistributionsOKBodyMmrSum struct {

	// count
	Count int64 `json:"count,omitempty"`
}

// Validate validates this get distributions o k body mmr sum
func (o *GetDistributionsOKBodyMmrSum) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get distributions o k body mmr sum based on context it is used
func (o *GetDistributionsOKBodyMmrSum) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetDistributionsOKBodyMmrSum) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetDistributionsOKBodyMmrSum) UnmarshalBinary(b []byte) error {
	var res GetDistributionsOKBodyMmrSum
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetDistributionsOKBodyRanks ranks
swagger:model GetDistributionsOKBodyRanks
*/
type GetDistributionsOKBodyRanks struct {

	// command
	Commmand string `json:"commmand,omitempty"`

	// fields
	Fields []*GetDistributionsOKBodyRanksFieldsItems0 `json:"fields"`

	// rowAsArray
	RowAsArray bool `json:"rowAsArray,omitempty"`

	// rowCount
	RowCount int64 `json:"rowCount,omitempty"`

	// rows
	Rows []*GetDistributionsOKBodyRanksRowsItems0 `json:"rows"`

	// sum
	Sum *GetDistributionsOKBodyRanksSum `json:"sum,omitempty"`
}

// Validate validates this get distributions o k body ranks
func (o *GetDistributionsOKBodyRanks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateFields(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRows(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSum(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetDistributionsOKBodyRanks) validateFields(formats strfmt.Registry) error {
	if swag.IsZero(o.Fields) { // not required
		return nil
	}

	for i := 0; i < len(o.Fields); i++ {
		if swag.IsZero(o.Fields[i]) { // not required
			continue
		}

		if o.Fields[i] != nil {
			if err := o.Fields[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getDistributionsOK" + "." + "ranks" + "." + "fields" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetDistributionsOKBodyRanks) validateRows(formats strfmt.Registry) error {
	if swag.IsZero(o.Rows) { // not required
		return nil
	}

	for i := 0; i < len(o.Rows); i++ {
		if swag.IsZero(o.Rows[i]) { // not required
			continue
		}

		if o.Rows[i] != nil {
			if err := o.Rows[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getDistributionsOK" + "." + "ranks" + "." + "rows" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetDistributionsOKBodyRanks) validateSum(formats strfmt.Registry) error {
	if swag.IsZero(o.Sum) { // not required
		return nil
	}

	if o.Sum != nil {
		if err := o.Sum.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getDistributionsOK" + "." + "ranks" + "." + "sum")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get distributions o k body ranks based on the context it is used
func (o *GetDistributionsOKBodyRanks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateFields(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateRows(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateSum(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetDistributionsOKBodyRanks) contextValidateFields(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Fields); i++ {

		if o.Fields[i] != nil {
			if err := o.Fields[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getDistributionsOK" + "." + "ranks" + "." + "fields" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetDistributionsOKBodyRanks) contextValidateRows(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Rows); i++ {

		if o.Rows[i] != nil {
			if err := o.Rows[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getDistributionsOK" + "." + "ranks" + "." + "rows" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetDistributionsOKBodyRanks) contextValidateSum(ctx context.Context, formats strfmt.Registry) error {

	if o.Sum != nil {
		if err := o.Sum.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getDistributionsOK" + "." + "ranks" + "." + "sum")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetDistributionsOKBodyRanks) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetDistributionsOKBodyRanks) UnmarshalBinary(b []byte) error {
	var res GetDistributionsOKBodyRanks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetDistributionsOKBodyRanksFieldsItems0 get distributions o k body ranks fields items0
swagger:model GetDistributionsOKBodyRanksFieldsItems0
*/
type GetDistributionsOKBodyRanksFieldsItems0 struct {

	// columnID
	ColumnID int64 `json:"columnID,omitempty"`

	// dataTypeID
	DataTypeID int64 `json:"dataTypeID,omitempty"`

	// dataTypeModifier
	DataTypeModifier string `json:"dataTypeModifier,omitempty"`

	// dataTypeSize
	DataTypeSize int64 `json:"dataTypeSize,omitempty"`

	// format
	Format string `json:"format,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// tableID
	TableID int64 `json:"tableID,omitempty"`
}

// Validate validates this get distributions o k body ranks fields items0
func (o *GetDistributionsOKBodyRanksFieldsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get distributions o k body ranks fields items0 based on context it is used
func (o *GetDistributionsOKBodyRanksFieldsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetDistributionsOKBodyRanksFieldsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetDistributionsOKBodyRanksFieldsItems0) UnmarshalBinary(b []byte) error {
	var res GetDistributionsOKBodyRanksFieldsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetDistributionsOKBodyRanksRowsItems0 get distributions o k body ranks rows items0
swagger:model GetDistributionsOKBodyRanksRowsItems0
*/
type GetDistributionsOKBodyRanksRowsItems0 struct {

	// bin
	Bin int64 `json:"bin,omitempty"`

	// bin_name
	BinName int64 `json:"bin_name,omitempty"`

	// count
	Count int64 `json:"count,omitempty"`

	// cumulative_sum
	CumulativeSum int64 `json:"cumulative_sum,omitempty"`
}

// Validate validates this get distributions o k body ranks rows items0
func (o *GetDistributionsOKBodyRanksRowsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get distributions o k body ranks rows items0 based on context it is used
func (o *GetDistributionsOKBodyRanksRowsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetDistributionsOKBodyRanksRowsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetDistributionsOKBodyRanksRowsItems0) UnmarshalBinary(b []byte) error {
	var res GetDistributionsOKBodyRanksRowsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetDistributionsOKBodyRanksSum sum
swagger:model GetDistributionsOKBodyRanksSum
*/
type GetDistributionsOKBodyRanksSum struct {

	// count
	Count int64 `json:"count,omitempty"`
}

// Validate validates this get distributions o k body ranks sum
func (o *GetDistributionsOKBodyRanksSum) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get distributions o k body ranks sum based on context it is used
func (o *GetDistributionsOKBodyRanksSum) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetDistributionsOKBodyRanksSum) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetDistributionsOKBodyRanksSum) UnmarshalBinary(b []byte) error {
	var res GetDistributionsOKBodyRanksSum
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}