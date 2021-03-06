// Code generated by go-swagger; DO NOT EDIT.

package constants

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

// NewGetConstantsParams creates a new GetConstantsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetConstantsParams() *GetConstantsParams {
	return &GetConstantsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetConstantsParamsWithTimeout creates a new GetConstantsParams object
// with the ability to set a timeout on a request.
func NewGetConstantsParamsWithTimeout(timeout time.Duration) *GetConstantsParams {
	return &GetConstantsParams{
		timeout: timeout,
	}
}

// NewGetConstantsParamsWithContext creates a new GetConstantsParams object
// with the ability to set a context for a request.
func NewGetConstantsParamsWithContext(ctx context.Context) *GetConstantsParams {
	return &GetConstantsParams{
		Context: ctx,
	}
}

// NewGetConstantsParamsWithHTTPClient creates a new GetConstantsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetConstantsParamsWithHTTPClient(client *http.Client) *GetConstantsParams {
	return &GetConstantsParams{
		HTTPClient: client,
	}
}

/* GetConstantsParams contains all the parameters to send to the API endpoint
   for the get constants operation.

   Typically these are written to a http.Request.
*/
type GetConstantsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get constants params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetConstantsParams) WithDefaults() *GetConstantsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get constants params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetConstantsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get constants params
func (o *GetConstantsParams) WithTimeout(timeout time.Duration) *GetConstantsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get constants params
func (o *GetConstantsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get constants params
func (o *GetConstantsParams) WithContext(ctx context.Context) *GetConstantsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get constants params
func (o *GetConstantsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get constants params
func (o *GetConstantsParams) WithHTTPClient(client *http.Client) *GetConstantsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get constants params
func (o *GetConstantsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetConstantsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
