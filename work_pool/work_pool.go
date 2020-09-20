package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("worker:%d start job:%d\n", id, j)
		time.Sleep(time.Second) //如果不中断一秒，该循环就会很快结束，导致goroutine重新挂起，无法达成那种穿插任务的现实
		fmt.Printf("worker:%d end job:%d\n", id, j)
		results <- j
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	// 开启3个goroutine，其内传入三个参数
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results) //这里的参数jobs是读取参数所以创建的goroutine会一直执行等到其有值
		//results则在函数内是一个写入参数，所以函数中不需要是否有值
	}
	// 5个任务
	for j := 1; j <= 5; j++ {
		jobs <- j //这里的通道会将参数交给上面挂起的goroutine，因为是五个参数所以三个goroutine执行完了之后会补上运行剩下的j参数。
	}
	close(jobs) //这里将jobs通道关闭的作用是避免函数内继续监听等待其值扩容
	// 输出结果
	for a := 1; a <= 5; a++ {
		<-results //这里需要执行五次，所以必须函数中的通道被打入五个数值
	}
}
