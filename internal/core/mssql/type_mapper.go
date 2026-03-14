package mssql

import (
	"fmt"
	"os"

	"github.com/smol-cat/nusqlcmd/internal/core"
	mssql_sqlcolumns "github.com/smol-cat/nusqlcmd/internal/core/mssql/sqlcolumns"
	"github.com/smol-cat/nusqlcmd/internal/core/sql_columns"
)

func MapTypeNameToSqlType(typeName string, nullable bool) core.SqlColumn {
	switch typeName {
	case "BIGINT":
		return sqlcolumns.Int64(nullable)
	case "INT":
		return sqlcolumns.Int32(nullable)
	case "SMALLINT":
		return sqlcolumns.Int16(nullable)
	case "TINYINT":
		return sqlcolumns.UInt8(nullable)
	case "FLOAT", "REAL", "DECIMAL", "MONEY", "SMALLMONEY":
		return sqlcolumns.Float64(nullable)
	case "BIT":
		return sqlcolumns.Bool(nullable)
	case "BINARY", "VARBINARY", "IMAGE", "GEOGRAPHY", "GEOMETRY":
		return sqlcolumns.ByteArray(nullable)
	case "UNIQUEIDENTIFIER":
		return mssql_sqlcolumns.UUID(nullable)
	case "CHAR", "DATE", "DATETIME", "DATETIME2", "DATETIMEOFFSET", "HIERARCHYID", "NCHAR", "NTEXT", "NVARCHAR", "SMALLDATETIME", "SQL_VARIANT", "TEXT", "TIME", "VARCHAR", "XML":
		return sqlcolumns.String(nullable)
	default:
		fmt.Fprintf(os.Stderr, "Warning: Unrecognized type '%s', defaulting to string\n", typeName)
		return sqlcolumns.String(nullable)
	}
}
