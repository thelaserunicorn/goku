package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Config holds database connection configuration
type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

// Resource represents a resource stored in the database
type Resource struct {
	ID        int64
	Data      map[string]interface{}
}

// DB wraps the database pool and provides CRUD operations
type DB struct {
	pool *pgxpool.Pool
}

// GetConfig returns database configuration from environment variables
func GetConfig() Config {
	return Config{
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     getEnv("DB_PORT", "5433"),
		User:     getEnv("DB_USER", "postgres"),
		Password: getEnv("DB_PASSWORD", "postgres"),
		DBName:   getEnv("DB_NAME", "goku"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// New creates a new DB instance and connects to the database
func New(cfg Config) (*DB, error) {
	connString := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName,
	)

	pool, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		return nil, fmt.Errorf("failed to create connection pool: %w", err)
	}

	// Test the connection
	if err := pool.Ping(context.Background()); err != nil {
		pool.Close()
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &DB{pool: pool}, nil
}

// InitSchema creates the resource_table if it doesn't exist
func (db *DB) InitSchema() error {
	query := `
		CREATE TABLE IF NOT EXISTS resource_table (
			id SERIAL PRIMARY KEY,
			data JSONB NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`
	_, err := db.pool.Exec(context.Background(), query)
	if err != nil {
		return fmt.Errorf("failed to create schema: %w", err)
	}
	log.Println("Database schema initialized")
	return nil
}

// Close releases the database connection pool
func (db *DB) Close() {
	db.pool.Close()
}