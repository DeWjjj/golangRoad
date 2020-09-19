package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan int, 2)
	ch1 <- 1
	ch1 <- 2
	close(ch1)
	for x := range ch1 {
		fmt.Println(x)
	} //即使关闭了通道，通道中的数据也是可以读到其数值的。
	var1 := <-ch1
	fmt.Println(var1) //0
	//对关闭的通道取值，可以取出其值类型的默认值。
}
