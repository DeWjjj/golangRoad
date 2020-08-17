package main

import (
	"fmt"
	"strings"
)

func main() {
	var (
		sentence string
	)
	sentence = "How do you do ?"
	words := strings.Split(sentence, " ")
	fmt.Println(sentence, words)
	total := make(map[string]int)

	for _, value := range words {
		total[value] = total[value] + 1
	}
	fmt.Println(total)
}
