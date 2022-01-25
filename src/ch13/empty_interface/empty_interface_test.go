package empty_interface

import (
	"fmt"
	"testing"
)

func DoSomeThing(p interface{}) {
	if i, ok := p.(int); ok {
		fmt.Println("Integer", i)
		return
	}
	if s, ok := p.(string); ok {
		fmt.Println("string", s)
		return
	}
	fmt.Println("Unknow Type")
}

func DoSomeThing1(p interface{}){
	switch v:=p.(type){
	case int:
		fmt.Println("Integer", v)
	case string:
		fmt.Println("String", v)
	default:
		fmt.Println("Unknow Type")
	}

	
}


func TestEmptyInterfaceAssertion(t *testing.T){
	DoSomeThing(10)
	DoSomeThing("10")
	DoSomeThing("hello")

	DoSomeThing1(10)
	DoSomeThing1("10")
	DoSomeThing1("hello")

}