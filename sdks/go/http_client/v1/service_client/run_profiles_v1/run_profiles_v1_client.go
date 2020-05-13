// Copyright 2018-2020 Polyaxon, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by go-swagger; DO NOT EDIT.

package run_profiles_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new run profiles v1 API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for run profiles v1 API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateRunProfile(params *CreateRunProfileParams, authInfo runtime.ClientAuthInfoWriter) (*CreateRunProfileOK, *CreateRunProfileNoContent, error)

	DeleteRunProfile(params *DeleteRunProfileParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteRunProfileOK, *DeleteRunProfileNoContent, error)

	GetRunProfile(params *GetRunProfileParams, authInfo runtime.ClientAuthInfoWriter) (*GetRunProfileOK, *GetRunProfileNoContent, error)

	ListRunProfileNames(params *ListRunProfileNamesParams, authInfo runtime.ClientAuthInfoWriter) (*ListRunProfileNamesOK, *ListRunProfileNamesNoContent, error)

	ListRunProfiles(params *ListRunProfilesParams, authInfo runtime.ClientAuthInfoWriter) (*ListRunProfilesOK, *ListRunProfilesNoContent, error)

	PatchRunProfile(params *PatchRunProfileParams, authInfo runtime.ClientAuthInfoWriter) (*PatchRunProfileOK, *PatchRunProfileNoContent, error)

	UpdateRunProfile(params *UpdateRunProfileParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateRunProfileOK, *UpdateRunProfileNoContent, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateRunProfile creates run profile
*/
func (a *Client) CreateRunProfile(params *CreateRunProfileParams, authInfo runtime.ClientAuthInfoWriter) (*CreateRunProfileOK, *CreateRunProfileNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateRunProfileParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateRunProfile",
		Method:             "POST",
		PathPattern:        "/api/v1/orgs/{owner}/run_profiles",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateRunProfileReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *CreateRunProfileOK:
		return value, nil, nil
	case *CreateRunProfileNoContent:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateRunProfileDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DeleteRunProfile deletes run profile
*/
func (a *Client) DeleteRunProfile(params *DeleteRunProfileParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteRunProfileOK, *DeleteRunProfileNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteRunProfileParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteRunProfile",
		Method:             "DELETE",
		PathPattern:        "/api/v1/orgs/{owner}/run_profiles/{uuid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteRunProfileReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *DeleteRunProfileOK:
		return value, nil, nil
	case *DeleteRunProfileNoContent:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteRunProfileDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetRunProfile gets run profile
*/
func (a *Client) GetRunProfile(params *GetRunProfileParams, authInfo runtime.ClientAuthInfoWriter) (*GetRunProfileOK, *GetRunProfileNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRunProfileParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetRunProfile",
		Method:             "GET",
		PathPattern:        "/api/v1/orgs/{owner}/run_profiles/{uuid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetRunProfileReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *GetRunProfileOK:
		return value, nil, nil
	case *GetRunProfileNoContent:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetRunProfileDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ListRunProfileNames lists run profiles names
*/
func (a *Client) ListRunProfileNames(params *ListRunProfileNamesParams, authInfo runtime.ClientAuthInfoWriter) (*ListRunProfileNamesOK, *ListRunProfileNamesNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListRunProfileNamesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListRunProfileNames",
		Method:             "GET",
		PathPattern:        "/api/v1/orgs/{owner}/run_profiles/names",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListRunProfileNamesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *ListRunProfileNamesOK:
		return value, nil, nil
	case *ListRunProfileNamesNoContent:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListRunProfileNamesDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ListRunProfiles lists run profiles
*/
func (a *Client) ListRunProfiles(params *ListRunProfilesParams, authInfo runtime.ClientAuthInfoWriter) (*ListRunProfilesOK, *ListRunProfilesNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListRunProfilesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListRunProfiles",
		Method:             "GET",
		PathPattern:        "/api/v1/orgs/{owner}/run_profiles",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListRunProfilesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *ListRunProfilesOK:
		return value, nil, nil
	case *ListRunProfilesNoContent:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListRunProfilesDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  PatchRunProfile patches run profile
*/
func (a *Client) PatchRunProfile(params *PatchRunProfileParams, authInfo runtime.ClientAuthInfoWriter) (*PatchRunProfileOK, *PatchRunProfileNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchRunProfileParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PatchRunProfile",
		Method:             "PATCH",
		PathPattern:        "/api/v1/orgs/{owner}/run_profiles/{run_profile.uuid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PatchRunProfileReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *PatchRunProfileOK:
		return value, nil, nil
	case *PatchRunProfileNoContent:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PatchRunProfileDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  UpdateRunProfile updates run profile
*/
func (a *Client) UpdateRunProfile(params *UpdateRunProfileParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateRunProfileOK, *UpdateRunProfileNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateRunProfileParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UpdateRunProfile",
		Method:             "PUT",
		PathPattern:        "/api/v1/orgs/{owner}/run_profiles/{run_profile.uuid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateRunProfileReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *UpdateRunProfileOK:
		return value, nil, nil
	case *UpdateRunProfileNoContent:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateRunProfileDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
