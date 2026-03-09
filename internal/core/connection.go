package core

import (
	"database/sql"
	_ "github.com/microsoft/go-mssqldb"
	_ "github.com/microsoft/go-mssqldb/integratedauth/krb5"
)

func ConnectToDb(connectionString string) (*sql.DB, error) {
	return sql.Open("sqlserver", connectionString)
}
