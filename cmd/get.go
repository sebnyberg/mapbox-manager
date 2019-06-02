package cmd

import (
	"fmt"

	"github.com/sebnyberg/mapboxcli/pkg/layer"
	"github.com/sebnyberg/mapboxcli/pkg/style"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get resources",
	Long:  `mapbox get - retrieve resources `,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var getStylesCmd = &cobra.Command{
	Use:   "styles",
	Short: "list styles",
	Long: `mapbox get styles - list styles
Note: listed styles contain less detail than styles retrieved in isolation.

For more detailed information about a style, use
	mapbox get style`,
	RunE: func(cmd *cobra.Command, args []string) error {
		bindGetFlags()

		err := exitIfMissing([]string{"username", "access-token"})
		if err != nil {
			return err
		}

		accessToken := viper.GetString("access-token")
		username := viper.GetString("username")
		outputFormat := viper.GetString("output")

		s, err := style.GetAll(outputFormat, accessToken, username)
		if err != nil {
			return err
		}
		fmt.Printf(s)

		return nil
	},
}

var getStyleCmd = &cobra.Command{
	Use:   "style",
	Short: "get style",
	Long:  `mapbox get style - get style`,
	RunE: func(cmd *cobra.Command, args []string) error {
		bindGetFlags()

		viper.BindPFlag("style-id", cmd.Flags().Lookup("style-id"))
		viper.BindPFlag("draft", cmd.Flags().Lookup("draft"))

		err := exitIfMissing([]string{"username", "access-token", "style-id"})
		if err != nil {
			return err
		}

		accessToken := viper.GetString("access-token")
		username := viper.GetString("username")
		styleId := viper.GetString("style-id")

		draft := viper.GetBool("draft")
		outputFormat := viper.GetString("output")

		s, err := style.Get(outputFormat, accessToken, username, styleId, draft)
		if err != nil {
			return err
		}

		fmt.Println(s)

		return nil
	},
}

var getLayerCmd = &cobra.Command{
	Use:   "layer",
	Short: "get layer",
	Long: `mapbox get layer - retrieves a layer from a style

The layer id is the name of the layer in Mapbox Studio.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		bindGetFlags()

		viper.BindPFlag("style-id", cmd.Flags().Lookup("style-id"))
		viper.BindPFlag("layer-id", cmd.Flags().Lookup("layer-id"))
		viper.BindPFlag("draft", cmd.Flags().Lookup("draft"))

		err := exitIfMissing([]string{"username", "access-token", "style-id", "layer-id"})
		if err != nil {
			return err
		}

		accessToken := viper.GetString("access-token")
		username := viper.GetString("username")
		styleId := viper.GetString("style-id")
		layerId := viper.GetString("layer-id")
		draft := viper.GetBool("draft")

		outputFormat := viper.GetString("output")

		s, err := layer.Get(outputFormat, accessToken, username, styleId, layerId, draft)
		if err != nil {
			return err
		}

		fmt.Println(s)

		return nil
	},
}

var getLayersCmd = &cobra.Command{
	Use:   "layers",
	Short: "get layers",
	Long: `mapbox get layers - retrieves all layers from a style

The layer ids are the names of the layers in Mapbox Studio.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		bindGetFlags()

		viper.BindPFlag("style-id", cmd.Flags().Lookup("style-id"))
		viper.BindPFlag("draft", cmd.Flags().Lookup("draft"))

		err := exitIfMissing([]string{"username", "access-token", "style-id"})
		if err != nil {
			return err
		}

		accessToken := viper.GetString("access-token")
		username := viper.GetString("username")
		styleId := viper.GetString("style-id")
		draft := viper.GetBool("draft")

		outputFormat := viper.GetString("output")

		s, err := layer.GetAll(outputFormat, accessToken, username, styleId, draft)
		if err != nil {
			return err
		}

		fmt.Println(s)

		return nil
	},
}

func bindGetFlags() {
	viper.BindPFlag("username", getCmd.PersistentFlags().Lookup("username"))
	viper.BindPFlag("access-token", getCmd.PersistentFlags().Lookup("access-token"))
	viper.BindPFlag("output", getCmd.PersistentFlags().Lookup("output"))
}

func init() {
	getCmd.PersistentFlags().StringP("username", "u", "", "username (required)")
	getCmd.PersistentFlags().String("access-token", "", "access token (required)")
	getCmd.PersistentFlags().StringP("output", "o", "table", "output format, default: 'table', options: 'table', 'id', yaml', 'json'")

	// get style
	getStyleCmd.Flags().StringP("style-id", "s", "", "style id")
	getStyleCmd.Flags().Bool("draft", false, "retrieve draft version")

	// get layer
	getLayerCmd.Flags().StringP("style-id", "s", "", "style id")
	getLayerCmd.Flags().StringP("layer-id", "l", "", "layer id")
	getLayerCmd.Flags().Bool("draft", false, "retrieve draft version of style")

	// get layer
	getLayersCmd.Flags().StringP("style-id", "s", "", "style id")
	getLayersCmd.Flags().Bool("draft", false, "retrieve draft version of style")

	getCmd.AddCommand(getStylesCmd)
	getCmd.AddCommand(getStyleCmd)
	getCmd.AddCommand(getLayerCmd)
	getCmd.AddCommand(getLayersCmd)

	rootCmd.AddCommand(getCmd)

	// styleCmd.AddCommand(getCmd)
}
