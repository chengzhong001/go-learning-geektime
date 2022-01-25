package panic_test

import (
	"errors"
	"fmt"
	"os"
	"testing"
)

func TestPanicVxExit(t *testing.T) {
	defer func() {
		// recover 捕获panic抛出的错误
		if err := recover(); err != nil {
			fmt.Println("recover from",err)
		}
	}()
	fmt.Println("start")
	// os.Exit(-1)
	panic(errors.New("Something wrong"))
}
func TestOsExit(t *testing.T) {
	defer func() {
		fmt.Println("finally")
	}()
	os.Exit(-1)
}
