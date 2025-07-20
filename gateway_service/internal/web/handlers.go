package web

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateRequest(ctx *gin.Context) {
	link, _ := ctx.Get("link")
	response, err := h.proxyRequest(ctx, link.(string))
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
