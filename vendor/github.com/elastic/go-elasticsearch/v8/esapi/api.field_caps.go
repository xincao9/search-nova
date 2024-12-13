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
//
// Code generated from specification version 8.16.0: DO NOT EDIT

package esapi

import (
	"context"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func newFieldCapsFunc(t Transport) FieldCaps {
	return func(o ...func(*FieldCapsRequest)) (*Response, error) {
		var r = FieldCapsRequest{}
		for _, f := range o {
			f(&r)
		}

		if transport, ok := t.(Instrumented); ok {
			r.instrument = transport.InstrumentationEnabled()
		}

		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

// FieldCaps returns the information about the capabilities of fields among multiple indices.
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/master/search-field-caps.html.
type FieldCaps func(o ...func(*FieldCapsRequest)) (*Response, error)

// FieldCapsRequest configures the Field Caps API request.
type FieldCapsRequest struct {
	Index []string

	Body io.Reader

	AllowNoIndices     *bool
	ExpandWildcards    string
	Fields             []string
	Filters            []string
	IgnoreUnavailable  *bool
	IncludeEmptyFields *bool
	IncludeUnmapped    *bool
	Types              []string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	instrument Instrumentation
}

// Do executes the request and returns response or error.
func (r FieldCapsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
		ctx    context.Context
	)

	if instrument, ok := r.instrument.(Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "field_caps")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	method = "POST"

	path.Grow(7 + 1 + len(strings.Join(r.Index, ",")) + 1 + len("_field_caps"))
	path.WriteString("http://")
	if len(r.Index) > 0 {
		path.WriteString("/")
		path.WriteString(strings.Join(r.Index, ","))
		if instrument, ok := r.instrument.(Instrumentation); ok {
			instrument.RecordPathPart(ctx, "index", strings.Join(r.Index, ","))
		}
	}
	path.WriteString("/")
	path.WriteString("_field_caps")

	params = make(map[string]string)

	if r.AllowNoIndices != nil {
		params["allow_no_indices"] = strconv.FormatBool(*r.AllowNoIndices)
	}

	if r.ExpandWildcards != "" {
		params["expand_wildcards"] = r.ExpandWildcards
	}

	if len(r.Fields) > 0 {
		params["fields"] = strings.Join(r.Fields, ",")
	}

	if len(r.Filters) > 0 {
		params["filters"] = strings.Join(r.Filters, ",")
	}

	if r.IgnoreUnavailable != nil {
		params["ignore_unavailable"] = strconv.FormatBool(*r.IgnoreUnavailable)
	}

	if r.IncludeEmptyFields != nil {
		params["include_empty_fields"] = strconv.FormatBool(*r.IncludeEmptyFields)
	}

	if r.IncludeUnmapped != nil {
		params["include_unmapped"] = strconv.FormatBool(*r.IncludeUnmapped)
	}

	if len(r.Types) > 0 {
		params["types"] = strings.Join(r.Types, ",")
	}

	if r.Pretty {
		params["pretty"] = "true"
	}

	if r.Human {
		params["human"] = "true"
	}

	if r.ErrorTrace {
		params["error_trace"] = "true"
	}

	if len(r.FilterPath) > 0 {
		params["filter_path"] = strings.Join(r.FilterPath, ",")
	}

	req, err := newRequest(method, path.String(), r.Body)
	if err != nil {
		if instrument, ok := r.instrument.(Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return nil, err
	}

	if len(params) > 0 {
		q := req.URL.Query()
		for k, v := range params {
			q.Set(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	if len(r.Header) > 0 {
		if len(req.Header) == 0 {
			req.Header = r.Header
		} else {
			for k, vv := range r.Header {
				for _, v := range vv {
					req.Header.Add(k, v)
				}
			}
		}
	}

	if r.Body != nil && req.Header.Get(headerContentType) == "" {
		req.Header[headerContentType] = headerContentTypeJSON
	}

	if ctx != nil {
		req = req.WithContext(ctx)
	}

	if instrument, ok := r.instrument.(Instrumentation); ok {
		instrument.BeforeRequest(req, "field_caps")
		if reader := instrument.RecordRequestBody(ctx, "field_caps", r.Body); reader != nil {
			req.Body = reader
		}
	}
	res, err := transport.Perform(req)
	if instrument, ok := r.instrument.(Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "field_caps")
	}
	if err != nil {
		if instrument, ok := r.instrument.(Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return nil, err
	}

	response := Response{
		StatusCode: res.StatusCode,
		Body:       res.Body,
		Header:     res.Header,
	}

	return &response, nil
}

// WithContext sets the request context.
func (f FieldCaps) WithContext(v context.Context) func(*FieldCapsRequest) {
	return func(r *FieldCapsRequest) {
		r.ctx = v
	}
}

// WithBody - An index filter specified with the Query DSL.
func (f FieldCaps) WithBody(v io.Reader) func(*FieldCapsRequest) {
	return func(r *FieldCapsRequest) {
		r.Body = v
	}
}

// WithIndex - a list of index names; use _all to perform the operation on all indices.
func (f FieldCaps) WithIndex(v ...string) func(*FieldCapsRequest) {
	return func(r *FieldCapsRequest) {
		r.Index = v
	}
}

// WithAllowNoIndices - whether to ignore if a wildcard indices expression resolves into no concrete indices. (this includes `_all` string or when no indices have been specified).
func (f FieldCaps) WithAllowNoIndices(v bool) func(*FieldCapsRequest) {
	return func(r *FieldCapsRequest) {
		r.AllowNoIndices = &v
	}
}

// WithExpandWildcards - whether to expand wildcard expression to concrete indices that are open, closed or both..
func (f FieldCaps) WithExpandWildcards(v string) func(*FieldCapsRequest) {
	return func(r *FieldCapsRequest) {
		r.ExpandWildcards = v
	}
}

// WithFields - a list of field names.
func (f FieldCaps) WithFields(v ...string) func(*FieldCapsRequest) {
	return func(r *FieldCapsRequest) {
		r.Fields = v
	}
}

// WithFilters - an optional set of filters: can include +metadata,-metadata,-nested,-multifield,-parent.
func (f FieldCaps) WithFilters(v ...string) func(*FieldCapsRequest) {
	return func(r *FieldCapsRequest) {
		r.Filters = v
	}
}

// WithIgnoreUnavailable - whether specified concrete indices should be ignored when unavailable (missing or closed).
func (f FieldCaps) WithIgnoreUnavailable(v bool) func(*FieldCapsRequest) {
	return func(r *FieldCapsRequest) {
		r.IgnoreUnavailable = &v
	}
}

// WithIncludeEmptyFields - include empty fields in result.
func (f FieldCaps) WithIncludeEmptyFields(v bool) func(*FieldCapsRequest) {
	return func(r *FieldCapsRequest) {
		r.IncludeEmptyFields = &v
	}
}

// WithIncludeUnmapped - indicates whether unmapped fields should be included in the response..
func (f FieldCaps) WithIncludeUnmapped(v bool) func(*FieldCapsRequest) {
	return func(r *FieldCapsRequest) {
		r.IncludeUnmapped = &v
	}
}

// WithTypes - only return results for fields that have one of the types in the list.
func (f FieldCaps) WithTypes(v ...string) func(*FieldCapsRequest) {
	return func(r *FieldCapsRequest) {
		r.Types = v
	}
}

// WithPretty makes the response body pretty-printed.
func (f FieldCaps) WithPretty() func(*FieldCapsRequest) {
	return func(r *FieldCapsRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f FieldCaps) WithHuman() func(*FieldCapsRequest) {
	return func(r *FieldCapsRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f FieldCaps) WithErrorTrace() func(*FieldCapsRequest) {
	return func(r *FieldCapsRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f FieldCaps) WithFilterPath(v ...string) func(*FieldCapsRequest) {
	return func(r *FieldCapsRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f FieldCaps) WithHeader(h map[string]string) func(*FieldCapsRequest) {
	return func(r *FieldCapsRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f FieldCaps) WithOpaqueID(s string) func(*FieldCapsRequest) {
	return func(r *FieldCapsRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}