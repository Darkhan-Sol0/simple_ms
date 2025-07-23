package web

import (
	"chat_service/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service service.ChatService
}

func NewHandler(service service.ChatService) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) RegistrationHandlers(r *gin.Engine) {

}
