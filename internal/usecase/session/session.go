package session

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/v1adhope/crypto-diary/pkg/rds"
)

type Session struct {
	*rds.Redis
}

func New(client *rds.Redis) *Session {
	return &Session{client}
}

func (s *Session) AddToBlackList(ctx context.Context, token string, exp time.Duration) error {
	exp *= time.Hour

	err := s.Client.Set(ctx, token, token, exp).Err()
	if err != nil {
		return fmt.Errorf("couldn't add to the blacklist: %w", err)
	}

	return nil
}

func (s *Session) CheckToken(ctx context.Context, token string) error {
	value, err := s.Client.Get(ctx, token).Result()
	if err == redis.Nil {
		return nil
	}

	if err != nil {
		return fmt.Errorf("session: CheckToken: Get: %w", err)
	}

	return fmt.Errorf("token in the blocklist: %s", value)
}
