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
