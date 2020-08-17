package main

import (
	"fmt"
)

func main() {
	//slice + map
	var s1 = make([]map[string]string, 10)
	s1[0] = make(map[string]string, 1)
	s1[0]["010"] = "beijing"
	fmt.Println(s1)

	//map + slice
	var m1 = make(map[string][]string, 10)
	m1["beijing"] = []string{"010", "010", "010"}
	fmt.Println(m1)
	//go语言中010会被默认识别成八进制数字，所以我们需要使用string类型
}
