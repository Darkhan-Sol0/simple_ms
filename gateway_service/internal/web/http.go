package web

import (
	"gateway/infrastructure/config"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Services  config.Services
	semaphore chan struct{}
}

func NewHandler() *Handler {
	return &Handler{
		Services: *config.GetCongigs(),
	}
}

func (h *Handler) MakeSemophore() {
	h.semaphore = make(chan struct{}, h.Services.SemophoreCount)
}

func (h *Handler) RegistrateHandler(r *gin.Engine) {
	r.GET("/", h.index)
	r.Any("/:service/*path", h.Timeout(), h.CheckService(), h.DecodeToken(), h.CreateRequest)
	// r.NoRoute(h.NoRouting)
}
