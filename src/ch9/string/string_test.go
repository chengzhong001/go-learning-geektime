package string_test

import "testing"

func TestString(t *testing.T){
	var s string
	t.Log(s)
	s = "hello"
	t.Log(len(s))

	s = "\xE4\xb8\xa5"
	t.Log(s, len(s))

	s = "发展中国家"
	c := []rune(s)

	t.Logf("c utf-8 %x, and c's length is %d charater", c, len(c))

	for i:=0; i<len(c); i++{
		t.Log(c[i])
	}
}

func TestStringToRune(t *testing.T){
	s := "中华人民共和国"
	for idx, c := range s{
		t.Logf("%[1]c %[1]x %d", c, idx)
	}
}