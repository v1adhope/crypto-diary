package config

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"github.com/v1adhope/crypto-diary/pkg/postgres"
)

type Config struct {
	Storage postgres.StorageConfig
}

var (
	configName = "config"
	configType = "yaml"
	configPath = "./configs"
	c          Config
)

// TODO: add logger
func GetConfig() (*Config, error) {
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)
	viper.AddConfigPath(configPath)

	if err := viper.ReadInConfig(); err != nil {
		return nil, errors.Wrap(err, "could not read config file")
	}

	if err := viper.Unmarshal(&c); err != nil {
		return nil, errors.Wrap(err, "could not decode config file into struct")
	}

	return &c, nil
}
