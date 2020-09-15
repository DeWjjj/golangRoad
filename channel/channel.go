package main

import (
	"fmt"
	"sync"
)

var a chan int //需要制定通道中的元素类型
var b chan int

var wg sync.WaitGroup

func noBufChannel() {
	fmt.Println(a)     // <nil>
	a = make(chan int) //不带缓冲期的通道初始化
	wg.Add(1)
	go func() {
		defer wg.Done()
		x := <-a
		fmt.Println("从通道中取到值是", x)
	}()
	a <- 10
	fmt.Println("发送到通道a的数值是10")
	wg.Wait()
}

func bufChannel() {
	fmt.Println(b)
	b = make(chan int, 1)
	b <- 10
	x := <-b
	fmt.Println("从通道中取到值是", x)
	close(b) //主动关闭通道，不过go语言有垃圾回收机制用完程序结束也会回收
}
func main() {
	bufChannel()
}
