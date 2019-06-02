package mapbox

import (
	"github.com/sebnyberg/mapboxcli/pkg/httpclient"
	"fmt"
)

var statusErrors map[int]string = map[int]string{
	401: "please verify usename and access token",
}

func GetErrorMessage(statusCode int, payload []byte) string {
	if errMsg, ok := statusErrors[statusCode]; ok {
		return fmt.Sprintf("%v - %v", httpclient.GetStatusString(statusCode), errMsg)
	} else {
		return fmt.Sprintf("Undefined error - %v\n", statusCode)
	}
}