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

// SimulateIngest type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4fcf747dfafc951e1dcf3077327e3dcee9107db3/specification/ingest/simulate/types.ts#L29-L37
type SimulateIngest struct {
	Pipeline  *string  `json:"pipeline,omitempty"`
	Redact_   *Redact  `json:"_redact,omitempty"`
	Timestamp DateTime `json:"timestamp"`
}

func (s *SimulateIngest) UnmarshalJSON(data []byte) error {

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

		case "pipeline":
			if err := dec.Decode(&s.Pipeline); err != nil {
				return fmt.Errorf("%s | %w", "Pipeline", err)
			}

		case "_redact":
			if err := dec.Decode(&s.Redact_); err != nil {
				return fmt.Errorf("%s | %w", "Redact_", err)
			}

		case "timestamp":
			if err := dec.Decode(&s.Timestamp); err != nil {
				return fmt.Errorf("%s | %w", "Timestamp", err)
			}

		}
	}
	return nil
}

// NewSimulateIngest returns a SimulateIngest.
func NewSimulateIngest() *SimulateIngest {
	r := &SimulateIngest{}

	return r
}
