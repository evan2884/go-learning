package fib

import (
	"fmt"
	"testing"
)

func TestFib1List(t *testing.T) {
	var (
		a int = 1
		b     = 1
	)
	t.Log(a)
	for i := 0; i < 10; i++ {
		t.Log(" ", b)
		tmp := a
		a = b
		b = tmp + a
	}

	fmt.Println()
}
