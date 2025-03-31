package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type DBInterface interface {
	Query(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	QueryRow(ctx context.Context, query string, args ...any) *sql.Row
	Exec(ctx context.Context, query string, args ...any) (sql.Result, error)
	Prepare(ctx context.Context, query string) (*sql.Stmt, error)
	Close() error
}

type Database struct {
	conn *sql.DB
}

func NewDatabase(connStr string) (*Database, error) {
	if connStr == "" {
		return nil, errors.New("DATABASE_URL is not set")
	}

	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	conn.SetMaxOpenConns(25)
	conn.SetMaxIdleConns(10)
	conn.SetConnMaxLifetime(5 * time.Minute)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := conn.PingContext(ctx); err != nil {
		conn.Close()
		return nil, err
	}

	log.Println("Database connected with pooling")
	return &Database{conn: conn}, nil
}

func (d *Database) Query(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	rows, err := d.conn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	return rows, nil
}

func (d *Database) QueryRow(ctx context.Context, query string, args ...any) *sql.Row {
	return d.conn.QueryRowContext(ctx, query, args...)
}

func (d *Database) Exec(ctx context.Context, query string, args ...any) (sql.Result, error) {
	result, err := d.conn.ExecContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute command: %w", err)
	}
	return result, nil
}

func (d *Database) Prepare(ctx context.Context, query string) (*sql.Stmt, error) {
	stmt, err := d.conn.PrepareContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to execute prepare: %w", err)
	}

	return stmt, nil
}

func (d *Database) Close() error {
	if d.conn != nil {
		log.Println("Closing database connection")
		return d.conn.Close()
	}
	return nil
}
