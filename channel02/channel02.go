package main

import (
	"fmt"
)

var ch chan int

func main() {

	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	x1 := <-ch
	x2 := <-ch
	fmt.Println(x1, x2)
}
