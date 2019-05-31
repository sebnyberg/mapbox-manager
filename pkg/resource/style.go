package resources

import (
	"fmt"
	"time"
)

var something resource = resource{
	endpoint: "styles",
	version:  "v1",
}

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

func GetStyles(apiKey string, username string) error {
	if username == "" {
		return fmt.Errorf("Username missing, please provide with --username/-u")
	}

	return nil
}
