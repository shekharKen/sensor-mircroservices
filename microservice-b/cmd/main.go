package main

import (
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"microservice-b/internal/db"
	grpcserver "microservice-b/internal/delivery/grpc"
	"microservice-b/internal/delivery/http"
	"microservice-b/internal/repository"
	"microservice-b/internal/usecase"
)

func main() {
	os.Setenv("DB_USER", "root")
	os.Setenv("DB_PASS", "Root@123")
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "3306")
	os.Setenv("DB_NAME", "interview_task_db")

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
