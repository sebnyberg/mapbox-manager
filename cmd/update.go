package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/sebnyberg/mapboxcli/pkg/layer"
	"github.com/sebnyberg/mapboxcli/pkg/style"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "update resources",
	Long:  `mapbox update - update resources `,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func bindUpdateFlags() {
	viper.BindPFlag("access-token", updateCmd.PersistentFlags().Lookup("access-token"))
	viper.BindPFlag("username", updateCmd.PersistentFlags().Lookup("username"))
	viper.BindPFlag("style-id", updateCmd.PersistentFlags().Lookup("style-id"))
	viper.BindPFlag("force", updateCmd.PersistentFlags().Lookup("force"))
	viper.BindPFlag("verbose", updateCmd.PersistentFlags().Lookup("verbose"))
}

func bindUpdateLayerFlags() {
	viper.BindPFlag("layer-id", updateLayerCmd.PersistentFlags().Lookup("layer-id"))
	viper.BindPFlag("draft", updateLayerCmd.PersistentFlags().Lookup("draft"))
}

func bindUpdateStyleFlags() {
	viper.BindPFlag("draft", updateStyleCmd.PersistentFlags().Lookup("draft"))
}

var updateLayerSetTilesetCmd = &cobra.Command{
	Use:   "set-tileset",
	Short: "update layer set-tileset",
	Long:  `mapbox update layer set-tileset - sets the layer tileset (source-layer) `,
	RunE: func(cmd *cobra.Command, args []string) error {
		bindUpdateFlags()
		bindUpdateLayerFlags()

		viper.BindPFlag("tileset-id", cmd.Flags().Lookup("tileset-id"))

		err := exitIfMissing([]string{"username", "access-token", "style-id", "layer-id", "tileset-id"})
		if err != nil {
			return err
		}

		accessToken := viper.GetString("access-token")
		username := viper.GetString("username")

		styleID := viper.GetString("style-id")
		layerID := viper.GetString("layer-id")
		draft := viper.GetBool("draft")
		printResponse := viper.GetBool("print-response")

		newTilesetID := viper.GetString("tileset-id")

<<<<<<< HEAD
		respBytes, err := layer.SetTileset(accessToken, username, styleID, layerID, draft, newTilesetID)
		if err != nil {
			return err
		}
		if printResponse {
			var prettyJSON bytes.Buffer
			err := json.Indent(&prettyJSON, respBytes, "", "  ")
			if err != nil {
				return err
			}
			fmt.Println(string(prettyJSON.Bytes()))
		} else {
			fmt.Printf("Successfully set the tileset of layer %v to %v\n", layerID, newTilesetID)
		}
=======
		updateResponse, err := layer.SetTileset(accessToken, username, styleID, layerID, draft, newTilesetID)
		if err != nil {
			return err
		}

		fmt.Println(updateResponse)
>>>>>>> 6974440... Improve update command to take into account sources

		return nil
	},
}

var updateStyleSetLayerTilesetCmd = &cobra.Command{
	Use:   "set-layer-tileset",
	Short: "set tilesets for many layers in a style",
	Long: `mapbox update style set-layer-tileset - sets layer-tileset mapping for many layers at one time

Example:
	mapbox update layer set-layer-tileset --layer-to-tileset "layer1=tileset1,layer2=tileset1"
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		bindUpdateFlags()
		bindUpdateStyleFlags()

		viper.BindPFlag("layer-to-tileset", cmd.Flags().Lookup("layer-to-tileset"))

		err := exitIfMissing([]string{"username", "access-token", "style-id", "layer-to-tileset"})
		if err != nil {
			return err
		}

		accessToken := viper.GetString("access-token")
		username := viper.GetString("username")

		styleID := viper.GetString("style-id")
		draft := viper.GetBool("draft")
		force := viper.GetBool("force")
		layerToTilesetsString := viper.GetString("layer-to-tileset")
		verbose := viper.GetBool("verbose")

		layersToTilesets, err := parseKeyValueStringAsMap(layerToTilesetsString)
		if err != nil {
			return err
		}

		updateResponse, err := style.SetLayerTilesets(accessToken, username, styleID, draft, layersToTilesets, force, verbose)
		if err != nil {
			return err
		}
		fmt.Println(updateResponse)

		return nil
	},
}

var updateStyleCmd = &cobra.Command{
	Use:   "style",
	Short: "update style",
	Long:  `mapbox update style - update a style`,
}

var updateLayerCmd = &cobra.Command{
	Use:   "layer",
	Short: "update layer",
	Long:  `mapbox update layer - update a layer`,
}

func init() {
	updateCmd.PersistentFlags().StringP("username", "u", "", "username (required)")
	updateCmd.PersistentFlags().String("access-token", "", "access token (required)")
	updateCmd.PersistentFlags().String("style-id", "", "style id")
	updateCmd.PersistentFlags().BoolP("force", "f", false, "force update in the face of warnings")
	updateCmd.PersistentFlags().BoolP("verbose", "v", false, "verbose")

	// update layer
	updateLayerCmd.PersistentFlags().String("layer-id", "", "layer id")
	updateLayerCmd.PersistentFlags().Bool("draft", false, "retrieve draft version")

	// update layer set-tileset
	updateLayerSetTilesetCmd.Flags().String("tileset-id", "", "new tileset id")

	updateLayerCmd.AddCommand(updateLayerSetTilesetCmd)
	updateCmd.AddCommand(updateLayerCmd)

	// update style
	updateStyleCmd.PersistentFlags().Bool("draft", false, "retrieve draft version")

	// update style set-layer-tileset
	updateStyleSetLayerTilesetCmd.Flags().String("layer-to-tileset", "", "map of layer=tileset values")

	updateStyleCmd.AddCommand(updateStyleSetLayerTilesetCmd)
	updateCmd.AddCommand(updateStyleCmd)

	rootCmd.AddCommand(updateCmd)
}
