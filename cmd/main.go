package main

import (
	"cerca-scraper/internal/config"
	"cerca-scraper/internal/handler"
	"cerca-scraper/internal/queue"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	rabbitMQ, err := queue.NewRabbitMQConfig("amqp://user:password@localhost:5672/")
	if err != nil {
		log.Fatal("cannot initialize RabbitMQ:", err)
	}
	defer rabbitMQ.Close()

	mux := http.NewServeMux()

	h := handler.NewHandler(rabbitMQ)

	mux.HandleFunc("GET /schedule/{stationNameSlug}", h.HandleSingleStation)
	mux.HandleFunc("GET /stations", handler.HandleAllStations)

	server := &http.Server{
		Addr:         config.GetServerAddress(),
		Handler:      mux,
		ReadTimeout:  config.Server.ReadTimeout,
		WriteTimeout: config.Server.WriteTimeout,
		IdleTimeout:  config.Server.IdleTimeout,
	}

	fmt.Printf("Server starting on http://localhost%s\n", config.GetServerAddress())
	log.Fatal(server.ListenAndServe())
}
