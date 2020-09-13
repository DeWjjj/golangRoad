package main

import (
	"fmt"
	"math/rand"
	"time"
)

func test() {
	rand.Seed(time.Now().UnixNano()) //定义随机种子，其内需要填入的是一个int64值类型
	for i := 0; i < 10; i++ {
		r := rand.Intn(10)
		fmt.Println(r, 0-r) //随机复数
	}
}

func main() {
	test()
}
