package domain

import "time"

type Plant struct {
	ID                      string
	Species                 string
	SpeciesType             string
	MinMoisture             float64
	MaxMoisture             float64
	LastWatered             time.Time
	CurrentStatus           string
	AssignedMicroController string
}
