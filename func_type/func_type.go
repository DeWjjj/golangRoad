package main

import (
	"fmt"
)

func plx() {
	fmt.Println("I am DeW!")
}

func plx1() int {
	return 10
}

func plx2(x func() int) {
	ret := x()
	fmt.Println(ret)
}

func plus(x, y int) int {
	return x + y
}

func plx3(cal int) int {
	val := cal
	fmt.Println(val)
	return val
}

func plx4(f func() int) func(int, int) int {
	x := f //函数如果不带参数，则其返回的是一个函数类型，如果带参数就会变成该参数的return类型。
	fmt.Printf("%T", x)
	return plus
}

func main() {
	a := plx
	b := plx1
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	plx2(plx1)
	plx3(plus(1, 2))
	plx4(plx1)
}
