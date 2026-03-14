package sqlcolumns

import (
	"database/sql"

	"github.com/smol-cat/nusqlcmd/internal/core"
)

func allocateNullBool() any {
	return new(sql.NullBool)
}

func allocateBool() any {
	return new(bool)
}

func scanNullBool(v any) any {
	val, _ := v.(*sql.NullBool).Value()
	return val
}

func scanBool(v any) any {
	return *v.(*bool)
}

func Bool(nullable bool) core.SqlColumn {
	if nullable {
		return core.SqlColumn{
			AllocateValue: allocateNullBool,
			Scan:          scanNullBool,
		}
	}

	return core.SqlColumn{
		AllocateValue: allocateBool,
		Scan:          scanBool,
	}
}
