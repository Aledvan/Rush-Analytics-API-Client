package config

import (
	"os"
	"ra-api-client/errors"

	"gopkg.in/yaml.v2"
)

type Config struct {
	API struct {
		BaseURL string `yaml:"base_url"`
		APIKey  string `yaml:"api_key"`
	} `yaml:"api"`

	HTTP struct {
		Timeout int `yaml:"timeout"`
		Retries int `yaml:"retries"`
	} `yaml:"http"`

	Logging struct {
		Level string `yaml:"level"`
	} `yaml:"logging"`
}

func Load() (*Config, error) {
	configPaths := []string{
		"config/config.yaml",
		"config.yaml",
	}

	var data []byte
	var err error

	for _, path := range configPaths {
		if _, statErr := os.Stat(path); statErr == nil {
			data, err = os.ReadFile(path)
			if err == nil {
				break
			}
		}
	}

	if data == nil {
		return nil, errors.NewConfigError("Config file not found in config/, config.yaml")
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, errors.NewConfigError(err.Error())
	}

	return &config, nil
}

func GetConfig() (*Config, error) {
	cfg, err := Load()
	if err != nil {
		return nil, errors.NewConfigError("Failed to load config file. Please check settings")
	}

	return cfg, nil
}
