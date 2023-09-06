package sqlite

import (
	"github.com/jmoiron/sqlx"
)

type Sqlite struct {
	db *sqlx.DB
}

type Tx struct {
	tx *sqlx.Tx
}

func ConnectSqlite() (*Sqlite, error) {
	// Connect a new database connection.
	db, err := sqlx.Open("sqlite3", "database/sqlite.db")
	if err != nil {
		return nil, err
	}

	return &Sqlite{
		db: db,
	}, nil
}

func (s *Sqlite) Close() error {
	return s.db.Close()
}

func (s *Sqlite) Ping() error {
	return s.db.Ping()
}

func (s *Sqlite) Begin() (*Tx, error) {
	tx, err := s.db.Beginx()
	if err != nil {
		return nil, err
	}

	return &Tx{tx: tx}, nil
}
