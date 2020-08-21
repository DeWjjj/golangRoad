package main

import (
	"fmt"
)

//闭包类型
func plus() func(x, y int) int {
	return func(x, y int) int {
		return x + y
	}
}

func main() {
	add := plus()
	sum := add(1, 2)
	fmt.Println(sum)
}
