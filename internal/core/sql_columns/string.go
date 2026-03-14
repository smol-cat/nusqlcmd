package sqlcolumns

import (
	"database/sql"

	"github.com/smol-cat/nusqlcmd/internal/core"
)

func allocateNullString() any {
	return new(sql.NullString)
}

func allocateString() any {
	return new(string)
}

func scanNullString(v any) any {
	val, _ := v.(*sql.NullString).Value()
	return val
}

func scanString(v any) any {
	return *v.(*string)
}

func String(nullable bool) core.SqlColumn {
	if nullable {
		return core.SqlColumn{
			AllocateValue: allocateNullString,
			Scan:          scanNullString,
		}
	}

	return core.SqlColumn{
		AllocateValue: allocateString,
		Scan:          scanString,
	}
}
