package empty_interface_test

import (
	"fmt"
	"testing"
)

func DoSomeThink(p interface{}) {
	/*if i, ok := p.(int); ok {
		fmt.Println("Integer", i)
		return
	}
	if s, ok := p.(string); ok {
		fmt.Println("String", s)
		return
	}
	fmt.Println("UnKnow Type")*/
	switch v := p.(type) {
	case int:
		fmt.Println("Integer", v)
	case string:
		fmt.Println("String", v)
	default:
		fmt.Println("UnKnow Type")
	}
}

func TestEmptyInterfaceAssertion(t *testing.T) {
	DoSomeThink(10)
	DoSomeThink("3")
}
