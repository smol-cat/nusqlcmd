package sqlcolumns

func ByteArray(nullable bool) SqlColumn {
	return SqlColumn {
		Value: new([]byte),
		Scan: func(v any) any {
			if v == nil {
				return nil
			}
			return *v.(*[]byte)
		},
	}
}
