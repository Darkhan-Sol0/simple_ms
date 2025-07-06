package web

import (
	"gateway/internal/service"

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

}
