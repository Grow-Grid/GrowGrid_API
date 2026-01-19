package ports

type plantService interface {
	RegisterPlant(plantID string, microcontrollerID string)
}
