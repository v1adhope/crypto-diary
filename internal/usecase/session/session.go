package session

import (
	"context"
	"fmt"
	"time"

	"github.com/v1adhope/crypto-diary/pkg/rds"
)

type Session struct {
	*rds.Redis
}

func New(client *rds.Redis) *Session {
	return &Session{client}
}

func (s *Session) AddToBlackList(ctx context.Context, token string, ttl time.Duration) error {
	ttl *= time.Hour

	err := s.Client.Set(ctx, token, "", ttl).Err()
	if err != nil {
		return fmt.Errorf("couldn't add to the blacklist: %w", err)
	}

	return nil
}

func (s *Session) CheckToken(ctx context.Context, token string) error {
	isExists, err := s.Client.Exists(ctx, token).Result()
	if err != nil {
		return fmt.Errorf("session: CheckToken: Exists: %w", err)
	}

	if isExists != 0 {
		return fmt.Errorf("token in the blocklist: %s", token)
	}

	return nil
}
