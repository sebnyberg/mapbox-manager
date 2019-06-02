package style

import (
	"fmt"
)


var availableFormats map[string]bool = map[string]bool {
	"table": true,
	"id": false,
	"yaml": false,
	"json": false,
}

func checkFormatAvailable(outputFormat string) error {
	if _, ok := availableFormats[outputFormat]; !ok {
		return fmt.Errorf("format not yet available: %v", outputFormat)
	}

	return nil
}