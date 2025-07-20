package web

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Data   any `json:"data"`
	Err    any `json:"error"`
	Status int
}

type Result struct {
	status int
	err    error
	data   any
}

func NewResult(data any, status int, err error) Result {
	return Result{
		data:   data,
		err:    err,
		status: status,
	}
}

func sendError(ctx *gin.Context, res Result) {
	ctx.JSON(res.status, gin.H{
		"error": res.err.Error(),
	})
}

func sendSucces(ctx *gin.Context, res Result) {
	ctx.JSON(res.status, gin.H{
		"data": res.data,
	})
}

func (h *Handler) sendMessage(ctx *gin.Context, res Result) {
	if res.err != nil {
		sendError(ctx, res)
		return
	}
	sendSucces(ctx, res)
}
