package main

import (
	"fmt"
)

//函数闭包，闭包等于啥？其实就是将一个为得到值的函数赋予一个变量，然后使用该变量函数并填充值得到结果的过程。
//简单的来说，调用函数，并且使用其的一个过程被称为闭包。
func plus() func(x, y int) int {
	return func(x, y int) int {
		return x + y
	}
}

func add10(x int) func() int {
	return func() int {
		x += 10
		return x
	}
}

func main() {
	plus := plus()
	sum := plus(1, 2)
	fmt.Println(sum)
	fmt.Println("")
	add101 := add10(1) //变量ad101得到函数add10(1)
	x := add101()
	fmt.Println(x) //变量x在函数的外面，所以每当运行一次内函数中的x都会去寻找外函数x的值。内函数\
	x = add101()   //在不断重复运行的过程中外函数x是一直在累加的。
	fmt.Println(x)
	x = add101()
	fmt.Println(x)
	x = add101()
	fmt.Println(x)
	fmt.Println("")

	add102 := add10(2) //变量ad102得到函数add10(2)这等于是换了一个新变量去继承函数，所以其内部的值反馈出来是不不同的。
	y := add102()
	fmt.Println(y) //变量x在函数的外面，所以每当运行一次内函数中的x都会去寻找外函数x的值。内函数\
	y = add102()   //在不断重复运行的过程中外函数x是一直在累加的。
	fmt.Println(y)
	y = add102()
	fmt.Println(y)
	y = add102()
	fmt.Println(y)
}
