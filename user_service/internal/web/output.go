package web

import "github.com/gin-gonic/gin"

type Result struct {
	data   interface{}
	err    error
	status int
}

func NewResult(data any, status int, err error) Result {
	return Result{
		data:   data,
		err:    err,
		status: status,
	}
}

func sendMessage(ctx *gin.Context, res Result) {
	if res.err != nil {
		ctx.JSON(res.status, gin.H{
			"status":  res.status,
			"error":   res.err,
			"details": res.err.Error(),
		})
	} else {
		ctx.JSON(res.status, gin.H{
			"status": res.status,
			"data":   res.data,
		})
	}
}
