package func_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// 返回多结果
func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

// 函数参数
func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)

		fmt.Println("time spent: ", time.Since(start).Seconds())

		return ret
	}
}

func slowFun(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

func TestFunc(t *testing.T) {
	a, b := returnMultiValues()
	t.Log(a, " => ", b)

	tsSF := timeSpent(slowFun)
	t.Log(tsSF(10))
}
