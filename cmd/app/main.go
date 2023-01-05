package main

import (
	"context"
	"log"

	"github.com/v1adhope/crypto-diary/internal/config"
	"github.com/v1adhope/crypto-diary/pkg/postgres"
)

// TODO: add logger
func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatalln(err)
	}

	postgresClient, err := postgres.NewClient(context.TODO(), cfg.Storage)
	if err != nil {
		log.Fatalln(err)
	}

	//TODO: use me
	_ = postgresClient
}
