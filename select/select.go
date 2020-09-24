package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 1) //创建一个容量为1的通道。
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Printf("通道中出来的值是%d，这是第%d次循环。\n", x, i+1)
		case ch <- i:
		}
		/*
			第一次通道中无法取出任意值，所以执行了case中存放值。
			通道中出来的值是0，这是第2次循环。
			通道中出来的值是2，这是第4次循环。
			通道中出来的值是4，这是第6次循环。
			通道中出来的值是6，这是第8次循环。
			通道中出来的值是8，这是第10次循环。
		*/
	}
}
