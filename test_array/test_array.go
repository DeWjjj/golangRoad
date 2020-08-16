package main

import (
	"fmt"
)

func main() {
	//求数组[1,3,5,7,8]所有元素的和
	var (
		n1  = [...]int{1, 3, 5, 7, 8}
		sum int //0
	)
	for _, v := range n1 {
		sum = sum + v //24
	}
	fmt.Println(sum)
	for x := 0; x < len(n1); x++ {
		for y := x + 1; y < len(n1); y++ {
			if n1[x]+n1[y] == 8 {
				fmt.Println(x, y)
			}
		}
	}
}
