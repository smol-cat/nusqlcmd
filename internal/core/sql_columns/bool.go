package sqlcolumns

import (
	"database/sql"

	"github.com/smol-cat/nusqlcmd/internal/core"
)

func Bool(nullable bool) core.SqlColumn {
	if nullable {
		return core.SqlColumn{
			Value: &sql.NullBool{},
			Scan: func(v any) any {
				val, _ := v.(*sql.NullBool).Value()
				return val
			},
		}
	}

	return core.SqlColumn{
		Value: new(bool),
		Scan: func(v any) any {
			if v == nil {
				return nil
			}
			return *v.(*bool)
		},
	}
}
