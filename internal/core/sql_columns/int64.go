package sqlcolumns

import "database/sql"

func Int64(nullable bool) SqlColumn {
	if nullable {
		return SqlColumn{
			Value: &sql.NullInt64{},
			Scan: func(v any) any {
				val, _ := v.(*sql.NullInt64).Value()
				return val
			},
		}
	}

	return SqlColumn{
		Value: new(int64),
		Scan: func(v any) any {
			return *v.(*int64)
		},
	}
}
