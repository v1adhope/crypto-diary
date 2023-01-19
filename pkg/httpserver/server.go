package httpserver

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/v1adhope/crypto-diary/internal/config"
)

// TODO: Decomposition
// TODO: Separate configure
// TODO: Logger
func New(handler http.Handler, cfg *config.Config) {
	srv := &http.Server{
		Addr:         cfg.Server.Address,
		Handler:      handler,
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
