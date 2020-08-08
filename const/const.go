package main

import (
	"fmt"
)

const pi = 3.1415926

const (
	notFound = 404
	statusOk = 200
)

const (
	n1 = 100
	n2 //n2会直接沿用n1的数值
	n3
)

func main() {
	fmt.Println("n1:", n1)
	fmt.Println("n2:", n2)
	fmt.Println("n3:", n3)
}
