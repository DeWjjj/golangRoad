package main

import (
	"fmt"
)

func main() {
	var (
		a = 5
		b = 3
		c = 2
	)
	fmt.Println(a + b)       //8
	fmt.Println(a - c)       //3
	fmt.Println(a * b)       //15
	fmt.Println(a / b)       //1
	fmt.Println(a % b)       //2
	fmt.Println(a + (b - c)) //6
	a++                      //a = a + 1
	b--                      //b = b - 1
	fmt.Println(a, b)        //6 2
	fmt.Println(a >= b)      //true
	fmt.Println(a <= b)      //false

	age := 22
	if age >= 18 && age <= 60 { //&& = and = 并且
		fmt.Println("生活生活，就是干活！")
	} else if age < 18 || age > 60 { //｜｜ = or = 其一
		fmt.Println("生活美好，日子真爽！")
	} else {
		fmt.Println("异次元")
	}

	var isBoy = true
	fmt.Println(isBoy)  // true
	fmt.Println(!isBoy) // false

	//位运算符，针对二进制数
	//5的二进制是：101
	//2的二进制是：010

	// &:两位均为1才为1，与
	fmt.Println(5 & 2) //0 00000000
	// ｜:两位其中有一个为1就为1，或
	fmt.Println(5 | 2) //7 00000111
	// ^:异或，两位不一样则为1
	fmt.Println(5 ^ 2) //7 00000111
	// <<:二进制移位
	fmt.Println(5 << 1) // 00000101 =>00001010 = 10
	fmt.Println(5 >> 1) // 00000101 =>00000010 = 2
	fmt.Println(5 >> 3) // 00000101 =>00000000 = 0

	var (
		x int
	)
	x = 10
	x += 2 //10 + 2 = 12
	fmt.Println(x)
	x <<= 2 //00001010 => 001010000 =48
	fmt.Println(x)
}
