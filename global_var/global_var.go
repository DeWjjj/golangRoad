package main

import (
	"fmt"
)

var x = 100 //定义一个全局变量

//定义一个函数(输出x)
func plx() {
	//现在函数内部查找
	//然后在查询外部，一直查到到全局
	x := 10
	fmt.Println(x)
}

func main() {
	plx()
	fmt.Println(x) //函数默认向外查找，不会查询函数内部的变量
}
