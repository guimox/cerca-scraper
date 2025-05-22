package schedule

type train struct {
	Time        string `json:"time"`
	Destination string `json:"destination"`
	TrainName   string `json:"name"`
	Via         string `json:"via"`
}

func NewTrain() train {
	return train{
		Time:        "",
		Destination: "",
		TrainName:   "",
		Via:         "",
	}
}
