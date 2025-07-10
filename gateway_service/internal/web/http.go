package web

import (
	"gateway/infrastructure/config"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Services config.Services
}

func NewHandler() *Handler {
	return &Handler{
		Services: *config.GetService(),
	}
}

func (h *Handler) RegistrateHandler(r *gin.Engine) {
	r.GET("/", h.MainHandler)
	r.GET("/auth/test/s", h.Test)
	r.GET("/auth/test/b", h.Test_bad)
	r.POST("/auth/sign_up", h.Registration)
	r.POST("/auth/sign_in", h.Authorization)
}
