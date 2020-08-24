package main

import (
	"fmt"
)

//Dog 标示符的第一个字母大写会使得该字段可以暴露出文件，例如Println。
type Dog struct {
	name string
}

func newName(name string) Dog {
	return Dog{
		name: name,
	}

}

func (x Dog) bay() {
	fmt.Printf("%s:bay!", x.name)
}

func main() {
	name := "大黄"
	ret := newName(name)
	ret.bay()
}
