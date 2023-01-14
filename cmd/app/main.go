package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/v1adhope/crypto-diary/internal/config"
	"github.com/v1adhope/crypto-diary/internal/entity"
	"github.com/v1adhope/crypto-diary/internal/usecase/repository"
	"github.com/v1adhope/crypto-diary/pkg/postgres"
)

// TODO: add logger
func main() {
	cfg := config.GetConfig()

	pgClient, err := postgres.NewClient(cfg)
	if err != nil {
		log.Fatalln(err)
	}
	defer pgClient.Close()

	//NOTE: pg test
	repo := repository.New(pgClient)

	a := entity.User{
		Email:    "nnn",
		Password: "aldfjad",
	}
	err = repo.Create(context.Background(), &a)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(a)

	users := make([]entity.User, 0)
	users, err = repo.FindAll(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(users)

	user := &entity.User{}
	user, err = repo.FindOne(context.Background(), "google@gmail.com")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(user)

	//TODO: handlers, route replace
	router := gin.Default()

	router.GET("/")

	srv := &http.Server{
		Addr:         cfg.Server.Address,
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("listen:", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("shutdown server ...")

	ctx, cancel := context.WithTimeout(context.Background(), cfg.Server.ShutdownTimeout*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("server shutdown:", err)
	}
	select {
	case <-ctx.Done():
		log.Printf("timeout of %d seconds", cfg.Server.ShutdownTimeout)
	}
	log.Println("server exiting")
}
