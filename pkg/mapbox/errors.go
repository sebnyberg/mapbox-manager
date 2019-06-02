package mapbox

import (
	"fmt"

	"github.com/sebnyberg/mapboxcli/pkg/httpclient"
)

var statusErrors map[int]string = map[int]string{
	401: "please verify usename and access token",
	422: "payload improperly formatted",
}

func GetErrorMessage(statusCode int, payload []byte) string {
	if errMsg, ok := statusErrors[statusCode]; ok {
		return fmt.Sprintf("%v - %v", httpclient.GetStatusString(statusCode), errMsg)
	} else {
		return fmt.Sprintf("Undefined error - %v\n", statusCode)
	}
}
