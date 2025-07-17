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
	r.POST("/sign_up", h.Registaration)
	r.POST("/sign_in", h.Authorization)

	r.POST("/check_auth", h.CheckAuthorization)

	r.GET("/user_list", h.GetUsersList)
}
