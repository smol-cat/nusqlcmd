package mssql

import (
	"database/sql"
	"strings"
)

func MapTypeNameToGoType(typeName string, nullable bool) (any, func(any) any) {
	switch strings.ToLower(typeName) {
	case "bigint":
		if nullable {
			return &sql.NullInt64{}, func(v any) any {
				val, _ := v.(*sql.NullInt64).Value()
				return val
			}
		}

		return new(int64), func(v any) any {
			return *v.(*int64)
		}
	case "int":
		if nullable {
			return &sql.NullInt32{}, func(v any) any {
				val, _ := v.(*sql.NullInt32).Value()
				return val
			}
		}
		return new(int32), func(v any) any {
			return *v.(*int32)
		}
	case "smallint":
		if nullable {
			return &sql.NullInt16{}, func(v any) any {
				val, _ := v.(*sql.NullInt16).Value()
				return val
			}
		}
		return new(int16), func(v any) any {
			return *v.(*int16)
		}
	case "tinyint":
		if nullable {
			return &sql.NullByte{}, func(v any) any {
				val, _ := v.(*sql.NullByte).Value()
				return val
			}
		}
		return new(int8), func(v any) any {
			return *v.(*int8)
		}
	case "float", "real":
		if nullable {
			return &sql.NullFloat64{}, func(v any) any {
				val, _ := v.(*sql.NullFloat64).Value()
				return val
			}
		}
		return new(float64), func(v any) any {
			return *v.(*float64)
		}
	case "bit":
		if nullable {
			return &sql.NullBool{}, func(v any) any {
				val, _ := v.(*sql.NullBool).Value()
				return val
			}
		}
		return new(bool), func(v any) any {
			return *v.(*bool)
		}
	case "binary", "varbinary", "image":
		return new([]byte), func(v any) any {
			if v == nil {
				return nil
			}
			return *v.(*[]byte)
		}
	default:
		if nullable {
			return &sql.NullString{}, func(v any) any {
				val, _ := v.(*sql.NullString).Value()
				return val
			}
		}
		return new(string), func(v any) any {
			return *v.(*string)
		}
	}
}
