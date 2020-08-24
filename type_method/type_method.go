package main

import (
	"fmt"
)

type myInt int

func (m myInt) hello() {
	fmt.Println("This is myInt.")
}

func main() {
	var m myInt = 100
	n := myInt(200) //以myInt为类型赋值
	fmt.Println(n)
	m.hello()

}
