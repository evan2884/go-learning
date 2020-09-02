package map_test

import "testing"

func TestMapWithFunValue(t *testing.T) {
	m := map[int]func(op int) int{}
	m[1] = func(op int) int {
		return op
	}
	m[2] = func(op int) int {
		return op * op
	}
	m[3] = func(op int) int {
		return op * op * op
	}
	num := 2
	t.Log(m[1](num), m[2](num), m[3](num))
}

func TestMapForSet(t *testing.T) {
	mySet := map[int]bool{}
	mySet[1] = true
	n := 3
	if mySet[n] {
		t.Logf("%d is existing.", n)
	} else {
		t.Logf("%d is not existing.", n)
	}
	t.Log(mySet, len(mySet))
	mySet[3] = true
	t.Log(mySet, len(mySet))

	delete(mySet, 1)
	t.Log(mySet, len(mySet))
}
