package style

import (
	"github.com/sebnyberg/mapboxcli/pkg/mapbox"
)

func GetAsTable(accessToken string, username string) (string, error) {
	_, err := mapbox.GetStyles(accessToken, username)
	if err != nil {
		return "", err
	}

	return "", nil
}
