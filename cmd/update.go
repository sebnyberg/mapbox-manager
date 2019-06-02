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

		viper.BindPFlag("new-tileset-id", updateLayerCmd.PersistentFlags().Lookup("tileset-id"))

		err := exitIfMissing([]string{"username", "access-token", "style-id", "layer-id", "new-tileset-id"})
		if err != nil {
			return err
		}

		accessToken := viper.GetString("access-token")
		username := viper.GetString("username")

		styleId := viper.GetString("style-id")
		layerId := viper.GetString("layer-id")
		draft := viper.GetBool("draft")

		newTilesetId := viper.GetString("new-tileset-id")

		err = layer.SetTileset(accessToken, username, styleId, layerId, draft, newTilesetId)
		if err != nil {
			return err
		}

		fmt.Printf("Successfully set the tileset of layer %v to %v\n", layerId, newTilesetId)

		return nil

		// s, err := style.GetAll(outputFormat, accessToken, username)
		// if err != nil {
		// 	return err
		// }
		// fmt.Printf(s)

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

	// update layer
	updateLayerCmd.PersistentFlags().StringP("style-id", "s", "", "style id")
	updateLayerCmd.PersistentFlags().StringP("layer-id", "l", "", "layer id")
	updateLayerCmd.PersistentFlags().Bool("draft", false, "retrieve draft version")

	// update layer set-tileset
	updateLayerSetTilesetCmd.PersistentFlags().Bool("new-tileset", false, "tileset id")

	updateLayerCmd.AddCommand(updateLayerSetTilesetCmd)

	updateCmd.AddCommand(updateLayerCmd)

	rootCmd.AddCommand(updateCmd)

	// styleCmd.AddCommand(getCmd)
}
