package layer

import (
	"fmt"
)

var availableFormats map[string]bool = map[string]bool{
	"table": true,
	"id":    false,
	"yaml":  false,
	"json":  true,
}

func checkFormatAvailable(outputFormat string) error {
	if available, ok := availableFormats[outputFormat]; !ok || !available {
		return fmt.Errorf("format not yet available: %v", outputFormat)
	}

	return nil
}
