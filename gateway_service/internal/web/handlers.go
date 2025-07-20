package web

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateRequest(ctx *gin.Context) {
	service := ctx.Param("service")
	serviceURL, exists := h.Services.Service[service]
	if !exists {
		h.sendMessage(ctx, NewResult(nil, http.StatusNotFound, fmt.Errorf("error: Not Found Service")))
		return
	}
	path := strings.TrimPrefix(ctx.Param("path"), "/")
	link, err := url.JoinPath(serviceURL, path)
	if err != nil {
		h.sendMessage(ctx, NewResult(nil, http.StatusNotFound, fmt.Errorf("error: Not Found Service")))
		return
	}
	response, err := h.proxyRequest(ctx, link)
	if err != nil {
		h.sendMessage(ctx, NewResult(nil, http.StatusNotFound, fmt.Errorf("error: Not Found Service")))
		return
	}
	res, err := h.GetResponse(response)
	if err != nil {
		h.sendMessage(ctx, NewResult(nil, res.Status, err))
		return
	}
	h.proxyResponse(ctx, res)
}

func (h *Handler) NoRouting(ctx *gin.Context) {
	h.sendMessage(ctx, NewResult(nil, http.StatusNotFound, fmt.Errorf("error: Not Found Service")))
}
