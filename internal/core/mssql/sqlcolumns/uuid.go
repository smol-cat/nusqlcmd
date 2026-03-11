package mssql_sqlcolumns

import (
	mssql_driver "github.com/microsoft/go-mssqldb"
	"github.com/smol-cat/nusqlcmd/internal/core"
)

func UUID(nullable bool) core.SqlColumn {
	if nullable {
		return core.SqlColumn{
			Value: &mssql_driver.NullUniqueIdentifier{},
			Scan: func(v any) any {
				if !v.(*mssql_driver.NullUniqueIdentifier).Valid {
					return nil
				}

				return v.(*mssql_driver.NullUniqueIdentifier).String()
			},
		}
	}

	return core.SqlColumn{
		Value: &mssql_driver.UniqueIdentifier{},
		Scan: func(v any) any {
			return v.(*mssql_driver.UniqueIdentifier).String()
		},
	}
}
