package my_map

import (
	"testing"
)

func TestInitMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 2, 3: 3}
	t.Log(m1, len(m1))
	m2 := map[string]int{"a": 1, "b": 2, "c": 3}
	t.Log(m2, len(m2))
	m3 := map[string]string{"a": "d", "b": "e", "c": "f"}
	t.Log(m3, len(m3))
	m4 := map[int]int{}
	m4[4] = 3
	t.Log(m4, len(m4))

	m5 := make(map[int]int, 10)
	t.Log(m5, len(m5))
}

//访问的key不存在时仍然返回0值，不能通过返回nil判断元素是否存在
func TestAccessNotExistingKey(t *testing.T) {
	m1 := map[int]int{}
	t.Log(m1[1])
	m1[2] = 0
	t.Log(m1[2])
	m1[3] = 4
	if v, ok := m1[3]; ok {
		t.Logf("key 3 exist value = %d", v)
	} else {
		t.Log("key 3 not exist")
	}
}

func TestTravelMap(t *testing.T) {
	m3 := map[string]string{"a": "d", "b": "e", "c": "f"}
	for k, v := range m3 {
		t.Log(k, v)
	}
}
