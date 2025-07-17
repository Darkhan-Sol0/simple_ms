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
	r.POST("/auth/sign_up", h.Registration)
	r.POST("/auth/sign_in", h.Authorization)

	validGroup := r.Group("/", h.ValidateToken())
	{
		validGroup.GET("/user", h.GetSelfUser)
		validGroup.PATCH("/user", h.UserUpdateInfo)

		adminGroup := validGroup.Group("/admin", h.RoleAccessor("admin"))
		{
			adminGroup.GET("/user_list", h.GetUserList)
		}
	}
}
