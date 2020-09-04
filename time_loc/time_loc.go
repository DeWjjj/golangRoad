package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now) //获取本地时间
	loc, err := time.LoadLocation("Asia/Shanghai")
	fmt.Println(loc) //Asia/Shanghai
	if err != nil {
		fmt.Printf("err: %v.\n", err)
	}
	time1, err := time.ParseInLocation("2006/01/02 15/04/05", "2020/08/30 00/00/00", loc)
	if err != nil {
		fmt.Printf("err: %v.\n", err)
	}
	fmt.Println(time1)
	time0 := time1.Sub(now)
	fmt.Println(time0) //6h54m9.180698s，有时区字符串转换其值带有时区。
	time2, err := time.Parse("2006/01/02 15/04/05", "2020/08/30 00/00/00")
	if err != nil {
		fmt.Printf("err: %v.\n", err)
	}
	time0 = time2.Sub(now)
	fmt.Println(time0) //14h54m9.180698s，无时区时其计算值将会会转化成无时区的值也就是伦敦时间来计算。
}
