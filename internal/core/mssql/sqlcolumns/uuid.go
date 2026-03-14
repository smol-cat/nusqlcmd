package mssql_sqlcolumns

import (
	mssql_driver "github.com/microsoft/go-mssqldb"
	"github.com/smol-cat/nusqlcmd/internal/core"
)

func allocateNullUUID() any {
	return new(mssql_driver.NullUniqueIdentifier)
}

func allocateUUID() any {
	return new(mssql_driver.UniqueIdentifier)
}

func scanNullUUID(v any) any {
	if !v.(*mssql_driver.NullUniqueIdentifier).Valid {
		return nil
	}

	return v.(*mssql_driver.NullUniqueIdentifier).String()
}

func scanUUID(v any) any {
	return v.(*mssql_driver.UniqueIdentifier).String()
}

func UUID(nullable bool) core.SqlColumn {
	if nullable {
		return core.SqlColumn{
			AllocateValue: allocateNullUUID,
			Scan:          scanNullUUID,
		}
	}

	return core.SqlColumn{
		AllocateValue: allocateUUID,
		Scan:          scanUUID,
	}
}
