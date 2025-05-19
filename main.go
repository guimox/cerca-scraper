package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/gocolly/colly"
)

type Train struct {
	Time        string `json:"time"`
	Destination string `json:"destination"`
	TrainID     string `json:"train_id"`
	Via         string `json:"via"`
}

type TableData struct {
	Trains []Train `json:"trains"`
}

func main() {
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36"),
	)

	var tableData TableData

	c.OnHTML("table.adif-table tr.horario-row", func(e *colly.HTMLElement) {
		train := Train{}

		timeText := e.ChildText("td.col-hora div span")
		train.Time = strings.TrimSpace(timeText)

		destinationText := e.ChildText("td.col-destino div")
		train.Destination = strings.TrimSpace(destinationText)

		trainIDText := e.ChildText("td.col-tren div span.lineColored")
		train.TrainID = strings.TrimSpace(trainIDText)

		viaText := e.ChildText("td.col-via div")
		train.Via = strings.TrimSpace(viaText)

		if train.Time != "" || train.Destination != "" || train.TrainID != "" || train.Via != "" {
			tableData.Trains = append(tableData.Trains, train)
		}
	})

	url := "https://www.adif.es/w/18000-madrid-atocha-c." // Adjust this URL as needed
	fmt.Printf("Scraping %s...\n", url)
	err := c.Visit(url)
	if err != nil {
		log.Fatalf("Error visiting website: %v", err)
	}

	jsonResult, err := json.MarshalIndent(tableData, "", "  ")
	if err != nil {
		log.Fatalf("Error creating JSON: %v", err)
	}

	fmt.Println(string(jsonResult))
}
