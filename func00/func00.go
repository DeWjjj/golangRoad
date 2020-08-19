package main

import (
	"fmt"
)

//返回值为两个
func double() (string, int) {
	return "白露", 767

}

//多个函数参数类型的简写
func more(x, y, z string, m, n int, k bool) {

}

//可边长参数
func expand(x string, y ...int) {
	fmt.Println(x)
	fmt.Println(y) //该类型为切片 []int
}
func main() {
	m, n := double()
	fmt.Println(m, n)
	expand("白露", 7, 6, 7)
}
