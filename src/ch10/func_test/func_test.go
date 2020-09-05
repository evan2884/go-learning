package func_test

import (
	"fmt"
	"testing"
)

//可变参数
func sum(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret
}

// 延迟执行函数
func Clear() {
	fmt.Println("Clear resources of Function.")
}

func TestVarParam(t *testing.T) {
	t.Log(sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
}

// 延迟执行函数(defer函数)
func TestDefer(t *testing.T) {
	defer Clear()

	defer func() {
		t.Log("Clear resources")
	}()

	t.Log("Started")
	panic("Fatal error") // defer仍会执行
}
