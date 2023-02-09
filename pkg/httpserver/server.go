package httpserver

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/v1adhope/crypto-diary/pkg/logger"
)

type Config struct {
	Socket          string        `mapstructure:"socket"`
	ShutdownTimeout time.Duration `mapstructure:"shutdown_timeout"`
	ReadTimeout     time.Duration `mapstructure:"read_timeout"`
	WriteTimeout    time.Duration `mapstructure:"write_timeout"`
}

// TODO: Decomposition
func New(h http.Handler, cfg *Config, logger *logger.Log) {
	srv := &http.Server{
		Addr:         cfg.Socket,
		Handler:      h,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal(err, "listen and serve")
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("shutdown server ...")

	ctx, cancel := context.WithTimeout(context.Background(), cfg.ShutdownTimeout*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatal(err, "server shutdown")
	}
	select {
	case <-ctx.Done():
		logger.Info("timeout of %d seconds", cfg.ShutdownTimeout)
	}
	logger.Info("server exiting")
}
