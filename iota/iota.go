package main

import (
	"fmt"
)

//简易是用iota

// const (
// 	a1 = iota //iota类似枚举
// 	a2 = iota
// 	a3 = iota
// )

//使用匿名变量去越过iota

// const (
// 	b1 = iota //0
// 	b2        //1
// 	_         // _ = iota = 2
// 	b3        //3
// )

//iota插队

const (
	c1 = iota //0
	c2 = 100  //100
	c3 = iota //2
	c4
)

const (
	d1, d2, d3 = iota, iota + 1, iota + 2     // 0 1 2
	d4, d5, d6 = iota + 3, iota + 4, iota + 5 // d4:1 + 3 = 4 d5:1 + 4 = 5 d6:1 + 5 =6
)

func main() {
	// fmt.Println(a1) //0
	// fmt.Println(a2) //1
	// fmt.Println(a3) //2

	// fmt.Println(b1)
	// fmt.Println(b2)
	// //<_>为匿名变量，意味为占位符
	// fmt.Println(b3)

	fmt.Println(d1)
	fmt.Println(d2)
	fmt.Println(d3)
	fmt.Println(d4)
	fmt.Println(d5)
}
