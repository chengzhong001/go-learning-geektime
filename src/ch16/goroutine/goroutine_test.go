package goroutine_test

import (
	"fmt"
	"testing"
	"time"
)

func Add(i int) {
	fmt.Println(i + i)
}

func TestGroutine(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
		// go Add(i)

	}
	time.Sleep(time.Millisecond * 50)
}
