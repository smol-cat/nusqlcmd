package sqlcolumns

import (
	"database/sql"

	"github.com/smol-cat/nusqlcmd/internal/core"
)

func allocateNullInt16() any {
	return new(sql.NullInt16)
}

func allocateInt16() any {
	return new(int16)
}

func scanNullInt16(v any) any {
	val, _ := v.(*sql.NullInt16).Value()
	return val
}

func scanInt16(v any) any {
	return *v.(*int16)
}

func Int16(nullable bool) core.SqlColumn {
	if nullable {
		return core.SqlColumn{
			AllocateValue: allocateNullInt16,
			Scan:          scanNullInt16,
		}
	}

	return core.SqlColumn{
		AllocateValue: allocateInt16,
		Scan:          scanInt16,
	}
}
