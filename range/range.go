package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	a := make(chan int, 2)
	a <- 1
	a <- 2
	wg.Add(2)
	go func() {
		for {
			x := <-a
			fmt.Println(x)
			wg.Done()
		}
	}()
	wg.Wait()

}
