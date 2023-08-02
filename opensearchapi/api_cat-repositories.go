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
	"net/http"

	"github.com/opensearch-project/opensearch-go/v2"
)

// CatRepositoriesReq represent possible options for the /_cat/repositories request
type CatRepositoriesReq struct {
	Header http.Header
	Params CatRepositoriesParams
}

// GetRequest returns the *http.Request that gets executed by the client
func (r CatRepositoriesReq) GetRequest() (*http.Request, error) {
	return opensearch.BuildRequest(
		"GET",
		"/_cat/repositories",
		nil,
		r.Params.get(),
		r.Header,
	)
}

// CatRepositoriesResp represents the returned struct of the /_cat/repositories response
type CatRepositoriesResp struct {
	Repositories []CatRepositorieResp
	response     *opensearch.Response
}

// CatRepositorieResp represents one index of the CatRepositoriesResp
type CatRepositorieResp struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

// Inspect returns the Inspect type containing the raw *opensearch.Reponse
func (r CatRepositoriesResp) Inspect() Inspect {
	return Inspect{
		Response: r.response,
	}
}
