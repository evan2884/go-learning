package string_test

import (
	"strconv"
	"strings"
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

func TestStringToRune(t *testing.T) {
	s := "中华人民共和国"
	for _, c := range s {
		t.Logf("%[1]c %[1]d", c)
	}
}

func TestStringFn(t *testing.T) {
	s := "A,B,C"
	sep := ","
	// 拆分
	parts := strings.Split(s, sep)
	for key, val := range parts {
		t.Log(key, val)
	}
	// 拼接
	t.Log(strings.Join(parts, "-"))
}

func TestStringConv(t *testing.T) {
	s := strconv.Itoa(10)
	t.Log("str => " + s)
	if sInt, error := strconv.Atoi("10"); error == nil {
		t.Log(10 + sInt)
	}
}
