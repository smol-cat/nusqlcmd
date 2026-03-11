package core

type SqlColumn struct {
	Value any
	Scan func(any) any
}
