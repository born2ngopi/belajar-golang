package main

import "fmt"
import "time"

func main() {

	now := time.Now()

	fmt.Printf("%v%02d%02d%02d%02d%02d%v\n", now.Year(), int(now.Month()), now.Day(), now.Hour(), now.Minute(), now.Second(), now.Nanosecond())

	fmt.Println("vim-go")
}
