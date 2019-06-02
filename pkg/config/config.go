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

var allowedOptions map[string]bool = map[string]bool{
	"access-token": true,
	"username":     true,
	"style-id":     true,
}

var sensitiveOptions map[string]bool = map[string]bool{
	"access-token": true,
}

func GetOptions() []string {
	optsList := make([]string, 0)
	for option := range allowedOptions {
		optsList = append(optsList, option)
	}
	return optsList
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

	for opt, _ := range allowedOptions {
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

func ToString(showSensitive bool) (string, error) {
	configDir := GetDir()

	v := viper.New()

	v.AddConfigPath(configDir)

	if err := v.ReadInConfig(); err != nil {
		return "", err
	}

	sb := strings.Builder{}

	for opt, _ := range allowedOptions {
		// Hide sensitive information
		if !showSensitive {
			if sensitive, ok := sensitiveOptions[opt]; ok && sensitive {
				sb.WriteString(fmt.Sprintf("%v: HIDDEN (show with --show-sensitive)\n", opt))
				continue
			}
		}

		stringVal := v.GetString(opt)
		if stringVal != "" {
			sb.WriteString(fmt.Sprintf("%v: %v\n", opt, stringVal))
		}
	}

	return sb.String(), nil
}
