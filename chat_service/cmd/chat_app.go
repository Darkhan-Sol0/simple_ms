package main

import (
	"chat_service/infrastructure/database"
	"chat_service/internal/datasource"
	"chat_service/internal/service"
	"chat_service/internal/web"
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

	if err := r.Run(":8383"); err != nil {
		log.Fatalln("Server failed: ", err)
	}
}
