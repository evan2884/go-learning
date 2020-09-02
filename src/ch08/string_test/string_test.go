package string_test

import (
	"testing"
	"unsafe"
)

func TestString(t *testing.T) {
	var str string
	t.Log(str) // 初始化为默认零值

	str = "hello"
	t.Log(str, len(str))
	// str[1] = '3' // string是不可变的byte slice
	str = "\xE4\xB8\xA5" // 可以存储任何二进制数据
	// str = "\xE4\xBA\xBB\xFF"
	t.Log(str, len(str))
	str = "中"
	t.Log(str, len(str)) // 是byte数

	// 1. Unicode是一种字符集（code point）
	// 2. UTF8是Unicode的存储实现(转换为字节序列的规则)
	c := []rune(str) // 取出字符串里的Unicode
	t.Log(c, len(c))
	t.Log("rune size:", unsafe.Sizeof(c[0]))
	t.Logf("中 unicode %x", c[0])
	t.Logf("中 UTF8 %x", str)
}
