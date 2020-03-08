package operator_test

import (
	"testing"
)

const (
	Readable = 1 << iota //0001 1
	Writable             //0010 2
	//0011 3
	Excutable //0100 4
	//0101 5
	//0110 6
)

func TestOperator(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	//b := [...]int{1, 2, 3, 4, 5}
	c := [...]int{1, 2, 3, 5}
	d := [...]int{1, 2, 3, 4}
	//t.Log(a == b)
	t.Log(a == c)
	t.Log(a == d)
}

func TestBitClear(t *testing.T) {
	a := 7 //1=0001 7=0111
	//a = a &^ Readable
	//a = a &^ Writable
	a = a &^ Excutable //按位清零   a=0111 Excutable=0100 把a的正序第二位清零得到 0011
	t.Log(a, Readable, Writable, Excutable)
	t.Log(a, a&Readable, a&Writable, a&Excutable)
	t.Log(a, a&Readable == Readable, a&Writable == Writable, a&Excutable == Excutable)
}
