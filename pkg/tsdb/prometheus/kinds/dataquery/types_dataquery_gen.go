// Code generated - EDITING IS FUTILE. DO NOT EDIT.
//
// Generated by:
//     public/app/plugins/gen.go
// Using jennies:
//     PluginGoTypesJenny
//
// Run 'make gen-cue' from repository root to regenerate.

package dataquery

// Defines values for PromQueryFormat.
const (
	PromQueryFormatHeatmap    PromQueryFormat = "heatmap"
	PromQueryFormatTable      PromQueryFormat = "table"
	PromQueryFormatTimeSeries PromQueryFormat = "time_series"
)

// Defines values for QueryEditorMode.
const (
	QueryEditorModeBuilder QueryEditorMode = "builder"
	QueryEditorModeCode    QueryEditorMode = "code"
)

// These are the common properties available to all queries in all datasources.
// Specific implementations will *extend* this interface, adding the required
// properties for the given context.
type DataQuery struct {
	// For mixed data sources the selected datasource is on the query level.
	// For non mixed scenarios this is undefined.
	// TODO find a better way to do this ^ that's friendly to schema
	// TODO this shouldn't be unknown but DataSourceRef | null
	Datasource *interface{} `json:"datasource,omitempty"`

	// Hide true if query is disabled (ie should not be returned to the dashboard)
	// Note this does not always imply that the query should not be executed since
	// the results from a hidden query may be used as the input to other queries (SSE etc)
	Hide *bool `json:"hide,omitempty"`

	// Specify the query flavor
	// TODO make this required and give it a default
	QueryType *string `json:"queryType,omitempty"`

	// A unique identifier for the query within the list of targets.
	// In server side expressions, the refId is used as a variable name to identify results.
	// By default, the UI will assign A->Z; however setting meaningful names may be useful.
	RefId string `json:"refId"`
}

// PromQueryFormat defines model for PromQueryFormat.
type PromQueryFormat string

// PrometheusDataQuery defines model for PrometheusDataQuery.
type PrometheusDataQuery struct {
	// DataQuery These are the common properties available to all queries in all datasources.
	// Specific implementations will *extend* this interface, adding the required
	// properties for the given context.
	DataQuery
	EditorMode *QueryEditorMode `json:"editorMode,omitempty"`

	// Execute an additional query to identify interesting raw samples relevant for the given expr
	Exemplar *bool `json:"exemplar,omitempty"`

	// The actual expression/query that will be evaluated by Prometheus
	Expr   *string          `json:"expr,omitempty"`
	Format *PromQueryFormat `json:"format,omitempty"`

	// Returns only the latest value that Prometheus has scraped for the requested time series
	Instant *bool `json:"instant,omitempty"`

	// @deprecated Used to specify how many times to divide max data points by. We use max data points under query options
	// See https://github.com/grafana/grafana/issues/48081
	IntervalFactor *float32 `json:"intervalFactor,omitempty"`

	// Series name override or template. Ex. {{hostname}} will be replaced with label value for hostname
	LegendFormat *string `json:"legendFormat,omitempty"`

	// Returns a Range vector, comprised of a set of time series containing a range of data points over time for each time series
	Range *bool `json:"range,omitempty"`
}

// QueryEditorMode defines model for QueryEditorMode.
type QueryEditorMode string
