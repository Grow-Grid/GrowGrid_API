package service

import (
	"growgrid/internal/core/ports"
	"growgrid/internal/core/service"
)

type PlantService struct {
	repo ports.PlantService
}

func (s *service) ProcessTelemetry(id string, moisture float64, temperature float64) error {
	return nil
}
