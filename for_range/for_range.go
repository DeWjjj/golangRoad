package main

import (
	"fmt"
)

func main() {
	str1 := "abc"
	str2 := []rune(str1)
	for key, value := range str2 { //key位置为下表 value位置则是循环的数值。
		//for range可以遍历数组、切片、字符串、map和通道(channel)后两者目前我也不知道，等待学习获得。
		fmt.Print(key)
		fmt.Println(string(value))
	}
}
