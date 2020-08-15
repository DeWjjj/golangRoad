package main

import (
	"fmt"
)

func main() {
	a1 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	n1 := a1[:]
	fmt.Println(n1)
	//删除数字3
	n1 = append(a1[:2], a1[3:]...)
	fmt.Println(n1)
	fmt.Println(a1)
}
