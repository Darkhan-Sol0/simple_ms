package main

import (
	"gateway/internal/web"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	handler := web.NewHandler()
	handler.RegistrateHandler(r)
	if err := r.Run(":8080"); err != nil {
		log.Fatalln("Server failed: ", err)
	}
}
