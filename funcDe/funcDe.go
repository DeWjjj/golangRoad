package main

import (
	"fmt"
)

func plus(x, sumY int, y ...int) int { //可增常参数y，实际上是一个数组
	for _, val := range y {
		sumY = sumY + val
	}
	sum := x + sumY
	return sum
}

func plus0(x int, y int) (sum, sub int) {
	sum = x + y
	sub = x - y
	return
}

//函数实际上，都可以写在main里面，实际上是我们把函数单独拿出来封装
func main() {
	value := plus(1, 2, 3, 4, 5)
	value0, value1 := plus0(1, 2)
	fmt.Println(value)
	fmt.Println(value0, value1)
}
