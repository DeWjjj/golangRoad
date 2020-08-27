package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	str := "DeW DeW\n"
	err := ioutil.WriteFile("./test.txt", []byte(str), 0755)
	if err != nil {
		fmt.Printf("Write file failed, err:%v.\n", err)
		return
	}
}
