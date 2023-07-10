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
	"net/http"
	"strings"

	"github.com/opensearch-project/opensearch-go/v2"
)

// NodesUsageReq represents possible options for the /_nodes request
type NodesUsageReq struct {
	Metrics []string
	NodeID  []string

	Header http.Header
	Params NodesUsageParams
}

// GetRequest returns the *http.Request that gets executed by the client
func (r NodesUsageReq) GetRequest() (*http.Request, error) {
	nodes := strings.Join(r.NodeID, ",")
	metrics := strings.Join(r.Metrics, ",")

	var path strings.Builder

	path.Grow(len("/_nodes//usage/") + len(nodes) + len(metrics))

	path.WriteString("/_nodes")
	if len(r.NodeID) > 0 {
		path.WriteString("/")
		path.WriteString(nodes)
	}
	path.WriteString("/usage")
	if len(r.Metrics) > 0 {
		path.WriteString("/")
		path.WriteString(metrics)
	}

	return opensearch.BuildRequest(
		"GET",
		path.String(),
		nil,
		r.Params.get(),
		r.Header,
	)
}

// NodesUsageResp represents the returned struct of the /_nodes response
type NodesUsageResp struct {
	NodesUsage struct {
		Total      int `json:"total"`
		Successful int `json:"successful"`
		Failed     int `json:"failed"`
	} `json:"_nodes"`
	ClusterName string                `json:"cluster_name"`
	Nodes       map[string]NodesUsage `json:"nodes"`
	response    *opensearch.Response
}

// Inspect returns the Inspect type containing the raw *opensearch.Reponse
func (r NodesUsageResp) Inspect() Inspect {
	return Inspect{Response: r.response}
}

// NodesUsage is a sub type of NodesUsageResp containing stats about rest api actions
type NodesUsage struct {
	Timestamp   int64 `json:"timestamp"`
	Since       int64 `json:"since"`
	RestActions struct {
		InternalUsersAPIAction              int `json:"InternalUsersApiAction"`
		NodesUsageAction                    int `json:"nodes_usage_action"`
		ClearScrollAction                   int `json:"clear_scroll_action"`
		DocumentUpdateAction                int `json:"document_update_action"`
		GetIndexTemplateAction              int `json:"get_index_template_action"`
		RemoteClusterInfoAction             int `json:"remote_cluster_info_action"`
		PrometheusMetricsAction             int `json:"prometheus_metrics_action"`
		GetMappingAction                    int `json:"get_mapping_action"`
		UpgradeStatusAction                 int `json:"upgrade_status_action"`
		GetIndicesAction                    int `json:"get_indices_action"`
		IndicesAliasesAction                int `json:"indices_aliases_action"`
		CreateIndexAction                   int `json:"create_index_action"`
		TenantInfoAction                    int `json:"Tenant Info Action"`
		NodesHotThreadsAction               int `json:"nodes_hot_threads_action"`
		DeleteIndexAction                   int `json:"delete_index_action"`
		NodesReloadAction                   int `json:"nodes_reload_action"`
		FieldCapabilitiesAction             int `json:"field_capabilities_action"`
		DocumentGetAction                   int `json:"document_get_action"`
		ClusterHealthAction                 int `json:"cluster_health_action"`
		CountAction                         int `json:"count_action"`
		ActionGroupsAPIAction               int `json:"ActionGroupsApiAction"`
		ExplainAction                       int `json:"explain_action"`
		DeleteDataStreamAction              int `json:"delete_data_stream_action"`
		ListDanglingIndices                 int `json:"list_dangling_indices"`
		CatPluginsAction                    int `json:"cat_plugins_action"`
		DocumentIndexAction                 int `json:"document_index_action"`
		CatIndicesAction                    int `json:"cat_indices_action"`
		ScriptContextAction                 int `json:"script_context_action"`
		GetPolicyAction                     int `json:"get_policy_action"`
		IndexPolicyAction                   int `json:"index_policy_action"`
		DeleteComponentTemplateAction       int `json:"delete_component_template_action"`
		SimulateTemplateAction              int `json:"simulate_template_action"`
		PermissionsInfoAction               int `json:"PermissionsInfoAction"`
		PutIndexTemplateAction              int `json:"put_index_template_action"`
		DocumentMgetAction                  int `json:"document_mget_action"`
		DataStreamStatsAction               int `json:"data_stream_stats_action"`
		RolesAPIAction                      int `json:"RolesApiAction"`
		BulkAction                          int `json:"bulk_action"`
		IngestGetPipelineAction             int `json:"ingest_get_pipeline_action"`
		DeleteComposableIndexTemplateAction int `json:"delete_composable_index_template_action"`
		GetComponentTemplateAction          int `json:"get_component_template_action"`
		DeleteStoredScriptAction            int `json:"delete_stored_script_action"`
		PutComponentTemplateAction          int `json:"put_component_template_action"`
		NodesInfoAction                     int `json:"nodes_info_action"`
		GetComposableIndexTemplateAction    int `json:"get_composable_index_template_action"`
		GetStoredScriptsAction              int `json:"get_stored_scripts_action"`
		DocumentMultiTermVectorsAction      int `json:"document_multi_term_vectors_action"`
		MainAction                          int `json:"main_action"`
		GetDataStreamsAction                int `json:"get_data_streams_action"`
		KibanaInfoAction                    int `json:"Kibana Info Action"`
		DocumentCreateAction                int `json:"document_create_action"`
		OpenSearchSecurityInfoAction        int `json:"OpenSearch Security Info Action"`
		ResolveIndexAction                  int `json:"resolve_index_action"`
		SearchAction                        int `json:"search_action"`
		ClearVotingConfigExclusionsAction   int `json:"clear_voting_config_exclusions_action"`
		GetAliasesAction                    int `json:"get_aliases_action"`
		AccountAPIAction                    int `json:"AccountApiAction"`
		PutComposableIndexTemplateAction    int `json:"put_composable_index_template_action"`
		RolloverIndexAction                 int `json:"rollover_index_action"`
		ClusterUpdateSettingsAction         int `json:"cluster_update_settings_action"`
		RefreshAction                       int `json:"refresh_action"`
		SearchScrollAction                  int `json:"search_scroll_action"`
		CatCountAction                      int `json:"cat_count_action"`
		RethrottleAction                    int `json:"rethrottle_action"`
		CatSegmentsAction                   int `json:"cat_segments_action"`
		CatHealthAction                     int `json:"cat_health_action"`
		ClusterAllocationExplainAction      int `json:"cluster_allocation_explain_action"`
		CatTemplatesAction                  int `json:"cat_templates_action"`
		CatShardsAction                     int `json:"cat_shards_action"`
		ClusterGetSettingsAction            int `json:"cluster_get_settings_action"`
		GetSettingsAction                   int `json:"get_settings_action"`
		CloseIndexAction                    int `json:"close_index_action"`
		DeleteDecommissionStateAction       int `json:"delete_decommission_state_action"`
		GetDecommissionStateAction          int `json:"get_decommission_state_action"`
		ValidateQueryAction                 int `json:"validate_query_action"`
		SplitIndexAction                    int `json:"split_index_action"`
		DeleteIndexTemplateAction           int `json:"delete_index_template_action"`
		CatThreadpoolAction                 int `json:"cat_threadpool_action"`
		CatNodeAttrsAction                  int `json:"cat_node_attrs_action"`
		CatClusterManagerAction             int `json:"cat_cluster_manager_action"`
		IndexPutAliasAction                 int `json:"index_put_alias_action"`
		ScriptLanguageAction                int `json:"script_language_action"`
		CloneIndexAction                    int `json:"clone_index_action"`
		CreateDataStreamAction              int `json:"create_data_stream_action"`
		DocumentGetSourceAction             int `json:"document_get_source_action"`
		AnalyzeAction                       int `json:"analyze_action"`
		PutMappingAction                    int `json:"put_mapping_action"`
		ShrinkIndexAction                   int `json:"shrink_index_action"`
		DocumentCreateActionAutoID          int `json:"document_create_action_auto_id"`
		CatPendingClusterTasksAction        int `json:"cat_pending_cluster_tasks_action"`
		IndexDeleteAliasesAction            int `json:"index_delete_aliases_action"`
		ClusterStatsAction                  int `json:"cluster_stats_action"`
		CatAllocationAction                 int `json:"cat_allocation_action"`
		FlushAction                         int `json:"flush_action"`
		ClearIndicesCacheAction             int `json:"clear_indices_cache_action"`
		CatAliasAction                      int `json:"cat_alias_action"`
		ClusterRerouteAction                int `json:"cluster_reroute_action"`
		AddIndexBlockAction                 int `json:"add_index_block_action"`
		NodesStatsAction                    int `json:"nodes_stats_action"`
		PutStoredScriptAction               int `json:"put_stored_script_action"`
		DocumentDeleteAction                int `json:"document_delete_action"`
		CatRecoveryAction                   int `json:"cat_recovery_action"`
		ForceMergeAction                    int `json:"force_merge_action"`
		IndicesShardStoresAction            int `json:"indices_shard_stores_action"`
		GetFieldMappingAction               int `json:"get_field_mapping_action"`
		IndicesSegmentsAction               int `json:"indices_segments_action"`
		CatFielddataAction                  int `json:"cat_fielddata_action"`
		DeleteByQueryAction                 int `json:"delete_by_query_action"`
		CatTasksAction                      int `json:"cat_tasks_action"`
		PendingClusterTasksAction           int `json:"pending_cluster_tasks_action"`
		CatRepositoriesAction               int `json:"cat_repositories_action"`
		IndicesStatsAction                  int `json:"indices_stats_action"`
		OpenIndexAction                     int `json:"open_index_action"`
		CatNodesAction                      int `json:"cat_nodes_action"`
		ScriptsPainlessExecute              int `json:"_scripts_painless_execute"`
		ClusterStateAction                  int `json:"cluster_state_action"`
		SimulateIndexTemplateAction         int `json:"simulate_index_template_action"`
		RecoveryAction                      int `json:"recovery_action"`
		UpdateSettingsAction                int `json:"update_settings_action"`
		CatMasterAction                     int `json:"cat_master_action"` // Deprecated field
		CreatePITAction                     int `json:"create_pit_action"`
		DeletePITAction                     int `json:"delete_pit_action"`
		GetAllPITAction                     int `json:"get_all_pit_action"`
		IngestDeletePipelineAction          int `json:"ingest_delete_pipeline_action"`
		IngestPutPipelineAction             int `json:"ingest_put_pipeline_action"`
		IngestSimulatePipelineAction        int `json:"ingest_simulate_pipeline_action"`
		IngestProcessorGrokGet              int `json:"ingest_processor_grok_get"`
		ListTasksAction                     int `json:"list_tasks_action"`
		ReindexAction                       int `json:"reindex_action"`
		CancelTasksAction                   int `json:"cancel_tasks_action"`
		GetTaskAction                       int `json:"get_task_action"`
		MSearchAction                       int `json:"msearch_action"`
		MultiSearchTemplateAction           int `json:"multi_search_template_action"`
		RankEvalAction                      int `json:"rank_eval_action"`
	} `json:"rest_actions"`
	Aggregations json.RawMessage `json:"aggregations"` // Can contain unknow fields
}
