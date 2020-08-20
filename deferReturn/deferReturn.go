package main

import (
	"fmt"
)

//go语言的func的Return语句并不是原子操作，而是分为给RET赋值，RET返回的操作。
//defer语句执行的时间，正好是在RET赋值之后，RET返回之前。
func f1() int {
	x := 5
	func() {
		x++
	}()
	return x // >> RET = x = 5 >> x++ >> x=6
}

func f2() (x int) {
	func() {
		x++
	}()
	return 5 // >> RET(x) = 5 >> x++ >> RET(x)++ = 6
}

func main() {
	fmt.Println(f1())
	fmt.Println(f2())
}
