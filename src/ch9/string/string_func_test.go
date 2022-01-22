package string_test

import (
	"strconv"
	"strings"
	"testing"
)

func TestString_func(t *testing.T) {
	s := "A,B,C"
	parts := strings.Split(s, ",")
	for _, part := range parts {
		t.Log(part)
	}

	s = "中 华 人 民 共 和 国"
	parts = strings.Split(s, " ")
	for _, part := range parts {
		t.Log(part)
	}
	t.Log(strings.Join(parts, "-"))
}

func TestConv(t *testing.T) {
	str := strconv.Itoa(10)
	t.Log("str" + str)
	if i, err := strconv.Atoi("10"); err == nil {
		t.Log(10 + i)
	}

}
