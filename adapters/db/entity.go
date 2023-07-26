package db

type SQL struct {
	Table  string
	Cols   []string
	Values []interface{}
}
