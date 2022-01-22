package string_test

import "testing"

func TestString(t *testing.T){
	var s string
	t.Log(s)
	s = "hello"
	t.Log(len(s))
	s = "中"
	t.Log(s, len(s))
	s = "\xE4\xb8\xa5"
	t.Log(s, len(s))


}