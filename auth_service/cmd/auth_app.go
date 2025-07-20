package main

import (
	"auth_service/infrastructure/database"
	"auth_service/internal/datasource"
	"auth_service/internal/service"
	"auth_service/internal/web"
	"context"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	log.Println("Connecting to database")
	db, err := database.ConnectDB(context.Background())
	if err != nil {
		log.Fatalln("DB connection failed: ", err)
	}
	defer db.Close()
	handler := web.NewHandler(service.NewService(datasource.NewRepository(db)))
	handler.RegistrationHandlers(r)

	if err := r.Run(":8181"); err != nil {
		log.Fatalln("Server failed: ", err)
	}
}
