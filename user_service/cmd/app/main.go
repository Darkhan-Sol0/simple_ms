package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"myApp/internal/config"
	"myApp/internal/datasource/database"
	"myApp/internal/datasource/repository"
	"myApp/internal/service"
	"myApp/internal/web"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	cfg := config.LoadConfig()
	s := web.NewServer(cfg)

	db, err := database.ConnectDB(context.Background(), cfg)
	if err != nil {
		log.Fatalf("Error connect database: %v", err)
	}

	r := web.NewRouting(service.NewService(repository.NewDatabase(db)))
	r.RegisterRoutes(e)

	go func() {
		s.Start(e)
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit
	s.Shutdown(e)
}
