package main

import (
	"fmt"
)

type nameAge struct {
	name string
	age  int
}

//当结构体比较大的时候，使用结构体指针以来减少内存地址的开销。
func newnameAge(name *string, age *int) *nameAge {
	fmt.Printf("函数结构体内所用参数name的值，外面name的内存地址：%p\n", name) //函数内的name等同于函数外的&name（name的地址）所以是可以直接用%p的。
	fmt.Printf("函数结构体内所用参数name的内存地址：%p\n", &name)         //函数内的name等同于函数外的&name（name的地址）所以是可以直接用%p的
	fmt.Printf("函数结构体内所用参数age的值，外面age的内存地址：%p\n", age)
	fmt.Printf("函数结构体内所用参数age的内存地址：%p\n", &age) //函数内的name等同于函数外的&name（name的地址）所以是可以直接用%p的
	*name = "1"                                 //name其实拿到的是函数外name的地址，所以其所指的*name是外面name的值。
	return &nameAge{
		name: *name,
		age:  *age,
	}

}

func main() {
	name := "DeW"
	fmt.Printf("name值的地址是：%p\n", &name)
	age := 22
	fmt.Printf("age值的地址是：%p\n", &age)
	ret := newnameAge(&name, &age)
	fmt.Printf("返回结构体的值：%+v\n", *ret)
	fmt.Printf("返回结构体的内存地址：%p\n", &ret)
	fmt.Printf("返回结构体中name的内存地址：%p\n", &ret.name)
	fmt.Printf("返回结构体中age的内存地址：%p\n", &ret.age)
}
