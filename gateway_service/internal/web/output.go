package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Data    interface{}
	Err     interface{}
	Status  int
	Details string
}

type Result struct {
	status int
	data   interface{}
}

func NewResult(data interface{}, status int) Result {
	return Result{
		data:   data,
		status: status,
	}
}

func sendError(ctx *gin.Context, res Result) {
	ctx.JSON(res.status, gin.H{
		"status": res.status,
		"error":  res.data,
	})
}

func sendSucces(ctx *gin.Context, res Result) {
	ctx.JSON(res.status, gin.H{
		"status": res.status,
		"data":   res.data,
	})
}

func (h *Handler) sendMessage(ctx *gin.Context, response *http.Response) {
	resp, err := h.GetResponse(response)
	if err != nil {
		sendError(ctx, NewResult(err.Error(), resp.Status))
		return
	}
	sendSucces(ctx, NewResult(resp.Data, resp.Status))
}
