package database

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func InitDB() error {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable&search_path=retail",
		getEnv("DB_USER", "admin"),
		getEnv("DB_PASSWORD", "admin123"),
		getEnv("DB_HOST", "localhost"),
		getEnv("DB_PORT", "5432"),
		getEnv("DB_NAME", "retailBrain"),
	)

	fmt.Println("Connecting to Postgres with DSN:", dsn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return fmt.Errorf("failed to connect to PostgreSQL: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		return fmt.Errorf("failed to ping DB: %w", err)
	}

	DB = pool
	return nil
}

func getEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallback
}
