package main // 包，表明代码所在的模块（包）

// 引入代码依赖
import (
	"fmt"
	"os"
)

// 功能实现
func main() {
	if len(os.Args) > 1 {
		fmt.Println("Hello, world!!!", os.Args[1])
	}
	fmt.Println("Hello, world!!!")
	os.Exit(-1)
}
