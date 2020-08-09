package main

import (
	"fmt"
	"unicode"
)

func main() {
	str1 := "讷们好，我系百鹿。"
	var isChineseCount int
	for _, t := range str1 {
		if unicode.Is(unicode.Han, t) {
			fmt.Println("Find chinese!")
			isChineseCount++
		}
	}
	fmt.Println(isChineseCount)
}
