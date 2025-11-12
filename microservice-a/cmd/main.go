package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"microservice-a/internal/delivery/http"
	usecase "microservice-a/internal/usercase"

	"github.com/labstack/echo/v4"
)

func main() {
	sensorType := os.Getenv("SENSOR_TYPE")
	if sensorType == "" {
		sensorType = "TEMPERATURE"
	}

	freqStr := os.Getenv("FREQUENCY")
	freqSec, err := strconv.Atoi(freqStr)
	if err != nil || freqSec <= 0 {
		freqSec = 2
	}

	gen := usecase.NewGeneratorUsecase(sensorType, time.Duration(freqSec)*time.Second)

	// Connect to Microservice B
	if err := gen.ConnectGRPC("localhost:50051"); err != nil {
		log.Printf("Failed to connect gRPC: %v", err)
	}

	gen.Start()

	e := echo.New()
	http.NewHandler(e, gen)

	log.Printf("Microservice A started | SensorType: %s | Frequency: %ds", sensorType, freqSec)
	e.Logger.Fatal(e.Start(":8080"))
}
