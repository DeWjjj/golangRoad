package main

import (
	"fmt"
)

type nameAge struct {
	name string
	age  int
}

func newNameAge(name string, age int) nameAge {
	return nameAge{
		name: "dew",
		age:  22,
	}
}

func (perInfor nameAge) newName0() { //perInfor为函数内重新定义的参数，其所代表的值和传入的值是不一样的。
	perInfor.name = "j"
}

func (perInfor *nameAge) newName() { //定义类型为传入指针类型的时候，修改其数据则是修改其源数据。
	perInfor.name = "jjj"

}

func main() {
	name := "dew"
	age := 22
	perInfor := newNameAge(name, age)
	fmt.Println(perInfor)
	perInfor.newName()
	fmt.Println(perInfor)
	perInfor.newName0() //其传入值并非指针等于其函数内的参数赋值。
	fmt.Println(perInfor)
}
