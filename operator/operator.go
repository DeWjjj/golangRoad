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
}
