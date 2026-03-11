package sqlcolumns

import "database/sql"

func String(nullable bool) SqlColumn {
	if nullable {
		return SqlColumn{
			Value: &sql.NullString{},
			Scan: func(v any) any {
				val, _ := v.(*sql.NullString).Value()
				return val
			},
		}
	}

	return SqlColumn{
		Value: new(string),
		Scan: func(v any) any {
			return *v.(*string)
		},
	}
}
