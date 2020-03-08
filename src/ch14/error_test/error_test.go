package error_test

import (
	"errors"
	"fmt"
	"strconv"
	"testing"
)

var LessThanTwoError = errors.New("n must be not less than 2")
var LargerThanHundredError = errors.New("n must be not larger than 100")

func GetFibonacci(n int) ([]int, error) {
	if n < 2 {
		//return nil, errors.New("n must be not less than 2")
		return nil, LessThanTwoError
	}
	if n > 100 {
		//return nil, errors.New("n must be not larger than 100")
		return nil, LargerThanHundredError
	}
	fibList := []int{1, 1}
	for i := 2; i < n; i++ {
		fmt.Println(i)
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil
}

func createError() error {
	return errors.New("you are wrong")
}

func GetFibonacci1(str string) {
	var (
		i    int
		err  error
		list []int
	)
	if i, err := strconv.Atoi(str); err == nil {
		if list, err := GetFibonacci(i); err == nil {
			fmt.Println(list)
		} else {
			fmt.Println("error", err)
		}
	} else {
		fmt.Println("error", err)
	}
}

func GetFibonacci2(str string) {
	var (
		i    int
		err  error
		list []int
	)
	if i, err := strconv.Atoi(str); err != nil {
		fmt.Println("error", err)
		return
	}
	if list, err := GetFibonacci(i); err != nil {
		fmt.Println("error", err)
		return
	}
	fmt.Println(list)
}

func TestError(t *testing.T) {
	if v, err := GetFibonacci(110); err != nil {
		if err == LessThanTwoError {
			t.Log("It is less")
		}
		if err == LargerThanHundredError {
			t.Log("It is larger")
		}
		//if v, err := GetFibonacci(10); err != nil {
		t.Error(err)
	} else {
		t.Log(v)
	}
	//fibList := GetFibonacci(-10)
	//t.Log(fibList)
	/*err := createError()
	t.Error(err)*/
}
