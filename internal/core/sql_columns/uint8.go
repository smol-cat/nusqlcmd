package sqlcolumns

import (
	"database/sql"

	"github.com/smol-cat/nusqlcmd/internal/core"
)

func allocateNullUInt8() any {
	return new(sql.NullByte)
}

func allocateUInt8() any {
	return new(uint8)
}

func scanNullUInt8(v any) any {
	val, _ := v.(*sql.NullByte).Value()
	return val
}

func scanUInt8(v any) any {
	return *v.(*uint8)
}

func UInt8(nullable bool) core.SqlColumn {
	if nullable {
		return core.SqlColumn{
			AllocateValue: allocateNullUInt8,
			Scan:          scanNullUInt8,
		}
	}

	return core.SqlColumn{
		AllocateValue: allocateUInt8,
		Scan:          scanUInt8,
	}
}
