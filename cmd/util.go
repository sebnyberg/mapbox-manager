package cmd

import (
	"errors"
	"strings"
)

func parseKeyValueStringAsMap(kvString string) (map[string]string, error) {
	m := make(map[string]string)

	commaItems := strings.Split(kvString, ",")
	if len(commaItems) <= 0 {
		return nil, errors.New("Failed to parse key-value string. String was empty")
	}
	for _, commaItem := range commaItems {
		equalItems := strings.Split(commaItem, "=")
		if len(equalItems) != 2 {
			return nil, errors.New("Failed to parse key-value string. Splitting by , then = did not result in pairs")
		}

		m[equalItems[0]] = equalItems[1]
	}

	return m, nil
}
