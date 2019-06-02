package cmd

import (
	"fmt"

	"github.com/sebnyberg/mapboxcli/pkg/style"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// import (
// 	"log"

// 	"github.com/spf13/viper"

// 	"github.com/sebnyberg/mapboxcli/pkg/resource"
// 	"github.com/spf13/cobra"
// )

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
	Run: func(cmd *cobra.Command, args []string) {
		
		exitIfMissing([]string{"username", "access-token"})

		username := viper.GetString("username")
		accessToken := viper.GetString("access-token")

		s, err := style.GetAsTable(accessToken, username)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf(s)
	},
}

// var getCmd = &cobra.Command{
// 	Use:   "get",
// 	Short: "Retrieves all styles",
// 	Long:  `Retrieves all styles`,
// 	Run: func(cmd *cobra.Command, args []string) {
// 		accessToken := viper.GetString("access-token")
// 		username := viper.GetString("username")

// 		err := viper.WriteConfigAs("config.yml")
// 		if err != nil {
// 			log.Fatal("Failed to write config")
// 		}

// 		_, err = resource.GetStyles(accessToken, username)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 	},
// }

func init() {
	getCmd.PersistentFlags().StringP("username", "u", "", "Username (required)")
	getCmd.MarkFlagRequired("username")
	getCmd.PersistentFlags().String("access-token", "", "Access token (required)")
	getCmd.MarkFlagRequired("access-token")

	viper.BindPFlag("username", getCmd.PersistentFlags().Lookup("username"))
	viper.BindPFlag("access-token", getCmd.PersistentFlags().Lookup("access-token"))

	getCmd.AddCommand(stylesCmd)

	rootCmd.AddCommand(getCmd)

	// styleCmd.AddCommand(getCmd)
}
