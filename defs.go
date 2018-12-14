package rtt

var baseURL = "https://api.rtt.io/api/v1/json"

type Station struct {
	Tiploc      string `json:"tiploc"`
	Description string `json:"description"`
	WorkingTime string `json:"workingTime"`
	PublicTime  string `json:"publicTime"`
}

type Services struct {
	Location        `json:"locationDetail"`
	ServiceUID      string `json:"serviceUid"`
	RunDate         string `json:"runDate"`
	TrainIdentity   string `json:"trainIdentity"`
	RunningIdentity string `json:"runningIdentity"`
	AtocCode        string `json:"atocCode"`
	AtocName        string `json:"atocName"`
	ServiceType     string `json:"serviceType"`
	IsPassenger     bool   `json:"isPassenger"`
}

type SimpleLocation struct {
	Name   string `json:"name"`
	Crs    string `json:"crs"`
	Tiploc string `json:"tiploc"`
}

type SearchResponse struct {
	SimpleLocation `json:"location"`
	Filter         interface{} `json:"filter"`
	Services       []Services  `json:"services"`
}
