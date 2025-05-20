package schedule

type train struct {
	Time        string `json:"time"`
	Destination string `json:"destination"`
	TrainID     string `json:"train_id"`
	Via         string `json:"via"`
}

func NewTrain() train {
	return train{
		Time:        "",
		Destination: "",
		TrainID:     "",
		Via:         "",
	}
}
