package web

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateRequest(ctx *gin.Context) {
	service := ctx.Param("service")
	serviceURL, exists := h.Services.Service[service]
	if !exists {
		sendError(ctx, NewResult("Not Found Service", http.StatusNotFound))
		return
	}
	path := strings.TrimPrefix(ctx.Param("path"), "/")
	link, err := url.JoinPath(serviceURL, path)
	if err != nil {
		sendError(ctx, NewResult("Not Found", http.StatusNotFound))
		return
	}
	response, err := h.proxyRequest(ctx, link)
	if err != nil {
		sendError(ctx, NewResult("Not Found", http.StatusNotFound))
		return
	}
	defer response.Body.Close()
	h.proxyResponse(ctx, response)
}

func (h *Handler) NoRouting(ctx *gin.Context) {
	sendError(ctx, NewResult("Not Found", http.StatusNotFound))
}
