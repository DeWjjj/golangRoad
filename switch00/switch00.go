package main

import (
	"fmt"
)

func main() {
	//IF
	var n = 5
	if n == 1 {
		fmt.Println("大拇指")
	} else if n == 2 {
		fmt.Println("食指")
	} else if n == 3 {
		fmt.Println("中指")
	} else if n == 4 {
		fmt.Println("无名指")
	} else if n == 5 {
		fmt.Println("小拇指")
	} else {
		fmt.Println("无效的数字")
	}
	//SWITCH
	switch a := 5; a {
	case 1, 3, 5:
		fmt.Println("奇数")
	case 2, 4:
		fmt.Println("偶数")
	}
}
