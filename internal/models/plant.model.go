package models

type Plant struct {
	ID                 string  `bson:"id"`
	Type               string  `bson:"type"`
	MinMoisture        float64 `bson:"min_moisture"`
	MaxMoisture        float64 `bson:"max_moisture"`
	MinTemperature     float64 `bson:"min_temperature"`
	MaxTemperature     float64 `bson:"max_temperature"`
	Microcontroller_ID string  `bson:"microcontroller_id"`
}
