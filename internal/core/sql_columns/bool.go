package sqlcolumns

import "database/sql"

func Bool(nullable bool) SqlColumn {
	if nullable {
		return SqlColumn{
			Value: &sql.NullBool{},
			Scan: func(v any) any {
				val, _ := v.(*sql.NullBool).Value()
				return val
			},
		}
	}

	return SqlColumn{
		Value: new(bool),
		Scan: func(v any) any {
			if v == nil {
				return nil
			}
			return *v.(*bool)
		},
	}
}
