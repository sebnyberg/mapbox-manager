package mapbox

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type Tileset struct {
	ID          string    `json:"id"`
	Type        string    `json:"type"`
	Name        string    `json:"name"`
	Description string    `json:"descriptionm"`
	Center      []float64 `json:"center,omitempty"`
	FileSize    int64     `json:"filesize"`
	Created     time.Time `json:"created"`
	Modified    time.Time `json:"modified"`
	Visibility  string    `json:"visibility"`
	Status      string    `json:"status"`
}

func GetTilesets(accessToken string, username string) ([]Tileset, error) {
	client := GetDefaultClient(accessToken)

	endpoint := fmt.Sprintf("/tilesets/v1/%v", username)

	res, err := client.Get(endpoint, nil, nil)
	if err != nil {
		log.Fatalf("Failed to fetch tilesets for user: %v", err)
	}

	if res.StatusCode > 200 {
		return nil, fmt.Errorf("failed to fetch tilesets: %v", GetErrorMessage(res.StatusCode, res.Payload))
	}

	var tilesets []Tileset
	if err := json.Unmarshal(res.Payload, &tilesets); err != nil {
		log.Fatalf("Failed to parse tilesets: %v", err)
	}

	return tilesets, nil
}
