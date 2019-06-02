package mapbox

type Layer struct {
	Id          string                 `json:"background"`
	Type        string                 `json:"type"`
	Paint       map[string]interface{} `json:"paint"`
	Source      string                 `json:"source"`
	SourceLayer string                 `json:"source-layer"`
	Filter      []interface{}          `json:"filter"`
	Layout      map[string]interface{} `json:"layout"`
	MinZoom     float32                `json:"minzoom"`
	MaxZoom     float32                `json:"maxzoom"`
	Metadata    map[string]interface{} `json:"metadata"`
}
