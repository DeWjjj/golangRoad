package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//waitGroup

func test() {
	rand.Seed(time.Now().UnixNano()) //定义随机种子，其内需要填入的是一个int64值类型
	//如果不添加随机数种子，其随机数种子是一样的，所随机出的数值也是一样的
	for i := 0; i < 10; i++ {
		r := rand.Intn(10)  //生成一个Int值小于10的数值，赋值给r
		fmt.Println(r, 0-r) //现实随机的数
	}
}

func test0(i int) {
	defer endVal.Done()
	time.Sleep(time.Second * time.Duration(rand.Intn(5)))
	fmt.Println(i)
}

var endVal sync.WaitGroup //默认的WaitGroup的值为0，直接调用其Wait()方法直接结束

func main() {
	// test()
	for i := 0; i < 10; i++ {
		endVal.Add(1)
		go test0(i)
	}
	endVal.Wait() //等待wg的计数器减为0
}
