package core

type TypeMapper interface {
	MapTypeNameToGoType(typeName string, nullable bool) any
	GetValueFromScanned(scanned any) any 
}
