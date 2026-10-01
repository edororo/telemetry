package model

import (
	"time"
	"uuid"
)

type Anomaly struct {
	ID          uuid.UUID `json:"id"`
	DroneID     uuid.UUID `json:"drone_id"`
	ZoneX       int       `json:"zone_x"`
	ZoneY       int       `json:"zone_y"`
	Type        string    `json:"type"`        // тип (батарея, помехи и т.д)
	Severity    string    `json:"severity"`    // степень серьёзности
	Description string    `json:"description"` // описание
	DetectedAt  time.Time `json:"detected_at"`
}
