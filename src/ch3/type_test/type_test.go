package typetest_test

import "testing"

type myInt int64
func TestImplicit(t *testing.T) {
	var a int = 1
	var b int64
	// b = a cannot use a (type int) as type int64 in assignment
	b = int64(a) 	// 显式类型转换
	var c myInt
	// var c int64
	c = myInt(b)
	t.Log(a, b, c)
}
