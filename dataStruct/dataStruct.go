package dataStruct

type DataStatus struct {
	StatusWater string `json:"StatusWater"`
	StatusWind  string `json:"StatusWind"`
	DataStatus  StatusW
}

type StatusW struct {
	Water int `json:"Water"`
	Wind  int `json:"Wind"`
}