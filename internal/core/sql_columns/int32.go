package sqlcolumns

import (
	"database/sql"

	"github.com/smol-cat/nusqlcmd/internal/core"
)

func Int32(nullable bool) core.SqlColumn {
	if nullable {
		return core.SqlColumn{
			Value: &sql.NullInt32{},
			Scan: func(v any) any {
				val, _ := v.(*sql.NullInt32).Value()
				return val
			},
		}
	}

	return core.SqlColumn{
		Value: new(int),
		Scan: func(v any) any {
			return *v.(*int)
		},
	}
}
