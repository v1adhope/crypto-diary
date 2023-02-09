package rds

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type Config struct {
	Socket   string `mapstructure:"socket"`
	Password string `mapstructure:"password"`
	Database int    `mapstructure:"database"`
}

type Redis struct {
	Client *redis.Client
}

func NewClient(ctx context.Context, cfg *Config) (*Redis, error) {
	rdb := &Redis{}

	rdb.Client = redis.NewClient(&redis.Options{
		Addr:     cfg.Socket,
		Password: cfg.Socket,
		DB:       cfg.Database,
	})

	_, err := rdb.Client.Ping(ctx).Result()
	if err != nil {
		return nil, fmt.Errorf("redis: NewClient: unable to create client %w", err)
	}

	return rdb, nil
}

func (r *Redis) Close() {
	if r.Client != nil {
		r.Client.Close()
	}
}
