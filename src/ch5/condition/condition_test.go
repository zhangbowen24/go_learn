package condition

import (
	"testing"
)

func TestIfMultiSec(t *testing.T) {
	if a := 1 != 1; a { //把 1 != 1运算后的值赋值给a 然后判断a
		t.Log("1 == 1")
	}

	/*if v, err := someFun(); err == nil {
		t.Log(v)
	} else {
		t.Log(err)
	}*/

	if v := someFun(); v == 1 {
		t.Log(v)
	} else {
		t.Log("error")
	}
}

func someFun() int {
	return 1
	//return 1, nil
}

func TestSwitchMultiSec(t *testing.T) {
	var job string = "writor"
	switch job {
	case "a", "b":
		t.Log("that")
	case "writor":
		t.Log("this")
	default:
		t.Log("any where")
	}

	for i := 0; i < 5; i++ {
		switch {
		case i%2 == 0:
			t.Log("even")
		case i%2 == 1:
			t.Log("odd")
		default:
			t.Log("unknow")
		}
	}
}
