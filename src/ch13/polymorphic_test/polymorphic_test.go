package polymorphic_test

import (
	"fmt"
	"testing"
)

type Code string

type Programmer interface {
	WriteHelloWorld() Code
}

type GoProgrammer struct {
}

func (p *GoProgrammer) WriteHelloWorld() Code {
	return "fmt.Println(\"hello world\")"
}

type JavaProgrammer struct {
}

func (p *JavaProgrammer) WriteHelloWorld() Code {
	return "System.out.Println(\"hello world\")"
}

//%T 输出值的类型
func WriteFirstProgrammer(p Programmer) {
	fmt.Printf("%T %v\n", p, p.WriteHelloWorld())
}

func TestPolymorphic(t *testing.T) {
	GoProg := new(GoProgrammer)
	JavaProg := new(JavaProgrammer)
	WriteFirstProgrammer(GoProg)
	WriteFirstProgrammer(JavaProg)
}
