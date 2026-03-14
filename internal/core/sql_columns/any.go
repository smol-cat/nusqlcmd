package sqlcolumns

import (
	"github.com/smol-cat/nusqlcmd/internal/core"
)

func allocateAny() any {
	return new(any)
}

func scanAny(v any) any {
	return *v.(*any)
}

func Any() core.SqlColumn {
	return core.SqlColumn{
		AllocateValue: allocateAny,
		Scan:          scanAny,
	}
}
