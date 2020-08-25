package main

import (
	"fmt"
)

type animals interface {
	move(feet int)
	eat(name string)
}

type chicken struct {
	name string
	feet int
}

type duck struct {
	name string
	feet int
}

func (c chicken) move(feet int) {
	fmt.Printf("Chicken is moving by %v feets!\n", feet)
}

func (c chicken) eat(name string) {
	fmt.Printf("%v is eating!\n", name)
}

func (d duck) move(feet int) {
	fmt.Printf("Duck is moving by %v feets!\n", feet)
}

func (d duck) eat(name string) {
	fmt.Printf("%v is eating!\n", name)
}

func main() {
	var (
		ani1 animals
		ani2 animals
	)
	rc := chicken{
		name: "coco",
		feet: 2,
	}
	yd := duck{
		name: "dada",
		feet: 2,
	}
	ani1 = rc
	ani2 = yd
	fmt.Printf("%T\n", ani1) //main.chicken
	ani1.eat(rc.name)
	ani1.move(rc.feet)
	ani2.eat(yd.name)
	ani2.move(rc.feet)
}
