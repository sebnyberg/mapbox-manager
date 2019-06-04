package layer

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/olekukonko/tablewriter"
	"github.com/sebnyberg/mapboxcli/pkg/mapbox"
)

func formatLayersList(layers []mapbox.Layer, outputFormat string) (string, error) {
	switch outputFormat {
	case "table":
		return LayersToTable(layers)
	case "json":
		return LayersToJSON(layers)
	}
	panic("Unsupported output format..")
}

func formatLayer(layer mapbox.Layer, outputFormat string) (string, error) {
	switch outputFormat {
	case "table":
		return LayersToTable([]mapbox.Layer{layer})
	case "json":
		return LayerToJSON(layer)
	}
	panic("Unsupported output format..")
}

func LayersToTable(layers []mapbox.Layer) (string, error) {
	var buf bytes.Buffer

	table := tablewriter.NewWriter(&buf)
	table.SetHeader([]string{"id", "type", "source", "source-layer", "minzoom", "maxzoom"})

	data := make([][]string, len(layers))

	for _, layer := range layers {
		rowData := []string{
			layer.ID,
			layer.Type,
			layer.Source,
			layer.SourceLayer,
			fmt.Sprintf(".2%f", layer.MinZoom),
			fmt.Sprintf(".2%f", layer.MaxZoom),
		}
		data = append(data, rowData)
	}
	table.AppendBulk(data)
	table.Render()

	return buf.String(), nil
}

func LayersToJSON(layers []mapbox.Layer) (string, error) {
	b, err := json.Marshal(layers)

	if err != nil {
		return "", fmt.Errorf("failed to parse layers: %v", err)
	}

	return string(b), nil
}

func LayerToJSON(layer mapbox.Layer) (string, error) {
	b, err := json.Marshal(layer)

	if err != nil {
		return "", fmt.Errorf("failed to parse layer: %v", err)
	}

	return string(b), nil
}

func GetAll(outputFormat string, accessToken string, username string, styleID string, draft bool) (string, error) {
	if err := checkFormatAvailable(outputFormat); err != nil {
		return "", err
	}

	style, err := mapbox.GetStyle(accessToken, username, styleID, draft)
	if err != nil {
		return "", err
	}

	s, err := formatLayersList(style.Layers, outputFormat)
	if err != nil {
		return "", err
	}

	return s, nil
}

func Get(outputFormat string, accessToken string, username string, styleID string, layerID string, draft bool) (string, error) {
	if err := checkFormatAvailable(outputFormat); err != nil {
		return "", err
	}

	style, err := mapbox.GetStyle(accessToken, username, styleID, draft)
	if err != nil {
		return "", err
	}

	for _, layer := range style.Layers {
		if layer.ID == layerID {
			return formatLayer(layer, outputFormat)
		}
	}

	return "", fmt.Errorf("could not find layer with id %v", layerID)
}

func SetTileset(accessToken string, username string, styleID string, layerID string, draft bool, newTilesetID string) ([]byte, error) {
	style, err := mapbox.GetStyle(accessToken, username, styleID, draft)
	if err != nil {
		return err
	}

	for index, layer := range style.Layers {
		if layer.ID == layerID {
			style.Modified = nil
			style.Created = nil
			style.Layers[index].SourceLayer = newTilesetID

			respBytes, err := mapbox.UpdateStyle(accessToken, username, styleID, draft, *style)
			if err != nil {
				return nil, err
			}
			return respBytes, nil
		}
	}

	return nil, fmt.Errorf("could not find layer with id %v", layerID)
}
