package main

import (
	"log"
	"net/http"

	httpAdapter "growgrid/internal/adapters/driving/http"
)

func main() {
	handler := httpAdapter.NewHandler()
	mux := http.NewServeMux()
	handler.RegisterRoutes(mux)
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
	log.Println("Server started succesfully")
}
