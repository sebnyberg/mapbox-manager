package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/viper"
)

// ! Explicit by design
var flagHelp map[string]string = map[string]string{
	"access-token": fmt.Sprintf("Access token is required, provide with --access-token, %v_ACCESS_TOKEN or config", ENV_PREFIX),
	"username":     fmt.Sprintf("Username is required, provide with --username, %v_USERNAME or config", ENV_PREFIX),
	"style-id":     fmt.Sprintf("Style ID is required, provide with --style-id, %v_STYLE_ID or config", ENV_PREFIX),
	"layer-id":     "Layer ID is required, provide with --layer-id",
	"dataset-id":   "Dataset ID is required, provide with --dataset-id",
	"tileset-id":   "Tileset ID is required, provide with --tileset-id",
	"layer-to-tileset": `Layer to tileset mapping is required, 
Example: --layer-to-tileset 'layer1=tileset1,layer2=tileset2'`,
}

func exitIfMissing(flagNames []string) error {
	for _, flagName := range flagNames {
		flagValue := viper.GetString(flagName)

		// if flag is not set, print help
		if flagValue == "" {
			helpMsg, ok := flagHelp[flagName]
			if !ok {
				return fmt.Errorf("Help for flag %v not defined\n", flagName)
			}
			return errors.New(helpMsg)
		}
	}

	return nil
}
