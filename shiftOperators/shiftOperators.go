package main

import (
	"fmt"
)

const (
	a = 1 //00000001
)

const b = a << 1 //00000001 => 00000010 = 2
const c = a << 2 //00000001 => 00000100 = 4

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
