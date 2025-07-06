package main

import (
	"gateway/internal/web"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	handler := web.NewHandler()
	handler.RegistrateHandler(r)
	r.Run(":8080")
}
