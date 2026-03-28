package core

import (
	"database/sql"

	_ "github.com/microsoft/go-mssqldb"
	_ "github.com/microsoft/go-mssqldb/integratedauth/krb5"
	"github.com/smol-cat/nusqlcmd/internal/config"
)

func ConnectToDb(runtimeConfig config.RuntimeConfig) (*sql.DB, error) {
	return sql.Open(runtimeConfig.Driver, runtimeConfig.ConnectionString)
}
