package main

import (
	"fmt"
)

type user struct {
	name string
	age  int
}

func changeAge(x *user) {
	x.age = 22
	x.name = "dew"
}

func main() {
	var x = new(user) //new()创建的结构体是指针类型
	fmt.Printf("%T %v\n", x, x)
	changeAge(x)
	fmt.Println(x)
	x.age = 23
	fmt.Printf("%+v", x)
}
