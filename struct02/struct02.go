package main

import (
	"fmt"
)

type nameAge struct {
	name string
	age  int
}

//批量初始化
func main() {

	var p1 = &nameAge{ //创建的是nameAge的指针结构体
		name: "DeW",
		age:  22,
	}
	fmt.Printf("%#v\n", p1)

	p2 := nameAge{ //创建的是nameAge结构体
		"DeW",
		23,
	}
	fmt.Printf("%#v", p2)
	//结构体内部的字段存放地址是连续的。
}
