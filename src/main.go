package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/microsoft/go-mssqldb"
)

type RowScanner struct {
}

func (r RowScanner) Scan(src any) error {
	return nil
}

func panicOnErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	dbConnection, err := sql.Open("sqlserver", "sqlserver://sa:password123!@localhost:1433?database=master")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if dbConnection == nil {
		fmt.Println("Failed to connect to the DB")
		return
	}

	fmt.Println(os.Args[1])
	rows, err := dbConnection.Query(os.Args[1])
	panicOnErr(err)

	colNames, err := rows.Columns()
	panicOnErr(err)

	for rows.Next() {
		var rawCols = make([]any, len(colNames))

		for i := range rawCols {
			var alloc string
			rawCols[i] = &alloc
		}

		err := rows.Scan(rawCols...)
		panicOnErr(err)

		for i := range rawCols {
			fmt.Printf("%s ", *(rawCols[i].(*string)))
		}

		fmt.Print("\n")
	}
}
