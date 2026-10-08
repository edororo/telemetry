package processor

import (
	"DroneTelemetry/internal/model"
	"DroneTelemetry/internal/storage"
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Processor struct {
	pool *pgxpool.Pool
}

func NewProcessor(pool *pgxpool.Pool) *Processor {
	return &Processor{pool: pool}
}

func (p *Processor) ProcessTelemetry(ctx context.Context, i int, t model.Telemetry) error {
	fmt.Printf("Воркер: %d, TelemetryID: %v\n", i, t.ID)
	return storage.CreateTelemetry(ctx, p.pool, t)
}

func Worker(ctx context.Context, ch <-chan model.Telemetry, id int, p *Processor) {
	for {
		select {
		case <-ctx.Done():
			return
		case telemetry := <-ch:
			if err := ValidateTelemetry(telemetry); err != nil {
				fmt.Printf("Ошибка валидации %v\n", err)
				continue
			}
			err := p.ProcessTelemetry(ctx, id, telemetry)
			if err != nil {
				fmt.Printf("Ошибка сохранения телеметрии: %v\n", err)
				continue
			}
			time.Sleep(100 * time.Millisecond)
		}
	}
}
