package service

import (
	"growgrid/internal/core/ports"
)

type PlantService struct {
	repo ports.PlantRepository
}

func NewPlantService(repo ports.PlantRepository) *PlantService {
	return &PlantService{
		repo: repo,
	}
}
