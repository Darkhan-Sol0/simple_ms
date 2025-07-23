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
	internal := r.Group("/internal").Use(h.Internal())
	{
		internal.POST("/create", h.CreateUser)
	}
	user := r.Group("/").Use(h.RoleChecker("user", "admin"))
	{
		user.GET("/", h.GetSelfUser)
		user.PATCH("/", h.UpdateUser)
		user.GET("/list", h.GetUsersList)
		user.GET("/:uuid", h.GetUser)
	}
	r.NoRoute(h.NotFound)
}
