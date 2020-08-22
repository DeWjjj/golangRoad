package main

import (
	"fmt"
)

//闭包案例，我们可以其实可以把要传入的数值单独拆出来，假如从网页上拿到的信息如果是一长条的。
//其实我们可以先将其取出，然后再对应，但是通过闭包我们可以直接把网页中的信息直接通过方法直接一次性输入进去。
func plus(x, y int) int {
	return x + y
}

func reduce(x, y int) int {
	return x - y
}

func times(x, y int) int {
	return x * y
}

func divide(x, y int) int {
	return x / y
}

func cal(oper func(int, int) int, x, y int) int {
	res := oper(x, y)
	return res
}

func main() {
	plus := cal(plus, 1, 2)
	fmt.Println(plus)
	times := cal(times, 1, 2)
	fmt.Println(times)
}
