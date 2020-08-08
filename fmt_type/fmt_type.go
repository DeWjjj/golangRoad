package main

import (
	"fmt"
)

func main() {
	var n = 100
	fmt.Printf("%T\n", n)
	fmt.Printf("%v\n", n)  //输出值
	fmt.Printf("%#v\n", n) //输出结构名字值，再输出结构体
	fmt.Printf("%+v", n)   //先输出字段类型，再输出该字段的值
	fmt.Printf("%b\n", n)
	fmt.Printf("%d\n", n)
	fmt.Printf("%o\n", n)
	fmt.Printf("%x\n", n)
	var s = "hello"
	fmt.Printf("%s\n", s)
}

//GOLANG字符串内部实现UTF-8，GO语言中字符串是用双引号包裹。 ex: s := "hello"。
//GO语言中单引号包裹的是字符。 ex s := 'h' //单独的字母、汉字、符号。
//字节：1字节=8bit(8个二进制位)。
//1个UTF-8的汉字字符“哈”，一般占3个字节。
//1个英文字符'a'，一般占1个字节。
