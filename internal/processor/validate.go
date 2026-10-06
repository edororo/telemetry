package processor

import (
	"DroneTelemetry/internal/model"
	"fmt"
)

func ValidateTelemetry(t model.Telemetry) error {
	if t.Battery > 100 || t.Battery < 0 {
		return fmt.Errorf("Заряд батареи внедопустимых значениях")
	}

	if t.Signal > 100 || t.Signal < 0 {
		return fmt.Errorf("Сигнал внедопустимых значениях")
	}

	if t.PacketLoss > 100 || t.PacketLoss < 0 {
		return fmt.Errorf("ПэкетЛосс внедопустимых значениях")
	}

	if t.Speed < 0 {
		return fmt.Errorf("Такой скорости не может быть")
	}
	return nil
}
