package Interface_test

import (
	"testing"
)

type Programmer interface {
	WritehelloWorld() string
}

type GoProgrammer struct {
}

func (g *GoProgrammer) WritehelloWorld() string {
	return "fmt.Println(\"Hello World\")"
}

/**
* Go接口与其他主要编程语言的差异：
*   1. 接口为非入侵性，实现不依赖于接口定义
*   2. 所有接口的定义可以包含在接口使用者包内
 */
func TestClient(t *testing.T) {
	var p Programmer
	p = new(GoProgrammer)
	t.Log(p.WritehelloWorld())
}
