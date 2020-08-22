package main

import (
	"fmt"
	"strings"
)

func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) { //HasSuffix的意思是指，查询后缀。
			return name + suffix
		}
		return name
	}
}

func main() {
	jpg := makeSuffix(".jpg")
	// text := makeSuffix(".txt")

	fmt.Println(jpg("test"))
	fmt.Println(jpg("test1.jpg"))
	fmt.Println(jpg("test2.txt"))
}
