package mapbox

import (
	"fmt"

	"github.com/sebnyberg/mapboxcli/pkg/httpclient"
)

var API_URL string = "https://api.mapbox.com"

func GetDefaultClient(accessToken string) httpclient.Client {
	if accessToken == "" {
		fmt.Println("Missing access token, please provide with --access-token or env: MAPBOX_ACCESS_TOKEN")
	}

	clientConfig := httpclient.Config{
		BaseURL: API_URL,
		DefaultQueryParams: map[string]interface{}{
			"access_token": accessToken,
		},
		DefaultHeaders: map[string]string{
			"Content-Type": "application/json",
		},
	}

	return httpclient.NewClient().WithConfig(&clientConfig)
}
