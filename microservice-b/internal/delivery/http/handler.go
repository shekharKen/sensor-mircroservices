package http

import (
	"net/http"
	"strconv"
	"time"

	"microservice-b/internal/auth"
	"microservice-b/internal/usecase"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	usecase *usecase.SensorUsecase
}

func NewHandler(e *echo.Echo, uc *usecase.SensorUsecase) {
	h := &Handler{usecase: uc}

	e.POST("/auth/login", Login)

	dataGroup := e.Group("/data")
	dataGroup.Use(auth.JWTMiddleware)
	dataGroup.GET("", h.GetData)
	dataGroup.DELETE("", h.DeleteData)
	dataGroup.PUT("", h.UpdateData)
}

func Login(c echo.Context) error {
	type Credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var cred Credentials
	if err := c.Bind(&cred); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	// default credentials
	if cred.Username != "admin" || cred.Password != "password123" {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid credentials"})
	}

	token, err := auth.GenerateToken(cred.Username)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "could not generate token"})
	}

	return c.JSON(http.StatusOK, map[string]string{"token": token})
}

func (h *Handler) GetData(c echo.Context) error {
	id1 := c.QueryParam("ID1")
	id2Str := c.QueryParam("ID2")
	startStr := c.QueryParam("start")
	endStr := c.QueryParam("end")
	pageStr := c.QueryParam("page")
	limitStr := c.QueryParam("limit")

	id2, _ := strconv.Atoi(id2Str)
	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)

	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 20
	}

	start, _ := time.Parse(time.RFC3339, startStr)
	end, _ := time.Parse(time.RFC3339, endStr)

	data, total, err := h.usecase.GetByFilter(id1, id2, start, end, page, limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]any{"message": "Data retrieved successfully", "data": data, "total": total})
}

func (h *Handler) DeleteData(c echo.Context) error {
	id1 := c.QueryParam("ID1")
	id2Str := c.QueryParam("ID2")
	if id1 == "" || id2Str == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID1 and ID2 are required"})
	}

	id2, _ := strconv.Atoi(id2Str)

	if err := h.usecase.DeleteByFilter(id1, id2); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Data deleted successfully"})
}

func (h *Handler) UpdateData(c echo.Context) error {
	id1 := c.QueryParam("ID1")
	id2Str := c.QueryParam("ID2")
	if id1 == "" || id2Str == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID1 and ID2 are required"})
	}

	var payload struct {
		SensorValue string `json:"sensor_value"`
	}
	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}
	valueStr := payload.SensorValue

	id2, _ := strconv.Atoi(id2Str)
	value, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid sensor value"})
	}

	if err := h.usecase.UpdateByFilter(id1, id2, value); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Data updated successfully"})
}
