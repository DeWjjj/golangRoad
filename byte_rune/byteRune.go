package main

import (
	"fmt"
)

func main() {
	s := "Hello,DeW!白露！백로!" //英语字符byte(ASCII,uint8)，其余UTF-8为rune(int32)。
	lenS := len(s)
	fmt.Println(lenS) //10

	for _, c := range s {
		fmt.Printf("%c\n", c)
	}

	str1 := "白露"
	str2 := []rune(str1) //str1字符串 => str2切片
	str2[0] = '黑'
	str3 := string(str2) //str2切片 => str字符串
	fmt.Println(str3)

	//"" 和 ''的区别
	a1 := "白" //string
	a2 := '黑' //rune
	fmt.Printf("%T %T", a1, a2)
}
