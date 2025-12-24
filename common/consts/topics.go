package consts

const (
	TruckTelemetry = "truck-telemetry"
	SpeedAlert     = "speed-alert"
	AccelAlert     = "accel-alert"
	OffRouteAlert  = "off-route-alert"
	EndTrip        = "end-trip"
)

type TopicConfig struct {
	Name          string
	ConsumerGroup string
}

var Topics = map[string]TopicConfig{
	TruckTelemetry: {Name: TruckTelemetry, ConsumerGroup: "telemetry-group"},
	SpeedAlert:     {Name: SpeedAlert, ConsumerGroup: "alert-group"},
	AccelAlert:     {Name: AccelAlert, ConsumerGroup: "alert-group"},
	OffRouteAlert:  {Name: OffRouteAlert, ConsumerGroup: "alert-group"},
	EndTrip:        {Name: EndTrip, ConsumerGroup: "trip-group"},
}
