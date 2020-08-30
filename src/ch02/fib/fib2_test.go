package fib

import (
	"testing"
)

var a int

func TestFib2List(t *testing.T) {
	a = 1
	b := 1
	t.Log(a)
	for i := 0; i < 10; i++ {
		t.Log(" ", b)
		tmp := a
		a = b
		b = tmp + a
	}
}

func TestExchange(t *testing.T) {
	a := 1
	b := 2
	// tmp := a
	// a = b
	// b = tmp
	a, b = b, a
	t.Log(a, b)
}
