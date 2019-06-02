package cmd

import (
	"fmt"

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

var stylesCmd = &cobra.Command{
	Use:   "styles",
	Short: "list styles",
	Long: `mapbox get styles - list styles
Note: listed styles contain less detail than styles retrieved in isolation.

For more detailed information about a style, use
	mapbox get style`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := exitIfMissing([]string{"username", "access-token"})
		if err != nil {
			return err
		}

		username := viper.GetString("username")
		accessToken := viper.GetString("access-token")
		outputFormat := viper.GetString("output")

		s, err := style.GetAll(accessToken, username, outputFormat)
		if err != nil {
			return err
		}
		fmt.Printf(s)

		return nil
	},
}

var styleCmd = &cobra.Command{
	Use:   "style",
	Short: "get style",
	Long:  `mapbox get style - get style`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := exitIfMissing([]string{"username", "access-token", "style-id"})
		if err != nil {
			return err
		}

		username := viper.GetString("username")
		accessToken := viper.GetString("access-token")
		outputFormat := viper.GetString("output")
		styleId := viper.GetString("style-id")
		draft := viper.GetBool("draft")

		s, err := style.Get(accessToken, username, styleId, outputFormat, draft)
		if err != nil {
			fmt.Println(err)
			return err
		}

		fmt.Println(s)

		return nil
	},
}

func init() {
	getCmd.PersistentFlags().StringP("username", "u", "", "username (required)")
	getCmd.PersistentFlags().String("access-token", "", "access token (required)")
	getCmd.PersistentFlags().StringP("output", "o", "table", "output format, default: 'table', options: 'table', 'id', yaml', 'json'")

	viper.BindPFlag("username", getCmd.PersistentFlags().Lookup("username"))
	viper.BindPFlag("access-token", getCmd.PersistentFlags().Lookup("access-token"))
	viper.BindPFlag("output", getCmd.PersistentFlags().Lookup("output"))

	styleCmd.Flags().String("style-id", "", "style id")
	styleCmd.Flags().Bool("draft", false, "retrieve draft version")

	viper.BindPFlag("style-id", styleCmd.Flags().Lookup("style-id"))
	viper.BindPFlag("draft", styleCmd.Flags().Lookup("draft"))

	getCmd.AddCommand(stylesCmd)
	getCmd.AddCommand(styleCmd)

	rootCmd.AddCommand(getCmd)

	// styleCmd.AddCommand(getCmd)
}
