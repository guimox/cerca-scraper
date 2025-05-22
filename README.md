# Cerca Scraper

A Go web scraper that fetches real-time train information from ADIF (Spanish Railway Infrastructure Administrator) stations.

## Description

Cerca Scraper is a web service that provides train schedule information for Madrid's Cercanías stations. It scrapes data from ADIF's website and returns it in a structured JSON format.

## Features

- Real-time train schedule information
- Support for all Madrid Cercanías stations
- RESTful API endpoint
- JSON formatted responses
- Station name and slug mapping
- Timestamp for each request

## API Usage

### Get Station Information

```http
GET /station/{station-name}
```

Example request:

```bash
curl http://localhost:8080/station/madrid-atocha-cercanias
```

Example response:

```json
{
  "timestamp": "2025-05-20T15:04:05+02:00",
  "trains": [
    {
      "time": "15:00",
      "destination": "Guadalajara",
      "train_id": "C2",
      "via": "1"
    }
  ],
  "station": "18000-madrid-atocha-c.",
  "station_name": "Madrid-Atocha Cercanías"
}
```

## Dependencies

- Colly

## Technical Details

- Written in Go
- Uses colly for web scraping
- Built-in HTTP server
- Supports concurrent requests
- Error handling for invalid stations and failed requests
