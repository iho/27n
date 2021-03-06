// Code generated by go-swagger; DO NOT EDIT.

package pro_players

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

// NewGetProPlayersParams creates a new GetProPlayersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetProPlayersParams() *GetProPlayersParams {
	return &GetProPlayersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetProPlayersParamsWithTimeout creates a new GetProPlayersParams object
// with the ability to set a timeout on a request.
func NewGetProPlayersParamsWithTimeout(timeout time.Duration) *GetProPlayersParams {
	return &GetProPlayersParams{
		timeout: timeout,
	}
}

// NewGetProPlayersParamsWithContext creates a new GetProPlayersParams object
// with the ability to set a context for a request.
func NewGetProPlayersParamsWithContext(ctx context.Context) *GetProPlayersParams {
	return &GetProPlayersParams{
		Context: ctx,
	}
}

// NewGetProPlayersParamsWithHTTPClient creates a new GetProPlayersParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetProPlayersParamsWithHTTPClient(client *http.Client) *GetProPlayersParams {
	return &GetProPlayersParams{
		HTTPClient: client,
	}
}

/* GetProPlayersParams contains all the parameters to send to the API endpoint
   for the get pro players operation.

   Typically these are written to a http.Request.
*/
type GetProPlayersParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get pro players params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProPlayersParams) WithDefaults() *GetProPlayersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get pro players params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProPlayersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get pro players params
func (o *GetProPlayersParams) WithTimeout(timeout time.Duration) *GetProPlayersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get pro players params
func (o *GetProPlayersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get pro players params
func (o *GetProPlayersParams) WithContext(ctx context.Context) *GetProPlayersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get pro players params
func (o *GetProPlayersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get pro players params
func (o *GetProPlayersParams) WithHTTPClient(client *http.Client) *GetProPlayersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get pro players params
func (o *GetProPlayersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetProPlayersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
