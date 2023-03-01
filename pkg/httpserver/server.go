package httpserver

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Config struct {
	Socket          string        `mapstructure:"socket"`
	ShutdownTimeout time.Duration `mapstructure:"shutdown_timeout"`
	ReadTimeout     time.Duration `mapstructure:"read_timeout"`
	WriteTimeout    time.Duration `mapstructure:"write_timeout"`
}

type Server struct {
	server          *http.Server
	shutdownTimeout time.Duration
}

func New(h http.Handler, cfg *Config) *Server {
	httpServer := &http.Server{
		Addr:         cfg.Socket,
		Handler:      h,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
	}

	return &Server{
		server:          httpServer,
		shutdownTimeout: cfg.ShutdownTimeout,
	}
}

func (s *Server) Run() {
	go func() {
		if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen and serve: %s", err)
		}
	}()

	s.gracefulShutdown()
}

func (s *Server) gracefulShutdown() {
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Printf("shutdown server ...")

	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	if err := s.server.Shutdown(ctx); err != nil {
		log.Fatalf("server shutdown: %s", err)
	}

	select {
	case <-ctx.Done():
		log.Printf("timeout of %d seconds", s.shutdownTimeout)
	}

	log.Printf("server exiting")
}
