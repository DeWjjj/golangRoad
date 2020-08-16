package main

import (
	"fmt"
)

func main() {
	number1 := 101
	fmt.Printf("%v\n", number1)
	fmt.Printf("%d\n", number1)
	fmt.Printf("%b\n", number1) //十进制 => 二进制
	fmt.Printf("%o\n", number1) //十进制 => 八进制
	fmt.Printf("%x\n", number1) //十进制 => 十六进制
}
