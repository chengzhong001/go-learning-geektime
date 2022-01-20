package constant_test

import "testing"

const (
	Monday = iota + 1 // const iota = 0 // Untyped int.
	Tueday
	Wednesday
)
const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestConstantTry(t *testing.T) {
	t.Log(Monday, Tueday, Wednesday)
}

func TestConstantTry1(t *testing.T) {
	a := 7 // 0111
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
}
