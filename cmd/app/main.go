package main

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/v1adhope/crypto-diary/internal/config"
	v1 "github.com/v1adhope/crypto-diary/internal/controller/http/v1"
	"github.com/v1adhope/crypto-diary/internal/usecase"
	"github.com/v1adhope/crypto-diary/internal/usecase/repository"
	"github.com/v1adhope/crypto-diary/internal/usecase/session"
	"github.com/v1adhope/crypto-diary/pkg/auth"
	"github.com/v1adhope/crypto-diary/pkg/hash"
	"github.com/v1adhope/crypto-diary/pkg/httpserver"
	"github.com/v1adhope/crypto-diary/pkg/logger"
	"github.com/v1adhope/crypto-diary/pkg/postgres"
	"github.com/v1adhope/crypto-diary/pkg/rds"
)

func main() {
	cfg := config.GetConfig()
	logger := logger.New(cfg.LogLevel)

	pgClient, err := postgres.NewClient(cfg.Storage)
	if err != nil {
		logger.Fatal(err, "main: pgClient")
	}
	defer pgClient.Close()

	repos := repository.New(pgClient)

	hasher := hash.New(cfg.PasswordSecret)

	validate := validator.New()

	auth := auth.New(cfg.Auth)

	redisClient, err := rds.NewClient(context.Background(), cfg.SessionStorage)
	if err != nil {
		logger.Fatal(err, "main: redisClient")
	}
	defer redisClient.Close()

	sessionStorage := session.New(redisClient)

	useCases := usecase.New(usecase.Deps{
		Repos:   repos,
		Hasher:  hasher,
		Auth:    auth,
		Session: sessionStorage,
	})

	if cfg.GinMode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	handler := gin.New()

	v1.NewRouter(&v1.Router{
		Handler:  handler,
		UseCases: useCases,
		Logger:   logger,
		Validate: validate,
	})

	srv := httpserver.New(handler, cfg.Server)

	srv.Run()
}
