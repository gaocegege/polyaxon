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

// RunsV1InvalidateRunsReader is a Reader for the RunsV1InvalidateRuns structure.
type RunsV1InvalidateRunsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RunsV1InvalidateRunsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRunsV1InvalidateRunsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewRunsV1InvalidateRunsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewRunsV1InvalidateRunsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRunsV1InvalidateRunsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewRunsV1InvalidateRunsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRunsV1InvalidateRunsOK creates a RunsV1InvalidateRunsOK with default headers values
func NewRunsV1InvalidateRunsOK() *RunsV1InvalidateRunsOK {
	return &RunsV1InvalidateRunsOK{}
}

/*RunsV1InvalidateRunsOK handles this case with default header values.

A successful response.
*/
type RunsV1InvalidateRunsOK struct {
}

func (o *RunsV1InvalidateRunsOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{project}/runs/invalidate][%d] runsV1InvalidateRunsOK ", 200)
}

func (o *RunsV1InvalidateRunsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRunsV1InvalidateRunsNoContent creates a RunsV1InvalidateRunsNoContent with default headers values
func NewRunsV1InvalidateRunsNoContent() *RunsV1InvalidateRunsNoContent {
	return &RunsV1InvalidateRunsNoContent{}
}

/*RunsV1InvalidateRunsNoContent handles this case with default header values.

No content.
*/
type RunsV1InvalidateRunsNoContent struct {
	Payload interface{}
}

func (o *RunsV1InvalidateRunsNoContent) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{project}/runs/invalidate][%d] runsV1InvalidateRunsNoContent  %+v", 204, o.Payload)
}

func (o *RunsV1InvalidateRunsNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *RunsV1InvalidateRunsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRunsV1InvalidateRunsForbidden creates a RunsV1InvalidateRunsForbidden with default headers values
func NewRunsV1InvalidateRunsForbidden() *RunsV1InvalidateRunsForbidden {
	return &RunsV1InvalidateRunsForbidden{}
}

/*RunsV1InvalidateRunsForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type RunsV1InvalidateRunsForbidden struct {
	Payload interface{}
}

func (o *RunsV1InvalidateRunsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{project}/runs/invalidate][%d] runsV1InvalidateRunsForbidden  %+v", 403, o.Payload)
}

func (o *RunsV1InvalidateRunsForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *RunsV1InvalidateRunsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRunsV1InvalidateRunsNotFound creates a RunsV1InvalidateRunsNotFound with default headers values
func NewRunsV1InvalidateRunsNotFound() *RunsV1InvalidateRunsNotFound {
	return &RunsV1InvalidateRunsNotFound{}
}

/*RunsV1InvalidateRunsNotFound handles this case with default header values.

Resource does not exist.
*/
type RunsV1InvalidateRunsNotFound struct {
	Payload interface{}
}

func (o *RunsV1InvalidateRunsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{project}/runs/invalidate][%d] runsV1InvalidateRunsNotFound  %+v", 404, o.Payload)
}

func (o *RunsV1InvalidateRunsNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *RunsV1InvalidateRunsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRunsV1InvalidateRunsDefault creates a RunsV1InvalidateRunsDefault with default headers values
func NewRunsV1InvalidateRunsDefault(code int) *RunsV1InvalidateRunsDefault {
	return &RunsV1InvalidateRunsDefault{
		_statusCode: code,
	}
}

/*RunsV1InvalidateRunsDefault handles this case with default header values.

An unexpected error response
*/
type RunsV1InvalidateRunsDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the runs v1 invalidate runs default response
func (o *RunsV1InvalidateRunsDefault) Code() int {
	return o._statusCode
}

func (o *RunsV1InvalidateRunsDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{project}/runs/invalidate][%d] RunsV1_InvalidateRuns default  %+v", o._statusCode, o.Payload)
}

func (o *RunsV1InvalidateRunsDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *RunsV1InvalidateRunsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}