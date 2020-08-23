package main

import (
	"fmt"
)

type nameAge struct {
	name string
	age  int
}

func changeAge(x nameAge) {
	x.age = 23
}

func changeAge1(x *nameAge) { //得到传参的值
	(*x).age = 24 //传参的值更改
	//x.age = 24 //go语言中不允许操作指针，所以传入的值类型为指针的时候，可以隐去*让go自动判断。
}

func main() {
	var dew nameAge
	dew.name = "DeW"
	dew.age = 22
	changeAge(dew) //go语言自建函数中的值传入，其实等同将原先的值赋值给函数内的定义参数。
	//修改其内部值无效。
	fmt.Println(dew) //{DeW 22}
	dew.age = 23     //有效操作
	fmt.Println("")
	fmt.Println(dew) //{DeW 23}
	fmt.Println("")
	changeAge1(&dew) //传入dew的内存地址，有效操作
	fmt.Println(dew)
}
