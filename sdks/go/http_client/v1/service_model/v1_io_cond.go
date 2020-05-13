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

package service_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1IoCond Input condition specification
//
// swagger:model v1IoCond
type V1IoCond struct {

	// Kind of condition, should be equal to "io"
	Kind string `json:"kind,omitempty"`

	// Param to condition on: e.g. builds.outputs.param1
	Param string `json:"param,omitempty"`

	// Trigger to check, trigger condition must conform to the iotype, e.g. value condition: v1|v2|V3
	Trigger string `json:"trigger,omitempty"`
}

// Validate validates this v1 io cond
func (m *V1IoCond) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1IoCond) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1IoCond) UnmarshalBinary(b []byte) error {
	var res V1IoCond
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
