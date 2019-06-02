package style

import (
	"fmt"
	"github.com/sebnyberg/mapboxcli/pkg/mapbox"
)

func GetAsTable(accessToken string, username string) (string, error) {
	fmt.Printf("accessToken: %v\n", accessToken)
	fmt.Printf("username: %v\n", username)

	_, err := mapbox.GetStyles(accessToken, username)
	if err != nil {
		return "", err
	}

	return "", nil
}
