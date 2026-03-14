package sqlcolumns

import "github.com/smol-cat/nusqlcmd/internal/core"

func allocateByteArray() any {
	return new([]byte)
}

func scanByteArray(v any) any {
	if v == nil {
		return nil
	}
	return *v.(*[]byte)
}

func ByteArray(nullable bool) core.SqlColumn {
	if nullable {
		return core.SqlColumn{
			AllocateValue: allocateByteArray,
			Scan:          scanByteArray,
		}
	}

	return core.SqlColumn{
		AllocateValue: allocateByteArray,
		Scan:          scanByteArray,
	}
}
