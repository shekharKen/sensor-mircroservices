package main

import (
	"log"

	"microservice-b/internal/db"
	grpcserver "microservice-b/internal/delivery/grpc"
	"microservice-b/internal/delivery/http"
	"microservice-b/internal/repository"
	"microservice-b/internal/usecase"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database := db.InitMySQL()
	repo := repository.NewSensorRepository(database)
	uc := usecase.NewSensorUsecase(repo)

	// Start gRPC server in background
	go grpcserver.StartGRPCServer(uc)

	// Start HTTP server
	e := echo.New()
	http.NewHandler(e, uc)

	log.Println("REST API running on :8081")
	e.Logger.Fatal(e.Start(":8081"))
}
