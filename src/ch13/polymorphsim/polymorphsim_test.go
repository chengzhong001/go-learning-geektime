package polymorphsim_test

import (
	"fmt"
	"testing"
)

type Code string
type Programmer interface {
	WriterHelloWorld() Code
}

type GoProgrammer struct {
}

func (p *GoProgrammer) WriterHelloWorld() Code {
	return "fmt.Println(\"hello world\")"
}

type JavaProgrammer struct {
}

func (p *JavaProgrammer) WriterHelloWorld() Code {
	return "System.out.Println(\"hello world\")"
}

func writeFirstProgram(p Programmer) {
	fmt.Printf("%T, %v\n", p, p.WriterHelloWorld())
}

func TestClient(t *testing.T) {
	var goProg Programmer = new(GoProgrammer)
	var javaProg Programmer = new(JavaProgrammer)
	writeFirstProgram(goProg)
	writeFirstProgram(javaProg)
	// t.Log(goProg.WriterHelloWorld())

}
