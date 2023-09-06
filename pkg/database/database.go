package database

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type Database interface {
	// Ping pings the database.
	Ping(ctx context.Context) error

	// Insert inserts the given data into the database.
	Insert(ctx context.Context, data []byte) error

	// Get gets the data from the database.
	Get(ctx context.Context) ([]byte, error)

	// Select selects the data from the database.
	Select(ctx context.Context) ([]byte, error)

	Begin(ctx context.Context) (sqlx.Tx, error)

	// Close closes the database.
	Close(ctx context.Context) error
}
