package main

import (
	"database/sql"
	"fmt"

	_ "github.com/microsoft/go-mssqldb"
)

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

	rows, err := dbConnection.Query("SELECT @@VERSION")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	rows.Next()
	var version string;
	if err := rows.Scan(&version); err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(version)
}
