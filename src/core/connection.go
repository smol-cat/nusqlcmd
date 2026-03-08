package core

import (
	"database/sql"
	_ "github.com/microsoft/go-mssqldb"
)

func ConnectToDb(connectionString string) (*sql.DB, error) {
	return sql.Open("sqlserver", connectionString)
}
