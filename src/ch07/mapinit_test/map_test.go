package mapinit_test

import "testing"

func TestMapInit(t *testing.T) {
	map1 := map[string]int{"one": 1, "two": 2, "three": 3}
	map2 := map[string]int{}
	map2["one"] = 1
	map3 := make(map[string]int, 10 /*Initial Capacity*/) // 为什么不初始化len？len单元格会赋为默认值0值
	map4 := map[int]int{1: 1, 2: 2, 3: 3}
	t.Log(map1, len(map1), map2, len(map2), map3, len(map3), map4, len(map4))
	t.Log(&map2, len(map2))
	map2["two"] = 2
	t.Log(&map2, len(map2))
}

func TestAccessNotExistingKey(t *testing.T) {
	map1 := map[int]int{}
	t.Log(map1[1])
	map1[2] = 0
	t.Log(map1[2])

	if val, ok := map1[3]; ok {
		t.Logf("Key 3'value is %d.", val)
	} else {
		t.Log("Key '3' is not existing.")
	}
}

func TestTravelMap(t *testing.T) {
	map1 := map[string]int{"one": 1, "two": 2, "three": 3}
	for key, val := range map1 {
		t.Log(key, val)
	}
}
