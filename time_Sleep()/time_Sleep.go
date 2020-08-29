package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	time1 := now.Format("2006/01/02 15/04/05")
	fmt.Println(time1)
	var t int = 5
	//如果函数Sleep()里面直接是数字的话，则是自动被转换成Duration型。如果要插入自己的变量，则需要在函数内做一次转换。
	time.Sleep(time.Duration(t) * time.Second)
	now = time.Now()
	time2 := now.Format("2006/01/02 15/04/05")
	fmt.Println(time2)
}
