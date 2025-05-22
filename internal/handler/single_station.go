package handler

import (
	"cerca-scraper/internal/constants"
	"cerca-scraper/internal/queue"
	"cerca-scraper/internal/scraper"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Handler struct {
	rabbitMQ *queue.RabbitMQConfig
}

func NewHandler(rabbitMQ *queue.RabbitMQConfig) *Handler {
	return &Handler{
		rabbitMQ: rabbitMQ,
	}
}

func (h *Handler) HandleSingleStation(w http.ResponseWriter, r *http.Request) {
	stationName := r.PathValue("stationNameSlug")

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	stationSlug, exists := constants.Stations[stationName]
	if !exists {
		http.Error(w, "Station not found", http.StatusNotFound)
		return
	}

	data, err := scraper.ScrapeStation(stationSlug)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error scraping station: %v", err), http.StatusInternalServerError)
		return
	}

	if err := h.rabbitMQ.PublishSchedule(data); err != nil {
		log.Printf("Error publishing to queue: %v", err)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("Error encoding JSON response: %v", err)
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}
