package reflect_test

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

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


func TestDeepEqual(t *testing.T) {
	a := map[int]string{1: "one", 2: "two", 3: "three"}
	b := map[int]string{1: "one", 2: "two", 4: "four"}
	fmt.Println("a==b?", reflect.DeepEqual(a, b))
	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	s3 := []int{2, 3, 1}
	t.Log("s1==s2?", reflect.DeepEqual(s1, s2))
	t.Log("s1==s3?", reflect.DeepEqual(s1, s3))

	c1 := Customer{"1", "Mike", 40}
	c2 := Customer{"2", "Mike", 40}
	fmt.Println(c1 == c2)
	fmt.Println(reflect.DeepEqual(c1, c2))
}

func FillBySettings(st interface{}, settings map[string]interface{}) error {
	if reflect.TypeOf(st).Kind() != reflect.Ptr {
		if reflect.TypeOf(st).Elem().Kind() != reflect.Struct {
			return errors.New("the first param shoul be a pointer to the struct")
		}
	}
	var (
		filed reflect.StructField
		ok    bool
	)
	if settings == nil {
		return errors.New("setting is nil")
	}
	for k, v := range settings {
		if filed, ok = (reflect.ValueOf(st)).Elem().Type().FieldByName(k); !ok {
			continue
		}
		if filed.Type == reflect.TypeOf(v) {
			vstr := reflect.ValueOf(st)
			vstr = vstr.Elem()
			vstr.FieldByName(k).Set(reflect.ValueOf(v))
		}
	}

	return nil
}

func TestFillNameAdnAge(t *testing.T) {
	settings := map[string]interface{}{"Name": "Mike", "Age": 40}
	e := Employee{}
	if err := FillBySettings(&e, settings); err != nil {
		t.Fatal(err)
	}
	t.Log(e)
	c := new(Customer)
	if err := FillBySettings(c, settings); err != nil {
		t.Fatal(err)
	}
	t.Log(*c)

}
