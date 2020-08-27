package main

import (
	"fmt"
	"io"
	"os"
)

func copyFile(xName string, yName string) error {
	xNametmp, err := os.Open(xName)
	if err != nil {
		return err
	}
	defer xNametmp.Close()
	yNametmp, err := os.OpenFile(yName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		return err
	}
	defer yNametmp.Close()
	io.Copy(yNametmp, xNametmp)
	return err
}

func main() {
	err := copyFile("test.txt", "test1.txt")
	if err != nil {
		fmt.Printf("Copy file failed err:%v.\n", err)
		return
	}
	fmt.Println("Copy file done.")
}
