package processor

import (
	"DroneTelemetry/internal/model"
	"context"
	"fmt"
	"time"
)

func ProcessTelemetry(i int, t model.Telemetry) {
	fmt.Printf("Воркер: %d, TelemetryID: %v\n", i, t)
}

func Worker(ctx context.Context, ch <-chan model.Telemetry, id int) {
	for {
		select {
		case <-ctx.Done():
			return
		case telemetry := <-ch:
			if err := ValidateTelemetry(telemetry); err != nil {
				fmt.Printf("Ошибка валидации %v\n", err)
				continue
			}
			ProcessTelemetry(id, telemetry)
			time.Sleep(100 * time.Millisecond)
		}
	}
}
