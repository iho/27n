// Code generated by go-swagger; DO NOT EDIT.

package constants

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new constants API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for constants API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetConstants(params *GetConstantsParams, opts ...ClientOption) (*GetConstantsOK, error)

	GetConstantsResource(params *GetConstantsResourceParams, opts ...ClientOption) (*GetConstantsResourceOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetConstants gs e t constants

  Gets an array of available resources.
*/
func (a *Client) GetConstants(params *GetConstantsParams, opts ...ClientOption) (*GetConstantsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetConstantsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetConstants",
		Method:             "GET",
		PathPattern:        "/constants",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetConstantsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetConstantsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetConstants: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetConstantsResource gs e t constants

  Get static game data mirrored from the dotaconstants repository.
*/
func (a *Client) GetConstantsResource(params *GetConstantsResourceParams, opts ...ClientOption) (*GetConstantsResourceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetConstantsResourceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetConstantsResource",
		Method:             "GET",
		PathPattern:        "/constants/{resource}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetConstantsResourceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetConstantsResourceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetConstantsResource: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
