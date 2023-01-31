package httpserver

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/v1adhope/crypto-diary/internal/config"
	"github.com/v1adhope/crypto-diary/pkg/logger"
)

// TODO: Decomposition
// TODO: Separate configure
// TODO: Logger
func New(handler http.Handler, cfg *config.Server, logger *logger.Logger) {
	srv := &http.Server{
		Addr:         cfg.Address,
		Handler:      handler,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal().Err(err).Msg("listen and serve")
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info().Msg("shutdown server ...")

	ctx, cancel := context.WithTimeout(context.Background(), cfg.ShutdownTimeout*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatal().Err(err).Msg("server shutdown")
	}
	select {
	case <-ctx.Done():
		logger.Info().Msg(fmt.Sprintf("timeout of %d seconds", cfg.ShutdownTimeout))
	}
	logger.Info().Msg("server exiting")
}
