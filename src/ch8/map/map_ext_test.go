package map_ext

import (
	"testing"
)

func TestMapWithFuncVal(t *testing.T) {
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
	t.Log(m[1](2), m[2](3), m[3](4))
}

//map 实现 set
func TestMapForSet(t *testing.T) {
	s := map[int]bool{}
	s[1] = true
	n := 1
	if s[n] {
		t.Logf("%d is exist", n)
	} else {
		t.Logf("%d is not exist", n)
	}
	s[3] = true
	t.Log(len(s))
	delete(s, 1)
	if s[n] {
		t.Logf("%d is exist", n)
	} else {
		t.Logf("%d is not exist", n)
	}
}
