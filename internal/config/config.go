// NOTE: SINGELTON
package config

import (
	"log"
	"sync"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Address         string        `yaml:"address"`
		ShutdownTimeout time.Duration `yaml:"shutdownTimeout"`
	} `yaml:"server"`

	Storage struct {
		Username     string        `yaml:"username"`
		Password     string        `yaml:"password"`
		Host         string        `yaml:"host"`
		Port         string        `yaml:"port"`
		Database     string        `yaml:"database"`
		ConnAttempts int           `yaml:"connAttempts"`
		ConnTimeout  time.Duration `yaml:"connTimeout"`
		PoolSize     int32         `yaml:"poolSize"`
	} `yaml:"storage"`
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
