package sqlcolumns

import (
	"database/sql"

	"github.com/smol-cat/nusqlcmd/internal/core"
)

func UInt8(nullable bool) core.SqlColumn {
	if nullable {
		return core.SqlColumn{
			Value: &sql.NullByte{},
			Scan: func(v any) any {
				val, _ := v.(*sql.NullByte).Value()
				return val
			},
		}
	}

	return core.SqlColumn{
		Value: new(uint8),
		Scan: func(v any) any {
			return *v.(*uint8)
		},
	}
}

