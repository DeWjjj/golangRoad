package main

import (
	"fmt"
)

type information struct {
	number        int
	nameAgeGender nameAgeGender
	address
}

type address struct {
	province string
}

type nameAgeGender struct {
	name   string
	age    int
	gender string
}

func main() {
	dew := information{
		number: 767,
		nameAgeGender: nameAgeGender{
			name:   "dew",
			age:    22,
			gender: "男",
		},
		address: address{
			province: "Zhejiang",
		},
	}
	fmt.Println(dew)
	fmt.Println(dew.number)
	fmt.Println(dew.nameAgeGender.name, dew.nameAgeGender.age, dew.nameAgeGender.gender) //非匿名结构体需要准确层级。
	fmt.Println(dew.address)                                                             //匿名结构体可以直接查询
	//匿名结构体会存在冲突，所以避免冲突字段尽量防止相同。

}
