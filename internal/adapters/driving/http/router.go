package http

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

type telemetryReq struct {
	DeviceID string  `json:"device_id"`
	Hum      float64 `json:"humidity,string"`
	Temp     float64 `json:"temperature,string"`
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /health", h.HandleHealth)
	mux.HandleFunc("POST /mqtt", h.HandleSensors)
}

func (h *Handler) HandleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, `{"status":"OK!"}`)
}

func (h *Handler) HandleSensors(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	var data telemetryReq
	if err := decoder.Decode(&data); err != nil {
		http.Error(w, "Invalid Body", http.StatusBadRequest)
		log.Fatal(err)
		return
	}
	fmt.Printf("Datos recibidos: Dispositivo %s, Hum: %.2f, Temp: %.2f\n", data.DeviceID, data.Hum, data.Temp)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	fmt.Fprintln(w, `{"status":"Data recieved!"}`)
}

//func (h *Handler) HandleTelemetry(w http.ResponseWriter, r *http.Request) {
//	var req telemtryReq
//	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
//		http.Error(w, "Invalid request body", http.StatusBadRequest)
//		return
//	}
//	config, err := h.service.ProcessTelemetry(r.Context(), req.DeviceID,
//		req.Hum, req.HumSlope, req.Temp, req.TempSlope)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//	w.Header().Set("Content-Type", "application/json")
//	json.NewEncoder(w).Encode(config)
//}
