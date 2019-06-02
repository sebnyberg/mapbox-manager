package cmd

import (
	"fmt"
	"strings"

	"github.com/sebnyberg/mapboxcli/pkg/config"

	"github.com/spf13/viper"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "manage configuration",
	Long: `mapbox config
Stores configuration in ~/.mapboxcli/config.yml

Supported commands:

mapbox config set --username myuser --access-token mytoken --style-id abc123
mapbox config reset
mapbox config show

Flags set in the config will be automatically passed to other commands.

Precedence ordering:
	1. flag
	2. env
	3. config
`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var setConfigCmd = &cobra.Command{
	Use:   "set",
	Short: "set configuration",
	Long: `mapbox config set --<flagname> <value>

Stores commandline flags for re-use in ~/.mapboxcli/config.yml

Supported flags:

--` + strings.Join(config.GetOptions(), "\n--"),
	Run: func(cmd *cobra.Command, args []string) {
		err := config.Write()
		if err != nil {
			fmt.Printf("Failed to set config: %v\n", err)
			fmt.Println("See `mapbox config set --help`")
		}
	},
}

var resetConfigCmd = &cobra.Command{
	Use:   "reset",
	Short: "reset configuration",
	Long: `mapbox config reset

Resets configuration by deleting ~/.mapboxcli/config.yml`,
	Run: func(cmd *cobra.Command, args []string) {
		config.Reset()
	},
}

var showConfigCmd = &cobra.Command{
	Use:   "show",
	Short: "show current configuration",
	Long: `mapbox config show 

Lists configuration options found in ~/.mapboxcli/config.yml`,
	Run: func(cmd *cobra.Command, args []string) {
		showSensitive := viper.GetBool("show-sensitive")

		s, err := config.ToString(showSensitive)
		if err != nil {
			fmt.Println("Configuration not set. See `mapbox config set`")
		} else {
			fmt.Print(s)
		}
	},
}

func init() {
	setConfigCmd.Flags().StringP("username", "u", "", "Username")
	setConfigCmd.Flags().String("access-token", "", "Access token")
	setConfigCmd.Flags().String("style-id", "", "Style id")

	viper.BindPFlag("access-token", setConfigCmd.Flags().Lookup("access-token"))
	viper.BindPFlag("username", setConfigCmd.Flags().Lookup("username"))
	viper.BindPFlag("style-id", setConfigCmd.Flags().Lookup("style-id"))


	showConfigCmd.Flags().Bool("show-sensitive", false, "show sensitive information, default: false")
	viper.BindPFlag("show-sensitive", showConfigCmd.Flags().Lookup("show-sensitive"))

	configCmd.AddCommand(setConfigCmd)
	configCmd.AddCommand(resetConfigCmd)
	configCmd.AddCommand(showConfigCmd)

	rootCmd.AddCommand(configCmd)
}
