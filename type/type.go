package main

import (
	"fmt"
)

type myInt int
type myInt1 = int

func main() {
	var x myInt
	var y myInt1
	x = 100
	y = 100
	fmt.Println(x)        //100
	fmt.Printf("%T\n", x) //main.myInt
	fmt.Println("")
	fmt.Println(y)        //100
	fmt.Printf("%T\n", y) //int
	fmt.Println("")
	var z rune //rune的底层是int32
	z = '超'
	fmt.Println(z)        //36229
	fmt.Printf("%T\n", z) //int32
}
