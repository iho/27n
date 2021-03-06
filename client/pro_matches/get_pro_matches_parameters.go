// Code generated by go-swagger; DO NOT EDIT.

package pro_matches

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetProMatchesParams creates a new GetProMatchesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetProMatchesParams() *GetProMatchesParams {
	return &GetProMatchesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetProMatchesParamsWithTimeout creates a new GetProMatchesParams object
// with the ability to set a timeout on a request.
func NewGetProMatchesParamsWithTimeout(timeout time.Duration) *GetProMatchesParams {
	return &GetProMatchesParams{
		timeout: timeout,
	}
}

// NewGetProMatchesParamsWithContext creates a new GetProMatchesParams object
// with the ability to set a context for a request.
func NewGetProMatchesParamsWithContext(ctx context.Context) *GetProMatchesParams {
	return &GetProMatchesParams{
		Context: ctx,
	}
}

// NewGetProMatchesParamsWithHTTPClient creates a new GetProMatchesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetProMatchesParamsWithHTTPClient(client *http.Client) *GetProMatchesParams {
	return &GetProMatchesParams{
		HTTPClient: client,
	}
}

/* GetProMatchesParams contains all the parameters to send to the API endpoint
   for the get pro matches operation.

   Typically these are written to a http.Request.
*/
type GetProMatchesParams struct {

	/* LessThanMatchID.

	   Get matches with a match ID lower than this value
	*/
	LessThanMatchID *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get pro matches params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProMatchesParams) WithDefaults() *GetProMatchesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get pro matches params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProMatchesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get pro matches params
func (o *GetProMatchesParams) WithTimeout(timeout time.Duration) *GetProMatchesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get pro matches params
func (o *GetProMatchesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get pro matches params
func (o *GetProMatchesParams) WithContext(ctx context.Context) *GetProMatchesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get pro matches params
func (o *GetProMatchesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get pro matches params
func (o *GetProMatchesParams) WithHTTPClient(client *http.Client) *GetProMatchesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get pro matches params
func (o *GetProMatchesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLessThanMatchID adds the lessThanMatchID to the get pro matches params
func (o *GetProMatchesParams) WithLessThanMatchID(lessThanMatchID *int64) *GetProMatchesParams {
	o.SetLessThanMatchID(lessThanMatchID)
	return o
}

// SetLessThanMatchID adds the lessThanMatchId to the get pro matches params
func (o *GetProMatchesParams) SetLessThanMatchID(lessThanMatchID *int64) {
	o.LessThanMatchID = lessThanMatchID
}

// WriteToRequest writes these params to a swagger request
func (o *GetProMatchesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.LessThanMatchID != nil {

		// query param less_than_match_id
		var qrLessThanMatchID int64

		if o.LessThanMatchID != nil {
			qrLessThanMatchID = *o.LessThanMatchID
		}
		qLessThanMatchID := swag.FormatInt64(qrLessThanMatchID)
		if qLessThanMatchID != "" {

			if err := r.SetQueryParam("less_than_match_id", qLessThanMatchID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
