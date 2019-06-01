package cmd

import (
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "mapbox config - store configuration for re-use",
	Long: `mapbox config
Stores commandline flags for re-use in ~/.mapboxcli/config.yml

Supported commands:

mapbox config set --username myuser --access-token mytoken --style-id abc123
mapbox config reset
mapbox config show

Any flags set in the config will be automatically passed to other commands.

Precedence is in following order:
	1. flag
	2. env
	3. config
`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		// if err := config.WriteConfig(); err != nil {
		// 	log.Fatalf("Failed to write config: %v", err)
		// }
	},
}

func init() {
	configCmd.PersistentFlags().StringP("username", "u", "", "Username")
	configCmd.PersistentFlags().String("access-token", "", "Access token")
	configCmd.PersistentFlags().String("style-id", "", "Style id")

	viper.BindPFlag("access-token", configCmd.Flags().Lookup("access-token"))
	viper.BindPFlag("username", configCmd.Flags().Lookup("username"))
	viper.BindPFlag("style-id", configCmd.Flags().Lookup("style-id"))

	rootCmd.AddCommand(configCmd)
}
