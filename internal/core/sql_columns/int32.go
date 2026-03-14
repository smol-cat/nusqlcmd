package sqlcolumns

import (
	"database/sql"

	"github.com/smol-cat/nusqlcmd/internal/core"
)

func allocateNullInt32() any {
	return new(sql.NullInt32)
}

func allocateInt32() any {
	return new(int)
}

func scanNullInt32(v any) any {
	val, _ := v.(*sql.NullInt32).Value()
	return val
}

func scanInt32(v any) any {
	return *v.(*int)
}

func Int32(nullable bool) core.SqlColumn {
	if nullable {
		return core.SqlColumn{
			AllocateValue: allocateNullInt32,
			Scan:          scanNullInt32,
		}
	}

	return core.SqlColumn{
		AllocateValue: allocateInt32,
		Scan:          scanInt32,
	}
}
