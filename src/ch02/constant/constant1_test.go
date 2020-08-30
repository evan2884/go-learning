package constant_test

import "testing"

const (
	Monday = iota + 1
	Tuesday
	Wednesday
)

const (
	Readable = 1 << iota
	Writeable
	Executeable
)

func TestConstant1(t *testing.T) {
	t.Log(Monday, Tuesday)
	t.Log(Readable, Writeable, Executeable)
	a := 7 // 0111
	t.Log(a&Readable == Readable, a&Writeable == Writeable, a&Executeable == Executeable)
}
