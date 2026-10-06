package model

import (
	"time"
	"uuid"
)

type Telemetry struct {
	ID          uuid.UUID `json:"id"`
	DroneID     uuid.UUID `json:"drone_id"`
	Timestamp   time.Time `json:"timestamp"`
	Latitude    float32   `json:"latitude"`  // широта
	Longitude   float32   `json:"longitude"` // долгота
	Altitude    float32   `json:"altitude"`  // высота
	Speed       float32   `json:"speed"`
	Battery     float32   `json:"battery"`
	Temperature float32   `json:"temperature"`
	Heading     float32   `json:"heading"`
	Signal      float32   `json:"signal"`
	Latency     int       `json:"latency"` // задержка
	PacketLoss  float32   `json:"packetLoss"`
}
