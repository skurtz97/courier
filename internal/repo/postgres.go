package repo

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresConfig struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Host     string `json:"host"`
	Table    string `json:"table"`
	Port     int    `json:"port"`
	SSL      string `json:"ssl"`
}

type PostgresPool struct {
	pool *pgxpool.Pool
}

func NewPostgresPool() (*PostgresPool, error) {
	pool, err := pgxpool.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	if err := pool.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}
	return &PostgresPool{pool: pool}, nil
}

func (p *PostgresPool) Close() {
	p.pool.Close()
}
