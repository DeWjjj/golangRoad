package main

import (
	"fmt"
)

func f6() (x int) {
	defer func(x *int) {
		(*x)++ //指针值++，也就是x++ >> 6
		fmt.Println("x的地址：", x)
		fmt.Printf("指针x值：%v", *x)
	}(&x)
	return 5 //x(ret) >> 5 >> func >> 6
}

func f4() (x int) {
	defer func(x int) { //函数内部的x 其实和外面的输出x并非一个数值，其所有的意义是不一样的。
		x++
		fmt.Println(x) //1
	}(x)
	return 5
}
func main() {
	fmt.Println(f4())
	//1
	//5
	fmt.Println(f6())
	//x的地址： 0xc0000b2020
	//指针x值：66
}
