package slice_test

import (
	"testing"
)

func TestSliceInit(t *testing.T) {
	var s0 []int
	t.Log(len(s0), cap(s0))
	s0 = append(s0, 2)
	t.Log(len(s0), cap(s0))

	s1 := []int{1, 2, 3, 4, 5}
	t.Log(len(s1), cap(s1))

	s2 := make([]int, 3, 5)
	t.Log(len(s2), cap(s2))
	t.Log(s2[0], s2[1], s2[2])
	s2 = append(s2, 1)
	s2 = append(s2, 1)
	s2 = append(s2, 1)
	t.Log(s2[0], s2[1], s2[2], s2[3], s2[4], s2[5])
	t.Log(len(s2), cap(s2))
	t.Log(s2)

	s3 := []int{}
	for i := 0; i < 10; i++ {
		s3 = append(s3, 1)
		t.Log(len(s3), cap(s3))
	}
}

//是否可比较，容量是否可伸缩
//数组 容量不可伸缩 相同长度可比较
//切片 容量可以伸缩 不可比较
func TestSliceShareMemory(t *testing.T) {
	year := []string{"jan", "feb", "mar", "apr", "may", "jun", "jul", "aug", "sep", "oct", "nov", "dec"}
	Q2 := year[3:6]
	t.Log(Q2, len(Q2), cap(Q2))
	summer := year[5:8]
	t.Log(summer, len(summer), cap(summer))
	summer[0] = "unknow"
	t.Log(Q2)
	t.Log(year)
}

func TestSliceComparing(t *testing.T) {
	a := []int{1, 2, 3, 4}
	//b := []int{1, 2, 3, 4}
	//if a == b {
	if a == nil {
		t.Log("equal")
	} else {
		t.Log("not equal")
	}
}
