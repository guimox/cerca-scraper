package scraper

import (
	"cerca-scraper/internal/constants"
	"cerca-scraper/internal/schedule"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

func ScrapeStation(stationSlug string) (schedule.TableData, error) {
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36"),
	)

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
		r.Headers.Set("Accept-Language", "pt-BR,pt;q=0.9,en-US;q=0.8,en;q=0.7,es;q=0.6,ak;q=0.5")
		r.Headers.Set("Accept-Encoding", "gzip, deflate, br, zstd")
		r.Headers.Set("Cache-Control", "max-age=0")
		r.Headers.Set("Sec-Ch-Ua", "\"Chromium\";v=\"136\", \"Google Chrome\";v=\"136\", \"Not.A/Brand\";v=\"99\"")
		r.Headers.Set("Sec-Ch-Ua-Mobile", "?0")
		r.Headers.Set("Sec-Ch-Ua-Platform", "\"Linux\"")
		r.Headers.Set("Sec-Fetch-Dest", "document")
		r.Headers.Set("Sec-Fetch-Mode", "navigate")
		r.Headers.Set("Sec-Fetch-Site", "none")
		r.Headers.Set("Sec-Fetch-User", "?1")
		r.Headers.Set("Upgrade-Insecure-Requests", "1")
		r.Headers.Set("Priority", "u=0, i")
		r.Headers.Set("Cookie", "COOKIE_SUPPORT=true; GUEST_LANGUAGE_ID=es_ES")
	})

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
		train.TrainName = strings.TrimSpace(trainIDText)

		viaText := e.ChildText("td.col-via div")
		train.Via = strings.TrimSpace(viaText)

		if train.Time != "" || train.Destination != "" || train.TrainName != "" || train.Via != "" {
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
