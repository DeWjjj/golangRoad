package main

import (
	"fmt"
)

func main() {
	//遇到5直接跳出循环
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}
	fmt.Println("Over!")

	//跳过5该次循环
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}
	fmt.Println("Over!")

	//tips:fallthrough是向前兼容的一种旧特性用于判断switch中1语句要用到2语句也要用到的时候。

	switch n := 7; n {
	case 0:
		fmt.Println("0")
	case 1, 5:
		fmt.Println("1,5")
		fallthrough
	case 10:
		fmt.Println("10")
	}
}
