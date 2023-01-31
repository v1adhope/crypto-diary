package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/v1adhope/crypto-diary/internal/config"
	v1 "github.com/v1adhope/crypto-diary/internal/controller/http/v1"
	"github.com/v1adhope/crypto-diary/internal/usecase"
	"github.com/v1adhope/crypto-diary/internal/usecase/repository"
	"github.com/v1adhope/crypto-diary/pkg/hash"
	"github.com/v1adhope/crypto-diary/pkg/httpserver"
	"github.com/v1adhope/crypto-diary/pkg/logger"
	"github.com/v1adhope/crypto-diary/pkg/postgres"
)

func main() {
	cfg := config.GetConfig()
	logger := logger.New(cfg.LogLevel)

	pgClient, err := postgres.NewClient(cfg.Storage)
	if err != nil {
		logger.Fatal().Err(err).Send()
	}
	defer pgClient.Close()

	repos := repository.New(pgClient)

	useCases := usecase.New(repos)

	hasher := hash.New(cfg.PasswordSecret)

	validate := validator.New()

	// TODO
	handler := gin.New()
	v1.NewRouter(&v1.Deps{
		Handler:  handler,
		UseCases: useCases,
		Logger:   logger,
		Hasher:   hasher,
		Validate: validate,
	})

	httpserver.New(handler, cfg.Server, logger)
}
