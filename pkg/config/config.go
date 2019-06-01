package config

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"path"

	"github.com/spf13/viper"
)

const (
	DEFAULT_APP_DIR     string = ".mapboxcli"
	DEFAULT_CONFIG_NAME string = "config.yml"
)

func GetConfigPath() string {
	user, err := user.Current()
	if err != nil {
		log.Fatalf("Failed to get current user: %v", err)
	}

	configPath := fmt.Sprintf("%v/%v/%v", user.HomeDir, DEFAULT_APP_DIR, DEFAULT_CONFIG_NAME)

	return configPath
}

func WriteConfig() error {
	configPath := GetConfigPath()

	configDir := path.Dir(configPath)

	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		os.MkdirAll(configDir, os.ModePerm)
	}
	fmt.Printf("access token: %v\n", viper.GetString("access-token"))
	fmt.Printf("username: %v\n", viper.GetString("username"))

	fmt.Printf("Saving to path: %v\n", configPath)

	if err := viper.WriteConfigAs(configPath); err != nil {
		log.Fatalf("Failed to write config: %v", err)
	}

	return nil
}
