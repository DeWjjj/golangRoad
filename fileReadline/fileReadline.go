package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func readFilebyBufil() {
	val, err := os.Open("./test.txt")
	fmt.Println(val)
	if err != nil {
		fmt.Printf("Open failed err:%v.\n", err)
		return
	}
	defer val.Close()
	tmp := bufio.NewReader(val)
	for {
		tmpLine, err := tmp.ReadString('\n')
		if err == io.EOF {
			fmt.Println("Read over.")
			return
		}
		if err != nil {
			fmt.Printf("Read failed err:%v.\n", err)
		}
		fmt.Print(tmpLine)
	}
}

func main() {
	readFilebyBufil()

}
