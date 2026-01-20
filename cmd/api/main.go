package main

import (
	"log"
	"net/http"

	"growgrid/internal/adapters/driven/db"
	httpAdapter "growgrid/internal/adapters/driving/http"
	"growgrid/internal/core/service"
)

func main() {
	var dbURI string = "mongodb://admin:admin@db"
	repo, err := db.NewMongoRepository(dbURI)
	if err != nil {
		log.Fatal(err)
	}
	log.Default()
	plantService := service.NewPlantService(repo)
	handler := httpAdapter.NewHandler(plantService)
	mux := http.NewServeMux()
	handler.SetupRoutes(mux)
}
