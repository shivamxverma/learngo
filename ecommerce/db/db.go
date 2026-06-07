package db

import (
	"context"
	_ "embed"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

//go:embed schema.sql
var SchemaSQL string

// DB holds our connection pool
var DB *pgxpool.Pool

// InitDB loads environment variables, connects to the database, and runs schemas
func InitDB(ctx context.Context) (*pgxpool.Pool, error) {
	// Load environment variables from .env file if it exists (ignored in production)
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on system environment variables")
	}

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		return nil, fmt.Errorf("DATABASE_URL environment variable is not set")
	}

	// Parse configuration
	config, err := pgxpool.ParseConfig(databaseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse database URL: %w", err)
	}

	// Set connection pool configurations
	config.MaxConns = 10
	config.MinConns = 2
	config.MaxConnIdleTime = 15 * time.Minute
	config.MaxConnLifetime = 1 * time.Hour

	// Establish connection pool
	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("failed to create connection pool: %w", err)
	}

	// Test connection
	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("Successfully connected to PostgreSQL database")
	DB = pool

	// Run migration to create tables if they do not exist
	if err := RunMigrations(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("failed to run database migrations: %w", err)
	}

	return pool, nil
}

// RunMigrations executes the embedded schema.sql script to ensure tables exist
func RunMigrations(ctx context.Context) error {
	log.Println("Running database migrations...")
	_, err := DB.Exec(ctx, SchemaSQL)
	if err != nil {
		return fmt.Errorf("error executing schema SQL: %w", err)
	}
	log.Println("Database migrations applied successfully")
	return nil
}
