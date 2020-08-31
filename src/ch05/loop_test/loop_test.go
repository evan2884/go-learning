package loop_test

import (
	"fmt"
	"runtime"
	"testing"
)

func TestWhileLoop(t *testing.T) {
	n := 0
	for n < 5 {
		t.Log(n)
		n++
	}

	for i := 0; i < 5; i++ {
		t.Log(i)
	}
}

func TestCondition(t *testing.T) {
	if a := 2 == 1; a {
		t.Log(a)
	} else {
		t.Log(a)
	}
}

func TestSwitch(t *testing.T) {
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
		// break
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s", os)
	}

	var num int = 5
	switch {
	case 0 <= num && num <= 3:
		fmt.Println("0 - 3")
	case 4 <= num && num <= 6:
		fmt.Println("4 - 6")
	case 7 <= num && num <= 9:
		fmt.Println("7 - 9")
	default:
		fmt.Println(" > 9")
	}

	// multi case
	for i := 0; i < 5; i++ {
		switch i {
		case 0, 2:
			t.Log("Even ", i)
		case 1, 3:
			t.Log("Odd ", i)
		default:
			t.Log("it is not 0-3 ", i)
		}
	}

	// case condition
	for i := 0; i < 5; i++ {
		switch {
		case i%2 == 0 && i < 4:
			t.Log("Even ", i)
		case i%2 == 1 && i < 4:
			t.Log("Odd ", i)
		default:
			t.Log("it is not 0-3 ", i)
		}
	}
}
