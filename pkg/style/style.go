package style

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/sebnyberg/mapboxcli/pkg/mapbox"
)

func GetAll(outputFormat string, accessToken string, username string) (string, error) {
	if err := checkFormatAvailable(outputFormat); err != nil {
		return "", err
	}

	styles, err := mapbox.GetStyles(accessToken, username)
	if err != nil {
		return "", err
	}

	s, err := formatStyleList(styles, outputFormat)
	if err != nil {
		return "", err
	}

	return s, nil
}

func Get(outputFormat string, accessToken string, username string, styleID string, draft bool) (string, error) {
	if err := checkFormatAvailable(outputFormat); err != nil {
		return "", err
	}

	style, err := mapbox.GetStyle(accessToken, username, styleID, draft)
	if err != nil {
		return "", err
	}

	s, err := formatStyle(*style, outputFormat)

	return s, err
}

func getSourcesFromSourcesURL(url string) ([]string, error) {
	mapboxPattern := regexp.MustCompile("mapbox://")

	parts := mapboxPattern.Split(url, -1)
	if len(parts) != 2 || parts[0] != "" || parts[1] == "" {
		return nil, errors.New("Failed to parse sources URL")
	}

	return strings.Split(parts[1], ","), nil
}

// getNonMapboxSources retrieves any sources listed in the sources URL that belongs to mapbox
func getMapboxSourcesFromURL(sourcesURL string) ([]string, error) {
	allSources, err := getSourcesFromSourcesURL(sourcesURL)
	if err != nil {
		return nil, err
	}

	mapboxSources := make([]string, 0)
	for _, source := range allSources {
		sourceParts := strings.Split(source, ".")
		if len(sourceParts) != 2 {
			return nil, fmt.Errorf("Encountered invalid source in the source URL: %v", source)
		}
		username := sourceParts[0]
		if username == "mapbox" {
			mapboxSources = append(mapboxSources, source)
		}
	}

	return mapboxSources, nil
}

// setTilsetsForLayers updates the layers of the passed style so that they point to
// the tilesets listed in the map
func setTilesetsForLayers(style *mapbox.Style, layerToTilesets map[string]string) error {
	foundLayers := 0
	for i, layer := range style.Layers {
		if tilesetName, exists := layerToTilesets[layer.ID]; exists {
			foundLayers++
			style.Layers[i].SourceLayer = tilesetName
		}
	}

	if foundLayers < len(layerToTilesets) {
		return errors.New("One or more layer names is in valid")
	}

	return nil
}

// SetLayerTilesets - update the underlying data sources for different layers in the style
func SetLayerTilesets(accessToken string, username string, styleID string, draft bool, layersToTilesets map[string]string, force bool, verbose bool) (string, error) {
	// Update source layers in the style
	style, err := mapbox.GetStyle(accessToken, username, styleID, draft)
	if err != nil {
		return "", err
	}
	setTilesetsForLayers(style, layersToTilesets)

	// Each source style that is listed in the layers of the map style may belong to the user or mapbox
	// If the source style belongs to a tileset owned by the user, the URL in style.Sources.Composite.URL
	// must contain the tileset ID in order for the source layer to work
	// Approach:
	// 1. Fetch all tilesets uploaded by the user, map tileset layer names to tileset ids
	// 2. For each layer in the update style, check whether the source layer is a user-uploaded layer name
	//    If it is: add the tileset id to a set of tileset ids that should be in the sources URL
	// 3. Concatenate the mapbox sources with the tileset ids to form the new sources URL
	tilesets, err := mapbox.GetTilesets(accessToken, username)
	if err != nil {
		return "", err
	}
	userTilesetLayerToTilesetIDs := make(map[string]string)
	for _, tileset := range tilesets {
		if tilesetID, exists := userTilesetLayerToTilesetIDs[tileset.Name]; exists {
			fmt.Printf("Two tilesets with id %v and %v have the same tileset layer name (%v), this may result in odd behaviour", tilesetID, tileset.ID, tileset.Name)
			if !force {
				return "", errors.New("Aborted update due to layer naming ambiguety, run with --force to force update")
			}
		}
		userTilesetLayerToTilesetIDs[tileset.Name] = tileset.ID
	}
	// fmt.Print("Tileset to layer map:")
	// for k, v := range userTilesetLayerToTilesetIDs {

	// }

	// List source layers in the style which are custom (user) tileset layers
	userTilesetsInStyle := make(map[string]bool)
	for _, layer := range style.Layers {
		if layer.SourceLayer != "" {
			if tilesetID, exists := userTilesetLayerToTilesetIDs[layer.SourceLayer]; exists {
				if verbose {
					fmt.Printf("Mapped layer to tileset id %v -> %v\n", layer.SourceLayer, tilesetID)
				}
				userTilesetsInStyle[tilesetID] = true
			} else {
				if verbose {
					fmt.Printf("Found unmapped layer: %v\n", layer.SourceLayer)
				}
			}
		}
	}

	// List mapbox sources in the style header
	if style.Sources.Composite == nil {
		return "", errors.New("Mapbox style does not contain any data sources, please add one in mapbox studio to initialize the object")
	}
	sourcesURL := style.Sources.Composite.URL
	mapboxSources, err := getMapboxSourcesFromURL(sourcesURL)
	if err != nil {
		return "", err
	}
	mapboxSourcesStr := strings.Join(mapboxSources, ",")
	userTilesetsInStyleList := make([]string, 0)
	for tileset := range userTilesetsInStyle {
		userTilesetsInStyleList = append(userTilesetsInStyleList, tileset)
	}
	userTilesetsInStyleStr := strings.Join(userTilesetsInStyleList, ",")
	style.Sources.Composite.URL = fmt.Sprintf("mapbox://%v,%v", mapboxSourcesStr, userTilesetsInStyleStr)

	// These are created by Mapbox
	style.Created = nil
	style.Modified = nil

	respBytes, err := mapbox.UpdateStyle(accessToken, username, styleID, draft, *style)

	return string(respBytes), err
}
