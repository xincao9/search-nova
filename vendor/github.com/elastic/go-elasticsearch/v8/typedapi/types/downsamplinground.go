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
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

// DownsamplingRound type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4fcf747dfafc951e1dcf3077327e3dcee9107db3/specification/indices/_types/DownsamplingRound.ts#L23-L32
type DownsamplingRound struct {
	// After The duration since rollover when this downsampling round should execute
	After Duration `json:"after"`
	// Config The downsample configuration to execute.
	Config DownsampleConfig `json:"config"`
}

func (s *DownsamplingRound) UnmarshalJSON(data []byte) error {

	dec := json.NewDecoder(bytes.NewReader(data))

	for {
		t, err := dec.Token()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}

		switch t {

		case "after":
			if err := dec.Decode(&s.After); err != nil {
				return fmt.Errorf("%s | %w", "After", err)
			}

		case "config":
			if err := dec.Decode(&s.Config); err != nil {
				return fmt.Errorf("%s | %w", "Config", err)
			}

		}
	}
	return nil
}

// NewDownsamplingRound returns a DownsamplingRound.
func NewDownsamplingRound() *DownsamplingRound {
	r := &DownsamplingRound{}

	return r
}
