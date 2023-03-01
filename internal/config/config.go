// NOTE: SINGELTON
// TODO: Replace to env
package config

import (
	"log"
	"sync"

	"github.com/spf13/viper"
	"github.com/v1adhope/crypto-diary/pkg/auth"
	"github.com/v1adhope/crypto-diary/pkg/httpserver"
	"github.com/v1adhope/crypto-diary/pkg/postgres"
	"github.com/v1adhope/crypto-diary/pkg/rds"
)

type Config struct {
	LogLevel       string             `mapstructure:"log_level"`
	GinMode        string             `mapstructure:"gin_mode"`
	Server         *httpserver.Config `mapstructure:"server"`
	Storage        *postgres.Config   `mapstructure:"storage"`
	PasswordSecret string             `mapstructure:"password_secret"`
	Auth           *auth.Config       `mapstructure:"auth"`
	SessionStorage *rds.Config        `mapstructure:"session_storage"`
}

var (
	cfgName = "config"
	cfgType = "yaml"
	cfgPath = "./configs"
	once    sync.Once
	cfg     *Config
)

func GetConfig() *Config {
	once.Do(
		func() {
			viper.SetConfigName(cfgName)
			viper.SetConfigType(cfgType)
			viper.AddConfigPath(cfgPath)

			if err := viper.ReadInConfig(); err != nil {
				log.Fatalf("could not read config file: %v", err)
			}

			if err := viper.Unmarshal(&cfg); err != nil {
				log.Fatalf("could not decode config file into struct: %v", err)
			}
		})

	return cfg
}
