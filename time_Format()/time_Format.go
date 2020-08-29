package main

import (
	"fmt"
	"time"
)

func main() {
	//Y:M:D H:M:S
	now := time.Now()
	//格式化时间，将时间中的字符串转换成字符串类型。
	fmt.Println(now.Format("2006-01-02")) //2020-08-29
	//time源码包中的信息 s := t.Format("2006-01-02 15:04:05.999999999 -0700 MST")
	//go语言中，其Format具体格式是使用2006 01 02 15 04 05，来分别代表 年 月 日 时 分 秒。
	fmt.Println(now.Format("2006/01/02 15/04/05"))        //2020/08/29 13/01/06
	fmt.Println(now.Format("2006/01/02 15/04/05 PM"))     //2020/08/29 13/01/06 PM
	fmt.Println(now.Format("2006/01/02 15/04/05.000 PM")) //2020/08/29 13/01/08.909 PM
	//go语言中按照对应格式，解析字符串时间
	timeStr, err := time.Parse("2006/01/02", "2020/08/29")
	if err != nil {
		fmt.Printf("err :%v.\n", err)
		return
	}

	fmt.Println(now.Unix()) //1598678247

	fmt.Println(timeStr)        //2020-08-29 00:00:00 +0000 UTC
	fmt.Println(timeStr.Unix()) //1598659200
}
