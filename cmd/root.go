package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "mapbox",
	Short: "mapbox is a CLI wrapper for the Mapbox API",
	Long: `mapbox - a CLI wrapper for the Mapbox API

The API key and username can be set as environment variables
MAPBOX_API_KEY and MAPBOX_USERNAME respectively.
	`,
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version",
	Long:  "Print the version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("gchat version %v", "0.1.0")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	viper.SetEnvPrefix("MAPBOX")
	viper.AutomaticEnv()

	replacer := strings.NewReplacer("-", "_")
	viper.SetEnvKeyReplacer(replacer)

	rootCmd.PersistentFlags().String("api-key", "", "Mapbox API key")

	viper.BindPFlag("api-key", rootCmd.PersistentFlags().Lookup("api-key"))

	rootCmd.AddCommand(versionCmd)
}
