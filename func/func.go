package main

import (
	"fmt"
	"strconv"
)

func sum(x int, y int) (ret int) { //函数，ret是指return
	ret = x + y
	return
}

func str(value int) string {
	str := strconv.Itoa(value)
	return str
}
func plus(x int, y int) {
	fmt.Println(x + y) //3
}

func line() { //没有参数和返回值，执行一次函数
	fmt.Println("")
}

func be3() int {
	return 3
}
func main() {
	sum := sum(1, 2)
	fmt.Println(sum)
	fmt.Printf("%T-%v", sum, sum) //int-3
	line()
	strSum := str(sum)
	fmt.Printf("%T-%v", strSum, strSum) //string-3
	line()
	numberChanged := be3()
	fmt.Println(numberChanged) //3
	plus(1, 2)                 //3
}
