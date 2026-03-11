package sqlcolumns

import "database/sql"

func UInt8(nullable bool) SqlColumn {
	if nullable {
		return SqlColumn{
			Value: &sql.NullByte{},
			Scan: func(v any) any {
				val, _ := v.(*sql.NullByte).Value()
				return val
			},
		}
	}

	return SqlColumn{
		Value: new(uint8),
		Scan: func(v any) any {
			return *v.(*uint8)
		},
	}
}

