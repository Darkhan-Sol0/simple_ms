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
	err    interface{}
}

func NewResult(data interface{}, status int, err interface{}) Result {
	return Result{
		data:   data,
		err:    err,
		status: status,
	}
}

func sendMessage(ctx *gin.Context, res Result) {
	if res.err != nil {
		ctx.JSON(res.status, gin.H{
			"status": res.status,
			"error":  res.data,
		})
	} else {
		ctx.JSON(res.status, gin.H{
			"status": res.status,
			"data":   res.data,
		})
	}
}
