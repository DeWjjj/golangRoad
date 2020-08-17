package main

import (
	"fmt"
	"strconv"
)

func sum(x int, y int) (ret int) { //函数，ret是指return
	return x + y
}

func str(value int) (ret string) {
	str := strconv.Itoa(value)
	return str
}
func main() {
	sum := sum(1, 2)
	fmt.Println(sum)
	fmt.Printf("%T-%v\n", sum, sum)
	strSum := str(sum)
	fmt.Printf("%T-%v\n", strSum, strSum)
}
