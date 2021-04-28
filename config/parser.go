package config

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const definitionsFolder = "./config/definitions/"

var environmentConfigFile = map[string]string{
	"development": "development.yaml",
	"staging":     "staging.yaml",
	"production":  "production.yaml",
}

// Parse function, parses the configuration and returns it.
func Parse() (*Config, error) {
	file, err := getConfigFile()
	if err != nil {
		pflag.PrintDefaults()
		return nil, err
	}

	// Read config from file
	viper.AddConfigPath(definitionsFolder)
	viper.SetConfigFile(definitionsFolder + file)
	viper.SetConfigType("yaml")
	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	// Read config from Env
	err = godotenv.Load()
	if err != nil {
		log.Printf("Failed to load .env: %v", err)
	}
	viper.SetEnvPrefix("app")
	viper.AutomaticEnv()

	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	// Validate
	if err := validateAPIConfig(&config); err != nil {
		return nil, err
	}
	if err := validateJWTConfig(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

// validateAPIConfig Validates API configuration.
func validateAPIConfig(config *Config) error {
	if config.API.Port <= 0 {
		return composeError("api.port", config.API.Port)
	}

	return nil
}

// validateJWTConfig Validates the JWT configuration.
func validateJWTConfig(config *Config) error {
	if config.JWT.Secret == "" {
		return composeError("jwt.secret", config.JWT.Secret)
	}

	return nil
}

// composeError Composes the error found for a config key.
func composeError(key string, value interface{}) error {
	return fmt.Errorf(`invalid config value "%v" for key "%s". please check the configuration files`, value, key)
}

// getConfigFile Gets the config file for the current environment.
func getConfigFile() (string, error) {
	result := pflag.StringP(
		"environment",
		"e",
		"development",
		`Sets the environment for the application.
This determines what's the config file that is going to be read.
Valid environments: "development", "staging", "production".`,
	)
	pflag.Parse()

	if result == nil || *result == "" {
		return "", fmt.Errorf("environment not found, check the command usage")
	}
	if config, exists := environmentConfigFile[*result]; exists {
		return config, nil
	}

	return "", fmt.Errorf("environment \"" + *result + "\" is invalid, check the command usage")
}
