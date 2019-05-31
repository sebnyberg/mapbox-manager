package cmd

import (
	"github.com/spf13/viper"

	"github.com/sebnyberg/mapboxcli/pkg/resource"
	"github.com/spf13/cobra"
)

var styleCmd = &cobra.Command{
	Use:   "style",
	Short: "style commands",
	Long:  `mapbox style`,
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieves all styles",
	Long:  `Retrieves all styles`,
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := viper.Get("api-key")
		username := viper.Get("username")

		resource.GetStyles(apiKey, username)
	},
}

func init() {
	getCmd.PersistentFlags().StringP("username", "u", "", "Username (required)")
	getCmd.MarkFlagRequired("username")

	viper.BindPFlag("username", getCmd.PersistentFlags().Lookup("username"))

	rootCmd.AddCommand(styleCmd)

	styleCmd.AddCommand(getCmd)
}
