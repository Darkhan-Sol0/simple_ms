package web

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
	Err     interface{} `json:"error"`
	Details string      `json:"details"`
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

func sendMessage(ctx *gin.Context, response Response) {
	if response.Err != nil {
		sendError(ctx, NewResult(response.Details, response.Status))
	} else {
		sendSucces(ctx, NewResult(response.Data, response.Status))
	}
}
