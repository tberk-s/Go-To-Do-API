package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

// Item ...
type Item struct {
	Task   string
	Status string
}

// DB ...
type DB struct {
	pool *pgxpool.Pool
}

// New ...
func New(user, password, host, dbname string, port int) (*DB, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", user, password, host, port, dbname)

	pool, err := pgxpool.Connect(context.Background(), connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to db: %w", err)
	}

	if err := pool.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("failed to ping the db: %w", err)
	}

	return &DB{pool: pool}, nil
}

// InsertItem ...
func (db *DB) InsertItem(ctx context.Context, item Item) error {
	query := `INSERT INTO todo_items (task, status) VALUES ($1, $2)`
	_, err := db.pool.Exec(ctx, query, item.Task, item.Status)

	return err
}

// GetAllItems ...
func (db *DB) GetAllItems(ctx context.Context) ([]Item, error) {
	query := "SELECT task, status FROM todo_items"
	rows, err := db.pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var items []Item
	for rows.Next() {
		var item Item
		err := rows.Scan(&item.Task, &item.Status)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return items, nil
}

// DeleteItem ...
func (db *DB) DeleteItem(ctx context.Context, task string) error {
	query := "DELETE FROM todo_items WHERE task = $1"
	_, err := db.pool.Exec(ctx, query, task)

	return err
}

// Close ...
func (db *DB) Close() {
	db.pool.Close()
}
