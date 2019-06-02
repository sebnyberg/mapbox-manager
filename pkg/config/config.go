package config

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/user"
	"path"
	"strings"

	"github.com/spf13/viper"
)

const (
	DEFAULT_APP_DIR     string = ".mapboxcli"
	DEFAULT_CONFIG_NAME string = "config.yml"
)

var allowedOptions []string = []string{
	"access-token",
	"username",
	"style-id",
}

func GetOptions() []string {
	return allowedOptions
}

func GetDir() string {
	p := GetPath()

	return path.Dir(p)
}

func GetPath() string {
	user, err := user.Current()
	if err != nil {
		log.Fatalf("Failed to get current user: %v", err)
	}

	configPath := fmt.Sprintf("%v/%v/%v", user.HomeDir, DEFAULT_APP_DIR, DEFAULT_CONFIG_NAME)

	return configPath
}

func Reset() {
	os.Remove(GetPath())
}

func Write() error {
	configPath := GetPath()
	configDir := GetDir()

	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		os.MkdirAll(configDir, os.ModePerm)
	}

	// config is empty
	writingValues := false

	v := viper.New()

	for _, opt := range allowedOptions {
		val := viper.GetString(opt)
		if val != "" {
			writingValues = true
		}
		v.Set(opt, val)
	}

	if !writingValues {
		return errors.New("no options passed to command")
	}

	if err := v.WriteConfigAs(configPath); err != nil {
		log.Fatalf("Failed to write config: %v", err)
	}

	fmt.Printf("Saved config to: %v\n", configPath)

	return nil
}

func String() (string, error) {
	configDir := GetDir()

	v := viper.New()

	v.AddConfigPath(configDir)

	if err := v.ReadInConfig(); err != nil {
		return "", err
	}

	sb := strings.Builder{}

	for _, opt := range allowedOptions {
		stringVal := v.GetString(opt)
		if stringVal != "" {
			sb.WriteString(fmt.Sprintf("%v: %v\n", opt, stringVal))
		}
	}

	return sb.String(), nil
}
