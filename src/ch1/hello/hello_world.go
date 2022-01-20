package main

import (
	"fmt"
	"os"
)

func main(){
	fmt.Println(os.Args)
	println(os.Args)
	if len(os.Args) > 1{
		fmt.Println("hello world",2)
	}
	os.Exit(1)
}