package web

import (
	"auth_service/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service service.AuthService
}

func NewHandler(service service.AuthService) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) RegistrationHandlers(r *gin.Engine) {
	r.POST("/sign_up", h.Registration)
	r.POST("/sign_in", h.Authorization)

	internal := r.Group("/internal").Use(h.Internal())
	{
		internal.POST("/check_auth", h.CheckAuthorization)
	}

	admin := r.Group("/admin").Use(h.RoleChecker("admin"))
	{
		admin.GET("/user_list", h.GetUsersList)
	}

	r.NoRoute(h.NotFound)
}
