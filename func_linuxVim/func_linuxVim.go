package main

import (
	"fmt"
	"os"
)

func linuxVim(fileName string, fileAppend string) error {
	tmpFile := "./" + fileName
	tmp, err := os.OpenFile(tmpFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return err
	}
	defer tmp.Close()
	tmp.WriteString(fileAppend)
	return err
}

func main() {
	err := linuxVim("test.txt", "123\n")
	if err != nil {
		fmt.Printf("err:%v.\n", err)
	}
}
