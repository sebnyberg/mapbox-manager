package resource

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type ListStyle struct {
	Version    int32     `json:"version"`
	Name       string    `json:"name"`
	Center     []float64 `json:"center"`
	Zoom       float64   `json:"zoom"`
	Bearing    float64   `json:"bearing"`
	Pitch      float64   `json:"pitch"`
	Created    time.Time `json:"created"`
	Id         string    `json:"id"`
	Modified   time.Time `json:"modified"`
	Owner      string    `json:"owner"`
	Visibility string    `json:"visibility"`
}

func GetStyles(accessToken string, username string) ([]ListStyle, error) {
	if username == "" {
		return nil, fmt.Errorf("Username missing, please provide with --username/-u")
	}

	if accessToken == "" {
		return nil, fmt.Errorf("Access token is missing, please provide with --access-token or env: MAPBOX_ACCESS_TOKEN")
	}

	client := GetDefaultClient(accessToken)

	endpoint := fmt.Sprintf("/styles/v1/%v", username)

	res, err := client.Get(endpoint, nil, nil)
	if err != nil {
		log.Fatalf("Failed to fetch styles for user: %v", err)
	}

	var styles []ListStyle
	if err := json.Unmarshal(res.Payload, &styles); err != nil {
		log.Fatalf("Failed to parse styles: %v", err)
	}

	out, err := json.Marshal(styles)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))

	return styles, nil
}
