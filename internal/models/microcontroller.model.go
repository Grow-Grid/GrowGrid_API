package models

type Microcontroller struct {
	ID          string  `bson:"id"`
	Humidity    float64 `bson:"humidity"`
	Temperature float64 `bson:"temperature"`
}
