package series

import (
	"errors"
	"fmt"
)

func Add(x, y int) int {
	return x + y
}

// func square(n int) int {
// 	return n * n
// }

func init(){
	fmt.Println("init1")
}

func init(){
	fmt.Println("init2")
}

func GetFibonacci(n int) ([]int, error) {
	if n < 0 || n > 100 {
		return nil, errors.New("p should be in [2,100]")
	}
	fibList := []int{1, 1}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil
}
