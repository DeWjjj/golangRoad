package main

import (
	"fmt"
)

func main() {
	setence := "黄山落叶松叶落山黄"
	seSlice := make([]rune, 0, len(setence))
	for _, val := range setence {
		seSlice = append(seSlice, val)
	}
	fmt.Printf("%c", seSlice)
	//setenceSlice := strings.Split(setence, "")
	for i := 0; i < len(seSlice)/2; i++ {
		if seSlice[i] != seSlice[len(seSlice)-1-i] {
			fmt.Println("Not palindrome!")
			return
		}
	}
	fmt.Println("It's palindrome.")
}
