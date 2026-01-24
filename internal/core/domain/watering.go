package domain

type WateringParameters struct {
	MinHumidityThreshold float64 `json:"min_humidity_threshold"`
	CheckIntervalSeconds int     `json:"check_interval_seconds"`
}

type WaterEngine struct{}

// THIS FUNCTION DETERMINES THE SAFE THRESHOLDS FOR A GIVEN PLANT
// SLOPE REFERS TO THE RATE AT WHICH THE PLANT CAN TAKE WATER
// MINTEMP IS THE MINIMAL SAFE TEMPERATURE FOR THE PLANT
func (e *WaterEngine) CalculateDynamicConstants(p Plant) (slope float64, minTemp float64) {
	if p.SpeciesType == "Tropical" {
		slope = 0.8
		minTemp = 12
	} else if p.SpeciesType == "Cactus" {
		slope = 0.2
		minTemp = 5
	} else {
		// GENERAL STANDARD
		slope = 0.5
		minTemp = 9
	}
	return slope, minTemp
}
