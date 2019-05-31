package resource

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
	}

	return httpclient.NewClient().WithConfig(&clientConfig)
}
