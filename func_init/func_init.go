package main

import (
	"fmt"
	in "golangRoad/import_init" //别名in
)

func init()  {
	fmt.Println("FUNC_INIT()")
}

func main() {
	ret := in.Plus(1,2)
	fmt.Println(ret)
}

//PACKAGE_INIT() 包中的init()先执行，因为
//FUNC_INIT()
//3