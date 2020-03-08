package type_test

import (
	"math"
	"testing"
)

//不支持隐式类型转换
type MyInt int64

func TestImplicit(t *testing.T) {
	//var a int = 1
	var a int32 = 1
	var b int64
	b = int64(a) //显示类型转换

	var c MyInt
	c = MyInt(b)
	t.Log(a, b, c, math.MaxInt32, math.MaxUint8)
}

func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a //获取指针
	//aPtr = aPtr + 1 //不支持指针运算
	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr) //格式化输出变量类型%T
}

func TestString(t *testing.T) {
	var s string
	t.Log("-" + s + "-")
	t.Log(len(s))
	if s == "" {
		t.Log("is empty")
	}
}
