package interface_test

import "testing"

type Programmer interface{
	WriterHelloWorld() string
}

type GoProgrammer struct{
}

func (p *GoProgrammer) WriterHelloWorld() string {
	return "fmt.Println(\"hello world\")"
}

func TestClient(t *testing.T){
	var p Programmer
	p = new(GoProgrammer)
	t.Log(p.WriterHelloWorld())

}