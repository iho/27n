// Code generated by go-swagger; DO NOT EDIT.

package rankings

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
)

// NewGetRankingsParams creates a new GetRankingsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetRankingsParams() *GetRankingsParams {
	return &GetRankingsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetRankingsParamsWithTimeout creates a new GetRankingsParams object
// with the ability to set a timeout on a request.
func NewGetRankingsParamsWithTimeout(timeout time.Duration) *GetRankingsParams {
	return &GetRankingsParams{
		timeout: timeout,
	}
}

// NewGetRankingsParamsWithContext creates a new GetRankingsParams object
// with the ability to set a context for a request.
func NewGetRankingsParamsWithContext(ctx context.Context) *GetRankingsParams {
	return &GetRankingsParams{
		Context: ctx,
	}
}

// NewGetRankingsParamsWithHTTPClient creates a new GetRankingsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetRankingsParamsWithHTTPClient(client *http.Client) *GetRankingsParams {
	return &GetRankingsParams{
		HTTPClient: client,
	}
}

/* GetRankingsParams contains all the parameters to send to the API endpoint
   for the get rankings operation.

   Typically these are written to a http.Request.
*/
type GetRankingsParams struct {

	/* HeroID.

	   Hero ID
	*/
	HeroID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get rankings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRankingsParams) WithDefaults() *GetRankingsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get rankings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRankingsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get rankings params
func (o *GetRankingsParams) WithTimeout(timeout time.Duration) *GetRankingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get rankings params
func (o *GetRankingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get rankings params
func (o *GetRankingsParams) WithContext(ctx context.Context) *GetRankingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get rankings params
func (o *GetRankingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get rankings params
func (o *GetRankingsParams) WithHTTPClient(client *http.Client) *GetRankingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get rankings params
func (o *GetRankingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithHeroID adds the heroID to the get rankings params
func (o *GetRankingsParams) WithHeroID(heroID string) *GetRankingsParams {
	o.SetHeroID(heroID)
	return o
}

// SetHeroID adds the heroId to the get rankings params
func (o *GetRankingsParams) SetHeroID(heroID string) {
	o.HeroID = heroID
}

// WriteToRequest writes these params to a swagger request
func (o *GetRankingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param hero_id
	qrHeroID := o.HeroID
	qHeroID := qrHeroID
	if qHeroID != "" {

		if err := r.SetQueryParam("hero_id", qHeroID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}