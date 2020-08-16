package main

import (
	"fmt"
)

//goto+label
func main() {
	for i := 0; i < 10; i++ {
		for j := 'A'; j < 'Z'; j++ {
			if j == 'C' {
				goto breakLabel
			}
			fmt.Printf("%v-%c\n", i, j)
		}
	}
breakLabel:
	fmt.Println("Over")
}
