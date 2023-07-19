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
//
//go:build integration

package opensearchapi_test

import (
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/opensearch-project/opensearch-go/v2/opensearchapi"
	osapitest "github.com/opensearch-project/opensearch-go/v2/opensearchapi/internal/test"
)

func TestCatClient(t *testing.T) {
	t.Run("Indices", func(t *testing.T) {
		client, err := opensearchapi.NewDefaultClient()
		require.Nil(t, err)
		_, err = client.Index(nil, opensearchapi.IndexReq{Index: "test-cat-indices", Body: strings.NewReader("{}")})
		require.Nil(t, err)

		t.Run("with nil request", func(t *testing.T) {
			res, err := client.Cat.Indices(nil, nil)
			assert.Nil(t, err)
			assert.NotEmpty(t, res.Indices)
			assert.NotEmpty(t, res.Inspect())
		})

		t.Run("with request", func(t *testing.T) {
			res, err := client.Cat.Indices(nil, &opensearchapi.CatIndicesReq{Index: []string{"*"}})
			assert.Nil(t, err)
			assert.NotEmpty(t, res.Indices)
			inspect := res.Inspect()
			assert.NotEmpty(t, inspect.Response)
			assert.Equal(t, http.StatusOK, inspect.Response.StatusCode)
		})

		t.Run("inspect", func(t *testing.T) {
			failingClient, err := osapitest.CreateFailingClient()
			require.Nil(t, err)

			res, err := failingClient.Cat.Indices(nil, nil)
			assert.NotNil(t, err)
			assert.NotNil(t, res)
			osapitest.VerifyInspect(t, res.Inspect())
		})
	})
}
