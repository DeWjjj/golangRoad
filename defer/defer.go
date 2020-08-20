package main

import (
	"fmt"
)

func deferDemo() {
	fmt.Println("First!")
	defer fmt.Println("Next!") //defer可以将语句延迟到函数即将返回时再执行
	fmt.Println("End!")
}

func deferDemo1() {
	fmt.Println("First!")
	defer fmt.Println("Next!") //defer的本质是将其队列置于最后一列，多个连续defer也是按照顺序再次将其加到队列最后
	defer fmt.Println("Next1!")
	defer fmt.Println("Next2!")
	defer fmt.Println("Next3!")
	fmt.Println("End!")
}

func main() {
	deferDemo()
	fmt.Println("")
	deferDemo1()
}
