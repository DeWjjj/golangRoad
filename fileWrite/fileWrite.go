package main

import (
	"fmt"
	"os"
)

func main() {
	val, err := os.OpenFile("./test.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY|os.O_TRUNC, 0755)
	if err != nil {
		fmt.Printf("Open file failed err:%v.\n", err)
		return
	}
	val.Write([]byte("11This is a test txt.\n"))
	val.WriteString("12This is a test txt.\n")
	defer val.Close()
}
