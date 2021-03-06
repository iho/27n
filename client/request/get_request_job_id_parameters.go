// Code generated by go-swagger; DO NOT EDIT.

package request

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

// NewGetRequestJobIDParams creates a new GetRequestJobIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetRequestJobIDParams() *GetRequestJobIDParams {
	return &GetRequestJobIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetRequestJobIDParamsWithTimeout creates a new GetRequestJobIDParams object
// with the ability to set a timeout on a request.
func NewGetRequestJobIDParamsWithTimeout(timeout time.Duration) *GetRequestJobIDParams {
	return &GetRequestJobIDParams{
		timeout: timeout,
	}
}

// NewGetRequestJobIDParamsWithContext creates a new GetRequestJobIDParams object
// with the ability to set a context for a request.
func NewGetRequestJobIDParamsWithContext(ctx context.Context) *GetRequestJobIDParams {
	return &GetRequestJobIDParams{
		Context: ctx,
	}
}

// NewGetRequestJobIDParamsWithHTTPClient creates a new GetRequestJobIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetRequestJobIDParamsWithHTTPClient(client *http.Client) *GetRequestJobIDParams {
	return &GetRequestJobIDParams{
		HTTPClient: client,
	}
}

/* GetRequestJobIDParams contains all the parameters to send to the API endpoint
   for the get request job ID operation.

   Typically these are written to a http.Request.
*/
type GetRequestJobIDParams struct {

	/* JobID.

	   The job ID to query.
	*/
	JobID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get request job ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRequestJobIDParams) WithDefaults() *GetRequestJobIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get request job ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRequestJobIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get request job ID params
func (o *GetRequestJobIDParams) WithTimeout(timeout time.Duration) *GetRequestJobIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get request job ID params
func (o *GetRequestJobIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get request job ID params
func (o *GetRequestJobIDParams) WithContext(ctx context.Context) *GetRequestJobIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get request job ID params
func (o *GetRequestJobIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get request job ID params
func (o *GetRequestJobIDParams) WithHTTPClient(client *http.Client) *GetRequestJobIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get request job ID params
func (o *GetRequestJobIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithJobID adds the jobID to the get request job ID params
func (o *GetRequestJobIDParams) WithJobID(jobID string) *GetRequestJobIDParams {
	o.SetJobID(jobID)
	return o
}

// SetJobID adds the jobId to the get request job ID params
func (o *GetRequestJobIDParams) SetJobID(jobID string) {
	o.JobID = jobID
}

// WriteToRequest writes these params to a swagger request
func (o *GetRequestJobIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param jobId
	if err := r.SetPathParam("jobId", o.JobID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
