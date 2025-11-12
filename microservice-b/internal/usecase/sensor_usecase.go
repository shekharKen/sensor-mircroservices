package usecase

import (
	"microservice-b/internal/domain"
	"microservice-b/internal/repository"
	"time"
)

type SensorUsecase struct {
	repo *repository.SensorRepository
}

func NewSensorUsecase(r *repository.SensorRepository) *SensorUsecase {
	return &SensorUsecase{repo: r}
}

func (u *SensorUsecase) Save(data domain.SensorData) error {
	return u.repo.Save(data)
}

func (u *SensorUsecase) GetByFilter(id1 string, id2 int, start, end time.Time, page, limit int) ([]domain.SensorData, int, error) {
	return u.repo.GetByFilter(id1, id2, start, end, page, limit)
}

func (u *SensorUsecase) DeleteByFilter(id1 string, id2 int) error {
	return u.repo.DeleteByFilter(id1, id2)
}

func (u *SensorUsecase) UpdateByFilter(id1 string, id2 int, value float64) error {
	return u.repo.UpdateByFilter(id1, id2, value)
}