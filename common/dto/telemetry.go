package dto

type Meta struct {
	TruckID   string  `json:"truck_id"`
	Timestamp float64 `json:"timestamp"`
}

// All data about the truck sent every 10 seconds
type Telemetry struct {
	Meta
	Lat   float64 `json:"lat"`
	Long  float64 `json:"long"`
	Speed float64 `json:"speed"`
}

// TODO: Consider combining speed and accel into MotionAlert to make logic easier
// maybe split things by
type SpeedAlert struct {
	Meta
	Speed float64 `json:"speed"`
}

type AccelAlert struct {
	Meta
	Acceleration float64 `json:"acceleration"`
}

type OffRouteAlert struct {
	Meta
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
}
