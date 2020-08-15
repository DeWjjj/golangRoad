package main

import (
	"fmt"
)

func main() {
	var m1 map[string]int
	m1 = make(map[string]int, 10) //make初始化map的时候要估算好该map容量，避免程序运行中程序检测后程序自动扩容
	fmt.Println(m1 == nil)        //还没有在内存中占有空间，现在是空的
	m1["dew"] = 18                //DeW永远18岁
	m1["wangchao"] = 22

	fmt.Println(m1)
	fmt.Println(m1["dew"]) //18
	fmt.Println(m1["abc"]) //0

	value, bool := m1["jackie"]
	if !bool {
		fmt.Println("查无此人")
	} else {
		fmt.Println(value)
	}

	//map的遍历
	for val, key := range m1 {
		fmt.Println(val, key)
	}
	//只遍历key
	for key := range m1 {
		fmt.Println(key)
	}
	//只遍历值
	for _, val := range m1 {
		fmt.Println(val)
	}

}
