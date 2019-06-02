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
	Center     []float64 `json:"center"`
	Zoom       float64   `json:"zoom"`
	Bearing    float64   `json:"bearing"`
	Pitch      float64   `json:"pitch"`
	Created    time.Time `json:"created"`
	Modified   time.Time `json:"modified"`
	Visibility string    `json:"visibility"`
}

type Style struct {
	Id         string                 `json:"id"`
	Name       string                 `json:"name"`
	Owner      string                 `json:"owner"`
	Version    int32                  `json:"version"`
	Metadata   map[string]interface{} `json:"metadata"`
	Zoom       float32                `json:"zoom"`
	Sprite     string                 `json:"sprite"`
	Center     []float64              `json:"center"`
	Bearing    float64                `json:"bearing"`
	Pitch      float64                `json:"pitch"`
	Draft      bool                   `json:"draft"`
	Created    time.Time              `json:"created"`
	Modified   time.Time              `json:"modified"`
	Visibility string                 `json:"visibility"`
	Layers     []Layer                `json:"layers"`
	Sources    Sources                `json:"sources"`
}

type NewStyle struct {
	Id         string                 `json:"id"`
	Name       string                 `json:"name"`
	Owner      string                 `json:"owner"`
	Version    int32                  `json:"version"`
	Metadata   map[string]interface{} `json:"metadata"`
	Zoom       float32                `json:"zoom"`
	Sprite     string                 `json:"sprite"`
	Center     []float64              `json:"center"`
	Bearing    float64                `json:"bearing"`
	Pitch      float64                `json:"pitch"`
	Draft      bool                   `json:"draft"`
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

func GetStyle(accessToken string, username string, styleId string) (*Style, error) {
	client := GetDefaultClient(accessToken)

	endpoint := fmt.Sprintf("/styles/v1/%v/%v", username, styleId)

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
