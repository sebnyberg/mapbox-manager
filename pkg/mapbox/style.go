package mapbox

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type ListStyle struct {
	Id         string    `json:"id"`
	Name       string    `json:"name"`
	Owner      string    `json:"owner"`
	Version    int32     `json:"version"`
	Center     []float64 `json:"center,omitempty"`
	Zoom       float64   `json:"zoom,omitempty"`
	Bearing    float64   `json:"bearing,omitempty"`
	Pitch      float64   `json:"pitch,omitempty"`
	Created    time.Time `json:"created"`
	Modified   time.Time `json:"modified"`
	Visibility string    `json:"visibility"`
}

// For more info, see https://docs.mapbox.com/mapbox-gl-js/style-spec/
type Style struct {
	Id       string                 `json:"id"`
	Name     string                 `json:"name,omitempty"`
	Owner    string                 `json:"owner"`
	Version  int32                  `json:"version"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	Sprite   string                 `json:"sprite,omitempty"`
	Glyphs   string                 `json:"glyphs,omitempty"`
	// View options
	Zoom       float32                `json:"zoom,omitempty"`
	Center     []float64              `json:"center,omitempty"`
	Bearing    float64                `json:"bearing,omitempty"`
	Pitch      float64                `json:"pitch,omitempty"`
	Light      float64                `json:"light,omitempty"`
	Transition map[string]interface{} `json:"transition,omitempty"`
	Draft      bool                   `json:"draft"`
	Created    time.Time              `json:"created,omitempty"`
	Modified   time.Time              `json:"modified,omitempty"`
	Visibility string                 `json:"visibility"`
	Layers     []Layer                `json:"layers"`
	Sources    Sources                `json:"sources"`
}

func GetStyles(accessToken string, username string) ([]ListStyle, error) {
	client := GetDefaultClient(accessToken)

	endpoint := fmt.Sprintf("/styles/v1/%v", username)

	res, err := client.Get(endpoint, nil, nil)
	if err != nil {
		log.Fatalf("failed to fetch styles for user: %v", err)
	}

	if res.StatusCode > 200 {
		return nil, fmt.Errorf("failed to fetch styles: %v", GetErrorMessage(res.StatusCode, res.Payload))
	}

	var styles []ListStyle
	if err := json.Unmarshal(res.Payload, &styles); err != nil {
		log.Fatalf("failed to parse styles: %v", err)
	}

	return styles, nil
}

func GetStyle(accessToken string, username string, styleId string, draft bool) (*Style, error) {
	client := GetDefaultClient(accessToken)

	endpoint := fmt.Sprintf("/styles/v1/%v/%v", username, styleId)

	if draft {
		endpoint += "/draft"
	}

	res, err := client.Get(endpoint, nil, nil)
	if err != nil {
		log.Fatalf("Failed to fetch styles for user: %v", err)
	}

	if res.StatusCode > 200 {
		return nil, fmt.Errorf("failed to fetch style: %v", GetErrorMessage(res.StatusCode, res.Payload))
	}

	var style Style
	if err := json.Unmarshal(res.Payload, &style); err != nil {
		log.Fatalf("Failed to parse styles: %v", err)
	}

	return &style, nil
}
