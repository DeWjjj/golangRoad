package main

import (
	"fmt"
)

type age struct {
	age int
}

func (a age) showAge() {
	fmt.Println("年龄：", a.age)
}

type information struct {
	name   string
	gender string
	age
}

func (i information) showInformation() {
	fmt.Printf("%v is a %v.", i.name, i.gender)
}

func main() {
	dew := information{
		name:   "dew",
		gender: "male",
		age: age{
			age: 22,
		},
	}
	fmt.Println(dew.name, dew.gender)
	dew.showAge()
	dew.showInformation()
	//age结构体使用的函数showAge，就可以直接被包括其结构体的information调用。
	//从而实现了对其功能的继承。
}
