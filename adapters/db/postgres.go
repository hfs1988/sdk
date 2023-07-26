package db

import (
	"database/sql"
	"fmt"
	"strings"
)

type postgresDB struct {
	host     string
	port     int
	user     string
	password string
	dbname   string
}

func GetPostgresInstance(host string, port int, user string, password string, dbname string) *postgresDB {
	return &postgresDB{
		host:     host,
		port:     port,
		user:     user,
		password: password,
		dbname:   dbname,
	}
}

func (p *postgresDB) Connect() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		p.host, p.port, p.user, p.password, p.dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	fmt.Println("Successfully connected!")
	return db, nil
}

func (p *postgresDB) Save(db *sql.DB, sql SQL) error {
	var vals []string
	for k, _ := range sql.Cols {
		vals = append(vals, fmt.Sprintf("$%d", k+1))
	}

	statement := fmt.Sprintf(`
	INSERT INTO %s (%s)
	VALUES (%s)`, sql.Table, strings.Join(sql.Cols, ", "), strings.Join(vals, ", "))

	_, err := db.Exec(statement, sql.Values...)
	if err != nil {
		return err
	}

	return nil
}

// func getArgs(db *sql.DB, sql SQL) []interface{} {
// 	var colsVals = make(map[string]string)
// 	for k, v := range sql.Cols {
// 		colsVals[v] = sql.Values[k]
// 	}

// 	var args []interface{}
// 	tcols, _ := schema.ColumnTypes(db, "", sql.Table)
// 	for _, v := range tcols {
// 		var value interface{} = colsVals[v.Name()]
// 		switch v.DatabaseTypeName() {
// 		case "INTEGER":
// 			value, _ = strconv.Atoi(value.(string))
// 		}

// 		args = append(args, value)
// 	}

// 	return args
// }
