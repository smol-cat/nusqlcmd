package sqlcolumns

import "database/sql"

func Int32(nullable bool) SqlColumn {
	if nullable {
		return SqlColumn{
			Value: &sql.NullInt32{},
			Scan: func(v any) any {
				val, _ := v.(*sql.NullInt32).Value()
				return val
			},
		}
	}

	return SqlColumn{
		Value: new(int),
		Scan: func(v any) any {
			return *v.(*int)
		},
	}
}
