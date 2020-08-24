package main

import (
	"fmt"
)

type dog struct {
	name string
}

func newDog(name string) dog {
	return dog{
		name: name,
	}
}

//method方法，用户限定函数传入，存在于函数名之前

func (x dog) bay() {
	fmt.Printf("%s:汪汪汪～", x.name)
}

func main() {
	name := "大黄"
	ret := newDog(name)
	fmt.Println(ret)
	ret.bay()
}
