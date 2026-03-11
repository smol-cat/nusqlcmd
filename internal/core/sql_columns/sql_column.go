package sqlcolumns

type SqlColumn struct {
	Value any
	Scan func(any) any
}
