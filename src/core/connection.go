package core

import "database/sql"

func ConnectToDb(connectionString string) (*sql.DB, error) {
	return sql.Open("sqlserver", connectionString)
}
