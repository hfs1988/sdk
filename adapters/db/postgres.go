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
	for k, _ := range sql.ColsVals.Cols {
		vals = append(vals, fmt.Sprintf("$%d", k+1))
	}

	statement := fmt.Sprintf(`
	INSERT INTO %s (%s)
	VALUES (%s)`, sql.Table, strings.Join(sql.ColsVals.Cols, ", "), strings.Join(vals, ", "))

	_, err := db.Exec(statement, sql.ColsVals.Values...)
	if err != nil {
		return err
	}

	return nil
}

func (p *postgresDB) Update(db *sql.DB, sql SQL) error {
	var args []any
	var filters []string
	counter := 0
	for k, v := range sql.Filters.Cols {
		counter++
		filters = append(filters, fmt.Sprintf("%s=$%d", v, k+1))
		args = append(args, sql.Filters.Values[k])
	}

	var sets []string
	for k, v := range sql.ColsVals.Cols {
		sets = append(sets, fmt.Sprintf("%s=$%d", v, counter+k+1))
		args = append(args, sql.ColsVals.Values[k])
	}

	statement := fmt.Sprintf(`
	UPDATE %s
	SET %s
	WHERE %s
	`, sql.Table, strings.Join(sets, ", "), strings.Join(filters, " and "))

	_, err := db.Exec(statement, args...)
	if err != nil {
		return err
	}

	return nil
}
