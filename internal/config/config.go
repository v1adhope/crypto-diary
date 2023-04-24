// NOTE: SINGELTON
// TODO: Replace to env
package config

import (
	"log"
	"sync"

	"github.com/spf13/viper"
	"github.com/v1adhope/crypto-diary/internal/pkg/auth"
	"github.com/v1adhope/crypto-diary/pkg/httpserver"
	"github.com/v1adhope/crypto-diary/pkg/logger"
	"github.com/v1adhope/crypto-diary/pkg/postgres"
	"github.com/v1adhope/crypto-diary/pkg/rds"
)

type Config struct {
	GinMode        string             `mapstructure:"gin_mode"`
	Logger         *logger.Config     `mapstructure:"logger"`
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

	defaults = map[string]any{
		"gin_mode": "release",

		"logger.log_level":             "info",
		"logger.file_name":             "logs/lumberjack.log",
		"logger.is_console_log_enable": "true",
		"logger.is_file_log_enable":    "true",
		"logger.is_compress":           "false",

		"server.socket":           "0.0.0.0:8081",
		"server.shutdown_timeout": "15s",
		"server.read_timeout":     "15s",
		"server.write_timeout":    "15s",

		"auth.issuer":               "crypto-diary.space",
		"auth.refresh_token_ttl":    "168h",
		"auth.refresh_token_secret": "7PFUSSxFjeVkbxqAn5ktapVE2PvbQgTgzsrztktawFbYK2TubTkp4JvgjpHP3Taa",
		"auth.access_token_ttl":     "15m",
		"auth.access_token_secret":  "KycVkWHPDygdJFgerbHjr7v7erWF2NA7uxEkgTaeQLgqS2939XjmcVJUQfKEApkS",

		"session_storage.socket":   "session:6379",
		"session_storage.password": "",
		"session_storage.database": "0",

		"password_secret": "pWu2EnUd3yqyP9mfK7uaqWdNFjWTeKEvVxFCLcxE2hhJuwY9jKRb2ZwHztjC3LmL",

		"storage.username":      "designer",
		"storage.password":      "designer",
		"storage.socket":        "storage:5432",
		"storage.database":      "crypto_diary",
		"storage.conn_attempts": "15",
		"storage.conn_timeout":  "3s",
		"storage.pool_size":     "100",
	}
)

func GetConfig() *Config {
	once.Do(
		func() {
			for k, v := range defaults {
				viper.SetDefault(k, v)
			}

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
