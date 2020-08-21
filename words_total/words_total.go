package main

import (
	"fmt"
	"strings"
)

func main() {
	sentence := "how do you do ?"
	words := strings.Split(sentence, " ")
	fmt.Println(words)

	Total := make(map[string]int, 10)
	for _, word := range words {
		Total[word] = Total[word] + 1
	}
	fmt.Println(Total)
}
