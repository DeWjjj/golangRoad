package main

import (
	"fmt"
)

var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func disPatchCoins(users []string) int {
	for _, name := range users {
		disCoins := 0
		for _, letter := range name {
			switch letter {
			case 'e', 'E':
				disCoins++
			case 'i', 'I':
				disCoins += 2
			case 'o', 'O':
				disCoins += 3
			case 'u', 'U':
				disCoins += 4
			}
		}
		distribution[name] = disCoins
		coins -= disCoins
	}
	fmt.Println(distribution)
	return coins
}

func main() {
	left := disPatchCoins(users)
	fmt.Println("Left:", left)
}
