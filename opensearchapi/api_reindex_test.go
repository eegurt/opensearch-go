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
	"context"
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/opensearch-project/opensearch-go/v2/opensearchapi"
	osapitest "github.com/opensearch-project/opensearch-go/v2/opensearchapi/internal/test"
	"github.com/opensearch-project/opensearch-go/v2/opensearchutil"
)

func TestReindex(t *testing.T) {
	client, err := opensearchapi.NewDefaultClient()
	require.Nil(t, err)

	sourceIndex := "test-reindex-source"
	destIndex := "test-reindex-dest"
	t.Cleanup(func() {
		client.Indices.Delete(
			nil,
			opensearchapi.IndicesDeleteReq{
				Indices: []string{sourceIndex, destIndex},
				Params:  opensearchapi.IndicesDeleteParams{IgnoreUnavailable: opensearchapi.ToPointer(true)},
			},
		)
	})

	bi, err := opensearchutil.NewBulkIndexer(opensearchutil.BulkIndexerConfig{
		Index:      sourceIndex,
		Client:     client,
		ErrorTrace: true,
		Human:      true,
		Pretty:     true,
	})
	for i := 1; i <= 10000; i++ {
		err := bi.Add(context.Background(), opensearchutil.BulkIndexerItem{
			Index:      sourceIndex,
			Action:     "index",
			DocumentID: strconv.Itoa(i),
			Body:       strings.NewReader(`{"title":"bar"}`),
		})
		if err != nil {
			t.Fatalf("Unexpected error: %s", err)
		}
	}
	if err := bi.Close(context.Background()); err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	t.Run("with request", func(t *testing.T) {
		resp, err := client.Reindex(
			nil,
			opensearchapi.ReindexReq{
				Body: strings.NewReader(fmt.Sprintf(`{"source":{"index":"%s"},"dest":{"index":"%s"}}`, sourceIndex, destIndex)),
			},
		)
		require.Nil(t, err)
		assert.NotEmpty(t, resp)
		osapitest.CompareRawJSONwithParsedJSON(t, resp, resp.Inspect().Response)
	})

	t.Run("with request but dont wait", func(t *testing.T) {
		resp, err := client.Reindex(
			nil,
			opensearchapi.ReindexReq{
				Body:   strings.NewReader(fmt.Sprintf(`{"source":{"index":"%s"},"dest":{"index":"%s"}}`, sourceIndex, destIndex)),
				Params: opensearchapi.ReindexParams{WaitForCompletion: opensearchapi.ToPointer(false)},
			},
		)
		require.Nil(t, err)
		assert.NotEmpty(t, resp)
		osapitest.CompareRawJSONwithParsedJSON(t, resp, resp.Inspect().Response)
	})

	t.Run("inspect", func(t *testing.T) {
		failingClient, err := osapitest.CreateFailingClient()
		require.Nil(t, err)

		res, err := failingClient.Reindex(nil, opensearchapi.ReindexReq{})
		assert.NotNil(t, err)
		assert.NotNil(t, res)
		osapitest.VerifyInspect(t, res.Inspect())
	})
}