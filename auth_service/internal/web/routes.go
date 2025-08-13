package web

import (
	"auth_service/internal/service"

	"github.com/labstack/echo/v4"
)

type (
	routingConfig struct {
		service service.Service
	}

	Routing interface {
		RegisterRoutes(e *echo.Echo)
	}
)

func NewRouting(service service.Service) Routing {
	return &routingConfig{
		service: service,
	}
}

func (r *routingConfig) RegisterRoutes(e *echo.Echo) {
	e.POST("/reg", r.PostNewUser)
	e.GET("/user_list", r.GetUsersList)
	e.GET("/user", r.GetUserByUuid)
}
