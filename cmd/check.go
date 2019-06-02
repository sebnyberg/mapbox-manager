package cmd

import (
	"os"
	"github.com/spf13/viper"
	"fmt"
)

// ! Explicit by design
var flagHelp map[string]string = map[string]string {
	"access-token": fmt.Sprintf("Access token is required, provide with --access-token, %v_ACCESS_TOKEN or config", ENV_PREFIX),
	"username": fmt.Sprintf("Username is required, provide with --username, %v_USERNAME or config", ENV_PREFIX),
	"style-id": fmt.Sprintf("Style ID is required, provide with --style-id, %v_STYLE_ID or config", ENV_PREFIX),
	"layer-id": "Layer ID is required, provide with --style-id",
	"dataset-id": "Layer ID is required, provide with --style-id",
}

func exitIfMissing(flagNames []string) {
	for _, flagName := range flagNames {
		flagValue := viper.GetString(flagName)

		// if flag is not set, print help
		if flagValue == "" {
			helpMsg, ok := flagHelp[flagName]
			if !ok {
				errMsg := fmt.Sprintf("Help for flag %v not defined\n", flagName)
				panic(errMsg)
			}
			fmt.Println(helpMsg)
			os.Exit(1)
		}
	}
}