package domain

import "time"

type SensorData struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	SensorValue float64   `json:"sensor_value"`
	SensorType  string    `json:"sensor_type"`
	ID1         string    `json:"ID1"`
	ID2         int       `json:"ID2"`
	Timestamp   time.Time `json:"timestamp"`
	CreatedAt   time.Time
}
