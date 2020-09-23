package main

import (
	"fmt"
)

func main() {
	n := 123
	var sum int
	for n > 0 {
		val := n % 10
		sum = sum + val
		n = n / 10
	}
	fmt.Println(sum)
}
