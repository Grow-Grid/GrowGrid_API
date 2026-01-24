package main

import (
	"log"
	"net/http"

	httpAdapter "growgrid/internal/adapters/driving/http"
	"growgrid/internal/core/service"
)

func main() {
	plantService := service.NewPlantService(repo)
	handler := httpAdapter.NewHandler(plantService)
	mux := http.NewServeMux()
	handler.SetupRoutes(mux)
}
