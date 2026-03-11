package sqlcolumns

import (
	"database/sql"

	"github.com/smol-cat/nusqlcmd/internal/core"
)

func Int16(nullable bool) core.SqlColumn {
	if nullable {
		return core.SqlColumn{
			Value: &sql.NullInt16{},
			Scan: func(v any) any {
				val, _ := v.(*sql.NullInt16).Value()
				return val
			},
		}
	}

	return core.SqlColumn{
		Value: new(int16),
		Scan: func(v any) any {
			return *v.(*int16)
		},
	}
}
