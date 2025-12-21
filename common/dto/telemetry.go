package dto

type Telemetry struct {
	TruckID   string  `json:"truck_id"`
	Lat       float64 `json:"lat"`
	Long      float64 `json:"long"`
	Speed     float64 `json:"speed"`
	Timestamp float64 `json:"timestamp"`
}
