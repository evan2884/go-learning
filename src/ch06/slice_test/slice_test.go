package slice_test

import "testing"

func TestSliceInit(t *testing.T) {
	var s0 []int
	t.Log(s0, len(s0), cap(s0))

	s0 = append(s0, 1)
	t.Log(s0, len(s0), cap(s0))

	s1 := []int{1, 2, 3, 4}
	t.Log(s1, len(s1), cap(s1))

	s2 := make([]int, 3, 5)
	t.Log(s2, len(s2), cap(s2))
	t.Log(s2[0], s2[1], s2[2])
	// t.Log(s2[0], s2[1], s2[2], s2[3], s2[4])
	s2 = append(s2, 1)
	t.Log(s2, len(s2), cap(s2))
	t.Log(s2[0], s2[1], s2[2], s2[3])
	// t.Log(s2[0], s2[1], s2[2], s2[3], s2[4])
	s2 = append(s2, 2)
	t.Log(s2, len(s2), cap(s2))
	t.Log(s2[0], s2[1], s2[2], s2[3], s2[4])
}

func TestSliceGrowing(t *testing.T) {
	s := []int{}
	for i := 0; i < 10; i++ {
		s = append(s, i)
		t.Log(len(s), cap(s))
	}
}

func TestSliceShareMemory(t *testing.T) {
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	Q1 := year[3:6]
	Q2 := year[5:8]
	t.Log(year, len(year), cap(year))
	t.Log(Q1, len(Q1), cap(Q1))
	t.Log(Q2, len(Q2), cap(Q2))
	Q2[0] = "Unknow"
	t.Log(year, len(year), cap(year))
	t.Log(Q1, len(Q1), cap(Q1))
	t.Log(Q2, len(Q2), cap(Q2))
}

func TestSliceComparing(t *testing.T) {
	a := []int{1, 2, 3, 4}
	b := []int{1, 2, 3, 4}

	if a == b {
		t.Log("equal")
	}
}
