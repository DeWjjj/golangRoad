package main

import (
	"fmt"
	"time"
)

func timeDemo() {
	now := time.Now()
	//2020-08-28 22:30:41.730344 +0800 CST m=+0.000075288
	fmt.Println(now)
	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()

	fmt.Printf("%04d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second) //2020-08-28 22:36:04

	fmt.Println(now.Unix())     //1598625364
	fmt.Println(now.UnixNano()) //1598625364385640000

	ret := time.Unix(1598625364, 0) //2020-08-28 22:36:04 +0800 CST
	fmt.Println(ret)
	fmt.Println(ret.Year()) //2020
	fmt.Println(ret.Day())  //28

	//Add()加一个小时之后的时间。
	fmt.Println(now.Add(1 * time.Hour))
}

func main() {
	timeDemo()
	fmt.Println(time.Second) //1s
}
