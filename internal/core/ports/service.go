package ports

import "context"

type WateringService interface {
	ProcessTelemetry(ctx context.Context, deviceID string, hum float64, temp float64, humSlope float64, tempSlope float64) (WateringConfig, error)
}

type WateringConfig struct {
	Threshold float64 `json:"threshold"`
	Interval  int     `json:"interval_seconds"`
}
