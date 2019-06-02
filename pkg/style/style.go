package style

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/olekukonko/tablewriter"
	"github.com/sebnyberg/mapboxcli/pkg/mapbox"
)

func toFormat(styles []mapbox.ListStyle, outputFormat string) (string, error) {
	switch outputFormat {
	case "table":
		return StylesListToTable(styles)
	case "json":
		return StylesListToJson(styles)
	}
	panic("Unsupported output format..")
}

func StylesListToTable(styles []mapbox.ListStyle) (string, error) {
	var buf bytes.Buffer

	table := tablewriter.NewWriter(&buf)
	table.SetHeader([]string{"id", "name", "owner"})

	data := make([][]string, len(styles))

	for _, style := range styles {
		rowData := []string{
			style.Id,
			style.Name,
			style.Owner,
		}
		data = append(data, rowData)
	}
	table.AppendBulk(data)
	table.Render()

	return buf.String(), nil
}

func StylesListToJson(styles []mapbox.ListStyle) (string, error) {
	jsonStr, err := json.Marshal(styles)

	if err != nil {
		return "", fmt.Errorf("failed parse response: %v", err)
	}

	return fmt.Sprintln(string(jsonStr)), nil
}

func StyleToTable() (string, error) {
	return "", errors.New("ERROR")
}

func GetAll(accessToken string, username string, outputFormat string) (string, error) {
	if err := checkFormatAvailable(outputFormat); err != nil {
		return "", err
	}

	styles, err := mapbox.GetStyles(accessToken, username)
	if err != nil {
		return "", err
	}

	s, err := toFormat(styles, outputFormat)

	return s, err
}
