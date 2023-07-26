package db

type SQL struct {
	Table    string
	ColsVals SQLColsVals
	Filters  SQLColsVals
}

type SQLColsVals struct {
	Cols   []string
	Values []any
}
