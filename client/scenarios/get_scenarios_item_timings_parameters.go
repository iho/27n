// Code generated by go-swagger; DO NOT EDIT.

package scenarios

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

// NewGetScenariosItemTimingsParams creates a new GetScenariosItemTimingsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetScenariosItemTimingsParams() *GetScenariosItemTimingsParams {
	return &GetScenariosItemTimingsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetScenariosItemTimingsParamsWithTimeout creates a new GetScenariosItemTimingsParams object
// with the ability to set a timeout on a request.
func NewGetScenariosItemTimingsParamsWithTimeout(timeout time.Duration) *GetScenariosItemTimingsParams {
	return &GetScenariosItemTimingsParams{
		timeout: timeout,
	}
}

// NewGetScenariosItemTimingsParamsWithContext creates a new GetScenariosItemTimingsParams object
// with the ability to set a context for a request.
func NewGetScenariosItemTimingsParamsWithContext(ctx context.Context) *GetScenariosItemTimingsParams {
	return &GetScenariosItemTimingsParams{
		Context: ctx,
	}
}

// NewGetScenariosItemTimingsParamsWithHTTPClient creates a new GetScenariosItemTimingsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetScenariosItemTimingsParamsWithHTTPClient(client *http.Client) *GetScenariosItemTimingsParams {
	return &GetScenariosItemTimingsParams{
		HTTPClient: client,
	}
}

/* GetScenariosItemTimingsParams contains all the parameters to send to the API endpoint
   for the get scenarios item timings operation.

   Typically these are written to a http.Request.
*/
type GetScenariosItemTimingsParams struct {

	/* HeroID.

	   Hero ID
	*/
	HeroID *int64

	/* Item.

	   Filter by item name e.g. "spirit_vessel"
	*/
	Item *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get scenarios item timings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetScenariosItemTimingsParams) WithDefaults() *GetScenariosItemTimingsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get scenarios item timings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetScenariosItemTimingsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get scenarios item timings params
func (o *GetScenariosItemTimingsParams) WithTimeout(timeout time.Duration) *GetScenariosItemTimingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get scenarios item timings params
func (o *GetScenariosItemTimingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get scenarios item timings params
func (o *GetScenariosItemTimingsParams) WithContext(ctx context.Context) *GetScenariosItemTimingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get scenarios item timings params
func (o *GetScenariosItemTimingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get scenarios item timings params
func (o *GetScenariosItemTimingsParams) WithHTTPClient(client *http.Client) *GetScenariosItemTimingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get scenarios item timings params
func (o *GetScenariosItemTimingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithHeroID adds the heroID to the get scenarios item timings params
func (o *GetScenariosItemTimingsParams) WithHeroID(heroID *int64) *GetScenariosItemTimingsParams {
	o.SetHeroID(heroID)
	return o
}

// SetHeroID adds the heroId to the get scenarios item timings params
func (o *GetScenariosItemTimingsParams) SetHeroID(heroID *int64) {
	o.HeroID = heroID
}

// WithItem adds the item to the get scenarios item timings params
func (o *GetScenariosItemTimingsParams) WithItem(item *string) *GetScenariosItemTimingsParams {
	o.SetItem(item)
	return o
}

// SetItem adds the item to the get scenarios item timings params
func (o *GetScenariosItemTimingsParams) SetItem(item *string) {
	o.Item = item
}

// WriteToRequest writes these params to a swagger request
func (o *GetScenariosItemTimingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.HeroID != nil {

		// query param hero_id
		var qrHeroID int64

		if o.HeroID != nil {
			qrHeroID = *o.HeroID
		}
		qHeroID := swag.FormatInt64(qrHeroID)
		if qHeroID != "" {

			if err := r.SetQueryParam("hero_id", qHeroID); err != nil {
				return err
			}
		}
	}

	if o.Item != nil {

		// query param item
		var qrItem string

		if o.Item != nil {
			qrItem = *o.Item
		}
		qItem := qrItem
		if qItem != "" {

			if err := r.SetQueryParam("item", qItem); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
