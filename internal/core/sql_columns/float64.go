package sqlcolumns

import "database/sql"

func Float64(nullable bool) SqlColumn {
	if nullable {
		return SqlColumn{
			Value: &sql.NullFloat64{},
			Scan: func(v any) any {
				val, _ := v.(*sql.NullFloat64).Value()
				return val
			},
		}
	}

	return SqlColumn{
		Value: new(float64),
		Scan: func(v any) any {
			return *v.(*float64)
		},
	}
}
