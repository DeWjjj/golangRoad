package main

import (
	"fmt"
)

func add(in chan<- int) { //此时函数传入的通道类型是单向写入通道
	for i := 0; i < 100; i++ {
		in <- i
	}
	close(in)
}

func time(out chan<- int, in <-chan int) { //此时第一个传入的通道为单向写入，第二个为单向读取通道
	for i := range in {
		out <- i * i

	}
	close(out)

}

func printer(in <-chan int) { //函数读取单向读取通道的值
	for i := range in {
		fmt.Println(i)
	}

}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go add(ch1)       //将通道ch1传入，给通道添加100个数值的函数
	go time(ch2, ch1) //将通道ch1中的值，平方之后传入ch2通道
	printer(ch2)      //打印ch2通道中的值
}
