package db

type SQLEntity struct {
	Table    string
	ColsVals SQLColsVals
	Filters  SQLColsVals
}

type SQLColsVals struct {
	Cols   []string
	Values []any
}
