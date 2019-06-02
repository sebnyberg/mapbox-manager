package httpclient

import (
	"fmt"
)

var statusString map[int]string = map[int]string{
	200: "OK",
	201: "Created",
	202: "Accepted",
	204: "No Content",
	301: "Moved Permanently",
	304: "Not Modified",
	400: "Bad Rewquest",
	401: "Unauthorized",
	402: "Payment Required",
	403: "Forbidden",
	404: "Not Found",
	405: "Method Not Allowed",
	408: "Request timed out",
	429: "Too Many Requests",
	500: "Internal Server Error",
	501: "Not Implemented",
	502: "Bad Gateway",
	503: "Service Unavailable",
}

func GetStatusString(statusCode int) string {
	if val, ok := statusString[statusCode]; ok {
		return val
	}

	return fmt.Sprint("Uknown", statusCode)
}