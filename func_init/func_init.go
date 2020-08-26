package main

import (
	i "golangRoad/import_init"i //别名
	"fmt"
)

func init()  {
	fmt.Println("FUNC_INIT()")
}

func main() {
	ret := .Plus(1,2)
	fmt.Println(ret)
}

//PACKAGE_INIT() 包中的init()先执行，因为
//FUNC_INIT()
//3