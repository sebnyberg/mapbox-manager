package cmd

import (
	"fmt"

	"github.com/sebnyberg/mapboxcli/pkg/layer"

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
}

func bindUpdateLayerFlags() {
	viper.BindPFlag("style-id", updateLayerCmd.PersistentFlags().Lookup("style-id"))
	viper.BindPFlag("layer-id", updateLayerCmd.PersistentFlags().Lookup("layer-id"))
	viper.BindPFlag("draft", updateLayerCmd.PersistentFlags().Lookup("draft"))
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

		respBytes, err := layer.SetTileset(accessToken, username, styleID, layerID, draft, newTilesetID)
		if err != nil {
			return err
		}
		if printResponse {
			fmt.Println(string(respBytes))
		}

		fmt.Printf("Successfully set the tileset of layer %v to %v\n", layerID, newTilesetID)

		return nil
	},
}

var updateLayerCmd = &cobra.Command{
	Use:   "layer",
	Short: "update layer",
	Long:  `mapbox update layer - update a layer`,
}

func init() {
	updateCmd.PersistentFlags().StringP("username", "u", "", "username (required)")
	updateCmd.PersistentFlags().String("access-token", "", "access token (required)")
	updateCmd.PersistentFlags().BoolP("print-response", false, "p", "print response from Mapbox API")

	// update layer
	updateLayerCmd.PersistentFlags().String("style-id", "", "style id")
	updateLayerCmd.PersistentFlags().String("layer-id", "", "layer id")
	updateLayerCmd.PersistentFlags().Bool("draft", false, "retrieve draft version")

	// update layer set-tileset
	updateLayerSetTilesetCmd.Flags().String("tileset-id", "", "new tileset id")

	updateLayerCmd.AddCommand(updateLayerSetTilesetCmd)

	updateCmd.AddCommand(updateLayerCmd)

	rootCmd.AddCommand(updateCmd)

	// styleCmd.AddCommand(getCmd)
}
