package main

import (
	"fmt"
	"golangRoad/import_init"
)

func init()  {
	fmt.Println("FUNC_INIT()")
}

func main() {
	ret := import_init.Plus(1,2)
	fmt.Println(ret)
}

//PACKAGE_INIT() 包中的init()先执行，因为
//FUNC_INIT()
//3