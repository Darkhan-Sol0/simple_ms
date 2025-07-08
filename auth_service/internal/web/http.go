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
	r.GET("/", h.Main)

	r.GET("/err", h.Erro)
	r.GET("/suc", h.Succes)

	r.POST("/sign_in", h.Registaration)
	r.POST("/sign_up", h.Authorization)
}
