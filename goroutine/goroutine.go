package main

import (
	"fmt"
	"time"
)

func hello(i int) {
	fmt.Println("hello", i)

}

func main() {
	//goroutine
	//并发：同一时间你跟你的女朋友和你朋友的女朋友·聊天
	//并行：同一时间你和你朋友都在和各自的女朋友聊天
	for i := 0; i < 100; i++ {
		go hello(i)
	}
	fmt.Println("main")
	//main函数结束了，所以main函数启动的goroutine也结束了
	time.Sleep(2 * time.Second) //使用time.Sleep()使得程序休眠了一秒减缓了程序结束

}
