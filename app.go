package main

import (
	"database/sql"
	"fmt"
	"sdk/adapters/db"

	_ "github.com/lib/pq"
)

type SQLDB interface {
	Connect() (*sql.DB, error)
	Save(db *sql.DB, sql db.SQL) error
	Update(db *sql.DB, sql db.SQL) error
	Get(db *sql.DB, sql db.SQL) []map[string]any
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
	// database.Save(dbSQL, db.SQL{
	// 	Table:  "users",
	// 	Cols:   []string{"name", "email", "age"},
	// 	Values: []interface{}{"Husni Firmansyah", "husni.firmansyah@gmail.com", 34},
	// })

	// database.Update(dbSQL, db.SQL{
	// 	Table: "users",
	// 	ColsVals: db.SQLColsVals{
	// 		Cols:   []string{"name", "email", "age"},
	// 		Values: []any{"Husni Firmansyah", "husni.firmansyah@gmail.com", 35},
	// 	},
	// 	Filters: db.SQLColsVals{
	// 		Cols:   []string{"id"},
	// 		Values: []any{1},
	// 	},
	// })

	results := database.Get(dbSQL, db.SQL{
		Table: "users",
		ColsVals: db.SQLColsVals{
			Cols: []string{"name", "email", "age"},
		},
		Filters: db.SQLColsVals{
			Cols:   []string{"id"},
			Values: []any{1},
		},
	})
	fmt.Println(results)
}
