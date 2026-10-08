package main

import (
	"DroneTelemetry/internal/model"
	"DroneTelemetry/internal/processor"
	"DroneTelemetry/internal/simulator"
	"DroneTelemetry/internal/storage"
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
	"uuid"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	pool, err := storage.NewPool(ctx)
	if err != nil {
		fmt.Printf("Ошибка подключения к БД: %v\n", err)
		return
	}
	defer pool.Close()
	wg := &sync.WaitGroup{}
	ch := make(chan model.Telemetry)
	p := processor.NewProcessor(pool)

	dronesim := simulator.DroneSimulator{ID: uuid.New(),
		Name:        "Вася",
		Temperature: 35,
		Latitude:    1,
		Longitude:   2,
		Altitude:    3,
		Speed:       4,
		Battery:     10,
		Heading:     6,
		Signal:      7,
		Latency:     8,
		PacketLoss:  9}
	drone := model.Drone{ID: dronesim.ID, Name: dronesim.Name, Status: "active", CreatedAt: time.Now(), LastSeen: time.Now()}
	err = storage.CreateDrone(ctx, pool, drone)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		return
	}

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			simulator.DroneRun(ctx, dronesim, ch)
		}()
	}
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			processor.Worker(ctx, ch, workerID, p)
		}(i + 1)
	}
	for {
		select {
		case <-ctx.Done():
			wg.Wait()
			fmt.Println("Завершено")
			return
		}
	}
}
