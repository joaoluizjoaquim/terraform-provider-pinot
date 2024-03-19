package models

import "github.com/hashicorp/terraform-plugin-framework/types"

type TableResourceModel struct {
	TableName        types.String      `tfsdk:"table_name"`
	Table            types.String      `tfsdk:"table"`
	TableType        types.String      `tfsdk:"table_type"`
	SegmentsConfig   *SegmentsConfig   `tfsdk:"segments_config"`
	TenantsConfig    *TenantsConfig    `tfsdk:"tenants"`
	TableIndexConfig *TableIndexConfig `tfsdk:"table_index_config"`
	UpsertConfig     *UpsertConfig     `tfsdk:"upsert_config"`
	IngestionConfig  *IngestionConfig  `tfsdk:"ingestion_config"`
	TierConfigs      []*TierConfig     `tfsdk:"tier_configs"`
	IsDimTable       types.Bool        `tfsdk:"is_dim_table"`
	Metadata         *Metadata         `tfsdk:"metadata"`
}

type TenantsConfig struct {
	Broker            types.String   `tfsdk:"broker"`
	Server            types.String   `tfsdk:"server"`
	TagOverrideConfig *types.MapType `tfsdk:"tag_override_config"`
}

type SegmentsConfig struct {
	TimeType           types.String `tfsdk:"time_type"`
	Replication        types.String `tfsdk:"replication"`
	TimeColumnName     types.String `tfsdk:"time_column_name"`
	RetentionTimeUnit  types.String `tfsdk:"retention_time_unit"`
	RetentionTimeValue types.String `tfsdk:"retention_time_value"`
}

type UpsertConfig struct {
	Mode                  types.String  `tfsdk:"mode"`
	PartialUpsertStrategy types.MapType `tfsdk:"partial_upsert_strategy"`
}

type SegmentPartitionConfig struct {
	ColumnPartitionMap map[string]map[string]string `tfsdk:"column_partition_map"`
}

type TableIndexConfig struct {
	SortedColumn                               []string                `tfsdk:"sorted_column"`
	LoadMode                                   types.String            `tfsdk:"load_mode"`
	NullHandlingEnabled                        types.Bool              `tfsdk:"null_handling_enabled"`
	CreateInvertedIndexDuringSegmentGeneration types.Bool              `tfsdk:"create_inverted_index_during_segment_generation"`
	EnableDynamicStarTree                      types.Bool              `tfsdk:"enable_dynamic_star_tree"`
	EnableDefaultStarTree                      types.Bool              `tfsdk:"enable_default_star_tree"`
	OptimizeDictionary                         types.Bool              `tfsdk:"optimize_dictionary"`
	OptimizeDictionaryForMetrics               types.Bool              `tfsdk:"optimize_dictionary_for_metrics"`
	NoDictionarySizeRatioThreshold             types.Float64           `tfsdk:"no_dictionary_size_ratio_threshold"`
	ColumnMinMaxValueGeneratorMode             types.String            `tfsdk:"column_min_max_value_generator_mode"`
	SegmentNameGeneratorType                   types.String            `tfsdk:"segment_name_generator_type"`
	AggregateMetrics                           types.Bool              `tfsdk:"aggregate_metrics"`
	StarTreeIndexConfigs                       []*StarTreeIndexConfigs `tfsdk:"star_tree_index_configs"`
	SegmentPartitionConfig                     *SegmentPartitionConfig `tfsdk:"segment_partition_config"`
}

type AggregationConfig struct {
	AggregationFunction types.String `tfsdk:"aggregation_function"`
	ColumnName          types.String `tfsdk:"column_name"`
	CompressionCodec    types.String `tfsdk:"compression_codec"`
}

type StarTreeIndexConfigs struct {
	MaxLeafRecords                  types.Int64          `tfsdk:"max_leaf_records"`
	SkipStarNodeCreationForDimNames []string             `tfsdk:"skip_star_node_creation_for_dim_names"`
	DimensionsSplitOrder            []string             `tfsdk:"dimensions_split_order"`
	FunctionColumnPairs             []string             `tfsdk:"function_column_pairs"`
	AggregationConfigs              []*AggregationConfig `tfsdk:"aggregation_configs"`
}

type IngestionConfig struct {
	SegmentTimeValueCheck types.Bool             `tfsdk:"segment_time_value_check"`
	RowTimeValueCheck     types.Bool             `tfsdk:"row_time_value_check"`
	ContinueOnError       types.Bool             `tfsdk:"continue_on_error"`
	StreamIngestionConfig *StreamIngestionConfig `tfsdk:"stream_ingestion_config"`
}

type StreamIngestionConfig struct {
	StreamConfigMaps []map[string]string `tfsdk:"stream_config_maps"`
}

type TierConfig struct {
	Name                types.String `tfsdk:"name"`
	StorageType         types.String `tfsdk:"storage_type"`
	SegmentSelectorType types.String `tfsdk:"segment_selector_type"`
	SegmentAge          types.String `tfsdk:"segment_age"`
	ServerTag           types.String `tfsdk:"server_tag"`
}

type Metadata struct {
	CustomConfigs map[string]string `tfsdk:"custom_configs"`
}
