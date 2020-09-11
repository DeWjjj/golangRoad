package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 1000; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
	// fmt.Println("main")
	time.Sleep(time.Second)
}
