// Code generated by go-swagger; DO NOT EDIT.

package leagues

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

// NewGetLeaguesLeagueIDTeamsParams creates a new GetLeaguesLeagueIDTeamsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetLeaguesLeagueIDTeamsParams() *GetLeaguesLeagueIDTeamsParams {
	return &GetLeaguesLeagueIDTeamsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetLeaguesLeagueIDTeamsParamsWithTimeout creates a new GetLeaguesLeagueIDTeamsParams object
// with the ability to set a timeout on a request.
func NewGetLeaguesLeagueIDTeamsParamsWithTimeout(timeout time.Duration) *GetLeaguesLeagueIDTeamsParams {
	return &GetLeaguesLeagueIDTeamsParams{
		timeout: timeout,
	}
}

// NewGetLeaguesLeagueIDTeamsParamsWithContext creates a new GetLeaguesLeagueIDTeamsParams object
// with the ability to set a context for a request.
func NewGetLeaguesLeagueIDTeamsParamsWithContext(ctx context.Context) *GetLeaguesLeagueIDTeamsParams {
	return &GetLeaguesLeagueIDTeamsParams{
		Context: ctx,
	}
}

// NewGetLeaguesLeagueIDTeamsParamsWithHTTPClient creates a new GetLeaguesLeagueIDTeamsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetLeaguesLeagueIDTeamsParamsWithHTTPClient(client *http.Client) *GetLeaguesLeagueIDTeamsParams {
	return &GetLeaguesLeagueIDTeamsParams{
		HTTPClient: client,
	}
}

/* GetLeaguesLeagueIDTeamsParams contains all the parameters to send to the API endpoint
   for the get leagues league ID teams operation.

   Typically these are written to a http.Request.
*/
type GetLeaguesLeagueIDTeamsParams struct {

	/* LeagueID.

	   League ID
	*/
	LeagueID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get leagues league ID teams params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLeaguesLeagueIDTeamsParams) WithDefaults() *GetLeaguesLeagueIDTeamsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get leagues league ID teams params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLeaguesLeagueIDTeamsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get leagues league ID teams params
func (o *GetLeaguesLeagueIDTeamsParams) WithTimeout(timeout time.Duration) *GetLeaguesLeagueIDTeamsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get leagues league ID teams params
func (o *GetLeaguesLeagueIDTeamsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get leagues league ID teams params
func (o *GetLeaguesLeagueIDTeamsParams) WithContext(ctx context.Context) *GetLeaguesLeagueIDTeamsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get leagues league ID teams params
func (o *GetLeaguesLeagueIDTeamsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get leagues league ID teams params
func (o *GetLeaguesLeagueIDTeamsParams) WithHTTPClient(client *http.Client) *GetLeaguesLeagueIDTeamsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get leagues league ID teams params
func (o *GetLeaguesLeagueIDTeamsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLeagueID adds the leagueID to the get leagues league ID teams params
func (o *GetLeaguesLeagueIDTeamsParams) WithLeagueID(leagueID int64) *GetLeaguesLeagueIDTeamsParams {
	o.SetLeagueID(leagueID)
	return o
}

// SetLeagueID adds the leagueId to the get leagues league ID teams params
func (o *GetLeaguesLeagueIDTeamsParams) SetLeagueID(leagueID int64) {
	o.LeagueID = leagueID
}

// WriteToRequest writes these params to a swagger request
func (o *GetLeaguesLeagueIDTeamsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param league_id
	if err := r.SetPathParam("league_id", swag.FormatInt64(o.LeagueID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
