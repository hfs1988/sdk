package main

import (
	"database/sql"
	"sdk/adapters/db"

	_ "github.com/lib/pq"
)

type SQLDB interface {
	Connect() (*sql.DB, error)
	Save(db *sql.DB, sql db.SQL) error
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbname   = "postgres"
)

// const (
// 	uri = "mongodb://localhost:27017"
// )

func main() {
	var database SQLDB = db.GetPostgresInstance(host, port, user, password, dbname)
	// var database DB = db.GetMongoDBInstance(uri)
	dbSQL, err := database.Connect()
	if err != nil {
		panic(err)
	}
	database.Save(dbSQL, db.SQL{
		Table:  "users",
		Cols:   []string{"name", "email", "age"},
		Values: []interface{}{"Husni Firmansyah", "husni.firmansyah@gmail.com", 34},
	})
}
