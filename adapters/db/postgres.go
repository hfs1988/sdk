package db

import (
	"database/sql"
	"fmt"
	"reflect"
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
	counter, filters, args := getFilters(sql)

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

func (p *postgresDB) Get(db *sql.DB, sql SQL) []map[string]any {
	var results []map[string]any

	_, filters, args := getFilters(sql)

	statement := fmt.Sprintf(`
	SELECT %s FROM %s WHERE %s
	`, strings.Join(sql.ColsVals.Cols, ", "), sql.Table, strings.Join(filters, " and "))

	rows, _ := db.Query(statement, args...)
	defer rows.Close()

	for rows.Next() {
		columns, _ := rows.ColumnTypes()
		valuePointer := make([]interface{}, len(columns))
		for i, column := range columns {
			valuePointer[i] = reflect.New(column.ScanType()).Interface()
		}

		rows.Scan(valuePointer...)

		result := make(map[string]any)
		for k, v := range valuePointer {
			switch fmt.Sprintf("%T", v) {
			case "*string":
				value := v.(*string)
				result[sql.ColsVals.Cols[k]] = *value
			case "*int32":
				value := v.(*int32)
				result[sql.ColsVals.Cols[k]] = *value
			}
		}

		results = append(results, result)
	}

	return results
}

func getFilters(sql SQL) (int, []string, []any) {
	counter := 0
	var filters []string
	var args []any

	for k, v := range sql.Filters.Cols {
		counter++
		filters = append(filters, fmt.Sprintf("%s=$%d", v, k+1))
		args = append(args, sql.Filters.Values[k])
	}

	return counter, filters, args
}
