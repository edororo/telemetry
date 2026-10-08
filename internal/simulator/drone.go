package simulator

import (
	"DroneTelemetry/internal/model"
	"context"
	"fmt"
	"time"
	"uuid"
)

type DroneSimulator struct {
	ID          uuid.UUID
	Name        string
	Temperature float32
	Latitude    float32
	Longitude   float32
	Altitude    float32
	Speed       float32
	Battery     float32
	Heading     float32
	Signal      float32
	Latency     int
	PacketLoss  float32
}

// симулятор Дрона
func DroneRun(ctx context.Context, simulator DroneSimulator, ch chan<- model.Telemetry) {
	// тикер, который срабатывает каждые 500 мс
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()
	for {
		select {
		// остановка дрона по сигналу
		case <-ctx.Done():
			fmt.Println("Симулятор дрона остановлен")
			return
			// отправка в канал данных телеметрии
		case <-ticker.C:
			telemetry := model.Telemetry{
				ID:          uuid.New(),
				DroneID:     simulator.ID,
				Temperature: simulator.Temperature,
				Timestamp:   time.Now(),
				Latitude:    simulator.Latitude,
				Longitude:   simulator.Longitude,
				Altitude:    simulator.Altitude,
				Speed:       simulator.Speed,
				Battery:     simulator.Battery,
				Heading:     simulator.Heading,
				Signal:      simulator.Signal,
				Latency:     simulator.Latency,
				PacketLoss:  simulator.PacketLoss,
			}
			select {
			case <-ctx.Done():
				return
			case ch <- telemetry:
				fmt.Println("Отправка телеметрии в канал - успешно")
			}
		}

	}
}
