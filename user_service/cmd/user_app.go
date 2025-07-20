package main

import (
	"context"
	"log"
	"user_service/infrastructure/database"
	"user_service/internal/datasource"
	"user_service/internal/service"
	"user_service/internal/web"

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

	if err := r.Run(":8282"); err != nil {
		log.Fatalln("Server failed: ", err)
	}
}
