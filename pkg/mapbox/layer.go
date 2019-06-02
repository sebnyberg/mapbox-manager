package mapbox

// For more info see https://docs.mapbox.com/mapbox-gl-js/style-spec/#layers
type Layer struct {
	Id          string                 `json:"id"`
	Type        string                 `json:"type"`
	Paint       map[string]interface{} `json:"paint"`
	Source      string                 `json:"source,omitempty"`
	SourceLayer string                 `json:"source-layer,omitempty"`
	Filter      []interface{}          `json:"filter,omitempty"`
	Layout      map[string]interface{} `json:"layout,omitempty"`
	MinZoom     float32                `json:"minzoom,omitempty"`
	MaxZoom     float32                `json:"maxzoom,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}
