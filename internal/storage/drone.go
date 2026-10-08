package storage

import (
	"DroneTelemetry/internal/model"
	"context"
	"uuid"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Добавляем нового дрона в таблицу
func CreateDrone(ctx context.Context, pool *pgxpool.Pool, drone model.Drone) error {
	_, err := pool.Exec(ctx, "INSERT INTO DRONES (id, name, status, created_at, last_seen) VALUES ($1, $2, $3, $4, $5)", drone.ID, drone.Name, drone.Status, drone.CreatedAt, drone.LastSeen)
	if err != nil {
		return err
	}
	return nil
}

func GetDroneByID(ctx context.Context, pool *pgxpool.Pool, id uuid.UUID) (model.Drone, error) {
	drone := model.Drone{}
	row := pool.QueryRow(ctx, "SELECT id, name, status, created_at, last_seen FROM DRONES WHERE id = $1", id)
	// берём результат и записываем
	err := row.Scan(&drone.ID, &drone.Name, &drone.Status, &drone.CreatedAt, &drone.LastSeen)
	if err != nil {
		return model.Drone{}, err
	}
	return drone, nil
}
