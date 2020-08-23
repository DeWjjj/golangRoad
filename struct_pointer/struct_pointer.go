package main

import (
	"fmt"
)

type nameAge struct {
	name string
	age  int
}

func main() {
	var x = new(nameAge)
	fmt.Println(x)         //函数赋予其的值
	fmt.Printf("%p\n", x)  //x的值是函数赋予其的内存地址
	fmt.Println(*x)        //通过指针x的值（内存地址）来查找到其值
	fmt.Printf("%p\n", &x) //x这个变量自己本身的地址
	//由此我们可以得到的结论是new函数其实是在一个空的区域开辟符合type nameAge的空间，并将其的地址赋值给x。
	(*x).name = "DeW"
	x.age = 22
	fmt.Println(x.age) //x.age的值
	fmt.Println(x.name)
	fmt.Printf("%p\n", &x.age)  //输出x.age值的地址
	fmt.Printf("%v\n", x.age)   //输出x.age值
	fmt.Printf("%p\n", &x.name) //输出x.name值的地址
	fmt.Printf("%v\n", x.name)  //输出x.name值
	//一个结构体内的字段其存放的内存地址是连续的。
}
