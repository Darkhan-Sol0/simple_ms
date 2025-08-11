package web

import (
	"github.com/labstack/echo/v4"
)

type (
	Service interface {
		Hello() string
		GetList() map[string]string
	}

	routingConfig struct {
		service Service
	}

	Routing interface {
		RegisterRoutes(e *echo.Echo)
	}
)

func NewRouting(service Service) Routing {
	return &routingConfig{
		service: service,
	}
}

func (r *routingConfig) RegisterRoutes(e *echo.Echo) {
	e.GET("/", r.indexHandler)
	e.GET("/a", r.listHandler)
}
