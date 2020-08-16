package main

import (
	"fmt"
)

func main() {
	ad := 22
	fmt.Printf("%T-%v\n", ad, ad)
	add := &ad //&:取出地址
	fmt.Printf("%T-%v\n", add, add)
	addd := *add //*:根据地址取值
	fmt.Printf("%T-%v\n", addd, addd)
}
