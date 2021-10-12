// Code generated by go-swagger; DO NOT EDIT.

package heroes

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

// NewGetHeroesHeroIDDurationsParams creates a new GetHeroesHeroIDDurationsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetHeroesHeroIDDurationsParams() *GetHeroesHeroIDDurationsParams {
	return &GetHeroesHeroIDDurationsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetHeroesHeroIDDurationsParamsWithTimeout creates a new GetHeroesHeroIDDurationsParams object
// with the ability to set a timeout on a request.
func NewGetHeroesHeroIDDurationsParamsWithTimeout(timeout time.Duration) *GetHeroesHeroIDDurationsParams {
	return &GetHeroesHeroIDDurationsParams{
		timeout: timeout,
	}
}

// NewGetHeroesHeroIDDurationsParamsWithContext creates a new GetHeroesHeroIDDurationsParams object
// with the ability to set a context for a request.
func NewGetHeroesHeroIDDurationsParamsWithContext(ctx context.Context) *GetHeroesHeroIDDurationsParams {
	return &GetHeroesHeroIDDurationsParams{
		Context: ctx,
	}
}

// NewGetHeroesHeroIDDurationsParamsWithHTTPClient creates a new GetHeroesHeroIDDurationsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetHeroesHeroIDDurationsParamsWithHTTPClient(client *http.Client) *GetHeroesHeroIDDurationsParams {
	return &GetHeroesHeroIDDurationsParams{
		HTTPClient: client,
	}
}

/* GetHeroesHeroIDDurationsParams contains all the parameters to send to the API endpoint
   for the get heroes hero ID durations operation.

   Typically these are written to a http.Request.
*/
type GetHeroesHeroIDDurationsParams struct {

	/* HeroID.

	   Hero ID
	*/
	HeroID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get heroes hero ID durations params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetHeroesHeroIDDurationsParams) WithDefaults() *GetHeroesHeroIDDurationsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get heroes hero ID durations params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetHeroesHeroIDDurationsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get heroes hero ID durations params
func (o *GetHeroesHeroIDDurationsParams) WithTimeout(timeout time.Duration) *GetHeroesHeroIDDurationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get heroes hero ID durations params
func (o *GetHeroesHeroIDDurationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get heroes hero ID durations params
func (o *GetHeroesHeroIDDurationsParams) WithContext(ctx context.Context) *GetHeroesHeroIDDurationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get heroes hero ID durations params
func (o *GetHeroesHeroIDDurationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get heroes hero ID durations params
func (o *GetHeroesHeroIDDurationsParams) WithHTTPClient(client *http.Client) *GetHeroesHeroIDDurationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get heroes hero ID durations params
func (o *GetHeroesHeroIDDurationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithHeroID adds the heroID to the get heroes hero ID durations params
func (o *GetHeroesHeroIDDurationsParams) WithHeroID(heroID int64) *GetHeroesHeroIDDurationsParams {
	o.SetHeroID(heroID)
	return o
}

// SetHeroID adds the heroId to the get heroes hero ID durations params
func (o *GetHeroesHeroIDDurationsParams) SetHeroID(heroID int64) {
	o.HeroID = heroID
}

// WriteToRequest writes these params to a swagger request
func (o *GetHeroesHeroIDDurationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param hero_id
	if err := r.SetPathParam("hero_id", swag.FormatInt64(o.HeroID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
