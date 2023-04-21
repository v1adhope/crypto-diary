package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/v1adhope/crypto-diary/internal/entity"
)

type Config struct {
	SecretKey            string        `mapstructure:"secret_key"`
	Issuer               string        `mapstructure:"issuer"`
	RefreshTokenLifetime time.Duration `mapstructure:"refresh_token_lifetime"`
	RefreshTokenSecret   string        `mapstructure:"refresh_token_secret"`
	AccessTokenLifetime  time.Duration `mapstructure:"access_token_lifetime"`
	AccessTokenSecret    string        `mapstructure:"access_token_secret"`
}

const (
	_kidRefreshToken = "refresh"
	_kidAccessToken  = "access"
)

type AuthManager interface {
	GenerateTokenPair(id string) (refreshToken string, accessToken string, err error)
	ValidateAccessToken(clientToken string) (string, error)
	ValidateRefreshToken(clientToken string) (string, time.Duration, error)
}

type Manager struct {
	issuer               string
	refreshTokenLifetime time.Duration
	refreshTokenSecret   string
	accessTokenLifetime  time.Duration
	accessTokenSecret    string
}

func New(cfg *Config) *Manager {
	return &Manager{
		issuer:               cfg.Issuer,
		refreshTokenLifetime: cfg.RefreshTokenLifetime,
		refreshTokenSecret:   cfg.RefreshTokenSecret,
		accessTokenLifetime:  cfg.AccessTokenLifetime,
		accessTokenSecret:    cfg.AccessTokenSecret,
	}
}

func (m *Manager) GenerateTokenPair(id string) (string, string, error) {
	refreshToken, err := m.generateRefreshToken(id)
	if err != nil {
		return "", "", err
	}

	accessToken, err := m.generateAccessToken(id)
	if err != nil {
		return "", "", err
	}

	return refreshToken, accessToken, nil
}

func (m *Manager) generateRefreshToken(id string) (string, error) {
	claims := &jwt.RegisteredClaims{
		Subject:   id,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(m.refreshTokenLifetime)),
		Issuer:    m.issuer,
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token.Header["kid"] = _kidRefreshToken

	signedToken, err := token.SignedString([]byte(m.refreshTokenSecret))
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	return signedToken, nil
}

func (m *Manager) generateAccessToken(id string) (string, error) {
	claims := &jwt.RegisteredClaims{
		Subject:   id,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(m.accessTokenLifetime)),
		Issuer:    m.issuer,
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token.Header["kid"] = _kidAccessToken

	signedToken, err := token.SignedString([]byte(m.accessTokenSecret))
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %w", err)
	}

	return signedToken, nil
}

func (m *Manager) ValidateRefreshToken(clientToken string) (string, time.Duration, error) {
	token, err := jwt.Parse(clientToken, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("%w: kid %v: unexpected signing method %v", entity.ErrTokenInvalid, token.Header["kid"], token.Header["alg"])
		}

		return []byte(m.refreshTokenSecret), nil
	})
	if err != nil {
		return "", 0, fmt.Errorf("%w: parse failed: %s", entity.ErrTokenInvalid, err)
	}

	if !token.Valid || token.Header["kid"] != _kidRefreshToken {
		return "", 0, entity.ErrTokenInvalid
	}

	id, err := m.extractClaimField(token, "sub")
	if err != nil {
		return "", 0, fmt.Errorf("%w: %s", entity.ErrTokenInvalid, err)
	}

	return id, m.refreshTokenLifetime, nil
}

func (m *Manager) ValidateAccessToken(clientToken string) (string, error) {
	token, err := jwt.Parse(clientToken, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("kid %v: unexpected signing method %v", token.Header["kid"], token.Header["alg"])
		}

		return []byte(m.accessTokenSecret), nil
	})
	if err != nil {
		return "", fmt.Errorf("parse failed: %w", err)
	}

	if !token.Valid || token.Header["kid"] != _kidAccessToken {
		return "", errors.New("invalid token")
	}

	id, err := m.extractClaimField(token, "sub")
	if err != nil {
		return "", err
	}

	return id, nil
}

func (m *Manager) extractClaimField(token *jwt.Token, key string) (string, error) {
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		switch claim := claims[key].(type) {
		default:
			return fmt.Sprintf("%v", claim), nil
		case nil:
			return "", errors.New("claim is empty")
		}
	}

	return "", errors.New("claims are empty")
}