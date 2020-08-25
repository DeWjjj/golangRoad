package main

import (
	"fmt"
)

func (c cat) bay() {
	fmt.Println("MeW MeW")
}

func (d dog) bay() {
	fmt.Println("Woof Woof")
}

func (d duck) bay() {
	fmt.Println("Quack Quack")
}

type cat struct{}
type dog struct{}
type duck struct{}

type all interface { //定义一个结构统帅所有函数为bay()的接口
	bay()
}

func bay(x all) { //以结构为穿参，再去调用各自的函数
	x.bay()
}

//计算机逻辑则是，先通过接口找到所有使用bay()的关联方法函数。
//然后再判断其值，符合哪一种函数。
//再调用其函数。
func main() {
	var (
		x cat
		y dog
		d duck
	)
	bay(x)
	bay(y)
	bay(d)

}
