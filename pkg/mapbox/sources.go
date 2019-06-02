package mapbox

type Composite struct {
	URL  string `json:"url"`
	Type string `json:"type"`
}

type Sources struct {
	Composite Composite `json:"composite"`
}
