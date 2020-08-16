package main

import (
	"fmt"
)

func main() {
	//for基本格式
	for i := 0; i < 3; i++ { //初始语句 条件表达式 结束语句
		fmt.Print(i)
	}
	fmt.Println("")

	//省略初始语句
	c := 5
	for ; c < 10; c++ {
		fmt.Print(c)
	}
	fmt.Println("")
	//将结束语句写入执行内

	b := 3
	for b < 5 {
		fmt.Print(b)
		b++
	}
}
