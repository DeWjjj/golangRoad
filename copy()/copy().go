package main

import (
	"fmt"
)

// copy
func main() {
	g1 := []int{1, 2, 3}
	g2 := g1
	var g3 = make([]int, 3)
	copy(g3, g1) //copy函数和普通切片赋值不一样，是属于将其地址内的数据赋值到另外一块内存地址上，所以数据会改变
	fmt.Println(g1, g2, g3)

	g1[0] = 0
	fmt.Println(g1, g2, g3)
}
