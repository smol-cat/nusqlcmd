package sqlcolumns

import "github.com/smol-cat/nusqlcmd/internal/core"

func ByteArray(nullable bool) core.SqlColumn {
	return core.SqlColumn {
		Value: new([]byte),
		Scan: func(v any) any {
			if v == nil {
				return nil
			}
			return *v.(*[]byte)
		},
	}
}
