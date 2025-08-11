package web

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type (
	ConfigServer interface {
		GetAddress() string
		GetSessionTimeout() time.Duration
		GetIdleTimeout() time.Duration
	}

	ServerHTTP struct {
		address        string
		sessionTimeout time.Duration
		idleTimeout    time.Duration
	}

	Server interface {
		Start(e *echo.Echo)
		Shutdown(e *echo.Echo)
	}
)

func NewServer(cfg ConfigServer) Server {
	return &ServerHTTP{
		address:        cfg.GetAddress(),
		sessionTimeout: cfg.GetSessionTimeout(),
		idleTimeout:    cfg.GetIdleTimeout(),
	}
}

func (s *ServerHTTP) Start(e *echo.Echo) {
	log.Printf("Server starting: %s...", s.address)
	if err := e.Start(s.address); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server failed to start: %v", err)
	}
}

func (s *ServerHTTP) Shutdown(e *echo.Echo) {
	ctx, cancel := context.WithTimeout(context.Background(), s.idleTimeout)
	defer cancel()
	log.Printf("Server shutting down...")
	if err := e.Shutdown(ctx); err != nil {
		log.Printf("Error during server shutdown: %v", err)
	}
	log.Println("Server by ended")
}
