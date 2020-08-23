package main

import (
	"fmt"
)

func waysCount(x int) int {
	if x == 1 {
		return 1
	} else if x == 2 {
		return 2
	}
	return waysCount(x-1) + waysCount(x-2)
}

func main() {
	ret := waysCount(4)
	fmt.Println(ret)
}
