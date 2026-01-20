package domain

import "time"

type MicroController struct {
	ID          string
	FirstAdded  time.Time
	LastCheckup time.Time
}
