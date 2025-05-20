package schedule

type TableData struct {
	Timestamp   string  `json:"timestamp"`
	Trains      []train `json:"trains"`
	Station     string  `json:"station"`
	StationName string  `json:"station_name"`
}

func NewTableData() TableData {
	return TableData{
		Timestamp:   "",
		Trains:      []train{},
		Station:     "",
		StationName: "",
	}
}
