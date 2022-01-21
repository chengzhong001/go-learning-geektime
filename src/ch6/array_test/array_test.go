package array_test

import "testing"

func TestArrayInit(t *testing.T) {
	var arr [3]int //声明一个长度为3默认值为0的数组
	arr1 := [4]int{1, 2, 3, 4}
	arr2 := [...]int{1, 2, 3, 4, 5}
	t.Log(arr)
	arr[1] = 10
	t.Log(arr)
	t.Log(arr1)
	t.Log(arr2)
}

func TestArrayTravel(t *testing.T) {
	arr3 := [...]int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr3); i++ {
		t.Log(arr3[i])
	}
	for idx, e := range arr3 {
		t.Log(idx, e)
	}
}

func TestArraySection(t *testing.T) {
	arr3 := [...]int{1,2,3,4,5}
	arr3_sec := arr3[3:]
	t.Log(arr3_sec)
}
