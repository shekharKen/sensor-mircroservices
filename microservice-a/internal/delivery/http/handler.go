package http

import (
	"log"
	"microservice-a/internal/data"
	usecase "microservice-a/internal/usercase"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	generator *usecase.GeneratorUsecase
}

func NewHandler(e *echo.Echo, gen *usecase.GeneratorUsecase) {
	h := &Handler{generator: gen}
	e.POST("/config/frequency", h.UpdateFrequency)
}

type FrequencyRequest struct {
	IntervalSeconds int `json:"interval_seconds"`
}

func (h *Handler) UpdateFrequency(c echo.Context) error {
	var req FrequencyRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	if req.IntervalSeconds <= 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Interval must be positive"})
	}

	h.generator.UpdateFrequency(time.Duration(req.IntervalSeconds) * time.Second)
	return c.JSON(http.StatusOK, map[string]string{"message": "Frequency updated successfully"})
}

func (h *Handler) GetFrequency(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"interval_seconds": int(h.generator.Frequency.Seconds()),
	})
}

func (h *Handler) GenerateOnce(c echo.Context) error {
	data := data.RandomSensorData(h.generator.SensorType)

	log.Printf("Manual generation: %+v\n", data)
	
	return c.JSON(http.StatusOK, data)
}
