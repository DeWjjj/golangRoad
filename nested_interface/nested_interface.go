package main

import (
	"fmt"
)

type animals interface { //用一个大接口统帅两个小接口
	// walk(feet int)
	// eat(feed string)
	walker
	eater
}

type walker interface {
	walk(feet int)
}

type eater interface {
	eat(feed string)
}

type cat struct {
	feed string
	feet int
}

type duck struct {
	feed string
	feet int
}

func (c *cat) walk(feet int) {
	fmt.Printf("Cat walk by %v feets.", feet)
}

func (c *cat) eat(feed string) {
	fmt.Printf("Cat eat %v.", feed)
}

func (d duck) walk(feet int) {
	fmt.Printf("Duck walk by %v feets.", feet)
}

func (d duck) eat(feed string) {
	fmt.Printf("Duck eat %v.", feed)
}

func main() {
	var (
		ani1 animals //大类型会覆盖其下的小类型，在寻找调用的时候其是一种动态类型。
		ani2 animals
	)

	bc := cat{
		feed: "fish",
		feet: 4,
	}
	yd := duck{ //值接受者无论传值还是传指针地址都可以使用。
		feed: "corn",
		feet: 2,
	}
	ani1 = &bc
	ani2 = yd

	ani1.walk(bc.feet)
	ani2.eat(bc.feed)

}
