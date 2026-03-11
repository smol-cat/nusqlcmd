package sqlcolumns

import (
	"database/sql"

	"github.com/smol-cat/nusqlcmd/internal/core"
)

func Float64(nullable bool) core.SqlColumn {
	if nullable {
		return core.SqlColumn{
			Value: &sql.NullFloat64{},
			Scan: func(v any) any {
				val, _ := v.(*sql.NullFloat64).Value()
				return val
			},
		}
	}

	return core.SqlColumn{
		Value: new(float64),
		Scan: func(v any) any {
			return *v.(*float64)
		},
	}
}
