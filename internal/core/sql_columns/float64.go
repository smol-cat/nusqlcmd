package sqlcolumns

import (
	"database/sql"

	"github.com/smol-cat/nusqlcmd/internal/core"
)

func allocateNullFloat64() any {
	return new(sql.NullFloat64)
}

func allocateFloat64() any {
	return new(float64)
}

func scanNullFloat64(v any) any {
	val, _ := v.(*sql.NullFloat64).Value()
	return val
}

func scanFloat64(v any) any {
	return *v.(*float64)
}

func Float64(nullable bool) core.SqlColumn {
	if nullable {
		return core.SqlColumn{
			AllocateValue: allocateNullFloat64,
			Scan:          scanNullFloat64,
		}
	}

	return core.SqlColumn{
		AllocateValue: allocateFloat64,
		Scan:          scanFloat64,
	}
}
