package sqlcolumns

import (
	"database/sql"
	"github.com/smol-cat/nusqlcmd/internal/core"
)

func allocateNullInt64() any {
	return new(sql.NullInt64)
}

func allocateInt64() any {
	return new(int64)
}

func scanNullInt64(v any) any {
	val, _ := v.(*sql.NullInt64).Value()
	return val
}

func scanInt64(v any) any {
	return *v.(*int64)
}

func Int64(nullable bool) core.SqlColumn {
	if nullable {
		return core.SqlColumn{
			AllocateValue: allocateNullInt64,
			Scan:          scanNullInt64,
		}
	}

	return core.SqlColumn{
		AllocateValue: allocateInt64,
		Scan:          scanInt64,
	}
}
