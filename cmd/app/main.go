package main

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/v1adhope/crypto-diary/internal/config"
	v1 "github.com/v1adhope/crypto-diary/internal/controller/http/v1"
	"github.com/v1adhope/crypto-diary/internal/entity"
	"github.com/v1adhope/crypto-diary/internal/usecase/repository"
	"github.com/v1adhope/crypto-diary/pkg/httpserver"
	"github.com/v1adhope/crypto-diary/pkg/logger"
	"github.com/v1adhope/crypto-diary/pkg/postgres"
)

// TODO: add logger
func main() {
	cfg := config.GetConfig()
	logger := logger.New(cfg.DebugLVL)

	pgClient, err := postgres.NewClient(cfg)
	if err != nil {
		logger.Fatal().Err(err).Send()
	}
	defer pgClient.Close()

	//TODO: pg test
	positionRepo := repository.NewPosition(pgClient)

	// NOTE: Position
	p := &entity.Position{
		OpenDate:        "2023.01.17",
		Pair:            "btc/usdt",
		Risk:            "1",
		Reason:          "Some reason",
		AccordingToPlan: "true",
		Direction:       "short",
		Deposit:         "100",
		OpenPrice:       "20000",
		StopLossPrice:   "19000",
		TakeProfitPrice: "23000",
		ClosePrice:      "23000",
		UserID:          "1",
	}
	p.ValidPosition()
	if err != nil {
		logger.Fatal().Err(err).Send()
	}
	err = positionRepo.Create(context.Background(), p)
	if err != nil {
		logger.Fatal().Err(err).Send()
	}
	fmt.Println(p)

	id := "3"
	err = positionRepo.Delete(context.Background(), &id)
	if err != nil {
		logger.Fatal().Err(err).Send()
	}

	positions := make([]entity.Position, 0)
	positions, err = positionRepo.FindAll(context.Background())
	if err != nil {
		logger.Fatal().Err(err).Send()
	}
	fmt.Println(positions)

	//NOTE: User
	userRepo := repository.NewUser(pgClient)

	u := &entity.User{
		Email:    "custom@custom.cu",
		Password: "password",
	}
	err = userRepo.CreateUser(context.Background(), u)
	if err != nil {
		logger.Fatal().Err(err).Send()

	}

	u2 := &entity.User{}
	email := "google@gmail.com"
	passwd := "password1"
	u2, err = userRepo.GetUser(context.Background(), &email, &passwd)
	if err != nil {
		logger.Fatal().Err(err).Send()
	}
	fmt.Println(u2)

	// user := &entity.User{}
	// user, err = repo.FindOne(context.Background(), "google@gmail.com")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(user)

	// TODO
	handler := gin.Default()
	v1.NewRouter(handler, logger)

	httpserver.New(handler, cfg)
}
