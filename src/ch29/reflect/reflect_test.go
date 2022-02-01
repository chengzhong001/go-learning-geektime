package reflect_test

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTypeAndValue(t *testing.T) {
	var f int64 = 10
	t.Log(reflect.TypeOf(f), reflect.ValueOf(f))
	t.Log(reflect.ValueOf(f).Type())
}

func CheckType(v interface{}) {
	t := reflect.TypeOf(v)
	switch t.Kind() {
	case reflect.Float32, reflect.Float64:
		fmt.Println("Float")
	case reflect.Int, reflect.Int32, reflect.Int64:
		fmt.Println("Int")
	default:
		fmt.Println("Unknow", t)
	}
}

func TestBasicType(t *testing.T) {
	var f float32 = 12
	CheckType(&f)
}

type Person struct {
	Name string
	Age  float32
}

type Employee struct {
	EmployeeID string
	Name       string `format:"normal"`
	Age        int
}

func (e *Employee) UpdateAge(newVal int) {
	e.Age = newVal
}

type Customer struct {
	CookieID string
	Name     string
	Age      int
}

func TestPersonStruct(t *testing.T) {
	person := Person{"bob", 18}
	t.Log(reflect.ValueOf(person).FieldByName("Name"))
	t.Log(reflect.ValueOf(person).FieldByName("Age"))
}
func TestEmployeeStruct(t *testing.T) {
	e := &Employee{"1", "Mike", 30}
	t.Logf("Name: Value(%[1]v), Type%[1]T", reflect.ValueOf(*e).FieldByName("Name"))
	if nameFiled, ok := reflect.TypeOf(*e).FieldByName("Name"); !ok {
		t.Error("Failed to get 'name' filed")
	} else {
		t.Log("Tag:format", nameFiled.Tag.Get("format"))
	}
	reflect.ValueOf(e).MethodByName("UpdateAge").Call([]reflect.Value{reflect.ValueOf(1)})
	t.Log("Updated Age:", e)
}
