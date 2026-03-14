package core

type SqlColumn struct {
	AllocateValue func() any
	Scan func(any) any
}
