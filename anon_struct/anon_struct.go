package main

import (
	"fmt"
)

//匿名结构体
type nameAge struct {
	string
	int //字段比较少的场景可以使用，而且确实不常用
}

func main() {
	dew := nameAge{
		"dew",
		22,
	}
	fmt.Println(dew.int, dew.string)

}
