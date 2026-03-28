package serialization

import (
	"errors"

	"github.com/smol-cat/nusqlcmd/internal/core"
	"github.com/smol-cat/nusqlcmd/internal/core/mssql"
	"github.com/smol-cat/nusqlcmd/internal/core/postgres"
)

func getMapper(driver string) ((func(string, bool) core.SqlColumn), error) {
	switch driver {
	case "sqlserver":
		return mssql.MapTypeNameToSqlType, nil
	case "postgres":
		return postgres.MapTypeNameToSqlType, nil
	default:
		return nil, errors.New("Unsupported")
	}
}
