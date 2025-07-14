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
	r.POST("/auth/sign_up", h.Registration)
	r.POST("/auth/sign_in", h.Authorization)

	useGroup := r.Group("/", h.ValidateToken())
	{
		useGroup.GET("/main/", h.MainHandler)
		useGroup.GET("/auth/test/s", h.RoleAccessor("admin"), h.Test)
		useGroup.GET("/auth/test/ss", h.RoleAccessor("user"), h.Test)
		useGroup.GET("/auth/test/b", h.Test_bad)
	}
}
