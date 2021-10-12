// Code generated by go-swagger; DO NOT EDIT.

package search

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

// NewGetSearchParams creates a new GetSearchParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetSearchParams() *GetSearchParams {
	return &GetSearchParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetSearchParamsWithTimeout creates a new GetSearchParams object
// with the ability to set a timeout on a request.
func NewGetSearchParamsWithTimeout(timeout time.Duration) *GetSearchParams {
	return &GetSearchParams{
		timeout: timeout,
	}
}

// NewGetSearchParamsWithContext creates a new GetSearchParams object
// with the ability to set a context for a request.
func NewGetSearchParamsWithContext(ctx context.Context) *GetSearchParams {
	return &GetSearchParams{
		Context: ctx,
	}
}

// NewGetSearchParamsWithHTTPClient creates a new GetSearchParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetSearchParamsWithHTTPClient(client *http.Client) *GetSearchParams {
	return &GetSearchParams{
		HTTPClient: client,
	}
}

/* GetSearchParams contains all the parameters to send to the API endpoint
   for the get search operation.

   Typically these are written to a http.Request.
*/
type GetSearchParams struct {

	/* Q.

	   Search string
	*/
	Q string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get search params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSearchParams) WithDefaults() *GetSearchParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get search params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSearchParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get search params
func (o *GetSearchParams) WithTimeout(timeout time.Duration) *GetSearchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get search params
func (o *GetSearchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get search params
func (o *GetSearchParams) WithContext(ctx context.Context) *GetSearchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get search params
func (o *GetSearchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get search params
func (o *GetSearchParams) WithHTTPClient(client *http.Client) *GetSearchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get search params
func (o *GetSearchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithQ adds the q to the get search params
func (o *GetSearchParams) WithQ(q string) *GetSearchParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the get search params
func (o *GetSearchParams) SetQ(q string) {
	o.Q = q
}

// WriteToRequest writes these params to a swagger request
func (o *GetSearchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param q
	qrQ := o.Q
	qQ := qrQ
	if qQ != "" {

		if err := r.SetQueryParam("q", qQ); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
