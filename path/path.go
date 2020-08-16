package main

import (
	"fmt"
)

func main() {
	path := "D:\\GO\\SRC\\GITHUB\\fmt"      // D:\GO\SRC\GITHUB\fmt
	path1 := "\"D:\\GO\\SRC\\GITHUB\\fmt\"" // "D:\GO\SRC\GITHUB\fmt"
	path2 := "'D:\\GO\\SRC\\GITHUB\\fmt'"   // D:\GO\SRC\GITHUB\fmt'
	fmt.Println(path)
	fmt.Println(path1)
	fmt.Println(path2)
}
