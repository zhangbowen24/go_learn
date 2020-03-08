package constant_test

import (
	"testing"
	//"fmt"
)

const (
	Monday = iota + 1
	Tuesday
	Wednesday
)

const (
	Readable  = 1 << iota //0001
	Writable              //0010
	Excutable             //0100
)

func TestConstant(t *testing.T) {
	t.Log(Monday, Tuesday, Wednesday)
}

func TestConstant1(t *testing.T) {
	a := 1 //1=0001 7=0111
	t.Log(Readable, Writable, Excutable)
	t.Log(a&Readable, a&Writable, a&Excutable)
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Excutable == Excutable)
}
