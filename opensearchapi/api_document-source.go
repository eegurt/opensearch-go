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
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/opensearch-project/opensearch-go/v2"
)

// DocumentSourceReq represents possible options for the /_doc request
type DocumentSourceReq struct {
	Index      string
	DocumentID string

	Header http.Header
	Params DocumentSourceParams
}

// GetRequest returns the *http.Request that gets executed by the client
func (r DocumentSourceReq) GetRequest() (*http.Request, error) {
	return opensearch.BuildRequest(
		"GET",
		fmt.Sprintf("/%s/_source/%s", r.Index, r.DocumentID),
		nil,
		r.Params.get(),
		r.Header,
	)
}

// DocumentSourceResp represents the returned struct of the /_doc response
type DocumentSourceResp struct {
	Source   json.RawMessage
	response *opensearch.Response
}

// Inspect returns the Inspect type containing the raw *opensearch.Reponse
func (r DocumentSourceResp) Inspect() Inspect {
	return Inspect{Response: r.response}
}
