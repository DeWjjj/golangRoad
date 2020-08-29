package main

import (
	"fmt"
	"time"
)

func main() {
	//Tick函数等于计时器，根据里面的时间给timer赋值，回来的是一种通道类型。
	timer := time.Tick(time.Second)
	for t := range timer {
		fmt.Println(t)
	}
}
