package style

import (
	"bytes"

	"github.com/olekukonko/tablewriter"
	"github.com/sebnyberg/mapboxcli/pkg/mapbox"
)

func toFormat(styles []mapbox.ListStyle, outputFormat string) string {
	switch outputFormat {
	case "table":
		return StylesListToTable(styles)
	}
	panic("Unsupported output format..")
}

func StylesListToTable(styles []mapbox.ListStyle) string {
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

	return buf.String()
}

func Get(accessToken string, username string, outputFormat string) (string, error) {
	if err := checkFormatAvailable(outputFormat); err != nil {
		return "", err
	}

	styles, err := mapbox.GetStyles(accessToken, username)
	if err != nil {
		return "", err
	}

	s := toFormat(styles, outputFormat)

	return s, nil
}
