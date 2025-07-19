package web

import (
	"user_service/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service service.UserService
}

func NewHandler(service service.UserService) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) RegistrationHandlers(r *gin.Engine) {
	r.POST("/", h.RoleChecker("user"), h.CreateUser)
	r.GET("/", h.RoleChecker("user"), h.GetSelfUser)
	r.PATCH("/", h.RoleChecker("user"), h.UpdateUser)

	r.GET("/:uuid", h.RoleChecker("user"), h.GetUser)

}
