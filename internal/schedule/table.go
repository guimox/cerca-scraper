package schedule

type TableData struct {
	Timestamp   string  `json:"timestamp"`
	Station     string  `json:"station"`
	StationName string  `json:"station_name"`
	Trains      []train `json:"trains"`
}

func NewTableData() TableData {
	return TableData{
		Timestamp:   "",
		Trains:      []train{},
		Station:     "",
		StationName: "",
	}
}
