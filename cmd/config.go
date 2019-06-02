package cmd

import (
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "mbstyle config - store configuration for re-use",
	Long: `mbstyle config
Stores commandline flags for re-use in ~/.mapboxcli/config.yml

Supported commands:

mbstyle config set --username myuser --access-token mytoken --style-id abc123
mbstyle config reset
mbstyle config show

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
	Use:   "config",
	Short: "mbstyle config set - set configuration for re-use",
	Long: `mbstyle config set --<flagname> <value>
Stores commandline flags for re-use in ~/.mapboxcli/config.yml

Supported flags:

--username
--access-token
--style-id
`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
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
