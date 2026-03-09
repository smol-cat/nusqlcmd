package mssql

import (
	"database/sql"
	"fmt"
	"strings"
)

type MssqlTypeMapper struct{}

func MapTypeNameToGoType(typeName string, nullable bool) any {
	switch strings.ToLower(typeName) {
	case "int", "bigint", "smallint", "tinyint":
		if nullable {
			return &sql.NullInt64{}
		}
		return new(int64)
	case "float", "real":
		if nullable {
			return &sql.NullFloat64{}
		}
		return new(float64)
	case "bit":
		if nullable {
			return &sql.NullBool{}
		}
		return new(bool)
	case "binary", "varbinary", "image":
		return new([]byte)
	default:
		if nullable {
			return &sql.NullString{}
		}
		return new(string)
	}
}

func GetValueFromScanned(scanned any) any {
	switch v := scanned.(type) {
	case *sql.NullInt64:
		if v.Valid {
			return v.Int64
		}
		return nil
	case *sql.NullFloat64:
		if v.Valid {
			return v.Float64
		}
		return nil
	case *sql.NullString:
		if v.Valid {
			return v.String
		}
		return nil
	case *sql.NullBool:
		if v.Valid {
			return v.Bool
		}
		return nil
	case *int64:
		return *v
	case *float64:
		return *v
	case *string:
		return *v
	case *bool:
		return *v
	case *[]byte:
		if v == nil {
			return nil
		}
		return *v
	case []byte:
		return v
	default:
		return fmt.Sprintf("%v", v)
	}
}
