package main

import (
	"fmt"
)

func calc(text string, x, y int) int {
	ret := x + y
	fmt.Println(text, x, y, ret)
	return ret

}

func main() {
	a := 1
	b := 2
	defer calc("text1", a, calc("text1_text10", a, b))
	a = 0
	defer calc("text2", a, calc("text2_text20", a, b))
	b = 1

	//a := 1
	//b := 2
	//defer calc("text1",1,calc("text1_text10",1,2))
	//calc("text1_text10",1,2) => fmt.Println(text,a,b,ret) => text1_text10 1 2 3
	//defer calc("text1",1,3)
	//a = 0
	//defer calc("text2", a, calc("text2_text20", a, b))
	//calc("text2_text20",0,2) => fmt.Println(text,a,b,ret) => text2_text20 0 2 2
	//defer calc("text2",0,2)
	//b = 1
	//calc("text2",0,2) >> text2 0 2 2
	//calc("text1",1,3) >> text1 1 3 4
}
