// SPDX-License-Identifier: Apache-2.0
//
// The OpenSearch Contributors require contributions made to
// this file be licensed under the Apache-2.0 license or a
// compatible open source license.
//
// Modifications Copyright OpenSearch Contributors. See
// GitHub history for details.

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package opensearchapi

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/opensearch-project/opensearch-go/v2"
)

// MTermvectors executes a /_mtermvectors request with the required MTermvectorsReq
func (c Client) MTermvectors(ctx context.Context, req MTermvectorsReq) (*MTermvectorsResp, error) {
	var (
		data MTermvectorsResp
		err  error
	)
	if data.response, err = c.do(ctx, req, &data); err != nil {
		return &data, err
	}

	return &data, nil
}

// MTermvectorsReq represents possible options for the /_mtermvectors request
type MTermvectorsReq struct {
	Index string

	Body io.Reader

	Header http.Header
	Params MTermvectorsParams
}

// GetRequest returns the *http.Request that gets executed by the client
func (r MTermvectorsReq) GetRequest() (*http.Request, error) {
	var path strings.Builder
	path.Grow(len("//_mtermvectors") + len(r.Index))
	if len(r.Index) > 0 {
		path.WriteString("/")
		path.WriteString(r.Index)
	}
	path.WriteString("/_mtermvectors")
	return opensearch.BuildRequest(
		"POST",
		path.String(),
		r.Body,
		r.Params.get(),
		r.Header,
	)
}

// MTermvectorsResp represents the returned struct of the /_mtermvectors response
type MTermvectorsResp struct {
	Docs []struct {
		Index       string          `json:"_index"`
		ID          string          `json:"_id"`
		Version     int             `json:"_version"`
		Found       bool            `json:"found"`
		Took        int             `json:"took"`
		Type        string          `json:"_type"` // Deprecated field
		TermVectors json.RawMessage `json:"term_vectors"`
	} `json:"docs"`
	response *opensearch.Response
}

// Inspect returns the Inspect type containing the raw *opensearch.Reponse
func (r MTermvectorsResp) Inspect() Inspect {
	return Inspect{Response: r.response}
}