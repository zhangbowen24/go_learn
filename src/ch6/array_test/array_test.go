package array_test

import (
	"testing"
)

func TestArrayInit(t *testing.T) {
	var arr [3]int
	arr1 := [3]int{1, 2, 3}
	arr2 := [...]int{1, 2, 3, 4, 5, 6}
	arr3 := [2][2]int{{1, 2}, {3, 4}}
	for _, e := range arr3 {
		t.Log(e)
		for idx, v := range e {
			t.Log(idx, v)
		}
	}
	t.Log(arr[0], arr[1], arr[2])
	t.Log(arr1, arr2)
}

func TestArrayTravel(t *testing.T) {
	arr2 := [...]int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr2); i++ {
		t.Log(arr2[i])
	}

	for idx, e := range arr2 {
		t.Log(idx, e)
	}

	for _, e := range arr2 {
		t.Log(e)
	}
}

func TestArraySubString(t *testing.T) {
	a := [...]int{1, 2, 3, 4, 5, 6}
	t.Log(a[1:2])
	t.Log(a[1:3])
	t.Log(a[1:])
	t.Log(a[1:len(a)])
	t.Log(a[:3])
}
