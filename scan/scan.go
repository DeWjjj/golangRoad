package main

import (
	"fmt"
)

func main() {
	var (
		name string
		age  int
		sex  string
	)
	// fmt.Scan(&name, &age, &sex)
	// fmt.Println(name, age, sex)
	// fmt.Scanf("%s %d %s\n", &name, &age, &sex)
	// fmt.Printf("Name:%s Age:%d Sex:%s\n", name, age, sex)
	fmt.Scanln(&name, &age, &sex)
	fmt.Printf("Name:%s Age:%d Sex:%s\n", name, age, sex)
}
