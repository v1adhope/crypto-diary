// NOTE: SINGELTON
package config

import (
	"log"
	"sync"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	LogLevel       string   `mapstructure:"log_level"`
	Server         *Server  `mapstructure:"server"`
	Storage        *Storage `mapstructure:"storage"`
	PasswordSecret string   `mapstructure:"password_secret"`
}

type Server struct {
	Address         string        `mapstructure:"socket"`
	ShutdownTimeout time.Duration `mapstructure:"shutdown_timeout"`
}

type Storage struct {
	Username     string        `mapstructure:"username"`
	Password     string        `mapstructure:"password"`
	Host         string        `mapstructure:"host"`
	Port         string        `mapstructure:"port"`
	Database     string        `mapstructure:"database"`
	ConnAttempts int           `mapstructure:"conn_attempts"`
	ConnTimeout  time.Duration `mapstructure:"conn_timeout"`
	PoolSize     int32         `mapstructure:"pool_size"`
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
