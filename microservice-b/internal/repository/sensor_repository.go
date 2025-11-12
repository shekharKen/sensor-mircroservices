package repository

import (
	"time"

	"microservice-b/internal/domain"

	"gorm.io/gorm"
)

type SensorRepository struct {
	db *gorm.DB
}

func NewSensorRepository(db *gorm.DB) *SensorRepository {
	return &SensorRepository{db: db}
}

func (r *SensorRepository) Save(data domain.SensorData) error {
	return r.db.Create(&data).Error
}

func (r *SensorRepository) GetByFilter(id1 string, id2 int, start, end time.Time, page, limit int) ([]domain.SensorData, int, error) {
	var data []domain.SensorData
	q := r.db.Model(&domain.SensorData{})

	if id1 != "" {
		q = q.Where("id1 = ?", id1)
	}
	if id2 != 0 {
		q = q.Where("id2 = ?", id2)
	}
	if !start.IsZero() && !end.IsZero() {
		q = q.Where("timestamp BETWEEN ? AND ?", start, end)
	}

	offset := (page - 1) * limit
	err := q.Offset(offset).Limit(limit).Find(&data).Error
	if err != nil {
		return nil, 0, err
	}

	var total int64
	countQ := r.db.Model(&domain.SensorData{})
	if id1 != "" {
		countQ = countQ.Where("id1 = ?", id1)
	}
	if id2 != 0 {
		countQ = countQ.Where("id2 = ?", id2)
	}
	if !start.IsZero() && !end.IsZero() {
		countQ = countQ.Where("timestamp BETWEEN ? AND ?", start, end)
	}
	if cErr := countQ.Count(&total).Error; cErr != nil {
		return nil, 0, cErr
	}

	return data, int(total), nil
}

func (r *SensorRepository) DeleteByFilter(id1 string, id2 int) error {
	q := r.db.Model(&domain.SensorData{})
	if id1 != "" {
		q = q.Where("id1 = ?", id1)
	}
	if id2 != 0 {
		q = q.Where("id2 = ?", id2)
	}
	return q.Delete(&domain.SensorData{}).Error
}

func (r *SensorRepository) UpdateByFilter(id1 string, id2 int, value float64) error {
	query := r.db.Model(&domain.SensorData{})

	if id1 != "" {
		query = query.Where("id1 = ?", id1)
	}
	if id2 != 0 {
		query = query.Where("id2 = ?", id2)
	}

	return query.Update("sensor_value", value).Error
}


