package scraper

import (
	"cercu-scraper/internal/constants"
	"cercu-scraper/internal/schedule"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

func ScrapeStation(stationSlug string) (schedule.TableData, error) {
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36"),
	)

	tableData := schedule.NewTableData()

	c.OnHTML("div.journal-content-article div.detalle-estacion h1", func(e *colly.HTMLElement) {
		stationNameText := e.Text
		stationNameText = strings.TrimSpace(stationNameText)
		if stationNameText != "" {
			tableData.StationName = stationNameText
		}
	})

	c.OnHTML("table.adif-table tr.horario-row", func(e *colly.HTMLElement) {
		train := schedule.NewTrain()

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

	url := constants.BaseURL + stationSlug

	err := c.Visit(url)
	if err != nil {
		return schedule.NewTableData(), err
	}

	tableData.Station = stationSlug
	tableData.Timestamp = time.Now().Format(time.RFC3339)

	return tableData, nil
}
