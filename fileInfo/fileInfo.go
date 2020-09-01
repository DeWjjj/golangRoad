package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	file0, err := os.OpenFile("./test.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0755)
	if err != nil {
		fmt.Printf("err: %v.\n", err)
		return
	}
	fmt.Printf("%T\n", file0) //*os.File
	for {
		_, err := file0.WriteString("1233211234567\n")
		time.Sleep(time.Second)
		if err != nil {
			fmt.Printf("err: %v.\n", err)
		}
		file0info, err := file0.Stat()
		if err != nil {
			fmt.Printf("err: %v.\n", err)
			return
		}
		fmt.Printf("File size is %d Byte.\n", file0info.Size())
		if file0info.Size() > 2000 {
			break
		}
	}

}
