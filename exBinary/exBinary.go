package main

import (
	"fmt"
)

const (
	eat   int = 4 //100
	sleep int = 2 //010
	walk  int = 1 //001
)

func ex1(arg int) {
	fmt.Printf("%b\n", arg)

}

func main() {
	ex1(eat | walk)         //101
	ex1(eat | walk | sleep) //111

}
