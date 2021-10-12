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

// NewGetHeroesHeroIDPlayersParams creates a new GetHeroesHeroIDPlayersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetHeroesHeroIDPlayersParams() *GetHeroesHeroIDPlayersParams {
	return &GetHeroesHeroIDPlayersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetHeroesHeroIDPlayersParamsWithTimeout creates a new GetHeroesHeroIDPlayersParams object
// with the ability to set a timeout on a request.
func NewGetHeroesHeroIDPlayersParamsWithTimeout(timeout time.Duration) *GetHeroesHeroIDPlayersParams {
	return &GetHeroesHeroIDPlayersParams{
		timeout: timeout,
	}
}

// NewGetHeroesHeroIDPlayersParamsWithContext creates a new GetHeroesHeroIDPlayersParams object
// with the ability to set a context for a request.
func NewGetHeroesHeroIDPlayersParamsWithContext(ctx context.Context) *GetHeroesHeroIDPlayersParams {
	return &GetHeroesHeroIDPlayersParams{
		Context: ctx,
	}
}

// NewGetHeroesHeroIDPlayersParamsWithHTTPClient creates a new GetHeroesHeroIDPlayersParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetHeroesHeroIDPlayersParamsWithHTTPClient(client *http.Client) *GetHeroesHeroIDPlayersParams {
	return &GetHeroesHeroIDPlayersParams{
		HTTPClient: client,
	}
}

/* GetHeroesHeroIDPlayersParams contains all the parameters to send to the API endpoint
   for the get heroes hero ID players operation.

   Typically these are written to a http.Request.
*/
type GetHeroesHeroIDPlayersParams struct {

	/* HeroID.

	   Hero ID
	*/
	HeroID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get heroes hero ID players params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetHeroesHeroIDPlayersParams) WithDefaults() *GetHeroesHeroIDPlayersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get heroes hero ID players params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetHeroesHeroIDPlayersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get heroes hero ID players params
func (o *GetHeroesHeroIDPlayersParams) WithTimeout(timeout time.Duration) *GetHeroesHeroIDPlayersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get heroes hero ID players params
func (o *GetHeroesHeroIDPlayersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get heroes hero ID players params
func (o *GetHeroesHeroIDPlayersParams) WithContext(ctx context.Context) *GetHeroesHeroIDPlayersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get heroes hero ID players params
func (o *GetHeroesHeroIDPlayersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get heroes hero ID players params
func (o *GetHeroesHeroIDPlayersParams) WithHTTPClient(client *http.Client) *GetHeroesHeroIDPlayersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get heroes hero ID players params
func (o *GetHeroesHeroIDPlayersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithHeroID adds the heroID to the get heroes hero ID players params
func (o *GetHeroesHeroIDPlayersParams) WithHeroID(heroID int64) *GetHeroesHeroIDPlayersParams {
	o.SetHeroID(heroID)
	return o
}

// SetHeroID adds the heroId to the get heroes hero ID players params
func (o *GetHeroesHeroIDPlayersParams) SetHeroID(heroID int64) {
	o.HeroID = heroID
}

// WriteToRequest writes these params to a swagger request
func (o *GetHeroesHeroIDPlayersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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