package main

import (
	"github.com/gin-gonic/gin"
	"messaging_service/internal/db"
	"messaging_service/internal/handlers"
	"messaging_service/internal/kafka"
	"messaging_service/internal/repositories"
	"messaging_service/internal/services"
	"os"
)

func main() {
	dbConnString := os.Getenv("DATABASE_URL")
	db.InitDB(dbConnString)
	defer db.CloseDB()

	kafkaBroker := os.Getenv("KAFKA_BROKER")
	kafkaTopic := os.Getenv("KAFKA_TOPIC")
	kafkaProducer := kafka.NewKafkaProducer(kafkaBroker, kafkaTopic)

	repo := repositories.NewMessageRepository()
	service := services.NewMessageService(repo, kafkaProducer)
	handler := handlers.NewMessageHandler(service)

	r := gin.Default()
	r.POST("/messages", handler.HandleMessage)
	r.GET("/statistics", handler.GetStatistics)

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
