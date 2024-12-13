// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated from the elasticsearch-specification DO NOT EDIT.
// https://github.com/elastic/elasticsearch-specification/tree/4fcf747dfafc951e1dcf3077327e3dcee9107db3

package types

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/dfiindependencemeasure"
)

// SettingsSimilarityDfi type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4fcf747dfafc951e1dcf3077327e3dcee9107db3/specification/indices/_types/IndexSettings.ts#L195-L198
type SettingsSimilarityDfi struct {
	IndependenceMeasure dfiindependencemeasure.DFIIndependenceMeasure `json:"independence_measure"`
	Type                string                                        `json:"type,omitempty"`
}

// MarshalJSON override marshalling to include literal value
func (s SettingsSimilarityDfi) MarshalJSON() ([]byte, error) {
	type innerSettingsSimilarityDfi SettingsSimilarityDfi
	tmp := innerSettingsSimilarityDfi{
		IndependenceMeasure: s.IndependenceMeasure,
		Type:                s.Type,
	}

	tmp.Type = "DFI"

	return json.Marshal(tmp)
}

// NewSettingsSimilarityDfi returns a SettingsSimilarityDfi.
func NewSettingsSimilarityDfi() *SettingsSimilarityDfi {
	r := &SettingsSimilarityDfi{}

	return r
}