package main

import (
	"fmt"
)

type animals interface {
	bay() //实现bay()的是func中的method，所以method的中的变量都变成了animals接口类型的组成部分。
}

type cat struct{}
type dog struct{}
type duck struct{}

func (c cat) bay() {
	fmt.Println("Mew")
}
func (d dog) bay() {
	fmt.Println("Woof")
}
func (d duck) bay() {
	fmt.Println("Quack")
}

func main() {
	var (
		cat1  cat
		dog1  dog
		duck1 duck
	)

	var (
		ani1 animals
		ani2 animals
		ani3 animals
	)
	ani1 = cat1
	ani2 = dog1
	ani3 = duck1
	ani1.bay()
	ani2.bay()
	ani3.bay()
}
