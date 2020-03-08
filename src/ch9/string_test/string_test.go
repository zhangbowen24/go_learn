package string_test

import (
	"testing"
	"unsafe"
)

func TestString(t *testing.T) {
	var s string
	t.Log(s)
	s = "hello"
	t.Log(len(s))
	//s[1] = 3 //不能修改的slice
	s = "\xE4\xB8\xA5"
	//s = "\xE4\xB8\xAD"
	//s = "\xE4\xBA\xBB\xFF"
	t.Log(s)
	s = "中国"
	t.Log(len(s))
	c := []rune(s)
	t.Log(c)
	t.Log(len(c))
	t.Log("rune size:", unsafe.Sizeof(c[0]))
	t.Logf("中 unicode %x", c[0])
	t.Logf("中 UTF8 %x", s)
}

func TestStringToRune(t *testing.T) {
	s := "中华人民共和国"
	for _, v := range s {
		t.Logf("%[1]c %[1]d", v)
		//上下相等
		//t.Logf("%c %d", v, v)
	}
}
