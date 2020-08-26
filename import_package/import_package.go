package main

import (
	"fmt"
	plus "golangRoad/plus"
)

func main() {
	ret := plus.Plus(1, 2) //第一个plus是文件的名字，第二个Plus是包里面暴露的函数名字，千万不要忘记签名的包名。
	fmt.Println(ret)
}
