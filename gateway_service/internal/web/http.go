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

	validGroup := r.Group("/", h.ValidateToken())
	{
		validGroup.GET("/main/", h.MainHandler)
		validGroup.GET("/auth/test/b", h.Test_bad)
		// validGroup.GET("/user_list", h.GetUserList)

		adminGroup := validGroup.Group("/admin", h.RoleAccessor("admin"))
		{
			adminGroup.GET("/user_list", h.GetUserList)
			adminGroup.GET("/test/s", h.Test)
		}

		userGroup := validGroup.Group("/user", h.RoleAccessor("user"))
		{
			userGroup.GET("/test/s", h.Test)
		}
	}
}
