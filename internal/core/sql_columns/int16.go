package sqlcolumns

import "database/sql"

func Int16(nullable bool) SqlColumn {
	if nullable {
		return SqlColumn{
			Value: &sql.NullInt16{},
			Scan: func(v any) any {
				val, _ := v.(*sql.NullInt16).Value()
				return val
			},
		}
	}

	return SqlColumn{
		Value: new(int16),
		Scan: func(v any) any {
			return *v.(*int16)
		},
	}
}
