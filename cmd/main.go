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
	dronesim := simulator.DroneSimulator{ID: uuid.New(),
		Name:       "Вася",
		Latitude:   1,
		Longitude:  2,
		Altitude:   3,
		Speed:      4,
		Battery:    10,
		Heading:    6,
		Signal:     7,
		Latency:    8,
		PacketLoss: 9}

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			simulator.DroneRun(ctx, dronesim, ch)
		}()
	}
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			processor.Worker(ctx, ch, i+1)
		}()
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
