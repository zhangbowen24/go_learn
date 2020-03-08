package main

import (
	"fmt"
	"testing"
)

func TestFibList(t *testing.T) {
	/*var x int = 1
	var y int = 1*/
	/*var (
		x int = 1
		y int = 1
	)*/
	x := 1 //赋值默认类型
	y := 1
	//fmt.Print(x, " ")
	t.Log(x)
	for i := 0; i < 5; i++ {
		//fmt.Print(y, " ")
		t.Log(y)
		tmp := x
		x = y
		y = tmp + x
		//fmt.Println(tmp)
	}
}

var (
	a   int
	b   int
	tmp int
)

func TestExchange(t *testing.T) {
	a := 1
	b := 2
	//旧的变量转换
	/*tmp := a
	a = b
	b = tmp*/
	a, b = b, a //变量转换新语法
	fmt.Println(a, b)
}
