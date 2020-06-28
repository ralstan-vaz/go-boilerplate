package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

var configModel *Config

// NewConfig gets the configuration based on the environment passed
func NewConfig(env string) (*Config, error) {
	configFile := "config/tier/" + env + ".yaml"
	bytes, err := ioutil.ReadFile(configFile)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(bytes, &configModel)
	if err != nil {
		return nil, err
	}
	return configModel, nil
}

// Get Retrives the config model without loading it from the disk
func Get() *Config {
	return configModel
}
