package main

import (
	"fmt"
)

func main() {
	var x int
	x = 100
	fmt.Println(x)  //输出x的值
	fmt.Println(&x) //输出x的地址
	y := &x         //y得到x的地址
	fmt.Println(y)  //y的值x的地址
	fmt.Println(*y) //y的值x的地址所代表的值
	fmt.Println(&y) //y的地址
}
