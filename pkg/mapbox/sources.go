package mapbox

type Composite struct {
	URL  string `json:"url,omitempty"`
	Type string `json:"type,omitempty"`
}

type Sources struct {
	Composite *Composite `json:"composite,omitempty"`
}
