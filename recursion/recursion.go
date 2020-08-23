package main

import (
	"fmt"
)

//recursion function（递归）自己调用自己的函数。
//递归一定要有一个适合的退出的条件，问题的规模越来越小的场景。

//fatorial（阶乘）
func fatorial(x uint) uint {
	if x == 1 {
		return 1
	}
	return x * fatorial(x-1)
}
func main() {
	ret := fatorial(7) //5040
	fmt.Println(ret)
}
