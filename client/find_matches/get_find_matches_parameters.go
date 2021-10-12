// Code generated by go-swagger; DO NOT EDIT.

package find_matches

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

// NewGetFindMatchesParams creates a new GetFindMatchesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetFindMatchesParams() *GetFindMatchesParams {
	return &GetFindMatchesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetFindMatchesParamsWithTimeout creates a new GetFindMatchesParams object
// with the ability to set a timeout on a request.
func NewGetFindMatchesParamsWithTimeout(timeout time.Duration) *GetFindMatchesParams {
	return &GetFindMatchesParams{
		timeout: timeout,
	}
}

// NewGetFindMatchesParamsWithContext creates a new GetFindMatchesParams object
// with the ability to set a context for a request.
func NewGetFindMatchesParamsWithContext(ctx context.Context) *GetFindMatchesParams {
	return &GetFindMatchesParams{
		Context: ctx,
	}
}

// NewGetFindMatchesParamsWithHTTPClient creates a new GetFindMatchesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetFindMatchesParamsWithHTTPClient(client *http.Client) *GetFindMatchesParams {
	return &GetFindMatchesParams{
		HTTPClient: client,
	}
}

/* GetFindMatchesParams contains all the parameters to send to the API endpoint
   for the get find matches operation.

   Typically these are written to a http.Request.
*/
type GetFindMatchesParams struct {

	/* TeamA.

	   Hero IDs on first team (array)
	*/
	TeamA *int64

	/* TeamB.

	   Hero IDs on second team (array)
	*/
	TeamB *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get find matches params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetFindMatchesParams) WithDefaults() *GetFindMatchesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get find matches params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetFindMatchesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get find matches params
func (o *GetFindMatchesParams) WithTimeout(timeout time.Duration) *GetFindMatchesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get find matches params
func (o *GetFindMatchesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get find matches params
func (o *GetFindMatchesParams) WithContext(ctx context.Context) *GetFindMatchesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get find matches params
func (o *GetFindMatchesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get find matches params
func (o *GetFindMatchesParams) WithHTTPClient(client *http.Client) *GetFindMatchesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get find matches params
func (o *GetFindMatchesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTeamA adds the teamA to the get find matches params
func (o *GetFindMatchesParams) WithTeamA(teamA *int64) *GetFindMatchesParams {
	o.SetTeamA(teamA)
	return o
}

// SetTeamA adds the teamA to the get find matches params
func (o *GetFindMatchesParams) SetTeamA(teamA *int64) {
	o.TeamA = teamA
}

// WithTeamB adds the teamB to the get find matches params
func (o *GetFindMatchesParams) WithTeamB(teamB *int64) *GetFindMatchesParams {
	o.SetTeamB(teamB)
	return o
}

// SetTeamB adds the teamB to the get find matches params
func (o *GetFindMatchesParams) SetTeamB(teamB *int64) {
	o.TeamB = teamB
}

// WriteToRequest writes these params to a swagger request
func (o *GetFindMatchesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.TeamA != nil {

		// query param teamA
		var qrTeamA int64

		if o.TeamA != nil {
			qrTeamA = *o.TeamA
		}
		qTeamA := swag.FormatInt64(qrTeamA)
		if qTeamA != "" {

			if err := r.SetQueryParam("teamA", qTeamA); err != nil {
				return err
			}
		}
	}

	if o.TeamB != nil {

		// query param teamB
		var qrTeamB int64

		if o.TeamB != nil {
			qrTeamB = *o.TeamB
		}
		qTeamB := swag.FormatInt64(qrTeamB)
		if qTeamB != "" {

			if err := r.SetQueryParam("teamB", qTeamB); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}