package sqlcolumns

import (
	"database/sql"
	"github.com/smol-cat/nusqlcmd/internal/core"
)

func Int64(nullable bool) core.SqlColumn {
	if nullable {
		return core.SqlColumn{
			Value: &sql.NullInt64{},
			Scan: func(v any) any {
				val, _ := v.(*sql.NullInt64).Value()
				return val
			},
		}
	}

	return core.SqlColumn{
		Value: new(int64),
		Scan: func(v any) any {
			return *v.(*int64)
		},
	}
}
