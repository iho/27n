// Code generated by go-swagger; DO NOT EDIT.

package explorer

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

// NewGetExplorerParams creates a new GetExplorerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetExplorerParams() *GetExplorerParams {
	return &GetExplorerParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetExplorerParamsWithTimeout creates a new GetExplorerParams object
// with the ability to set a timeout on a request.
func NewGetExplorerParamsWithTimeout(timeout time.Duration) *GetExplorerParams {
	return &GetExplorerParams{
		timeout: timeout,
	}
}

// NewGetExplorerParamsWithContext creates a new GetExplorerParams object
// with the ability to set a context for a request.
func NewGetExplorerParamsWithContext(ctx context.Context) *GetExplorerParams {
	return &GetExplorerParams{
		Context: ctx,
	}
}

// NewGetExplorerParamsWithHTTPClient creates a new GetExplorerParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetExplorerParamsWithHTTPClient(client *http.Client) *GetExplorerParams {
	return &GetExplorerParams{
		HTTPClient: client,
	}
}

/* GetExplorerParams contains all the parameters to send to the API endpoint
   for the get explorer operation.

   Typically these are written to a http.Request.
*/
type GetExplorerParams struct {

	/* SQL.

	   The PostgreSQL query as percent-encoded string.
	*/
	SQL *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get explorer params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetExplorerParams) WithDefaults() *GetExplorerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get explorer params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetExplorerParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get explorer params
func (o *GetExplorerParams) WithTimeout(timeout time.Duration) *GetExplorerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get explorer params
func (o *GetExplorerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get explorer params
func (o *GetExplorerParams) WithContext(ctx context.Context) *GetExplorerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get explorer params
func (o *GetExplorerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get explorer params
func (o *GetExplorerParams) WithHTTPClient(client *http.Client) *GetExplorerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get explorer params
func (o *GetExplorerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSQL adds the sql to the get explorer params
func (o *GetExplorerParams) WithSQL(sql *string) *GetExplorerParams {
	o.SetSQL(sql)
	return o
}

// SetSQL adds the sql to the get explorer params
func (o *GetExplorerParams) SetSQL(sql *string) {
	o.SQL = sql
}

// WriteToRequest writes these params to a swagger request
func (o *GetExplorerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.SQL != nil {

		// query param sql
		var qrSQL string

		if o.SQL != nil {
			qrSQL = *o.SQL
		}
		qSQL := qrSQL
		if qSQL != "" {

			if err := r.SetQueryParam("sql", qSQL); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
