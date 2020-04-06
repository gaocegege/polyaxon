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

package runs_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// RunsV1DeleteRunsReader is a Reader for the RunsV1DeleteRuns structure.
type RunsV1DeleteRunsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RunsV1DeleteRunsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRunsV1DeleteRunsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewRunsV1DeleteRunsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewRunsV1DeleteRunsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRunsV1DeleteRunsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewRunsV1DeleteRunsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRunsV1DeleteRunsOK creates a RunsV1DeleteRunsOK with default headers values
func NewRunsV1DeleteRunsOK() *RunsV1DeleteRunsOK {
	return &RunsV1DeleteRunsOK{}
}

/*RunsV1DeleteRunsOK handles this case with default header values.

A successful response.
*/
type RunsV1DeleteRunsOK struct {
}

func (o *RunsV1DeleteRunsOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/{owner}/{project}/runs/delete][%d] runsV1DeleteRunsOK ", 200)
}

func (o *RunsV1DeleteRunsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRunsV1DeleteRunsNoContent creates a RunsV1DeleteRunsNoContent with default headers values
func NewRunsV1DeleteRunsNoContent() *RunsV1DeleteRunsNoContent {
	return &RunsV1DeleteRunsNoContent{}
}

/*RunsV1DeleteRunsNoContent handles this case with default header values.

No content.
*/
type RunsV1DeleteRunsNoContent struct {
	Payload interface{}
}

func (o *RunsV1DeleteRunsNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/{owner}/{project}/runs/delete][%d] runsV1DeleteRunsNoContent  %+v", 204, o.Payload)
}

func (o *RunsV1DeleteRunsNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *RunsV1DeleteRunsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRunsV1DeleteRunsForbidden creates a RunsV1DeleteRunsForbidden with default headers values
func NewRunsV1DeleteRunsForbidden() *RunsV1DeleteRunsForbidden {
	return &RunsV1DeleteRunsForbidden{}
}

/*RunsV1DeleteRunsForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type RunsV1DeleteRunsForbidden struct {
	Payload interface{}
}

func (o *RunsV1DeleteRunsForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/{owner}/{project}/runs/delete][%d] runsV1DeleteRunsForbidden  %+v", 403, o.Payload)
}

func (o *RunsV1DeleteRunsForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *RunsV1DeleteRunsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRunsV1DeleteRunsNotFound creates a RunsV1DeleteRunsNotFound with default headers values
func NewRunsV1DeleteRunsNotFound() *RunsV1DeleteRunsNotFound {
	return &RunsV1DeleteRunsNotFound{}
}

/*RunsV1DeleteRunsNotFound handles this case with default header values.

Resource does not exist.
*/
type RunsV1DeleteRunsNotFound struct {
	Payload interface{}
}

func (o *RunsV1DeleteRunsNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/{owner}/{project}/runs/delete][%d] runsV1DeleteRunsNotFound  %+v", 404, o.Payload)
}

func (o *RunsV1DeleteRunsNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *RunsV1DeleteRunsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRunsV1DeleteRunsDefault creates a RunsV1DeleteRunsDefault with default headers values
func NewRunsV1DeleteRunsDefault(code int) *RunsV1DeleteRunsDefault {
	return &RunsV1DeleteRunsDefault{
		_statusCode: code,
	}
}

/*RunsV1DeleteRunsDefault handles this case with default header values.

An unexpected error response
*/
type RunsV1DeleteRunsDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the runs v1 delete runs default response
func (o *RunsV1DeleteRunsDefault) Code() int {
	return o._statusCode
}

func (o *RunsV1DeleteRunsDefault) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/{owner}/{project}/runs/delete][%d] RunsV1_DeleteRuns default  %+v", o._statusCode, o.Payload)
}

func (o *RunsV1DeleteRunsDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *RunsV1DeleteRunsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}