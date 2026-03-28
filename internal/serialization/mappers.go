package serialization

import (
	"errors"
	"github.com/smol-cat/nusqlcmd/internal/core"
	"github.com/smol-cat/nusqlcmd/internal/core/mssql"
)

func getMapper(driver string) ((func(string, bool) core.SqlColumn), error) {
	switch driver {
	case "sqlserver":
		return mssql.MapTypeNameToSqlType, nil
	default:
		return nil, errors.New("Unsupported")
	}
}
