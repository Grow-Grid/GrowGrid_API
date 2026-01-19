package http

import (
	"growgrid/internal/core/ports"
	"net/http"
)

type Handler struct {
	plantService ports.PlantService
}

func NewHandler(service ports.PlantService) *Handler {
	return &Handler{plantService: service}
}
