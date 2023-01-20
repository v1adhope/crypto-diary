package main

import (
	"github.com/gin-gonic/gin"
	"github.com/v1adhope/crypto-diary/internal/config"
	v1 "github.com/v1adhope/crypto-diary/internal/controller/http/v1"
	"github.com/v1adhope/crypto-diary/internal/usecase/repository"
	"github.com/v1adhope/crypto-diary/pkg/httpserver"
	"github.com/v1adhope/crypto-diary/pkg/logger"
	"github.com/v1adhope/crypto-diary/pkg/postgres"
)

func main() {
	cfg := config.GetConfig()
	logger := logger.New(cfg.LogLevel)

	pgClient, err := postgres.NewClient(cfg)
	if err != nil {
		logger.Fatal().Err(err).Send()
	}
	defer pgClient.Close()

	positionRepo := repository.NewPosition(pgClient)

	//TODO: Use me
	_ = positionRepo

	// TODO
	handler := gin.Default()
	v1.NewRouter(handler, logger)

	httpserver.New(handler, cfg, logger)
}
