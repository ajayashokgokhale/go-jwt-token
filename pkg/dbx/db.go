// Package dbx provides database connection management for MySQL using a connection pool.
package dbx

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// Default connection pool settings.
const (
	defaultMaxIdleConns    = 5
	defaultMaxOpenConns    = 20
	defaultConnMaxLifetime = 60 * time.Minute
	defaultPingRetries     = 3
	defaultPingInterval    = 5 * time.Second
)

var db *sql.DB

// Init initializes the MySQL database connection using environment variables.
// It configures a connection pool and verifies connectivity with retries.
// Required environment variables: DB_USER, DB_PASS, DB_HOST, DB_PORT, DB_NAME.
// Optional: DB_MAX_IDLE_CONNS, DB_MAX_OPEN_CONNS, DB_CONN_MAX_LIFETIME_MINUTES.
// Returns an error if initialization fails.
func Init() error {
	// Load .env file only if critical variables are not set.
	if os.Getenv("DB_USER") == "" || os.Getenv("DB_HOST") == "" || os.Getenv("DB_NAME") == "" {
		if err := godotenv.Load(); err != nil {
			return fmt.Errorf("failed to load .env file, ensure variables are set: %w", err)
		}
	}

	// Validate required environment variables.
	requiredVars := []string{"DB_USER", "DB_PASS", "DB_HOST", "DB_PORT", "DB_NAME"}
	for _, v := range requiredVars {
		if os.Getenv(v) == "" {
			return fmt.Errorf("missing required environment variable: %s", v)
		}
	}

	// Create DSN for MySQL connection.
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	// Open database connection.
	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("failed to open database connection: %w", err)
	}

	// Configure connection pool.
	maxIdleConns := defaultMaxIdleConns
	if v := os.Getenv("DB_MAX_IDLE_CONNS"); v != "" {
		if n, err := strconv.Atoi(v); err == nil && n > 0 {
			maxIdleConns = n
		}
	}
	maxOpenConns := defaultMaxOpenConns
	if v := os.Getenv("DB_MAX_OPEN_CONNS"); v != "" {
		if n, err := strconv.Atoi(v); err == nil && n > 0 {
			maxOpenConns = n
		}
	}
	connMaxLifetime := defaultConnMaxLifetime
	if v := os.Getenv("DB_CONN_MAX_LIFETIME_MINUTES"); v != "" {
		if n, err := strconv.Atoi(v); err == nil && n > 0 {
			connMaxLifetime = time.Duration(n) * time.Minute
		}
	}
	db.SetMaxIdleConns(maxIdleConns)
	db.SetMaxOpenConns(maxOpenConns)
	db.SetConnMaxLifetime(connMaxLifetime)

	// Verify connectivity with retries.
	for i := 0; i < defaultPingRetries; i++ {
		if err := db.Ping(); err == nil {
			return nil
		}
		if i < defaultPingRetries-1 {
			time.Sleep(defaultPingInterval)
		}
	}
	return fmt.Errorf("failed to ping database after %d retries: %w", defaultPingRetries, err)
}

// GetDB returns the initialized database connection.
// Panics if the database is not initialized (Init must be called first).
func GetDB() *sql.DB {
	if db == nil {
		panic("database not initialized; call dbx.Init first")
	}
	return db
}
