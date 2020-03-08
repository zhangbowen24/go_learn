package series

import (
	"fmt"
)

//类似构造函数  init方法可以存在多次，按外部依赖顺序执行各个包的init方法，同一包内按顺序执行
func init() {
	fmt.Println("init2...")
}

func init() {
	fmt.Println("init...")
}

//小写方法名 包外无法访问
func square(n int) int {
	return n * n
}

func GetFibonacci(n int) ([]int, error) {
	fibList := []int{1, 1}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil
}
