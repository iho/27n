// Code generated by go-swagger; DO NOT EDIT.

package players_by_rank

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

// NewGetPlayersByRankParams creates a new GetPlayersByRankParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetPlayersByRankParams() *GetPlayersByRankParams {
	return &GetPlayersByRankParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetPlayersByRankParamsWithTimeout creates a new GetPlayersByRankParams object
// with the ability to set a timeout on a request.
func NewGetPlayersByRankParamsWithTimeout(timeout time.Duration) *GetPlayersByRankParams {
	return &GetPlayersByRankParams{
		timeout: timeout,
	}
}

// NewGetPlayersByRankParamsWithContext creates a new GetPlayersByRankParams object
// with the ability to set a context for a request.
func NewGetPlayersByRankParamsWithContext(ctx context.Context) *GetPlayersByRankParams {
	return &GetPlayersByRankParams{
		Context: ctx,
	}
}

// NewGetPlayersByRankParamsWithHTTPClient creates a new GetPlayersByRankParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetPlayersByRankParamsWithHTTPClient(client *http.Client) *GetPlayersByRankParams {
	return &GetPlayersByRankParams{
		HTTPClient: client,
	}
}

/* GetPlayersByRankParams contains all the parameters to send to the API endpoint
   for the get players by rank operation.

   Typically these are written to a http.Request.
*/
type GetPlayersByRankParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get players by rank params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPlayersByRankParams) WithDefaults() *GetPlayersByRankParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get players by rank params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPlayersByRankParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get players by rank params
func (o *GetPlayersByRankParams) WithTimeout(timeout time.Duration) *GetPlayersByRankParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get players by rank params
func (o *GetPlayersByRankParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get players by rank params
func (o *GetPlayersByRankParams) WithContext(ctx context.Context) *GetPlayersByRankParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get players by rank params
func (o *GetPlayersByRankParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get players by rank params
func (o *GetPlayersByRankParams) WithHTTPClient(client *http.Client) *GetPlayersByRankParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get players by rank params
func (o *GetPlayersByRankParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetPlayersByRankParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}