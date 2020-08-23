package main

import (
	"fmt"
)

//结构体
type perInformation struct {
	name   string
	age    int
	hobby  []string
	gender string
}

func main() {
	var dew perInformation
	dew.name = "DeW"
	dew.age = 22
	dew.hobby = []string{"computer", "music", "vedio games"}
	dew.gender = "male"
	fmt.Println(dew)
	fmt.Println(dew.name)
	fmt.Printf("%+v\n", dew)
	fmt.Printf("%#v\n", dew)
	//匿名结构体
	var jjj struct {
		name string
		age  int
	}
	jjj.name = "Jackie"
	jjj.age = 22
	fmt.Printf("type:%T value:%v\n", jjj, jjj)
}
