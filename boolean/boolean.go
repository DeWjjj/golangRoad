package main

import "fmt"

func main() {
	a := true
	b := false
	fmt.Printf("%T\n", a)
	fmt.Printf("type:%T value:%v\n ", b, b)
}

//tips:布尔值在GO语言中不可以转换成数值类型
