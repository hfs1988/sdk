package client

import (
	"database/sql"
	"net/http"

	"github.com/hfs1988/sdk/adapters/db"
)

// client interface
type SQLDB interface {
	Connect() (*sql.DB, error)
	Save(db *sql.DB, sql db.SQLEntity) error
	Update(db *sql.DB, sql db.SQLEntity) error
	Get(db *sql.DB, sql db.SQLEntity) []map[string]any
}

type Router interface {
	Get(path string, f http.HandlerFunc)
	Post(path string, f http.HandlerFunc)
	ListenAndServe()
}
