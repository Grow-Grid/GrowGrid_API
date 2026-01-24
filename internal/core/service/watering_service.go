package service

import (
	"growgrid/internal/core/ports"
)

type wateringService struct {
	repo ports.PlantRepository
}

func NewWateringService(r ports.PlantRepository) ports.WateringService {
	return &wateringService{
		repo: r,
	}
}
