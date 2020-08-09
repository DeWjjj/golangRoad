package main

import (
	"fmt"
)

var (
	age int
)

func main() {
	fmt.Scanln(&age)
	if age > 18 {
		fmt.Println("Adult")
	} else {
		fmt.Println("Teenager")
	}
}
