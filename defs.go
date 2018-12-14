package rtt

var baseURL = "https://api.rtt.io/api/v1/json"

type SearchResponse struct {
	Location struct {
		Name   string `json:"name"`
		Crs    string `json:"crs"`
		Tiploc string `json:"tiploc"`
	} `json:"location"`
	Filter   interface{} `json:"filter"`
	Services []struct {
		LocationDetail struct {
			RealtimeActivated   bool   `json:"realtimeActivated"`
			Tiploc              string `json:"tiploc"`
			Crs                 string `json:"crs"`
			Description         string `json:"description"`
			WttBookedArrival    string `json:"wttBookedArrival"`
			WttBookedDeparture  string `json:"wttBookedDeparture"`
			GbttBookedArrival   string `json:"gbttBookedArrival"`
			GbttBookedDeparture string `json:"gbttBookedDeparture"`
			Origin              []struct {
				Tiploc      string `json:"tiploc"`
				Description string `json:"description"`
				WorkingTime string `json:"workingTime"`
				PublicTime  string `json:"publicTime"`
			} `json:"origin"`
			Destination []struct {
				Tiploc      string `json:"tiploc"`
				Description string `json:"description"`
				WorkingTime string `json:"workingTime"`
				PublicTime  string `json:"publicTime"`
			} `json:"destination"`
			IsCall                  bool   `json:"isCall"`
			IsPublicCall            bool   `json:"isPublicCall"`
			RealtimeArrival         string `json:"realtimeArrival"`
			RealtimeArrivalActual   bool   `json:"realtimeArrivalActual"`
			RealtimeDeparture       string `json:"realtimeDeparture"`
			RealtimeDepartureActual bool   `json:"realtimeDepartureActual"`
			Platform                string `json:"platform"`
			PlatformConfirmed       bool   `json:"platformConfirmed"`
			PlatformChanged         bool   `json:"platformChanged"`
			DisplayAs               string `json:"displayAs"`
		} `json:"locationDetail"`
		ServiceUID      string `json:"serviceUid"`
		RunDate         string `json:"runDate"`
		TrainIdentity   string `json:"trainIdentity"`
		RunningIdentity string `json:"runningIdentity"`
		AtocCode        string `json:"atocCode"`
		AtocName        string `json:"atocName"`
		ServiceType     string `json:"serviceType"`
		IsPassenger     bool   `json:"isPassenger"`
	} `json:"services"`
}
