package main

import (
	"fmt"
	"time"
)

func main() {

	var a int = 0

	for {
		a++
		fmt.Println(a)
		time.Sleep(3 * time.Second)
	}
	fmt.Println("vim-go")
}
