package ports

import "growgrid/internal/core/domain"

type PlantRepository interface {
	SaveSensorData(data domain.SensorData) error
	GetThresholds(plantID string) (domain.Thresholds, error)
}
