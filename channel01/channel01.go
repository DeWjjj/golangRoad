package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var once sync.Once

func f1(ch1 chan int) { //给通道中填入0～100
	defer wg.Done()
	for i := 0; i < 100; i++ {
		ch1 <- i
	}
	close(ch1)

}

func f2(ch1, ch2 chan int) {
	defer wg.Done()
	for {
		x, ok := <-ch1
		if !ok {
			break
		}
		ch2 <- x * x
	}
	// once.Do(func() { close(ch2) }) //确保通道只关闭一次
}

func main() {
	a := make(chan int, 100)
	b := make(chan int, 100)
	wg.Add(3)
	go f1(a)
	go f2(a, b)
	go f2(a, b)
	for ret := range b {
		fmt.Println(ret)
	}
	wg.Wait()
}
