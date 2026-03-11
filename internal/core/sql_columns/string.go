package sqlcolumns

import (
	"database/sql"

	"github.com/smol-cat/nusqlcmd/internal/core"
)

func String(nullable bool) core.SqlColumn {
	if nullable {
		return core.SqlColumn{
			Value: &sql.NullString{},
			Scan: func(v any) any {
				val, _ := v.(*sql.NullString).Value()
				return val
			},
		}
	}

	return core.SqlColumn{
		Value: new(string),
		Scan: func(v any) any {
			return *v.(*string)
		},
	}
}
