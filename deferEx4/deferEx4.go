package main

import (
	"fmt"
)

func f4() (x int) {
	defer func(x int) { //函数内部的x 其实和外面的输出x并非一个数值，其所有的意义是不一样的。
		x++
		fmt.Println(x)
	}(x)
	return 5
}
func main() {
	fmt.Println(f4())

}
