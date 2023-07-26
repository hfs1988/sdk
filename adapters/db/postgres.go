package db

import (
	"database/sql"
	"fmt"
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

func (p *postgresDB) Connect() error {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		p.host, p.port, p.user, p.password, p.dbname)
	_, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return err
	}

	fmt.Println("Successfully connected!")
	return nil
}
