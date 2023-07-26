package client

import (
	"database/sql"
	"sdk/adapters/db"
)

// client interface
type SQLDB interface {
	Connect() (*sql.DB, error)
	Save(db *sql.DB, sql db.SQLEntity) error
	Update(db *sql.DB, sql db.SQLEntity) error
	Get(db *sql.DB, sql db.SQLEntity) []map[string]any
}