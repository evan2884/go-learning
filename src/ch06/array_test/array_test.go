package array_test

import "testing"

func TestArrayInit(t *testing.T) {
	var arr1 [3]int
	t.Log(arr1[1], arr1[2])

	arr2 := [4]int{1, 2, 3, 4}
	arr3 := [...]int{1, 2, 3, 4, 5}

	arr4 := [3][3]int{{1, 3, 4}, {2, 5, 9}, {3, 0, 8}}
	arr5 := [...][3]int{{1, 3, 4}, {2, 5, 9}, {3, 0, 8}}

	arr2[1] = 5
	t.Log(arr1[1], arr1[2])
	t.Log(arr2, arr3, arr4, arr5)

	a := [...]int{1, 2, 3, 4, 5}
	// 数组截取： a[开始索引(包含), 结束索引(不包含)]
	t.Log(a[1:2], a[1:3], a[1:len(a)], a[1:], a[:3])
}

func TestArrayTravel(t *testing.T) {
	arr := [...]int{1, 3, 4, 5}
	for i := 0; i < len(arr); i++ {
		t.Log(arr[i])
	}

	for idx, val := range arr {
		t.Log(idx, "===>", val)
	}

	for _, val := range arr {
		t.Log("===>", val)
	}
}
