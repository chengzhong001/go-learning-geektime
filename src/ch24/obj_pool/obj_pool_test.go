package objpool

import (
	"fmt"
	"testing"
	"time"
)

func TestOjbPool(t *testing.T) {
	pool := NewOjbPool(10)
	// if err:= pool.RelaseObj(&ReusableObj{}); err!=nil{
	// 	t.Error(err)
	// }
	for i := 0; i < 11; i++ {
		if v, err := pool.GetObj(time.Second * 1); err != nil {
			t.Error(err)
		} else {
			fmt.Printf("%T\n", v)
			if err := pool.RelaseObj(v); err != nil {
				t.Error(err)
			}
		}
	}
	fmt.Println("Done")
}
