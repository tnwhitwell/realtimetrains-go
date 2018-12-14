package rtt

type RealtimeGbttDepartureLateness interface{}

type Location struct {
	RealtimeActivated             bool      `json:"realtimeActivated"`
	Tiploc                        string    `json:"tiploc"`
	Crs                           string    `json:"crs"`
	Description                   string    `json:"description"`
	WttBookedDeparture            string    `json:"wttBookedDeparture,omitempty"`
	GbttBookedDeparture           string    `json:"gbttBookedDeparture,omitempty"`
	Origin                        []Station `json:"origin"`
	Destination                   []Station `json:"destination"`
	IsCall                        bool      `json:"isCall"`
	IsPublicCall                  bool      `json:"isPublicCall"`
	RealtimeDeparture             string    `json:"realtimeDeparture,omitempty"`
	RealtimeDepartureActual       bool      `json:"realtimeDepartureActual,omitempty"`
	RealtimeGbttDepartureLateness `json:"realtimeGbttDepartureLateness,omitempty"`
	Platform                      string `json:"platform"`
	PlatformConfirmed             bool   `json:"platformConfirmed"`
	PlatformChanged               bool   `json:"platformChanged"`
	Line                          string `json:"line,omitempty"`
	LineConfirmed                 bool   `json:"lineConfirmed,omitempty"`
	DisplayAs                     string `json:"displayAs"`
	WttBookedArrival              string `json:"wttBookedArrival,omitempty"`
	GbttBookedArrival             string `json:"gbttBookedArrival,omitempty"`
	RealtimeArrival               string `json:"realtimeArrival,omitempty"`
	RealtimeArrivalActual         bool   `json:"realtimeArrivalActual,omitempty"`
	ServiceLocation               string `json:"serviceLocation,omitempty"`
}
type ServiceResponse struct {
	ServiceUID           string     `json:"serviceUid"`
	RunDate              string     `json:"runDate"`
	ServiceType          string     `json:"serviceType"`
	IsPassenger          bool       `json:"isPassenger"`
	TrainIdentity        string     `json:"trainIdentity"`
	PowerType            string     `json:"powerType"`
	TrainClass           string     `json:"trainClass"`
	AtocCode             string     `json:"atocCode"`
	AtocName             string     `json:"atocName"`
	PerformanceMonitored bool       `json:"performanceMonitored"`
	Origin               []Station  `json:"origin"`
	Destination          []Station  `json:"destination"`
	Locations            []Location `json:"locations"`
	RealtimeActivated    bool       `json:"realtimeActivated"`
	RunningIdentity      string     `json:"runningIdentity"`
}
