package main

import (
	"fmt"
)

//匿名函数是函数内部的函数
//函数内部没有办法声明带名字的函数
func main() {
	f1 := func() int {
		return 767
	}
	fmt.Println(f1())
	func(x, y int) {
		fmt.Println(x + y)
	}(1, 766) //匿名函数的写法如果其中需要传值的话，则是在其最后放数值
	//func(val type) (returnVal type) {
	//<语句块>
	//}(val)
}
