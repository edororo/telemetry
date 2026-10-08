package storage

import (
	"DroneTelemetry/internal/model"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateTelemetry(ctx context.Context, pool *pgxpool.Pool, telemetry model.Telemetry) error {
	_, err := pool.Exec(ctx, "INSERT INTO TELEMETRY (id, drone_id, timestamp, latitude, longitude,altitude,speed,battery,temperature, heading,signal,latency,packet_loss) VALUES ($1, $2, $3, $4, $5,$6, $7, $8, $9, $10,$11, $12, $13)", telemetry.ID, telemetry.DroneID, telemetry.Timestamp, telemetry.Latitude, telemetry.Longitude, telemetry.Altitude, telemetry.Speed, telemetry.Battery, telemetry.Temperature, telemetry.Heading, telemetry.Signal, telemetry.Latency, telemetry.PacketLoss)
	if err != nil {
		return err
	}
	return nil
}
