package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h Handler) MainHandler(c *gin.Context) {
	text := "Main Text"
	c.JSON(http.StatusOK, gin.H{
		"test": text,
	})
}
