package testing

import (
	//"ch32/unit_test_pool"
	"fmt"
	"github.com/stretchr/testify/assert" //断言
	"testing"
)

func TestSquare(t *testing.T) {
	//t.Log(unit_test_pool.Square(3))
	inputs := [...]int{1, 2, 3}
	expected := [...]int{1, 4, 9}
	for i := 0; i < len(inputs); i++ {
		//ret := unit_test_pool.Square(inputs[i])
		ret := Square(inputs[i])
		if ret != expected[i] {
			t.Errorf("input is %d, the expect is %d, the actual %d", inputs[i], expected[i], ret)
		}
	}
}

//Fail Error 该测试失败，该测试继续， 其他测试继续执行
//FailNow Fatal该测试失败，该测试中止， 其他测试继续执行

func TestErrorInCode(t *testing.T) {
	fmt.Println("start")
	t.Error("Error")
	fmt.Println("end")
}

func TestFailInCode(t *testing.T) {
	fmt.Println("start")
	t.Fatal("Error")
	fmt.Println("end")
}

func TestSquareWithAssert(t *testing.T) {
	//t.Log(unit_test_pool.Square(3))
	inputs := [...]int{1, 2, 3}
	expected := [...]int{1, 4, 9}
	for i := 0; i < len(inputs); i++ {
		//ret := unit_test_pool.Square(inputs[i])
		ret := Square(inputs[i])
		//ret = 99
		assert.Equal(t, expected[i], ret)
	}
}
