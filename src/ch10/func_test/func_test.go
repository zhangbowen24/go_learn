package func_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent", time.Since(start).Seconds())
		return ret
	}
}

func slowFun(op int) int {
	time.Sleep(time.Second * 1)
	return op
	/*time.Sleep(time.Second*1)
	return op*/
}

func Sum(ops ...int) int {
	ret := 0
	for _, v := range ops {
		ret += v
	}
	return ret
}

func TestFn(t *testing.T) {
	aa, bb := returnMultiValues()
	t.Log(aa, bb)
	c, _ := returnMultiValues()
	t.Log(c)
	tsSF := timeSpent(slowFun)
	t.Log(tsSF(10))
}

func TestSum(t *testing.T) {
	t.Log(Sum(1, 3, 6, 8))
	t.Log(Sum(1, 3, 6, 8, 15))
}

func Clear() {
	fmt.Println("Clear Resources")
}

//defer特性 延迟执行
func TestDefer(t *testing.T) {
	defer func() {
		fmt.Println("Clear connections")
	}()
	defer Clear()
	fmt.Println("start")
	panic("error") //程序异常后释放资源链接

}
