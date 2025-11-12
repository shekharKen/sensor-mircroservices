package data

import (
	"math/rand"
	"strings"
	"time"

	"microservice-a/internal/domain"
)

func RandomSensorData(sensorType string) domain.SensorData {
	rand.Seed(time.Now().UnixNano())

	return domain.SensorData{
		SensorValue: rand.Float64() * 100, // random float 0–100
		SensorType:  sensorType,
		ID1:         randomID1(),
		ID2:         rand.Intn(100),
		Timestamp:   time.Now().UTC(),
	}
}

func randomID1() string {
	letters := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	var sb strings.Builder
	for i := 0; i < 3; i++ {
		sb.WriteRune(letters[rand.Intn(len(letters))])
	}
	return sb.String()
}
