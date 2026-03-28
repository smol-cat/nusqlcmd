package postgres

import (
	"github.com/smol-cat/nusqlcmd/internal/common"
	"github.com/smol-cat/nusqlcmd/internal/core"
	sqlcolumns "github.com/smol-cat/nusqlcmd/internal/core/sql_columns"
)

func MapTypeNameToSqlType(typeName string, nullable bool) core.SqlColumn {
	switch typeName {
	default:
		common.WarnUnrecognizedType(typeName)
		return sqlcolumns.String(nullable)
	}
}
