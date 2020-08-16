package main

import (
	"fmt"
)

var (
	age int
)

func main() {
	fmt.Scanln(&age) //fmt包中自带Scanln，类似于python中的Input函数。
	if age >= 18 {
		fmt.Println("Adult")
	} else if age > 60 {
		fmt.Println("Older")
	} else if age > 100 {
		fmt.Println("Goder")
	} else {
		fmt.Println("Children")
	}
}
