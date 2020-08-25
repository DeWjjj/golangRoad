package main

import (
	"fmt"
)

type animals interface {
	walk(feet int)
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
		ani1 animals
		ani2 animals
	)

	bc := cat{
		feed: "fish",
		feet: 4,
	}
	yd := &duck{ //值接受者无论传值还是传指针地址都可以使用。
		feed: "corn",
		feet: 2,
	}
	ani1 = &bc //如果在函数内设置了只接受指针，底下的变量体也要赋予以指针的方法赋予。
	ani2 = yd
	ani1.eat(bc.feed)
	ani1.walk(bc.feet)
	ani2.eat(yd.feed)
	ani2.walk(yd.feet)

}
