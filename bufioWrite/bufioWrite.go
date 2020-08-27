package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("./test.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0755)
	if err != nil {
		fmt.Printf("Open file failed err:%v.\n", err)
		return
	}
	defer file.Close()
	tmp1 := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		tmp1.WriteString("ohoh~\n")
	}
	tmp1.Flush()
}
