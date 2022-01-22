package encap_test

import (
	"fmt"
	"testing"
	"unsafe"
)

type Employlee struct {
	Id   string
	Name string
	Age  int
}

func TestCreateEmployeeObj(t *testing.T) {
	e := Employlee{"0", "Bob", 20}
	e1 := Employlee{Name: "Mike", Age: 30}
	e2 := new(Employlee)
	t.Log(e)
	t.Log(e1)
	t.Log(e2)
	e2.Id = "2"
	e2.Age = 22
	e2.Name = "Rose"
	t.Log(e2)
	t.Logf("e is %T", e)
	t.Logf("e2 is %T", e2)
}

func (e Employlee) String() string {
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID: %s-Name: %s-Age:%d", e.Id, e.Name, e.Age)
}
func (e *Employlee) String1() string {
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID: %s-Name: %s-Age:%d", e.Id, e.Name, e.Age)
}

func String(e Employlee) string {
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID: %s-Name: %s-Age:%d", e.Id, e.Name, e.Age)
}

func TestStructOperation(t *testing.T) {
	e := Employlee{"0", "Bob", 20}
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
	t.Log(e.String())
	t.Log(e.String1()) // 没有复制产生
	t.Log(String(e))
}
