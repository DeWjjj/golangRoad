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
	m.hello()

}
