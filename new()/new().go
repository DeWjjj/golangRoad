package main

import (
	"fmt"
)

func main() {
	var a1 *int     //*int = int类型的指针
	fmt.Println(a1) //<nil>
	var a2 = new(int)
	fmt.Println(a2)  //0xc000014088
	fmt.Println(*a2) //0
	*a2 = 100
	fmt.Println(a2, *a2)
	//new函数申请一个内存地址
}
