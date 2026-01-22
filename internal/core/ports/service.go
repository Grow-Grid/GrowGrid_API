package ports

type PlantService interface {
	RegisterPlant(plantID string, microcontrollerID string)
}
