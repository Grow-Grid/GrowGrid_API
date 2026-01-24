package http

import (
	"encoding/json"
	"growgrid/internal/core/ports"
	"net/http"
)

type Handler struct {
	service ports.WateringService
}

func NewHandler(service ports.WateringService) *Handler {
	return &Handler{service: service}
}

type telemtryReq struct {
	DeviceID  string  `json:"device_id"`
	Hum       float64 `json:"humidity"`
	Temp      float64 `json:"temperature"`
	HumSlope  float64 `json:"humidity_slope"`
	TempSlope float64 `json:"temperature_slope"`
}

func (h *Handler) HandleTelemetry(w http.ResponseWriter, r *http.Request) {
	var req telemtryReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	config, err := h.service.ProcessTelemetry(r.Context(), req.DeviceID,
		req.Hum, req.HumSlope, req.Temp, req.TempSlope)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(config)
}
